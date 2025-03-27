package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

type Todo struct {
	ID        int    `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func main() {
	app := fiber.New()
	todos := []Todo{}

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")

	// Get all todos
	app.Get("/api/todos", func(c fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"msg": "Success", "todos": todos})
	})

	// Create Todo
	app.Post("/api/todos", func(c fiber.Ctx) error {
		todo := &Todo{}

		if err := c.Bind().Body(todo); err != nil {
			return err
		}

		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"msg": "Bad Request. Body is Empty"})
		}

		todo.ID = len(todos) + 1
		todos = append(todos, *todo)

		return c.Status(201).JSON(fiber.Map{
			"msg":  "Todo created successfully",
			"todo": todo,
		})

	})

	// Update Todo
	app.Patch("/api/todos/:id", func(c fiber.Ctx) error {
		id := c.Params("id")

		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				todos[i].Completed = !todos[i].Completed
				return c.Status(200).JSON(fiber.Map{
					"msg":  "Todo updated successfully.",
					"todo": todo,
				})
			}
		}

		return c.Status(400).JSON(fiber.Map{"msg": "Todo is not found."})
	})

	// Delete Todo
	app.Delete("/api/todos/:id", func(c fiber.Ctx) error {
		id := c.Params("id")

		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				todos = append(todos[:i], todos[i+1:]...)
				return c.Status(200).JSON(fiber.Map{"msg": fmt.Sprintf("Todo with id %s is deleted successfully.", id)})
			}
		}
		return c.Status(400).JSON(fiber.Map{"msg": "Todo is not found."})

	})
	log.Fatal(app.Listen(":" + PORT))
}
