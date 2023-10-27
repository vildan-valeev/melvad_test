package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"github.com/vildan-valeev/melvad_test/internal/transport/dto"
	"net/http"
)

func (t *Transport) UserUpdate(c *fiber.Ctx) error {
	user := new(dto.UserUpdateDtoRequest)

	if err := c.BodyParser(user); err != nil {
		log.Error().Msgf("Ошибка парсинга входящих данных: %v ", err)
		return c.SendStatus(http.StatusBadRequest)
	}

	age, err := t.user.UpdateUserInCache(c.Context(), user)
	if err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}

	response := dto.UserUpdateDtoResponse{Value: age}

	return c.Status(http.StatusOK).JSON(response)
}
