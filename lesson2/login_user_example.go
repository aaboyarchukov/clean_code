package lesson2

import (
	authenticationservice "backend/authentication_service"
	"backend/database"
	"encoding/json"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func LoginUser(ctx *fiber.Ctx) error {
	type request struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	type answer struct {
		Ok    bool   `json:"ok"`
		Error error  `json:"error"`
		JWT   string `json:"jwt"`
	}

	var requestBody request

	// 6.4
	// old name: err
	// new name: errJSONUnmarshal
	if errJSONUnmarshal := json.Unmarshal(ctx.Request().Body(), &requestBody); errJSONUnmarshal != nil {
		return ctx.JSON(answer{
			Ok:    false,
			Error: errJSONUnmarshal,
		})
	}

	user, errExistUser := database.User(requestBody.Login)
	if errExistUser != nil {
		return ctx.JSON(answer{
			Ok:    false,
			Error: errExistUser,
		})
	}

	errCompare := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requestBody.Password))
	if errCompare != nil {
		return ctx.JSON(answer{
			Ok:    false,
			Error: errCompare,
		})
	}

	jwt, errGenerateJWT := authenticationservice.GenerateJWT(user, time.Duration(time.Hour*8))
	if errGenerateJWT != nil {
		return ctx.JSON(answer{
			Ok:    false,
			Error: errGenerateJWT,
		})
	}

	ctx.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    jwt,
		HTTPOnly: true,
		Secure:   true,
		Expires:  time.Now().Add(24 * time.Hour),
	})

	return ctx.JSON(answer{
		Ok:    true,
		Error: nil,
		JWT:   jwt,
	})
}
