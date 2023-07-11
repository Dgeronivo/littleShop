package main

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	htmlEngine := html.New("./project/view", ".html")
	app := fiber.New(fiber.Config{
		Views: htmlEngine,
	})
	ProductList := [5]int{1, 2, 3, 4, 5}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title":       "Hello, World!",
			"ProductList": ProductList,
		})
	})

	app.Use([]string{"/delivery", "/about-us"}, func(c *fiber.Ctx) error {
		page := "pages/" + strings.TrimLeft(c.Path(), "/")

		return c.Render(page, fiber.Map{})
	})

	app.Static("/img", "./project/img")
	app.Static("/style", "./project/style")

	log.Fatal(app.Listen(":3000"))
}
