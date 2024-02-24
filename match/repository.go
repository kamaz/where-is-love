package match

import "context"

var (
	_ MatchRepository = (*sqlMatchRepository)(nil)
	_ MatchRepository = (*mockMatchRepository)(nil)
)

// MatchEntity represents a match entity
type MatchEntity struct {
	Id             uint
	Name           string
	Gender         string
	Age            uint
	DistanceFromMe uint
}

// MatchCriteria represents criteria for getting matches
type MatchCriteria struct {
	UserId    uint
	Longitude float64
	Latitude  float64
	Age       string
	Gender    string
}

// MatchPreferenceEntity represents a match preference entity between two users and their preference 'YES' and 'NO'
type MatchPreferenceEntity struct {
	FromId     uint
	ToId       uint
	Preference string
}

// MatchPreferenceCriteria represents criteria for creating and finding match preference
type MatchPreferenceCriteria struct {
	UserId  uint
	MatchId uint
	// keeping string because in future we may have a new status e.g. `maybe` or `blocked`
	Preference string
}

// Sort represents sorting criteria
type Sort struct {
	Property string
	Asc      bool
}

// MatchRepository is an interface for match repository
type MatchRepository interface {
	FindMatches(context.Context, *MatchCriteria, *Sort) ([]*MatchEntity, error)
	CreatePreference(context.Context, *MatchPreferenceCriteria) (*MatchPreferenceEntity, error)
	FindPreference(context.Context, *MatchPreferenceCriteria) (*MatchPreferenceEntity, error)
}
