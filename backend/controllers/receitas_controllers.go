// Package controllers Lida com as requisições relacionadas às receitas
package controllers

import (
	"receitasfitness/backend/config"
	"receitasfitness/backend/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func CreateReceita(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	// Apenas admins podem criar receitas
	if claims["role"] != "admin" {
		return c.Status(403).JSON(fiber.Map{"error": "Apenas admins"})
	}

	var receita models.Receita
	if err := c.BodyParser(&receita); err != nil {
		return err
	}

	// Define o campo CreatedBy com o ID do usuário autenticado
	receita.CreatedBy = uint(claims["user_id"].(float64))

	config.DB.Create(&receita)
	return c.JSON(receita)
}


func ListaReceitas(c *fiber.Ctx) error {
	var receita []models.Receita
	config.DB.Find(&receita)
	return c.JSON(receita)
}
