//Package controllers Funções para manipulação de receitas e usuários
package controllers

import (
	"receitasfitness/backend/config"
	"receitasfitness/backend/models"
	"receitasfitness/backend/utils"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var body models.User
	if err := c.BodyParser(&body); err != nil {
		return err
}

hash , err := utils.HashPassword(body.Password)
if err != nil {
	return c.Status(500).JSON(fiber.Map{
		"error": "Erro ao criptografar a senha",
	})
}

body.Password = hash
body.Role = "user" // Padrão para novos usuários

if err := config.DB.Create(&body).Error; err != nil {
	return c.Status(500).JSON(fiber.Map{"error": "Erro ao criar o usuário",
	})
}
return c.Status(201).JSON(fiber.Map{
	"message": "Usuário criado com sucesso",
})
}

func Login (c *fiber.Ctx) error {

	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&body); err != nil {
		return err
}

	var user models.User

	if err := config.DB.Where("email = ?", body.Email).First(&user).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Usuário não encontrado"})
	}

	if err := utils.CheckPasswordHash(body.Password, user.Password); err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Senha incorreta"})
	}

	token, err := utils.GenerateToken(user.ID, user.Role)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Erro ao gerar token"})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    token,
		HTTPOnly: true,
		Secure:   false,
		Path:     "/",
	})

	return c.JSON(fiber.Map{"message": "Logado com sucesso"})
}
