package match

import "context"

var (
	_ MatchRepository = (*SQLMatchRepository)(nil)
	_ MatchRepository = (*MockMatchRepository)(nil)
)

type MatchEntity struct {
	Id     uint
	Name   string
	Gender string
	Age    uint
}

type MatchCriteria struct {
	UserId uint
	Age    string
	Gender string
}

type MatchPreferenceEntity struct {
	FromId     uint
	ToId       uint
	Preference string
}

type MatchPreferenceCriteria struct {
	UserId  uint
	MatchId uint
	// keeping string because in future we may have a new status e.g. `maybe` or `blocked`
	Preference string
}

type MatchRepository interface {
	FindMatches(context.Context, *MatchCriteria) ([]*MatchEntity, error)
	CreatePreference(context.Context, *MatchPreferenceCriteria) (*MatchPreferenceEntity, error)
	FindPreference(context.Context, *MatchPreferenceCriteria) (*MatchPreferenceEntity, error)
}
