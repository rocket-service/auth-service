package router

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"go-service-template/internal/router/socket"
)

func Serve() (app *fiber.App) {
	app = fiber.New(fiber.Config{
		DisableStartupMessage: true,
		EnablePrintRoutes:     true,
	})

	app.Use("/ws", socket.Initialize)
	app.Get("/ws/:id", websocket.New(socket.Handle))

	return
}
