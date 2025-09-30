//Package controllers contém as funções para Registrar , logar e listar usuários , com JWT e bcrypt salvando no banco de dados.
package controllers

import (
	"errors"
	"log"
	"receitas_app/backend/config"
	"receitas_app/backend/models"
	"receitas_app/backend/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Register(c *fiber.Ctx) error {

	type Input struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Role     string `json:"role"` // Permite definir o papel via JSON
	}

	var body Input

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Dados inválidos"})
	}

	if body.Email == "" || body.Password == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Email e senha obrigatórios"})
	}

	// Verifica se já existe usuário com esse email
	var existingUser models.User

	if err := config.DB.Where("email = ?", body.Email).First(&existingUser).Error; err == nil {
		return c.Status(400).JSON(fiber.Map{"error": "Email já cadastrado"})
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Println("Erro ao verificar email:", err)
		return c.Status(500).JSON(fiber.Map{"error": "Erro ao verificar email"})
	}

	// Gera hash da senha

	hash, err := utils.HashPassword(body.Password)
	if err != nil {
		log.Println("Erro ao gerar hash:", err)
		return c.Status(500).JSON(fiber.Map{"error": "Erro ao processar senha"})
	}

	// Cria usuário

	user := models.User{
		Username: body.Username,
		Email:    body.Email,
		Password: hash,
		Role:     body.Role, // permite definir admin ou user via JSON
	}

	if err := config.DB.Create(&user).Error; err != nil {
		log.Println("Erro ao criar usuário:", err)
		return c.Status(500).JSON(fiber.Map{"error": "Erro ao criar usuário"})
	}

	return c.JSON(fiber.Map{"message": "Usuário criado com sucesso"})
}

// Login autentica o usuário e retorna um token JWT
func Login(c *fiber.Ctx) error {

	// Estrutura para receber os dados de login
	type Input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var body Input

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Dados inválidos"})
	}

	if body.Email == "" || body.Password == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Email e senha obrigatórios"})
	}

	// Busca usuário pelo email
	var user models.User
	if err := config.DB.Where("email = ?", body.Email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(401).JSON(fiber.Map{"error": "Credenciais inválidas"})
		}
		log.Println("Erro ao buscar usuário:", err)
		return c.Status(500).JSON(fiber.Map{"error": "Erro interno"})
	}

	// Verifica senha
	if !utils.CheckPasswordHash(body.Password, user.Password) {
		return c.Status(401).JSON(fiber.Map{"error": "Credenciais inválidas"})
	}

	// gerar um token JWT
	token, err := utils.GenerateToken(user.ID, user.Role)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Erro ao gerar token"})
	}

    // Retorna dados do usuário e token
	return c.JSON(fiber.Map{
		"message": "Login realizado com sucesso",
		"user": fiber.Map{
			"id":    user.ID,
			"email": user.Email,
            "role":  user.Role,
		},

        "token": token,
	})
}

// GetAllUsers retorna todos os usuários
func GetAllUsers(c *fiber.Ctx) error {
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Erro ao buscar usuários"})
	}
	return c.JSON(users)
}

// GetUserByID retorna um usuário pelo ID
func GetUserByID(c *fiber.Ctx) error {
    id := c.Params("id")
    var user models.User

    if err:= config.DB.First(&user,id).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return c.Status(404).JSON(fiber.Map{"error": "Usuário não encontrado"})
    }
        return c.Status(500).JSON(fiber.Map{"error": "Erro ao buscar usuário"})
    }
    return c.JSON(user)
}
