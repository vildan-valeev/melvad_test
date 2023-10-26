package v1

import (
	"github.com/rs/zerolog/log"
	"github.com/vildan-valeev/melvad_test/internal/transport/dto"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (t *Transport) Sign(c *fiber.Ctx) error {
	s := new(dto.SignDtoRequest)

	if err := c.BodyParser(s); err != nil {
		log.Error().Msgf("Ошибка парсинга входящих данных: %v ", err)
		return c.SendStatus(http.StatusBadRequest)
	}

	sign, err := t.sign.Sign(c.Context(), *s)
	if err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}

	return c.Status(http.StatusOK).JSON(dto.ToDTO(sign))
}
