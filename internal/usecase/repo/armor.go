package repo

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vmoltaemcrkonrgcechd/pocu/internal/entities"
	"github.com/vmoltaemcrkonrgcechd/pocu/pkg/psql"
	"log"
	"net/http"
)

type ArmorRepo struct {
	*psql.PSQL
}

func NewArmorRepo(pg *psql.PSQL) ArmorRepo {
	return ArmorRepo{pg}
}

func (r ArmorRepo) Add(armor entities.ArmorDTO) error {
	if _, err := r.Sq.Insert("armor").
		Columns("name", "protection", "weight").
		Values(armor.Name, armor.Protection, armor.Weight).Exec(); err != nil {
		log.Println(err)
		return fiber.NewError(http.StatusInternalServerError,
			"произошла ошибка при добавлении брони")
	}

	return nil
}
