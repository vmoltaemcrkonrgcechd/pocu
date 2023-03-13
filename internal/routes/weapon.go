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

	app.Post("/weapons", r.add)
	app.Get("/weapons", r.all)
	app.Patch("/weapons/:id", r.edit)
	app.Delete("/weapons/:id", r.delete)
}

// @tags		weapons
// @param		weapon	body	entities.WeaponDTO	true	"оружие"
// @success	201
// @router		/weapons [post]
func (r weaponRoutes) add(ctx *fiber.Ctx) error {
	var weapon entities.WeaponDTO

	err := ctx.BodyParser(&weapon)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest,
			"неверный json")
	}

	if err = r.uc.Add(weapon); err != nil {
		return err
	}

	ctx.Status(http.StatusCreated)

	return nil
}

// @tags		weapons
// @param		attack	query		string	false	"атака. пример: [мин],[макс]"
// @param		weight	query		string	false	"масса. пример: [мин],[макс]"
// @param		orderBy	query		string	false	"сортировать по. пример: [имя поля],[asc/desc]"
// @param		limit	query		int		false	"ограничение"
// @param		offset	query		int		false	"смещение"
// @success	200		{object}	entities.AllWeaponsDTO
// @router		/weapons [get]
func (r weaponRoutes) all(ctx *fiber.Ctx) error {
	qp := new(repo.AllWeaponsQP)

	err := ctx.QueryParser(qp)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest,
			"неверные параметры запроса")
	}

	var weapons entities.AllWeaponsDTO
	if weapons, err = r.uc.All(*qp); err != nil {
		return err
	}

	return ctx.JSON(weapons)
}

// @tags		weapons
// @param		id		path	int					true	"идентификатор оружия"
// @param		weapon	body	entities.WeaponDTO	true	"оружие"
// @success	200
// @router		/weapons/{id} [patch]
func (r weaponRoutes) edit(ctx *fiber.Ctx) error {
	var (
		weapon entities.WeaponDTO
		id     int
	)

	err := ctx.BodyParser(&weapon)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest,
			"неверный json")
	}

	if id, err = ctx.ParamsInt("id"); err != nil {
		return fiber.NewError(http.StatusBadRequest,
			"неверный идентификатор")
	}

	if err = r.uc.Edit(weapon, uint16(id)); err != nil {
		return err
	}

	return nil
}

// @tags		weapons
// @param		id	path	int	true	"идентификатор оружия"
// @success	200
// @router		/weapons/{id} [delete]
func (r weaponRoutes) delete(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return fiber.NewError(http.StatusBadRequest,
			"неверный идентификатор")
	}

	if err = r.uc.Delete(uint16(id)); err != nil {
		return err
	}

	return nil
}
