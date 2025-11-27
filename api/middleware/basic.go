package middleware

import (
	"encoding/base64"
	"errors"
	"strings"

	env "github.com/IberChambiU/MD2PDF/env"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

var e = env.Env()

var (
	ErrMailformedAuthHeader = errors.New("mailformed authorization header")
)

const (
	basic = "Basic"
)

func BasicAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {

		// decode base64 string
		user, pass, err := basicAuth(utils.CopyString(c.Get(fiber.HeaderAuthorization)))

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"error":   "Invalid Authorization header",
			})
		}

		if user != e.GetMiddlewareUser() && pass != e.GetMiddlewarePassword() {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"error":   "Invalid username or password",
			})
		}
		return c.Next()
	}
}

func basicAuth(auth string) (username, password string, err error) {

	cred := strings.Split(auth, " ")

	if len(cred) != 2 || cred[0] != basic {
		return "", "", ErrMailformedAuthHeader
	}
	// Split the base64 encoded string
	credentials, err := base64.StdEncoding.DecodeString(cred[1])

	if err != nil {
		return "", "", err
	}

	// Split the string into username and password
	credss := strings.Split(string(credentials), ":")

	if len(credss) != 2 {
		return "", "", err
	}

	return credss[0], credss[1], nil
}
