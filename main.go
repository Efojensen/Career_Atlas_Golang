package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	Id        int    `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func main() {
	fmt.Println("Hello world")
	app := fiber.New()

	todos := []Todo{}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"msg": "Hello world"})
	})

	// GET all todoItems
	app.Get("/api/todos", func (c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"items": todos})
	})

	// CREATE todo
	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := &Todo{}

		if err := c.BodyParser(todo); err != nil {
			return err
		}

		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Todo body is required"})
		}

		todo.Id = len(todos) + 1
		todos = append(todos, *todo)

		return c.Status(201).JSON(todo)
	})

	// UPDATE a todo
	app.Patch("/api/todos/:id", func (c *fiber.Ctx) error {
		id := c.Params("id")

		for i, todo := range todos {
			if fmt.Sprint(todo.Id) == id {
				todos[i].Completed = true
				return c.Status(201).JSON(todos[i])
			}
		}

		return c.Status(404).JSON(fiber.Map{"msg": "Todo not found"})
	})

	log.Fatal(app.Listen(":4000"))
}
