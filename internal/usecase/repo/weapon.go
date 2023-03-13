package repo

import (
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/vmoltaemcrkonrgcechd/pocu/internal/entities"
	"github.com/vmoltaemcrkonrgcechd/pocu/pkg/psql"
	"github.com/vmoltaemcrkonrgcechd/pocu/pkg/sqhelper"
	"log"
	"net/http"
)

type WeaponRepo struct {
	*psql.PSQL
}

func NewWeaponRepo(pg *psql.PSQL) WeaponRepo {
	return WeaponRepo{pg}
}

func (r WeaponRepo) Add(weapon entities.WeaponDTO) error {
	if _, err := r.Sq.Insert("weapon").
		Columns("name", "attack", "weight").
		Values(weapon.Name, weapon.Attack, weapon.Weight).Exec(); err != nil {
		log.Println(err)
		return fiber.NewError(http.StatusInternalServerError,
			"произошла ошибка при добавлении оружия")
	}

	return nil
}

type AllWeaponsQP struct {
	Attack  []float32 `query:"attack"`
	Weight  []float32 `query:"weight"`
	OrderBy []string  `query:"orderBy"`
	Limit   uint64    `query:"limit"`
	Offset  uint64    `query:"offset"`
}

func (r WeaponRepo) All(qp AllWeaponsQP) (weapons entities.AllWeaponsDTO, err error) {
	subQ, _ := r.Sq.Select("weapon_id", "min(attack) OVER () min_attack",
		"max(attack) OVER () max_attack", "min(weight) OVER () min_weight",
		"max(weight) OVER () max_weight").
		From("weapon").Prefix("(").Suffix(")").MustSql()

	sb := r.Sq.Select("weapon_id", "name", "attack", "weight",
		"min_attack", "max_attack", "min_weight", "max_weight", "count(*) OVER ()").
		From("weapon").Join(fmt.Sprintf("%s w USING (weapon_id)", subQ))

	sb = sqhelper.MoreLess(sb, qp.Attack, "attack")
	sb = sqhelper.MoreLess(sb, qp.Weight, "weight")
	sb = sqhelper.OrderBy(sb, qp.OrderBy)
	sb = sqhelper.LimitOffset(sb, qp.Limit, qp.Offset)

	var (
		errAll = fiber.NewError(http.StatusInternalServerError,
			"произошла ошибка при получении оружия")
		rows *sql.Rows
		w    entities.Weapon
	)

	if rows, err = sb.Query(); err != nil {
		log.Println(err)
		return weapons, errAll
	}

	for rows.Next() {
		if err = rows.Scan(&w.ID, &w.Name, &w.Attack, &w.Weight, &weapons.MinAttack,
			&weapons.MaxAttack, &weapons.MinWeight, &weapons.MaxWeight,
			&weapons.Quantity); err != nil {
			log.Println(err)
			return weapons, errAll
		}

		weapons.Weapons = append(weapons.Weapons, w)
	}

	return weapons, nil
}

func (r WeaponRepo) Edit(weapon entities.WeaponDTO, id uint16) error {
	if _, err := r.Sq.Update("weapon").
		Set("name", weapon.Name).
		Set("attack", weapon.Attack).
		Set("weight", weapon.Weight).
		Where("weapon_id = ?", id).Exec(); err != nil {
		log.Println(err)
		return fiber.NewError(http.StatusInternalServerError,
			"произошла ошибка при редактировании оружия")
	}

	return nil
}

func (r WeaponRepo) Delete(id uint16) error {
	if _, err := r.Sq.Delete("weapon").
		Where("weapon_id = ?", id).Exec(); err != nil {
		log.Println(err)
		return fiber.NewError(http.StatusInternalServerError,
			"произошла ошибка при удалении оружия")
	}

	return nil
}
