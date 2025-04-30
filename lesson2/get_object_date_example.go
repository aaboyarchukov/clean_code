package lesson2

import (
	"backend/database"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

func GetObjectDate(ctx *fiber.Ctx) error {
	type request struct {
		ObjectID string `json:"object_id"`
	}

	// 6.4
	// old name: result
	// new name: requestBody
	var requestBody request

	// 6.4
	// old name: err
	// new name: errJSONUnmarshal
	if errJSONUnmarshal := json.Unmarshal(ctx.Request().Body(), &requestBody); errJSONUnmarshal != nil {
		return errJSONUnmarshal
	}

	// 6.4
	// old name: date
	// new name: objectDate
	objectDate := database.GetObjectDate(requestBody.ObjectID)

	return ctx.JSON(struct {
		Date string `json:"date"`
	}{
		Date: objectDate,
	})
}
