package sign

import (
	"context"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"github.com/rs/zerolog/log"
	"github.com/vildan-valeev/melvad_test/internal/domain"
	"github.com/vildan-valeev/melvad_test/internal/transport/dto"
	"io"
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

// Sign Создание подписи.
func (c Service) Sign(ctx context.Context, s dto.SignDtoRequest) (domain.Sign, error) {
	var sign domain.Sign

	sign.Hash = encode(s.Text, s.Key)

	return sign, nil

}

// To HMAC, then to HEX
func encode(text, key string) string {
	hash := hmac.New(sha512.New, []byte(key))
	_, err := io.WriteString(hash, text)
	if err != nil {
		log.Info().Err(err)
	}
	return hex.EncodeToString(hash.Sum(nil))
}
