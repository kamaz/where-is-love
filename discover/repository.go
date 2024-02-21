package discover

import "context"

var (
	_ DiscoverRepository = (*SQLDiscoverRepository)(nil)
	_ DiscoverRepository = (*MockDiscoverRepository)(nil)
)

type MatchEntity struct {
	Id     int
	Name   string
	Gender string
	Age    uint
}

type MatchCriteria struct {
	UserId  int
	AgeFrom uint
	AgeTo   uint
	Gender  string
}

type DiscoverRepository interface {
	FindMatches(context.Context, *MatchCriteria) ([]*MatchEntity, error)
}
