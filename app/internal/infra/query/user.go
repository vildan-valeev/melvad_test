package query

import (
	"context"
	"github.com/vildan-valeev/melvad_test/internal/domain"
	"github.com/vildan-valeev/melvad_test/internal/infra"
)

func UpdateUser(ctx context.Context, db infra.DB, id int64) error {
	return nil
}

func InsertUser(ctx context.Context, db infra.DB, u domain.User) (int64, error) {
	var id int64

	tx, err := db.Begin(ctx)
	if err != nil {
		return id, err
	}

	defer func() {
		_ = tx.Rollback(ctx)
	}()

	if err := tx.QueryRow(ctx,
		`INSERT INTO users (name, age) VALUES ($1, $2) ON CONFLICT ON CONSTRAINT users_id_key DO UPDATE SET name=EXCLUDED.name, age=EXCLUDED.age RETURNING id`,
		u.Name,
		u.Age,
	).Scan(&id); err != nil {
		return id, err
	}

	return id, tx.Commit(ctx)
}
