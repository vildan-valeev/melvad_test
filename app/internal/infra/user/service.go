package user

import (
	"github.com/vildan-valeev/melvad_test/internal/infra"
)

type Repository struct {
	db infra.DB
	rs infra.RedisCache
}

func New(db infra.DB, rs infra.RedisCache) *Repository {
	return &Repository{
		db: db,
		rs: rs,
	}
}
