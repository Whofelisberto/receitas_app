//Package utils Funções utilitárias para manipulação de JWT
package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"os"
)

var Secret = []byte(os.Getenv("JWT_SECRET"))

// GenerateToken Gera um token JWT para o usuário fornecido
func GenerateToken(userID uint, role string) (string, error) {
	claims:= jwt.MapClaims{
		"user_id": userID,
		"role": role,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}

 token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(Secret)
}
