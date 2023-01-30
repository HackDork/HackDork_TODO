package routes

import (
	"github.com/HackDork/HackDork_TODO/controller"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", controller.ListTasks)

	app.Get("/task", controller.NewTaskView)
	app.Post("/task", controller.CreateTask)
	//app.Put("/update", controller.UpdateTask)
	//app.Delete("/delete", DeleteTask)
}
