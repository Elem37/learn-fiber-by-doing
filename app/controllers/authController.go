package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/koybigino/learn-fiber/app/models"
	"github.com/koybigino/learn-fiber/app/oauth2"
	"github.com/koybigino/learn-fiber/app/utils"
	"github.com/koybigino/learn-fiber/app/validation"
)

func Login(c *fiber.Ctx) error {
	body := new(models.UserLogin)

	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}

	errors := validation.ValidateStruct(body)

	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": errors,
		})
	}

	user := new(models.User)

	if err := db.Where("email = ?", body.Email).First(user).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "bad Credentials !",
		})
	}

	if err := utils.Verify([]byte(body.Password), []byte(user.Password)); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "bad Credentials !",
		})
	}

	token := oauth2.CreateJWTToken(user.Id, user.UserName, user.Email)

	userResponse := new(models.UserResponse)
	models.ParseToUserResponse(*user, userResponse)

	return c.JSON(fiber.Map{
		"token":      token,
		"token_type": "Bearer",
	})

}
