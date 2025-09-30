package main

import (

	"github.com/gofiber/fiber/v2"
	"receitas_app/backend/config"
	"receitas_app/backend/models"
	"receitas_app/backend/routes"
	"fmt"
	"log"
)

func main( ) {

 config.ConnectDatabase()
 config.DB.AutoMigrate(&models.User{}, &models.Receita{})

 app := fiber.New()
 routes.Setup(app)
 log.Fatal(app.Listen(":3000"))
 fmt.Println("Servidor rodando na porta 3000")

}
