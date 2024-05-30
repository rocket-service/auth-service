package socket

import (
	"fmt"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"log"
)

func Initialize(c *fiber.Ctx) error {
	fmt.Println("User:", c.IP(), "is try to connect")
	if websocket.IsWebSocketUpgrade(c) {
		fmt.Println("User:", c.IP(), "success connected")

		c.Locals("Allowed", true)
		return c.Next()
	}
	fmt.Println("User:", c.IP(), "failed to connect")

	return fiber.ErrUpgradeRequired
}

func Handle(c *websocket.Conn) {
	log.Println(c.Locals("Allowed"))

	var (
		mt  int
		msg []byte
		err error
	)
	for {
		if mt, msg, err = c.ReadMessage(); err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", msg)

		if err = c.WriteMessage(mt, msg); err != nil {
			log.Println("write:", err)
			break
		}
	}
}
