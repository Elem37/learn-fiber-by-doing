package controllers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/koybigino/learn-fiber/app/models"
	"github.com/koybigino/learn-fiber/app/utils"
	"github.com/koybigino/learn-fiber/app/validation"
)

func CreateUser(c *fiber.Ctx) error {
	newUser := new(models.User)
	newUserRequest := new(models.UserRequest)
	newUserResponse := new(models.UserResponse)

	if err := c.BodyParser(newUserRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	errors := validation.ValidateStruct(newUserRequest)

	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": errors,
		})
	}

	models.ParseToUser(newUser, *newUserRequest)
	hashPassword := utils.Hash(newUser.Password)
	newUser.Password = string(hashPassword)

	if err := db.Create(newUser).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Error to createthe element !",
		})
	}

	models.ParseToUserResponse(*newUser, newUserResponse)

	return c.JSON(fiber.Map{
		"User": newUserResponse,
	})
}

func GetUserByID(c *fiber.Ctx) error {
	strId := c.Params("id")

	id, err := strconv.Atoi(strId)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "error to parse the id parameter",
		})
	}

	User := new(models.User)
	UserResponse := new(models.UserResponse)

	if err := db.First(User, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   err.Error(),
			"message": fmt.Sprintf("Any User correspond to this id = %d", id),
		})
	}

	models.ParseToUserResponse(*User, UserResponse)

	return c.JSON(fiber.Map{
		"User": UserResponse,
	})
}
