package lead

import (
	"github.com/AlperSeyman/fiber-crm-basic/database"
	"github.com/AlperSeyman/fiber-crm-basic/models"
	"github.com/gofiber/fiber"
)

func GetLeads(c *fiber.Ctx) {
	db := database.DBconn
	var leads []models.Lead
	db.Find(&leads)
	c.JSON(leads)
}

func GetLead(c *fiber.Ctx) {

	id := c.Params("id")
	db := database.DBconn
	var lead models.Lead
	db.Find(&lead, id)
	c.JSON(lead)

}

func NewLead(c *fiber.Ctx) {

	db := database.DBconn
	lead = new(models.Lead)

}

func UpdateLead(c *fiber.Ctx) {

	id := c.Params("id")
}

func DeleteLead(c *fiber.Ctx) {

	id := c.Params("id")
}
