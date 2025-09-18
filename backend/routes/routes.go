//Package routes Rotas da aplicação
package routes

import (
	"github.com/gofiber/fiber/v2"
	//"receitasfit/backend/controllers"
	)

	func Setup(app *fiber.App) {
		app.Post("/registrar")
		app.Post("/login" )
		app.Get("/usuarios" ) // Rota protegida, apenas para admins
		app.Get("/usuarios/:id" )
		app.Post("/receitas" )
		app.Get("/receitas/:id" )
	}
