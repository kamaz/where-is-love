package swipe

import (
	"net/http"

	"github.com/labstack/echo"
)

/*
/swipe

You should specify the other user id + a preference (YES or NO).
It should store and return if there was a match (both users swipe YES).

todo: renamed `results` to `result` to be consistent with other endpoints
	{
	    "results": {
	        "matched": <bool>,
	        "matchID": <integer>
	    }
	}

Note: matchID must only be included if matched is true.
*/

type SwipeRequest struct {
	UserId     int    `json:"userId"`
	Preference string `json:"preference"`
}

type SwipeResponse struct {
	Matched bool `json:"matched,omitempty"`
	MatchID int  `json:"matchID,omitempty"`
}

type SwipeResult struct {
	Result *SwipeResponse `json:"result"`
}

type SwipeEndpoint struct {
	middlewares []echo.MiddlewareFunc
}

func (u *SwipeEndpoint) Process(e echo.Context) error {
	var swipeRequest SwipeRequest
	if err := e.Bind(&swipeRequest); err != nil {
		return err
	}

	result := SwipeResult{Result: &SwipeResponse{}}
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
