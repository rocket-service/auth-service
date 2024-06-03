package socket

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"io"
	"log"
	"net/http"
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
	method := c.Query("m")

	var (
		mt  int
		msg []byte
		err error
	)

	type AuthModel struct {
		UserKey  string `json:"k"`
		UserHwid string `json:"h"`
	}

	for {
		if method == "auth" {
			if mt, msg, err = c.ReadMessage(); err != nil {
				log.Println("read:", err)
				break
			}

			authModel := AuthModel{}

			err := json.Unmarshal(msg, &authModel)
			if err != nil {
				return
			}

			// @note: don;t use request to web-api. Make communication with database
			resp, _ := http.Get(fmt.Sprintf("http://localhost:3070/api/v1/auth?key=%s&hwid=%s", authModel.UserKey, authModel.UserHwid))
			body, err := io.ReadAll(resp.Body)

			if err = c.WriteMessage(mt, body); err != nil {
				log.Println("write:", err)
				break
			}
		}
	}
}
