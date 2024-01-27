package controller

import (
	"github.com/alnoviantirs/contactsAPI/config"
	inimodel "github.com/alnoviantirs/contactsAPI/package"
	"github.com/gofiber/fiber/v2"

	"net/http"
	"strconv"
)

func Home(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"github_repo": "https://github.com/alnoviantirs/be_catatan",
		"message":     "You are at the root endpoint 😉",
		"success":     true,
	})
}

func GetAllCatatan(c *fiber.Ctx) error {
	ps := inimodel.GetAllCatatan(config.Ulbimongoconn, "catatan")
	return c.JSON(fiber.Map{
		"status": http.StatusOK,
		"data":   ps,
	})
}


func GetCatatanByUsername(c *fiber.Ctx) error {
	username := c.Params("username")
	if username == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Username is required",
		})
	}

	catatanByUsername := inimodel.GetCatatanByUsername(config.Ulbimongoconn, "catatan", username)
	return c.JSON(fiber.Map{
		"status": http.StatusOK,
		"data":   catatanByUsername,
	})
}


func InsertCatatan(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var kontak inimodel.Catatan
	if err := c.BodyParser(&kontak); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Get last Catatan
	lastCatatan, err := inimodel.GetLastCatatan(db, "catatan")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Increment ID2
	newID2 := lastCatatan.ID2 + 1

	// Assign incremented ID2 to the catatan
	kontak.ID2 = newID2

	catatan := inimodel.Catatan{
		ID2:         kontak.ID2,
		Title:       kontak.Title,
		Note:        kontak.Note,
		Date:        kontak.Date,
		StartTime:   kontak.StartTime,
		EndTime:     kontak.EndTime,
		Remind:      kontak.Remind,
		Repeat:      kontak.Repeat,
		IsCompleted: kontak.IsCompleted,
		CompletedAt: kontak.CompletedAt,
		CreatedAt:   kontak.CreatedAt,
		UpdatedAt:   kontak.UpdatedAt,
		Color:       kontak.Color,
		User:       kontak.User,
	}

	insertedID, err := inimodel.InsertCatatan(db, "catatan", catatan)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":      http.StatusOK,
		"message":     "Data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}

func UpdateCatatan(c *fiber.Ctx) error {
	db := config.Ulbimongoconn

	// Get Catatan ID from request parameters
	//catatanID, err := primitive.ObjectIDFromHex(c.Params("id"))
	catatanIDStr := c.Params("id")
	catatanID, err := strconv.Atoi(catatanIDStr)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
	}

	// Get updated Catatan from request body
	var updatedCatatan inimodel.Catatan
	if err := c.BodyParser(&updatedCatatan); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Update Catatan in the database
	if err := inimodel.UpdateCatatan(db, "catatan", catatanID, updatedCatatan); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Data berhasil diupdate.",
	})
}

func UpdateStatus(c *fiber.Ctx) error {
	db := config.Ulbimongoconn

	// Get Catatan ID from request parameters
	//catatanID, err := primitive.ObjectIDFromHex(c.Params("id"))
	catatanIDStr := c.Params("id")
	catatanID, err := strconv.Atoi(catatanIDStr)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
	}

	// Get updated Catatan from request body
	var updatedCatatan inimodel.Catatan
	if err := c.BodyParser(&updatedCatatan); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Update Catatan in the database
	if err := inimodel.UpdateStatus(db, "catatan", catatanID, updatedCatatan); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Data berhasil diupdate.",
	})
}

func DeleteCatatan(c *fiber.Ctx) error {
	db := config.Ulbimongoconn

	// Get Catatan ID from request parameters
	catatanIDStr := c.Params("id")
	catatanID, err := strconv.Atoi(catatanIDStr)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
	}

	// Delete Catatan from the database
	if err := inimodel.DeleteCatatan(db, "catatan", catatanID); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Data berhasil dihapus.",
	})
}
