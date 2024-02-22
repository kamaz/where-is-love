package swipe

import (
	"context"
)

type MockSwipeRepository struct {
	users map[uint]*MatchPreferenceEntity
}

func CreateMockSwipeRepository() *MockSwipeRepository {
	return &MockSwipeRepository{
		users: map[uint]*MatchPreferenceEntity{
			2: {
				FromId:     2,
				ToId:       1,
				Preference: "NO",
			},
			3: {
				FromId:     3,
				ToId:       1,
				Preference: "YES",
			},
		},
	}
}

func (u *MockSwipeRepository) CreatePreference(
	ctx context.Context,
	criteria *MatchCriteria,
) (*MatchPreferenceEntity, error) {
	return nil, nil
}

func (u *MockSwipeRepository) FindPreference(
	ctx context.Context,
	criteria *MatchCriteria,
) (*MatchPreferenceEntity, error) {
	return u.users[criteria.UserId], nil
}
