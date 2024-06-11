package main

import (
	"os"
	"github.com/gofiber/fiber/v2"
)

type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var tasks []Task
var idCounter = 0

func main() {
	app := fiber.New()
	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	// Ruta para obtener todas las tareas
	app.Get("/tasks", func(c *fiber.Ctx) error {
		return c.JSON(tasks)
	})

	// Ruta para agregar una nueva tarea
	app.Post("/tasks", func(c *fiber.Ctx) error {
		var newTask Task
		if err := c.BodyParser(&newTask); err != nil {
			return err
		}
		idCounter++
		newTask.ID = idCounter
		tasks = append(tasks, newTask)
		return c.JSON(newTask)
	})

	app.Listen(":" + port)
}
