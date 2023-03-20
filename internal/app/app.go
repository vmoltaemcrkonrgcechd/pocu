package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/vmoltaemcrkonrgcechd/pocu/config"
	"github.com/vmoltaemcrkonrgcechd/pocu/internal/routes"
	"github.com/vmoltaemcrkonrgcechd/pocu/internal/usecase"
	"github.com/vmoltaemcrkonrgcechd/pocu/internal/usecase/repo"
	"github.com/vmoltaemcrkonrgcechd/pocu/pkg/psql"
)

func Run(cfg *config.Config) error {
	app := fiber.New()

	app.Use(cors.New(cors.ConfigDefault))

	pg, err := psql.New(cfg)
	if err != nil {
		return err
	}

	routes.WithRouter(
		app,
		usecase.NewWeaponUseCase(repo.NewWeaponRepo(pg)),
		usecase.NewArmorUseCase(repo.NewArmorRepo(pg)))

	return app.Listen(":80")
}
