package match

import (
	"errors"
	"strings"
)

var ErrInvalidPreference = errors.New("invalid preference")

const (
	PreferenceYes = "YES"
	PreferenceNo  = "NO"
)

func toCreateMatchCriteria(userId uint, match *SwipeRequest) (*MatchPreferenceCriteria, error) {
	if match.Preference != PreferenceYes && match.Preference != PreferenceNo {
		return nil, ErrInvalidPreference
	}
	criteria := &MatchPreferenceCriteria{
		UserId:     userId,
		MatchId:    match.UserId,
		Preference: match.Preference,
	}
	return criteria, nil
}

func toSwipeResponse(myPreference, datePreference *MatchPreferenceEntity) *SwipeResponse {
	response := &SwipeResponse{}

	if myPreference.Preference == PreferenceYes && datePreference.Preference == PreferenceYes {
		response.Matched = true
		response.MatchId = datePreference.FromId
	}

	return response
}

func toSort(sort string) *Sort {
	ascSort := strings.Split(sort, "+")
	if len(ascSort) == 2 && ascSort[1] == "distanceFromMe" {
		return &Sort{
			Property: ascSort[1],
			Asc:      true,
		}
	}

	descSort := strings.Split(sort, "-")
	if len(descSort) == 2 && descSort[1] == "distanceFromMe" {
		return &Sort{
			Property: descSort[1],
			Asc:      false,
		}
	}

	return nil
}
