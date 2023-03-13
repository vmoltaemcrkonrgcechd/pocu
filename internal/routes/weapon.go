package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vmoltaemcrkonrgcechd/pocu/internal/entities"
	"github.com/vmoltaemcrkonrgcechd/pocu/internal/usecase"
	"github.com/vmoltaemcrkonrgcechd/pocu/internal/usecase/repo"
	"net/http"
)

type weaponRoutes struct {
	uc usecase.WeaponUseCase
}

func withWeaponRoutes(app *fiber.App, uc usecase.WeaponUseCase) {
	r := weaponRoutes{uc}

	app.Get("/weapons", r.all)
}

//	@tags		weapons
//	@param		attack	query		string	false	"атака. пример: [мин],[макс]"
//	@param		weight	query		string	false	"масса. пример: [мин],[макс]"
//	@param		orderBy	query		string	false	"сортировать по. пример: [имя поля],[asc/desc]"
//	@param		limit	query		int		false	"ограничение"
//	@param		offset	query		int		false	"смещение"
//	@success	200		{object}	entities.AllWeaponsDTO
//	@router		/weapons [get]
func (r weaponRoutes) all(ctx *fiber.Ctx) error {
	qp := new(repo.AllWeaponsQP)

	err := ctx.QueryParser(qp)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest,
			"")
	}

	var weapons entities.AllWeaponsDTO
	if weapons, err = r.uc.All(*qp); err != nil {
		return err
	}

	return ctx.JSON(weapons)
}
