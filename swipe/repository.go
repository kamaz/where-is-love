package swipe

import "context"

// maybe swipe and discover should be in the same package e.g. matching
var _ SwipeRepository = (*SQLSwipeRepository)(nil)

type MatchPreferenceEntity struct {
	FromId     uint
	ToId       uint
	Preference string
}

type MatchCriteria struct {
	UserId  uint
	MatchId uint
	// keeping string because in future we may have a new status e.g. `maybe` or `blocked`
	Preference string
}

type SwipeRepository interface {
	CreatePreference(context.Context, *MatchCriteria) (*MatchPreferenceEntity, error)
	FindPreference(context.Context, *MatchCriteria) (*MatchPreferenceEntity, error)
}
