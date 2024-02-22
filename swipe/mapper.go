package swipe

import "errors"

var ErrInvalidPreference = errors.New("invalid preference")

const (
	PreferenceYes = "YES"
	PreferenceNo  = "No"
)

func toCreateMatchCriteria(userId uint, match *SwipeRequest) (*MatchCriteria, error) {
	if match.Preference != PreferenceYes && match.Preference != PreferenceNo {
		return nil, ErrInvalidPreference
	}
	criteria := &MatchCriteria{
		UserId:     userId,
		MatchId:    match.UserId,
		Preference: match.Preference,
	}
	return criteria, nil
}

func toSwipeResponse(preference *MatchPreferenceEntity) *SwipeResponse {
	response := &SwipeResponse{}

	if preference.Preference == PreferenceYes {
		response.Matched = true
		response.MatchId = preference.FromId
	}

	return response
}
