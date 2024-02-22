package match

import "context"

type MockMatchRepository struct {
	users map[uint]*MatchPreferenceEntity
}

func CreateMockMatchRepository() *MockMatchRepository {
	return &MockMatchRepository{
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

func (u *MockMatchRepository) FindMatches(
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

func (u *MockMatchRepository) CreatePreference(
	ctx context.Context,
	criteria *MatchPreferenceCriteria,
) (*MatchPreferenceEntity, error) {
	return nil, nil
}

func (u *MockMatchRepository) FindPreference(
	ctx context.Context,
	criteria *MatchPreferenceCriteria,
) (*MatchPreferenceEntity, error) {
	return u.users[criteria.UserId], nil
}
