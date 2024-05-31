package router

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"go-service-template/internal/router/socket"
)

func Serve() (app *fiber.App) {
	app = fiber.New(fiber.Config{
		AppName: "Rocket authorization service",
	})

	app.Use("/handle", socket.Initialize)
	app.Get("/handle", websocket.New(socket.Handle))

	return
}
