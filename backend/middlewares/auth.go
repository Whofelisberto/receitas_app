//Package middlewares Middleware de autenticação
package middlewares

import (
	"receitasfit/backend/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/jwt/v3"
)

func JWTProtected() fiber.Handler {
	return jwt.New(jwt.Config{
		SigningKey:   utils.GetSecret(),
		TokenLookup:  "cookie:token",
	})
}
