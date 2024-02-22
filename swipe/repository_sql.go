package swipe

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// todo: rename to MatchRepository
func CreateSQLSwipeRepository(pool *pgxpool.Pool) SwipeRepository {
	return &SQLSwipeRepository{db: pool}
}

type SQLSwipeRepository struct {
	db *pgxpool.Pool
}

// save match
// get match
func (u *SQLSwipeRepository) CreatePreference(
	ctx context.Context,
	criteria *MatchCriteria,
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

func (u *SQLSwipeRepository) FindPreference(
	ctx context.Context,
	criteria *MatchCriteria,
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
