package handlers

import (
	"github.com/efrenfuentes/todo-backend-golang-fiber/database"
	"github.com/efrenfuentes/todo-backend-golang-fiber/internals/models"
	"github.com/gofiber/fiber/v2"
)

// Get all todos
func GetAllTodos(c *fiber.Ctx) error {
	db := database.DB

	var todos []models.Todo

	db.Find(&todos)

	if len(todos) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "No todos found",
			"todos":   []models.Todo{},
		})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "todos found", "todos": todos})
}

func GetTodoById(c *fiber.Ctx) error {
	db := database.DB

	id := c.Params("id")

	var todo models.Todo

	db.Find(&todo, id)

	if todo.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "No todo found with this ID",
			"todo":    nil,
		})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "todo found", "todo": todo})
}

func CreateNewTodo(c *fiber.Ctx) error {
	db := database.DB

	todo := new(models.Todo)

	err := c.BodyParser(todo)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot parse JSON",
			"todo":    nil,
		})
	}

	err = db.Create(&todo).Error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot create todo",
			"todo":    nil,
		})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Created todo", "todo": todo})
}

func UpdateTodoById(c *fiber.Ctx) error {
	db := database.DB

	id := c.Params("id")

	todo := new(models.UpdateTodo)

	err := c.BodyParser(todo)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot parse JSON",
			"todo":    nil,
		})
	}

	var todoFromDB models.Todo

	db.First(&todoFromDB, id)

	if todoFromDB.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "No todo found with this ID",
			"todo":    nil,
		})
	}

	if todo.Title != nil {
		todoFromDB.Title = *todo.Title
	}

	if todo.Completed != nil {
		todoFromDB.Completed = *todo.Completed
	}

	if todo.Order != nil {
		todoFromDB.Order = *todo.Order
	}

	if todo.Url != nil {
		todoFromDB.Url = *todo.Url
	}

	db.Save(&todoFromDB)

	return c.JSON(fiber.Map{"status": "success", "message": "Todo successfully updated", "todo": todoFromDB})
}

func DeleteTodoById(c *fiber.Ctx) error {
	db := database.DB

	id := c.Params("id")

	var todo models.Todo

	db.First(&todo, id)

	if todo.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "No todo found with this ID",
			"todo":    nil,
		})
	}

	err := db.Delete(&todo).Error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot delete todo",
			"todo":    nil,
		})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Todo successfully deleted", "todo": nil})
}
