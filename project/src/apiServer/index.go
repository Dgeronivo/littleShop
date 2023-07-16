package apiserver

import "github.com/gofiber/fiber/v2"

func handleIndex(c *fiber.Ctx) error {
	// todo products mock
	ProductList := [5]int{1, 2, 3, 4, 5}

	return c.Render("index", fiber.Map{
		"ProductList": ProductList,
	})
}
