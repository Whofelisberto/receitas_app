// Package utils Funções utilitárias para hashing de senhas
package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword Gera um hash bcrypt para a senha fornecida
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password),10)
	return string(bytes), err
}

// CheckPasswordHash Compara a senha fornecida com o hash armazenado
func CheckPasswordHash(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
