package bootstrap

import (
	"io"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
	"github.com/kooroshh/fiber-boostrap/app/ws"
	"github.com/kooroshh/fiber-boostrap/pkg/database"
	"github.com/kooroshh/fiber-boostrap/pkg/env"
	"github.com/kooroshh/fiber-boostrap/pkg/router"
	"go.elastic.co/apm/module/apmfiber"
)

func NewApplication() *fiber.App {
	env.SetupEnvFile()
	SetupLogger()

	database.SetupDatabase()
	database.SetupMongoDB()
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{Views: engine})
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(apmfiber.Middleware())
	app.Get("/dashboard", monitor.New())

	go ws.ServerWSMessage(app)

	router.InstallRouter(app)

	return app
}

func SetupLogger() {
	file, err := os.OpenFile("logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("Error opening log file:", err)
	}

	mw := io.MultiWriter(os.Stdout, file)
	log.SetOutput(mw)
}