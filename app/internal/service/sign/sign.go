package sign

import (
	"context"
	"github.com/vildan-valeev/melvad_test/internal/domain"
)

// Repository - методы для работы с БД (интерфейс реализован в инфре)
type Repository interface {
	SaveSign(ctx context.Context, sign domain.Sign) error // Не задействован!!!
}

type Service struct {
	db Repository
}

func New(db Repository) *Service {
	return &Service{
		db: db,
	}
}

// Sign Создание запроса на платеж.
func (c Service) Sign(ctx context.Context, sign domain.Sign) error {
	return c.db.SaveSign(ctx, sign)
}
