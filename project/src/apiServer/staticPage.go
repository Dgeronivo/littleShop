package apiserver

import (
	"littleShop/project/src/staticParts"

	"github.com/gofiber/fiber/v2"
)

const staticPageDir = "staticPages"

func handleDelivery(c *fiber.Ctx) error {
	deliveryBreadcrumbs := staticParts.Breadcrumbs{Href: "#", Name: "Доставка та оплата", Active: true}
	breadcrmbHtml := staticParts.GenerateBreadcrumbsHtml(deliveryBreadcrumbs)

	return c.Render(staticPageDir+"/delivery", fiber.Map{
		"Breadcrumbs": breadcrmbHtml,
		"Page":        "delivery",
	})
}

func handleAboutUs(c *fiber.Ctx) error {
	deliveryBreadcrumbs := staticParts.Breadcrumbs{Href: "#", Name: "Про магазин", Active: true}
	breadcrumbsHtml := staticParts.GenerateBreadcrumbsHtml(deliveryBreadcrumbs)

	return c.Render(staticPageDir+"/aboutUs", fiber.Map{
		"Breadcrumbs": breadcrumbsHtml,
		"Page":        "aboutUs",
	})
}

func handlePolicy(c *fiber.Ctx) error {
	deliveryBreadcrumbs := staticParts.Breadcrumbs{Href: "#", Name: "Політика конфіденційності", Active: true}
	breadcrumbsHtml := staticParts.GenerateBreadcrumbsHtml(deliveryBreadcrumbs)

	return c.Render(staticPageDir+"/policy", fiber.Map{
		"Breadcrumbs": breadcrumbsHtml,
		"Page":        "policy",
	})
}
