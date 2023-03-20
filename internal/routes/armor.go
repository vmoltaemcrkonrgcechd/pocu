package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vmoltaemcrkonrgcechd/pocu/internal/entities"
	"github.com/vmoltaemcrkonrgcechd/pocu/internal/usecase"
	"net/http"
)

type armorRoutes struct {
	uc usecase.ArmorUseCase
}

func withArmorRoutes(app *fiber.App, uc usecase.ArmorUseCase) {
	r := armorRoutes{uc}

	app.Post("/armor", r.add)
}

//	@tags		armor
//	@param		armor	body	entities.ArmorDTO	true	"броня"
//	@success	201
//	@router		/armor [post]
func (r armorRoutes) add(ctx *fiber.Ctx) error {
	var armor entities.ArmorDTO

	err := ctx.BodyParser(&armor)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "неверный json")
	}

	if err = r.uc.Add(armor); err != nil {
		return err
	}

	ctx.Status(http.StatusCreated)

	return nil
}
