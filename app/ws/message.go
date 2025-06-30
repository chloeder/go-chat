package ws

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/kooroshh/fiber-boostrap/app/models"
	"github.com/kooroshh/fiber-boostrap/app/repositories"
	"github.com/kooroshh/fiber-boostrap/pkg/env"
)

func ServerWSMessage(app *fiber.App) {
	var clients = make(map[*websocket.Conn]bool)
	var broadcast = make(chan models.MessagePayload)

	app.Get("/message/send", websocket.New(func(c *websocket.Conn) {
		defer func() {
			c.Close()
			delete(clients, c)
		}()
		clients[c] = true
		for {
			var msg models.MessagePayload
			if err := c.ReadJSON(&msg); err != nil {
				log.Println("Error reading message:", err)
				break
			}
			msg.Date = time.Now()
			err := repositories.InsertMessage(context.Background(), msg)
			if err != nil {
				log.Println("Error inserting message:", err)
				break
			}
			broadcast <- msg
		}
	}))

	go func() {
		for {
			msg := <-broadcast
			for client := range clients {
				if err := client.WriteJSON(msg); err != nil {
					log.Println("Error writing message:", err)
				}
			}
		}
	}()

	log.Fatal(app.Listen(fmt.Sprintf("%s:%s", env.GetEnv("APP_HOST", "localhost"), env.GetEnv("APP_PORT_SOCKET", "8080"))))
}
