package controllers

import (
	"errors"
	"net/http"
	"webapp/initialisers"
	"webapp/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		c.Status(http.StatusBadRequest)
		return err
	}

	initialisers.DB.Create(&user)
	return c.Status(http.StatusCreated).JSON(user)
}

func CreateFollowing(c *fiber.Ctx) error {
	following := new(models.Following)

	if err := c.BodyParser(following); err != nil {
		c.Status(http.StatusBadRequest)
		return err
	}

	initialisers.DB.Create(&following)
	return c.Status(http.StatusCreated).JSON(following)
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user := new(models.User)

	err := initialisers.DB.First(&user, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(http.StatusNotFound).SendString("User not found")
	}

	return c.Status(http.StatusOK).JSON(user)

}

func GetUsers(c *fiber.Ctx) error {
	users := new([]models.User)
	initialisers.DB.Find(&users)

	return c.Status(http.StatusOK).JSON(users)
}

func GetUserFollowers(c *fiber.Ctx) error {
	id := c.Params("id")

	user := new(models.User)
	err := initialisers.DB.First(&user, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(http.StatusNotFound).SendString("User not found")
	}

	followers := new([]models.Follower)
	initialisers.DB.Table("users").
		Joins("INNER JOIN followings ON users.id = followings.follower").
		Select("users.first_name", "users.last_name", "users.email").
		Where("followings.followee = ?", id).
		Order("users.first_name, users.last_name").
		Find(&followers)

	return c.Status(http.StatusOK).JSON(followers)
}
