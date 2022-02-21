package main

import (
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	pusherClient := pusher.Client{
		AppID:   "",
		Key:     "",
		Secret:  "",
		Cluster: "",
		Secure:  true,
	}

	app.Post("/api/messages", func(c *fiber.Ctx) error {
		var data map[string]string

		if err := c.BodyParser(&data); err != nil {
			return err
		}

		pusherClient.Trigger("chat", "message", data)

		return c.JSON([]string{})
	})

	app.Listen(":8000")
}
