package controllers

import (
	"auth/database"
	"auth/helpers"
	"auth/middleware"
	"auth/models"

	"github.com/gofiber/fiber/v2"
)

type TokenRequest struct {
	Token string `json:"token"`
}

func SignUp(c *fiber.Ctx) error {
	var user models.User
	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"Status": "error", "Message": "ERROR : S-U-1"})
	}
	mailControl, _ := helpers.MailControl(user.Mail)
	if mailControl == true {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "S-U-2"})
	}
	user.Password = helpers.HashPass(user.Password)
	err = database.DB.Db.Create(&user).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "S-U-3"})
	}
	return c.Status(200).JSON(fiber.Map{"status": "Success", "message": "Success"})
}

func SignIn(c *fiber.Ctx) error {
	var user models.User
	var session models.Session
	var login models.SignIn

	err := c.BodyParser(&login)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "S-I-1"})
	}
	login.Password = helpers.HashPass(login.Password)
	err = database.DB.Db.Where("mail=? and password=?", login.Mail, login.Password).First(&user).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "S-I-2"})
	}
	session.UserID = user.ID
	token, err := middleware.GenerateToken(user.Mail)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "S-I-3"})
	}
	session.Token = token
	err = database.DB.Db.Create(&session).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "S-I-4"})
	}

	return c.Status(200).JSON(fiber.Map{"status": "Success", "message": "Success"})
}

func UpdatePassword(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(models.User)
	if !ok {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "U-P-1"})
	}
	var updatedpassword models.UpdatePassword
	err := c.BodyParser(&updatedpassword)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "U-P-2"})
	}
	user.Password = helpers.HashPass(user.Password)
	updatedpassword.OldPassword = user.Password
	if updatedpassword.NewPassword1 != updatedpassword.NewPassword2 {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "U-P-3"})
	}
	if updatedpassword.OldPassword == updatedpassword.NewPassword1 {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "U-P-3"})
	}
	user.Password = updatedpassword.NewPassword1
	user.Password = helpers.HashPass(user.Password)
	err = database.DB.Db.Where("id=?", user.ID).Updates(&user).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "U-P-4"})
	}
	return c.Status(200).JSON(fiber.Map{"status": "Success", "message": "Success"})
}

func TokenControlHandler(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" || len(token) < 7 || token[:7] != "Bearer " {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid token format",
		})
	}
	token = token[7:]

	var session models.Session
	err := database.DB.Db.Where("token = ?", token).First(&session).Error
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid token",
		})
	}

	var user models.User
	err = database.DB.Db.Where("id = ?", session.UserID).First(&user).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "User not found",
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Token is valid",
		"data":    user,
	})
}
