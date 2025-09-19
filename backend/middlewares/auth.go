//Package middlewares Middleware de autenticação
package middlewares

import (
	"receitasfitness/backend/utils"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func JWTProtected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   utils.GetSecret(),
		TokenLookup:  "cookie:token",
	})
}
