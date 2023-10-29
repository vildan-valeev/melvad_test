package v1

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/vildan-valeev/melvad_test/internal/domain"
	"github.com/vildan-valeev/melvad_test/internal/transport/dto"
)

type Transport struct {
	user User
	sign Sign
}
type DI struct {
	User User
	Sign Sign
}

func NewTransport(di DI) *Transport {
	return &Transport{
		user: di.User,
		sign: di.Sign,
	}
}

func (t *Transport) Register() *fiber.App {
	app := fiber.New()

	app.Post("/redis/incr", t.UserUpdate)
	app.Post("/postgres/users", t.UserCreate)
	app.Post("/sign/hmacsha512", t.Sign)

	return app
}

/*
Интерфейсы от бизнес слоя.
*/

type Sign interface {
	Sign(ctx context.Context, s dto.SignDtoRequest) (domain.Sign, error)
}

type User interface {
	CreateUser(ctx context.Context, u dto.UserCreateDtoRequest) (int64, error)
	UpdateUserInCache(ctx context.Context, u *dto.UserUpdateDtoRequest) (uint8, error)
}
