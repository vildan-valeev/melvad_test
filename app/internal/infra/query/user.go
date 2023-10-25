package query

import (
	"context"
	"github.com/vildan-valeev/melvad_test/internal/domain"
	"github.com/vildan-valeev/melvad_test/internal/infra"

	"github.com/jackc/pgx/v5/pgtype"
)

func UpdateUser(ctx context.Context, db infra.DB, id int64) error {
	return nil
}

func InsertUser(ctx context.Context, db infra.DB, u domain.User) error {
	tx, err := db.Begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		_ = tx.Rollback(ctx)
	}()

	var (
		uid      int64
		updateAt pgtype.Timestamp
	)

	if err := tx.QueryRow(ctx,
		`INSERT INTO sgn_user (profile_id, last_name, first_name, middle_name, phone, email, position, department, avatar, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) ON CONFLICT ON CONSTRAINT sgn_user_profile_id_key DO UPDATE SET last_name=EXCLUDED.last_name, first_name=EXCLUDED.first_name, middle_name=EXCLUDED.middle_name, phone=EXCLUDED.phone, email=EXCLUDED.email, position=EXCLUDED.position, department=EXCLUDED.department, updated_at=$11 RETURNING uid, updated_at`,
		u.ProfileID,
		u.LastName,
		u.FirstName,
		u.MiddleName,
		u.Phone,
		u.Email,
		u.Position,
		u.Department,
		u.Avatar,
		u.CreatedAt,
		u.UpdatedAt,
	).Scan(&uid, &updateAt); err != nil {
		return dto.User{}, err
	}

	user := dto.User{
		UID:        uid,
		ProfileID:  u.ProfileID,
		LastName:   u.LastName,
		FirstName:  u.FirstName,
		MiddleName: u.MiddleName,
		Phone:      u.Phone,
		Email:      u.Email,
		Department: u.Department,
		Position:   u.Position,
		Avatar:     u.Avatar,
		CreatedAt:  u.CreatedAt,
		UpdatedAt:  updateAt,
	}

	return user, tx.Commit(ctx)
}
