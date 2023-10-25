package user

import (
	"github.com/vildan-valeev/melvad_test/internal/infra"
)

type Repository struct {
	db infra.DB
}

func New(db infra.DB) *Repository {
	return &Repository{db: db}
}
