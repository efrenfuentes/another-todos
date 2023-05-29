package router

import (
	"github.com/efrenfuentes/todo-backend-golang-fiber/internals/handlers"
	"github.com/gofiber/fiber/v2"
)

func TodosRoute(route fiber.Router) {
	route.Get("", handlers.GetAllTodos)
	route.Get("/:id", handlers.GetTodoById)
	route.Post("", handlers.CreateNewTodo)
	route.Put("/:id", handlers.UpdateTodoById)
	route.Delete("/:id", handlers.DeleteTodoById)
}
