// Package middlewares provê middleware de autenticação JWT.
package middlewares
import (
    "receitas_app/backend/utils"
    "github.com/gofiber/fiber/v2"
    "github.com/golang-jwt/jwt/v5"
)
func JWTProtected() fiber.Handler {
    return func(c *fiber.Ctx) error {
        tokenString := c.Get("Authorization")
        if tokenString == "" {
            return c.Status(401).JSON(fiber.Map{"error": "Token ausente"})
        }

        // Remove "Bearer "
        if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
            tokenString = tokenString[7:]
        }

        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fiber.NewError(401, "Método de assinatura inválido")
            }
            return []byte(utils.Secret), nil
        })
        if err != nil || !token.Valid {
            return c.Status(401).JSON(fiber.Map{"error": "Token inválido"})
        }

        // Pega as claims e salva no Locals
        if claims, ok := token.Claims.(jwt.MapClaims); ok {
            c.Locals("user", claims)
            role, ok := claims["role"].(string)
						if !ok || role != "admin" {
            return c.Status(403).JSON(fiber.Map{"error": "Apenas administradores podem criar receitas"})
            }
        } else {
            return c.Status(401).JSON(fiber.Map{"error": "Claims inválidas"})
        }

        return c.Next()
    }
}
