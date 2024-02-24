package match

import (
	"context"
	"strconv"
)

type mockMatchRepository struct {
	users map[uint]*MatchPreferenceEntity
}

func CreateMockMatchRepository() *mockMatchRepository {
	return &mockMatchRepository{
		users: map[uint]*MatchPreferenceEntity{
			1: {
				FromId:     2,
				ToId:       1,
				Preference: "YES",
			},
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

func (u *mockMatchRepository) FindMatches(
	ctx context.Context,
	criteria *MatchCriteria,
	sort *Sort,
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
		Age:    25,
	}
	entries := []*MatchEntity{entry1, entry2}
	result := []*MatchEntity{}
	// fitler by criteria
	if criteria.Gender != "" {
		for _, entry := range entries {
			if entry.Gender == criteria.Gender {
				result = append(result, entry)
			}
		}
	} else {
		result = entries
	}

	finalResult := []*MatchEntity{}
	if criteria.Age != "" {
		for _, entry := range result {
			value, err := strconv.Atoi(criteria.Age)
			if err != nil {
				continue
			}
			if entry.Age == uint(value) {
				finalResult = append(finalResult, entry)
			}
		}
	} else {
		finalResult = result
	}

	return finalResult, nil
}

func (u *mockMatchRepository) CreatePreference(
	ctx context.Context,
	criteria *MatchPreferenceCriteria,
) (*MatchPreferenceEntity, error) {
	return u.users[criteria.UserId], nil
}

func (u *mockMatchRepository) FindPreference(
	ctx context.Context,
	criteria *MatchPreferenceCriteria,
) (*MatchPreferenceEntity, error) {
	return u.users[criteria.UserId], nil
}
