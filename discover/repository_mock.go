package discover

import "context"

type MockDiscoverRepository struct{}

func (u *MockDiscoverRepository) FindMatches(
	ctx context.Context,
	criteria *MatchCriteria,
) ([]*MatchEntity, error) {
	entry1 := &MatchEntity{
		Id:     1,
		Name:   "Mark",
		Gender: "male",
		Age:    23,
	}
	entry2 := &MatchEntity{
		Id:     2,
		Name:   "Joanna",
		Gender: "female",
		Age:    23,
	}

	return []*MatchEntity{entry1, entry2}, nil
}
