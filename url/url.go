package url

import (
	"github.com/alnoviantirs/contactsAPI/controller"
	"github.com/gofiber/fiber/v2"
)

func Web(page *fiber.App) {
	page.Get("/", controller.Home)
	page.Get("/catatan", controller.GetAllCatatan)
	page.Post("/insertcatatan", controller.InsertCatatan)
	page.Put("/editcatatan/:id", controller.UpdateCatatan)
	page.Put("/setstatus/:id", controller.UpdateStatus)
	page.Delete("/deletecatatan/:id", controller.DeleteCatatan)
}
