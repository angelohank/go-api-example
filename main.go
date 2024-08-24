package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func main() {
	app := fiber.New()

	app.Post("/login", Login)

	app.Listen(":3000")
}

func Login(c *fiber.Ctx) error {
	user := c.FormValue("user")
	pass := c.FormValue("password")

	if user != "angelo" || pass != "123" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	claims := jwt.MapClaims{
		"name":  "Angelo",
		"admin": true,
		"exp":   time.Now().Add(time.Hour * 72).Unix(), //definindo validade de 72 horas pro token
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret-key"))

	if err != nil {
		c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}
