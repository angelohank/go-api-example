package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func main() {
	app := fiber.New()

	app.Get("/user", func(c *fiber.Ctx) error {

		var apiResponse map[string]string

		res, err := http.Get("https://api.chucknorris.io/jokes/random")
		if err != nil {
			c.Status(fiber.StatusInternalServerError).SendString("Erro ao fazer requisicao")
		}

		json.NewDecoder(res.Body).Decode(&apiResponse)
		return c.SendString(apiResponse["value"])
	})

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
