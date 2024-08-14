package controllers

import (
	"log"
	"post/config"
	"post/database"
	"post/helpers"
	"post/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func AddTaste(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "A-T-1"})
	}

	// Bearer token'ı ayıklıyoruz
	token := authHeader[len("Bearer "):]
	log.Println(token)

	// Token kontrolü yapıyoruz
	tokenResponse, err := helpers.CheckToken(token)
	if err != nil {
		return c.Status(403).JSON(fiber.Map{"status": "error", "message": "Invalid token", "data": err.Error()})
	}

	// Eğer userInfo boşsa, token geçerli değil, işlemi durduruyoruz
	// if tokenResponse != nil || tokenResponse.Data != nil {
	// 	return c.Status(403).JSON(fiber.Map{"status": "error", "message": "Unauthorized"})
	// }
	log.Println(tokenResponse)

	food := new(models.Food)

	// err = c.BodyParser(&food)

	// if err != nil {
	// 	return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error parsing request body", "data": err.Error()})
	// }

	// Dosya yükleme işlemleri
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "not loading file", "data": err.Error()})
	}

	fileBytes, err := file.Open()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "not reading file", "data": err.Error()})
	}
	defer fileBytes.Close()

	imageBytes := make([]byte, file.Size)
	_, err = fileBytes.Read(imageBytes)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "failed to read file", "data": err.Error()})
	}

	id, url, err := config.CloudConnect(imageBytes)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "failed to upload cloud", "data": err.Error()})
	}
	food.Image = id
	food.ImageUrl = url

	// Diğer form alanlarını alıyoruz
	foodname := c.FormValue("foodname")
	materials := c.FormValue("materials")
	eatpersonStr := c.FormValue("eatperson")
	specification := c.FormValue("specification")
	guesspriceStr := c.FormValue("guessprice")
	preparationtimeStr := c.FormValue("preparationtime")

	if len(foodname) != 0 {
		food.FoodName = foodname
	}

	if len(materials) != 0 {
		food.Materials = materials
	}

	if eatpersonStr != "" {
		eatperson, err := strconv.Atoi(eatpersonStr)
		if err != nil {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Invalid eat person value", "data": err.Error()})
		}
		if eatperson > 0 {
			food.EatPerson = eatperson
		} else {
			return c.Status(402).JSON(fiber.Map{"status": "error", "message": "Eat person must be greater than zero"})
		}
	}

	if len(specification) != 0 {
		food.Specification = specification
	}

	if guesspriceStr != "" {
		guessprice, err := strconv.Atoi(guesspriceStr)
		if err != nil {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Invalid guess price", "data": err.Error()})
		}
		if guessprice > 0 {
			food.GuessPrice = guessprice
		} else {
			return c.Status(402).JSON(fiber.Map{"status": "error", "message": "Guess price must be greater than zero"})
		}
	}
	if preparationtimeStr != "" {
		preparationtime, err := strconv.Atoi(preparationtimeStr)
		if err != nil {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Invalid preparation time", "data": err.Error()})
		}
		if preparationtime > 0 {
			food.PreparationTime = preparationtime
		} else {
			return c.Status(402).JSON(fiber.Map{"status": "error", "message": "Preparation time must be greater than zero"})
		}
	}

	err = database.DB.Db.Create(&food).Error
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "error create food step", "data": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "your food created successfully"})
}

func UpdateTaste(c *fiber.Ctx) error {
	id := c.Params("id")
	var foods models.Food
	err := database.DB.Db.First(&foods, id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "foods not found"})
	}
	updateData := foods
	err = c.BodyParser(&updateData)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "error bodyparsing step"})
	}
	err = database.DB.Db.Model(foods).Updates(updateData).Error
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "error update step"})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "your foods updated successfully", "data": foods})
}

func DeleteTaste(c *fiber.Ctx) error {
	id := c.Params("id")
	food := new(models.Food)
	err := database.DB.Db.Where("id=?", id).First(&food).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "error not found the food post"})
	}
	err = database.DB.Db.Delete(&food).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "error delete step", "data": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "your food post deleted successfully"})
}
