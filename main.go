package main

import (
	"context"
	"medotlink/data"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django/v3"
	"github.com/redis/go-redis/v9"
)

func main() {
	data.Init("./me.json")

	ctx := context.Background()
	engine := django.New("./views", ".django")

	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	_, err := client.Get(ctx, "views").Result()
	if err != nil {
		client.Set(ctx, "views", 0, 0)
	}

	app.Static("/static", "./static")

	app.Get("/", func(c *fiber.Ctx) error {
		value, _ := client.Get(ctx, "views").Result()

		views, _ := strconv.Atoi(value)
		views += 1

		client.Set(ctx, "views", views, 0)

		return c.Render("index", fiber.Map{
			"Title":   data.Data.Title,
			"Name":    data.Data.Name,
			"Short":   data.Data.Description.Short,
			"Long":    data.Data.Description.Long,
			"Views":   views,
			"Buttons": data.Data.Buttons,
			"Link":    data.Data.Link,
		})
	})

	app.Listen(":5890")
}
