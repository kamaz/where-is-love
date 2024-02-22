package discover

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateSQLDiscoveryRepository(pool *pgxpool.Pool) DiscoverRepository {
	return &SQLDiscoverRepository{db: pool}
}

type SQLDiscoverRepository struct {
	db *pgxpool.Pool
}

// It should return other profiles that are potential matches for this user.
// Exclude profiles you’ve already swiped on.
// select * app_user join swipes on

// matches
// { gender, age  }
func (u *SQLDiscoverRepository) FindMatches(
	ctx context.Context,
	criteria *MatchCriteria,
) ([]*MatchEntity, error) {
	rows, err := u.db.Query(
		ctx,
		"SELECT id, name, gender, age FROM app_user WHERE id != $1",
		criteria.UserId,
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
