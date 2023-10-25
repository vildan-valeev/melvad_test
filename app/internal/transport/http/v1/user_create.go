package v1

import (
	"github.com/rs/zerolog/log"
	"github.com/vildan-valeev/melvad_test/internal/transport/dto"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (t *Transport) UserCreate(c *fiber.Ctx) error {
	user := new(dto.UserCreateDtoRequest)

	if err := c.BodyParser(user); err != nil {
		log.Error().Msgf("Ошибка парсинга входящих данных: %v ", err)
		return c.SendStatus(http.StatusBadRequest)
	}

	id, err := t.user.CreateUser(c.Context(), *user)
	if err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}
	response := dto.UserCreateDtoResponse{ID: id}
	return c.Status(http.StatusOK).JSON(response)
}
