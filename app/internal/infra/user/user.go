package user

import (
	"context"
	"github.com/vildan-valeev/melvad_test/internal/domain"
	"github.com/vildan-valeev/melvad_test/internal/infra/pg"
	"github.com/vildan-valeev/melvad_test/internal/infra/rs"
)

func (r Repository) InsertUser(ctx context.Context, u domain.User) (id int64, err error) {
	id, err = pg.InsertUser(ctx, r.db, u)
	if err != nil {
		return id, err
	}

	return id, nil
}

func (r Repository) UpdateUser(ctx context.Context, u domain.User) error {
	err := pg.UpdateUser(ctx, r.db, u)
	if err != nil {
		return err
	}

	return nil
}

func (r Repository) UpdateUserInCache(ctx context.Context, u domain.User) error {
	err := rs.UpdateUserInCache(ctx, r.rs, u)
	if err != nil {
		return err
	}

	return nil
}
