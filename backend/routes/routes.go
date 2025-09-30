// Package routes Rotas da aplicação
package routes

import (
	"receitas_app/backend/controllers"
	"receitas_app/backend/middlewares"

	"github.com/gofiber/fiber/v2"
)

	func Setup(app *fiber.App) {
		app.Post("/registrar", controllers.Register)
		app.Post("/login", controllers.Login)
		app.Post("/receitas",middlewares.JWTProtected(), controllers.CreateReceita )

		app.Get("/users", controllers.GetAllUsers)
		app.Get("/users/:id", controllers.GetUserByID)
		app.Get("/receitas", controllers.ListaReceitas )
	}
