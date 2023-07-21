package main

import (
	"os"
	"webapp/controllers"
	"webapp/initialisers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html/v2"
)

func init() {
	initialisers.GetEnvVars()
	initialisers.ConnectToDatabase()
	initialisers.SyncDB()
}

func setupRoutes(app *fiber.App) {
	app.Get("/", controllers.Home)
	users := app.Group("/users")
	users.Get("/", controllers.GetUsers)
	users.Get("/:id", controllers.GetUser)
	users.Get("/:id/followers", controllers.GetUserFollowers)
	users.Post("/register", controllers.CreateUser)
	users.Post("/following/register", controllers.CreateFollowing)
}

func main() {
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{Views: engine})

	app.Static("/", "./public")
	app.Use(cors.New())

	setupRoutes(app)

	app.Listen(":" + os.Getenv("PORT"))
}
