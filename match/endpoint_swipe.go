package match

import (
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
) *SwipeEndpoint {
	return &SwipeEndpoint{
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

type SwipeEndpoint struct {
	middlewares []echo.MiddlewareFunc
	repository  MatchRepository
}

func (u *SwipeEndpoint) Process(e echo.Context) error {
	var swipeRequest SwipeRequest
	if err := e.Bind(&swipeRequest); err != nil {
		return err
	}

	ctx := e.Request().Context()
	user := ctx.Value(user.UserKey).(*user.UserToken)

	// what should happen if you swipe again we should just return error
	// for simplicity we will just DB to throw error but ideally we would have
	// some validation
	matchCriteria, err := toCreateMatchCriteria(user.Id, &swipeRequest)
	_, err = u.repository.CreatePreference(ctx, matchCriteria)
	if err != nil {
		return err
	}

	matchResult, err := u.repository.FindPreference(
		ctx,
		&MatchPreferenceCriteria{
			UserId:     swipeRequest.UserId,
			MatchId:    user.Id,
			Preference: PreferenceYes,
		},
	)
	if err != nil {
		return err
	}

	response := toSwipeResponse(matchResult)

	result := SwipeResult{Result: response}
	e.JSON(http.StatusOK, result)
	return nil
}

func (u *SwipeEndpoint) Method() string {
	return "POST"
}

func (u *SwipeEndpoint) Path() string {
	return "/swipe"
}

func (u *SwipeEndpoint) Middlewares() []echo.MiddlewareFunc {
	return u.middlewares
}
