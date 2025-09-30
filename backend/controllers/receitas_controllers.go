// Package controllers Lida com as requisições relacionadas às receitas
package controllers

import (
	"receitas_app/backend/config"
	"receitas_app/backend/models"
	"github.com/golang-jwt/jwt/v5"

	"github.com/gofiber/fiber/v2"
)

func CreateReceita(c *fiber.Ctx) error {
	// Pega as claims diretamente do Locals (definido no middleware)
	claims, ok := c.Locals("user").(jwt.MapClaims)
if !ok {
    return c.Status(401).JSON(fiber.Map{"error": "Usuário não autenticado"})
}

userIDFloat, ok := claims["user_id"].(float64)
if !ok {
    return c.Status(400).JSON(fiber.Map{"error": "user_id inválido"})
}
userID := uint(userIDFloat)

// Verifica se o usuário é admin
role, ok := claims["role"].(string)
if !ok || role != "admin" {
    return c.Status(403).JSON(fiber.Map{"error": "Apenas administradores podem criar receitas"})
}

	// Parse do corpo da requisição
	var receita models.Receita
	if err := c.BodyParser(&receita); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	// Define o usuário que criou a receita
	receita.CreatedBy = userID

	// Cria a receita no banco
	if err := config.DB.Create(&receita).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(receita)
}

func ListaReceitas(c *fiber.Ctx) error {
	var receita []models.Receita
	if err := config.DB.Find(&receita).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(receita)
}
