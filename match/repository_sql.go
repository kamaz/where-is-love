package match

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateSQLMatchRepository(pool *pgxpool.Pool) MatchRepository {
	return &SQLMatchRepository{db: pool}
}

type SQLMatchRepository struct {
	db *pgxpool.Pool
}

// It should return other profiles that are potential matches for this user.
// Exclude profiles you’ve already swiped on.
// select * app_user join swipes on

func (u *SQLMatchRepository) createQueryAndParams(
	criteria *MatchCriteria,
) (string, []any) {
	filters := []string{}
	filterValues := []any{criteria.UserId}
	if criteria.Age != "" && criteria.Gender != "" {
		filters = append(filters, "AND age = $2 AND gender = $3")
		filterValues = append(filterValues, criteria.Age, criteria.Gender)
	} else {
		if criteria.Age != "" {
			filters = append(filters, "AND age = $2")
			filterValues = append(filterValues, criteria.Age)
		} else if criteria.Gender != "" {
			filters = append(filters, "AND gender = $2")
			filterValues = append(filterValues, criteria.Gender)
		}
	}
	query := fmt.Sprintf(
		"SELECT id, name, gender, age FROM app_user WHERE id != $1 %s AND id NOT IN (SELECT to_id FROM user_preference WHERE from_id = $1)",
		strings.Join(filters, " "),
	)

	return query, filterValues
}

// matches
// { gender, age  }
func (u *SQLMatchRepository) FindMatches(
	ctx context.Context,
	criteria *MatchCriteria,
) ([]*MatchEntity, error) {
	query, filterValues := u.createQueryAndParams(criteria)
	rows, err := u.db.Query(
		ctx,
		query,
		filterValues...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var matches []*MatchEntity
	for rows.Next() {
		var match MatchEntity
		if err := rows.Scan(&match.Id, &match.Name, &match.Gender,
			&match.Age); err != nil {
			return nil, err
		}
		matches = append(matches, &match)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return matches, nil
}

func (u *SQLMatchRepository) CreatePreference(
	ctx context.Context,
	criteria *MatchPreferenceCriteria,
) (*MatchPreferenceEntity, error) {
	_, err := u.db.Exec(
		ctx,
		"INSERT INTO user_preference(from_id, to_id, preference) VALUES ($1, $2, $3)",
		criteria.UserId,
		criteria.MatchId,
		criteria.Preference,
	)
	if err != nil {
		return nil, err
	}

	return &MatchPreferenceEntity{
		FromId:     criteria.UserId,
		ToId:       criteria.MatchId,
		Preference: criteria.Preference,
	}, nil
}

func (u *SQLMatchRepository) FindPreference(
	ctx context.Context,
	criteria *MatchPreferenceCriteria,
) (*MatchPreferenceEntity, error) {
	var preference MatchPreferenceEntity
	err := u.db.QueryRow(
		ctx,
		"SELECT from_id, to_id, preference FROM user_preference WHERE from_id = $1 AND to_id = $2 AND preference = $3",
		criteria.UserId,
		criteria.MatchId,
		criteria.Preference,
	).Scan(&preference.FromId, &preference.ToId, &preference.Preference)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return &MatchPreferenceEntity{
				FromId:     criteria.UserId,
				ToId:       criteria.MatchId,
				Preference: PreferenceNo,
			}, nil
		}
		return nil, err
	}
	return &preference, nil
}
