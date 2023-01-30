package main

import (
	. "github.com/HackDork/HackDork_TODO/database"
	. "github.com/HackDork/HackDork_TODO/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {
	ConnectDb()

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
	})

	SetupRoutes(app)

	app.Static("/", "./public")

	app.Listen(":3000")
}
