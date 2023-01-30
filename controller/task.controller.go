package controller

import (
	"github.com/HackDork/HackDork_TODO/database"
	"github.com/HackDork/HackDork_TODO/models"
	"github.com/gofiber/fiber/v2"
)

func ListTasks(c *fiber.Ctx) error {
	var tasks []models.Task
	database.DB.Db.Find(&tasks)

	return c.Render("index", fiber.Map{
		"Title":    "HackDork TODO App",
		"Subtitle": "Daily !",
		"Tasks":    tasks,
	})
}

func NewTaskView(c *fiber.Ctx) error {
	return c.Render("new", fiber.Map{
		"Title":    "New Task",
		"Subtitle": "Add a cool Task!",
	})
}

func CreateTask(c *fiber.Ctx) error {
	task := new(models.Task)
	if err := c.BodyParser(task); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&task)

	return ConfirmationView(c)
}

func ConfirmationView(c *fiber.Ctx) error {
	return c.Render("confirmation", fiber.Map{
		"Title":    "Task added successfully",
		"Subtitle": "Add more wonderful tasks to the list!",
	})
}
