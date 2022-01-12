package basic

import (
	basicService "go-backend-application/services/basic"

	"github.com/gofiber/fiber/v2"
)

func GetMsg(c *fiber.Ctx) error {
	return basicService.GetMsg(c)
}

func GetMsgName(c *fiber.Ctx) error {
	return basicService.GetMsgName(c)
}

func AddMsg(c *fiber.Ctx) error {
	return basicService.AddMsg(c)
}
