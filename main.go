package main

import (
	"log"

	"github.com/alnoviantirs/contactsAPI/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/whatsauth/whatsauth"

	"github.com/alnoviantirs/contactsAPI/url"
)

// @title TES SWAG
// @version 1.0
// @description This is a sample swagger for Fiber

// @contact.name API Support
// @contact.url https://github.com/alnoviantirs
// @contact.email indra@ulbi.ac.id

// @host https://github.com/alnoviantirs
// @BasePath /
// @schemes https http
func main() {
	go whatsauth.RunHub()
	site := fiber.New(config.Iteung)
	site.Use(cors.New(config.Cors))
	site.Use(logger.New())
	url.Web(site)
	log.Fatal(site.Listen(musik.Dangdut()))
}
