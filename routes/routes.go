package routes

import (
	"go-react-todo/controllers"

	"github.com/gofiber/fiber/v2"
)

var UseRoutes = func(app *fiber.App) {

	app.Get("/api/todos", controllers.GetAllTodo)
	app.Post("/api/todo", controllers.CreateTodo)
	app.Patch("/api/todo/:id", controllers.UpdateTodo)
	app.Delete("/api/todo/:id", controllers.DeleteTodo)

}
