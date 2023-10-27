package sign

import (
	"context"
	"github.com/vildan-valeev/melvad_test/internal/domain"
	"github.com/vildan-valeev/melvad_test/internal/transport/dto"
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
func (c Service) Sign(ctx context.Context, s dto.SignDtoRequest) (domain.Sign, error) {
	var sign domain.Sign
	//return c.db.SaveSign(ctx, sign)

	return sign, nil

}
func toHMAC(text, key string) string {

}

func toHEX(string2 string) string {
	return string2
}
