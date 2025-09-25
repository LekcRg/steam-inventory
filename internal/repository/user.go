package repository

import (
	"context"
	"errors"

	"github.com/LekcRg/steam-inventory/internal/models"
)

func (r *Repo) CreateOrUpdateUser(
	ctx context.Context, user *models.User,
) (*models.User, error) {
	sql := `INSERT INTO users(
		steamid, personaname, avatar, realname,
		communityvisibilitystate, lastlogoff_steam,
		timecreated_steam
	)
	VALUES (
		:steamid, :personaname, :avatar, :realname,
		:communityvisibilitystate, :lastlogoff_steam,
		:timecreated_steam
	)
	ON CONFLICT (steamid)
	DO UPDATE SET
		personaname = :personaname,
		avatar = :avatar,
		realname = :realname,
		communityvisibilitystate = :communityvisibilitystate,
		lastlogoff_steam = :lastlogoff_steam,
		updated_at = NOW()
	RETURNING *`

	rows, err := r.db.NamedQueryContext(ctx, sql, user)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, err
		}

		return nil, errors.New("test")
	}

	var u models.User

	err = rows.StructScan(&u)
	if err != nil {
		return nil, err
	}

	return &u, nil
}
