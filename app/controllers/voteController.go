package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/koybigino/learn-fiber/app/models"
	"github.com/koybigino/learn-fiber/app/validation"
)

func CreateVote(c *fiber.Ctx) error {
	newVote := new(models.Vote)
	newVoteRequest := new(models.VoteRequest)

	if err := c.BodyParser(newVoteRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	errors := validation.ValidateStruct(newVoteRequest)

	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": errors,
		})
	}

	models.ParseToVote(newVote, *newVoteRequest)

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	id, _ := claims["user_id"].(float64)
	userId := int(id)

	if err := db.Where("user_id = ? AND post_id = ?", userId, newVoteRequest.PostId).First(newVote).Error; err != nil {
		newVote.PostId = newVoteRequest.PostId
		newVote.UserId = userId
		if err := db.Create(newVote).Error; err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":   err.Error(),
				"message": "Error to createthe element !",
			})
		}

		return c.JSON(fiber.Map{
			"vote": newVote,
		})
	}

	db.Delete(newVote, newVote.UserId, newVote.PostId)

	return c.SendStatus(fiber.StatusNoContent)

}
