package main

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

func FiberIdGET(c *fiber.Ctx) error {
	return c.SendString(strconv.Itoa(server.add_client(&server)))
}
func FiberConfigGET(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(http.StatusUnprocessableEntity)
	}
	client, ok := server.clients[id]
	if ok {
		return c.JSON(client.Config)
	} else {
		return c.SendStatus(http.StatusNotFound)
	}
}
func FiberConfigPOST(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(http.StatusUnprocessableEntity)
	}
	client, ok := server.clients[id]
	if ok {
		err := c.BodyParser(&client.Config)
		if err != nil {
			return c.SendStatus(http.StatusNotAcceptable)
		}
		return c.SendStatus(http.StatusOK)
	}
	return c.SendStatus(http.StatusNotFound)
}
func FiberResultGET(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(http.StatusUnprocessableEntity)
	}
	client, ok := server.clients[id]
	if ok {
		err = getRoutine(client)
		if err != nil {
			return c.SendStatus(http.StatusInternalServerError)
		}
		if client.World == nil {
			return c.SendStatus(http.StatusInternalServerError)
		}
		return c.JSON(&client.World)
	} else {
		return c.SendStatus(http.StatusNotFound)
	}
}
func FiberLogGET(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(http.StatusUnprocessableEntity)
	}
	client, ok := server.clients[id]
	if ok {
		if client.World == nil {
			return c.SendStatus(http.StatusInternalServerError)
		}
		if !client.World.end {
			return c.SendStatus(http.StatusNotAcceptable)
		}
		return c.JSON(struct {
			Log string `json:"log"`
		}{
			Log: client.World.log,
		})
	} else {
		return c.SendStatus(http.StatusNotFound)
	}
}
