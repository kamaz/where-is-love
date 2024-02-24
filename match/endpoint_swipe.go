package match

import (
	"fmt"
	"net/http"

	"github.com/kamaz/where-is-love/user"
	"github.com/labstack/echo"
)

/*
/swipe

You should specify the other user id + a preference (YES or NO).
It should store and return if there was a match (both users swipe YES).

todo: renamed `results` to `result` to be consistent with other endpoints
	{
	    "result": {
	        "matched": <bool>,
	        "matchID": <integer>
	    }
	}

Note: matchID must only be included if matched is true.
*/

func CreateSwipeEndpoint(
	repo MatchRepository,
	middlewares ...echo.MiddlewareFunc,
) *swipeEndpoint {
	return &swipeEndpoint{
		repository:  repo,
		middlewares: middlewares,
	}
}

type SwipeRequest struct {
	UserId     uint   `json:"userID"`
	Preference string `json:"preference"`
}

// todo: discuss the idea of match especially when we don't know if someone has matched yet
// it is possible that we don't know if they matched yet as we are waiting for the other user to swipe
type SwipeResponse struct {
	Matched bool `json:"matched"`
	MatchId uint `json:"matchID,omitempty"`
}

type SwipeResult struct {
	Result *SwipeResponse `json:"result"`
}

type swipeEndpoint struct {
	middlewares []echo.MiddlewareFunc
	repository  MatchRepository
}

func (u *swipeEndpoint) Process(e echo.Context) error {
	var swipeRequest SwipeRequest
	if err := e.Bind(&swipeRequest); err != nil {
		return fmt.Errorf("failed to read payload %w", err)
	}

	ctx := e.Request().Context()
	user := ctx.Value(user.UserKey).(*user.UserToken)

	// what should happen if you swipe again we should just return error
	// for simplicity we will just DB to throw error but ideally we would have
	// some validation
	matchCriteria, err := toCreateMatchCriteria(user.Id, &swipeRequest)
	if err != nil {
		return fmt.Errorf("failed to create match %w", err)
	}

	myPreference, err := u.repository.CreatePreference(ctx, matchCriteria)
	if err != nil {
		return fmt.Errorf("failed to create preference %w", err)
	}

	datePreference, err := u.repository.FindPreference(
		ctx,
		&MatchPreferenceCriteria{
			UserId:     swipeRequest.UserId,
			MatchId:    user.Id,
			Preference: PreferenceYes,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to find preference %w", err)
	}

	response := toSwipeResponse(myPreference, datePreference)

	result := SwipeResult{Result: response}
	e.JSON(http.StatusOK, result)
	return nil
}

func (u *swipeEndpoint) Method() string {
	return "POST"
}

func (u *swipeEndpoint) Path() string {
	return "/swipe"
}

func (u *swipeEndpoint) Middlewares() []echo.MiddlewareFunc {
	return u.middlewares
}
