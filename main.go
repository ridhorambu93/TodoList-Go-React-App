package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
    "github.com/joho/godotenv"
    _ "github.com/lib/pq" // PostgreSQL driver
)

type Todo struct {
    ID        int    `json:"id,omitempty"`
    Completed bool   `json:"completed"`
    Body      string `json:"body"`
}

var db *sql.DB

func main() {
    // Load environment variables
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // Connect to PostgreSQL
    POSTGRES_URI := os.Getenv("POSTGRES_URI")
    db, err = sql.Open("postgres", POSTGRES_URI)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Connected to PostgreSQL")

    app := fiber.New()

    // Setup CORS
    app.Use(cors.New(cors.Config{
        AllowOrigins:     "http://localhost:5173", // Ganti dengan URL frontend Anda
        AllowMethods:     "GET,POST,PATCH,DELETE,OPTIONS",
        AllowHeaders:     "Content-Type, Authorization",
        AllowCredentials: true,
    }))

    // Define routes
    app.Get("/api/todos", getTodos)
    app.Post("/api/todos", createTodo)
    app.Patch("/api/todos/:id", updateTodo)
    app.Delete("/api/todos/:id", deleteTodo)

    port := os.Getenv("PORT")
    if port == "" {
        port = "3000"
    }

    log.Fatal(app.Listen("0.0.0.0:" + port))
}

func getTodos(c *fiber.Ctx) error {
    rows, err := db.Query("SELECT id, completed, body FROM todos")
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch todos"})
    }
    defer rows.Close()

    var todos []Todo
    for rows.Next() {
        var todo Todo
        if err := rows.Scan(&todo.ID, &todo.Completed, &todo.Body); err != nil {
            return err
        }
        todos = append(todos, todo)
    }

    return c.JSON(todos)
}

func createTodo(c *fiber.Ctx) error {
    todo := new(Todo)
    if err := c.BodyParser(todo); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Invalid request payload"})
    }

    if todo.Body == "" {
        return c.Status(400).JSON(fiber.Map{"error": "Todo body cannot be empty"})
    }

    var id int
    err := db.QueryRow(
        "INSERT INTO todos (completed, body) VALUES ($1, $2) RETURNING id",
        todo.Completed, todo.Body).Scan(&id)

    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Failed to create todo"})
    }

    todo.ID = id
    return c.Status(201).JSON(todo)
}

func updateTodo(c *fiber.Ctx) error {
    id := c.Params("id")
    completed := true // Ini bisa diubah untuk menerima input dari body

    _, err := db.Exec("UPDATE todos SET completed = $1 WHERE id = $2", completed, id)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Failed to update todo"})
    }

    return c.Status(200).JSON(fiber.Map{"success": true})
}

func deleteTodo(c *fiber.Ctx) error {
    id := c.Params("id")

    _, err := db.Exec("DELETE FROM todos WHERE id = $1", id)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Failed to delete todo"})
    }

    return c.Status(200).JSON(fiber.Map{"success": true})
}