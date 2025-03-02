package lead

import (
	"net/http"

	"github.com/Aman-Shitta/crm-go-fiber/database"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
}

func GetLeads(c *fiber.Ctx) error {
	var leads []Lead

	db := database.DBConn

	db.Find(&leads)

	return c.JSON(leads)
}

func GetLead(c *fiber.Ctx) error {
	id := c.Params("id")

	db := database.DBConn

	var lead Lead
	db.Find(&lead, id)
	if lead.Name != "" {
		return c.JSON(lead)
	}
	return c.Status(http.StatusBadRequest).JSON(map[string]string{"error": "Data missing"})
}

func NewLead(c *fiber.Ctx) error {
	db := database.DBConn
	lead := new(Lead)

	if err := c.BodyParser(lead); err != nil {
		c.Status(http.StatusBadRequest).Send([]byte(err.Error()))
	}

	db.Create(lead)
	return c.JSON(lead)
}

func DeleteLead(c *fiber.Ctx) error {
	id := c.Params("id")

	db := database.DBConn

	var lead Lead

	db.First(&lead, id)

	if lead.Name == "" {
		c.Status(http.StatusBadRequest).Send([]byte("no lead found"))
	}

	db.Delete(&lead)

	return c.Send([]byte("Lead deleted successfully"))
}
