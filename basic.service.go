package basic

import (
	"fmt"

	"strings"

	"github.com/Jeffail/gabs/v2"
	"github.com/gofiber/fiber/v2"
)

func GetMsg(c *fiber.Ctx) error {
	c.JSON(&fiber.Map{"data": "hello freshers"})
	return nil
}

func GetMsgName(c *fiber.Ctx) error {
	str := fmt.Sprintf("%s", c.Params("name"))
	str = strings.Join(strings.Split(str, "+"), " ")
	return c.SendString(str)
}

func AddMsg(c *fiber.Ctx) error {
	apiBody, _ := gabs.ParseJSON([]byte(c.Body()))
	return c.JSON(apiBody)
}
