package user

import (
	"context"
	"github.com/vildan-valeev/melvad_test/internal/domain"
	"github.com/vildan-valeev/melvad_test/internal/infra/pg"
)

func (r Repository) InsertUser(ctx context.Context, u domain.User) (id int64, err error) {
	id, err = pg.InsertUser(ctx, r.db, u)
	if err != nil {
		return id, err
	}

	return id, nil
}

func (r Repository) UpdateUser(ctx context.Context, id int64) error {
	err := pg.UpdateUser(ctx, r.db, id)
	if err != nil {
		return err
	}

	return nil
}
