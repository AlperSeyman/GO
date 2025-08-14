package main

import (
	"github.com/AlperSeyman/fiber-crm-basic/database"
	"github.com/AlperSeyman/fiber-crm-basic/lead"
	"github.com/gofiber/fiber"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Put("/api/v1/lead/:id", lead.UpdateLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func main() {

	app := fiber.New()
	database.InitDatabase()
	setupRoutes(app)

	defer func() {
		sqlDB, _ := database.DBconn.DB()
		_ = sqlDB.Close()
	}()

	app.Listen(":3000")
}
