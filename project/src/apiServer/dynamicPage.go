package apiserver

import (
	"littleShop/project/src/staticParts"

	"github.com/gofiber/fiber/v2"
)

func handleDynamic(c *fiber.Ctx) error {
	// c.Path() - return /path
	var products []int
	for i := 1; i < 22; i++ {
		products = append(products, i)
	}

	bc := staticParts.Breadcrumbs{Href: "#", Name: "Категорія 1", Active: true}
	bcHtml := staticParts.GenerateBreadcrumbsHtml(bc)

	return c.Render("dynamic/catalog", fiber.Map{
		"Breadcrumbs": bcHtml,
		"Products":    products,
	})
}
