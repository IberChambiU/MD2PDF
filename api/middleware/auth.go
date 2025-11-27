package middleware

import (
	"strings"

	// p "github.com/bruno5200/TSM/api/presenter"
	d "github.com/IberChambiU/MD2PDF/api/token/domain"
	"github.com/IberChambiU/MD2PDF/api/token/entities"
	p "github.com/IberChambiU/MD2PDF/api/token/infrastructure/presenter"
	"github.com/IberChambiU/MD2PDF/env"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

var (
	ADMIN = "admin"
)

const (
	tokenBearer = "Bearer"
)

func Admin() fiber.Handler {
	return func(c *fiber.Ctx) error {

		payload, err := Token(utils.CopyString(c.Get(fiber.HeaderAuthorization)))

		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(p.ErrorResponse(fiber.ErrNotFound))
		}

		// log.Printf("UserType: %s UserId: %s", payload.Username, payload.Id)

		if payload.Username != ADMIN {
			return c.Status(fiber.StatusNotFound).JSON(p.ErrorResponse(fiber.ErrNotFound))
		}

		return c.Next()
	}
}

func Token(auth string) (*entities.Payload, error) {

	e := env.Env()

	bearer := strings.Split(auth, " ")

	if len(bearer) != 2 || bearer[0] != tokenBearer {
		return nil, fiber.ErrNotFound
	}

	token := bearer[1]
	// log.Printf("token: %s", token)
	maker, err := d.NewPasetoMaker(e.GetSimetricKey())

	if err != nil {
		return nil, fiber.ErrNotFound
	}

	return maker.ValidateToken(token)
}
