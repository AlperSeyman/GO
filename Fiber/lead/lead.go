package lead

import (
	"github.com/AlperSeyman/fiber-crm-basic/database"
	"github.com/AlperSeyman/fiber-crm-basic/models"
	"github.com/gofiber/fiber"
)

func GetLeads(c *fiber.Ctx) {
	db := database.DBconn
	var leads []models.Lead
	if err := db.Find(&leads); err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not retrieve leads"})
	}
	c.JSON(leads)
}

func GetLead(c *fiber.Ctx) {

	id := c.Params("id")
	db := database.DBconn
	var lead models.Lead
	if err := db.Find(&lead, id).Error; err != nil {
		c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Lead not found"})
	}
	c.JSON(lead)

}

func UpdateLead(c *fiber.Ctx) {

	id := c.Params("id")
	db := database.DBconn
	var updateLead models.Lead

	db.Model(&models.Lead{}).Where("id=?", id).Updates(updateLead)
	c.Send("Lead Successfully updated")
}

func NewLead(c *fiber.Ctx) {

	var lead models.Lead

	db := database.DBconn

	if err := c.BodyParser(&lead); err != nil {
		c.Status(503).Send(err)
		return
	}
	db.Create(&lead)
	c.JSON(lead)
}

func DeleteLead(c *fiber.Ctx) {

	id := c.Params("id")
	db := database.DBconn

	var lead models.Lead
	db.First(&lead, id)

	if lead.Name == "" {
		c.Status(500).Send("No lead found with ID")
		return
	}
	db.Delete(&lead)
	c.SendString("Lead Successfully deleted")
}
