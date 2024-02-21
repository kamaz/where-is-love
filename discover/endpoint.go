package discover

import (
	"net/http"

	"github.com/labstack/echo"
)

// /discover
// It should return other profiles that are potential matches for this user.
// Exclude profiles you’ve already swiped on.
// {
//     "results": [
//         {
//             "id": <integer>,
//             "name": <string>,
//             "gender": <string>,
//             "age": <integer>
//         },
//         ...
//     ]
// }

// i) Extend /discover to filter results by age and or gender.
//
// ii) Extend /discover to sort profiles by distance from the authenticated user.
// You will need to add location to the user model.
// Add a "distanceFromMe" to the /discover results.
//
// iii) Bonus: Extend /discover to sort profiles by attractiveness.
// You will need to come up with a ranking based on swipe statistics.
//

type MatchResult struct {
	Results []*MatchResponse `json:"results"`
}

type MatchResponse struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Age    int    `json:"age"`
}

type DiscoverEndpoint struct {
	middlewares []echo.MiddlewareFunc
}

func (u *DiscoverEndpoint) Process(e echo.Context) error {
	// as part of context we will get user but at the moment lets return response
	result := &MatchResult{[]*MatchResponse{}}
	e.JSON(http.StatusOK, result)
	return nil
}

func (u *DiscoverEndpoint) Method() string {
	return "GET"
}

func (u *DiscoverEndpoint) Path() string {
	return "/discover"
}

func (u *DiscoverEndpoint) Middlewares() []echo.MiddlewareFunc {
	return u.middlewares
}
