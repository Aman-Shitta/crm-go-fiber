package main

import (
	"fmt"

	"github.com/Aman-Shitta/crm-go-fiber/database"
	"github.com/Aman-Shitta/crm-go-fiber/lead"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)

	defer database.DBConn.Close()
	app.Listen(":8000")
}

func initDatabase() {
	var err error

	database.DBConn, err = gorm.Open("sqlite3", "leads.db")

	if err != nil {
		panic("failed to connect DB")
	}

	fmt.Println("[+] DB connection established [+]")
	database.DBConn.AutoMigrate(&lead.Lead{})
}

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}
