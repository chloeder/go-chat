package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/kooroshh/fiber-boostrap/app/controllers"
	"github.com/kooroshh/fiber-boostrap/pkg/middleware"
)

type ApiRouter struct{}

func (h ApiRouter) InstallRouter(app *fiber.App) {
	api := app.Group("/api", limiter.New())
	api.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Hello from api",
		})
	})

	userGroup := app.Group("/user")
	userGroup.Post("/register", controllers.Register)
	userGroup.Post("/login", controllers.Login)
	userGroup.Delete("/logout", middleware.MiddlewareValidateAuth, controllers.Logout)
}

func NewApiRouter() *ApiRouter {
	return &ApiRouter{}
}
