package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	_ "github.com/vmoltaemcrkonrgcechd/pocu/docs"
	"github.com/vmoltaemcrkonrgcechd/pocu/internal/usecase"
)

func WithRouter(app *fiber.App, weaponUC usecase.WeaponUseCase) {
	app.Get("/swagger-ui/*", swagger.New(swagger.ConfigDefault))

	withWeaponRoutes(app, weaponUC)
}
