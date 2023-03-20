package usecase

import (
	"github.com/vmoltaemcrkonrgcechd/pocu/internal/entities"
	"github.com/vmoltaemcrkonrgcechd/pocu/internal/usecase/repo"
	"github.com/vmoltaemcrkonrgcechd/pocu/pkg/validator"
)

type ArmorUseCase struct {
	repo repo.ArmorRepo
}

func NewArmorUseCase(repo repo.ArmorRepo) ArmorUseCase {
	return ArmorUseCase{repo}
}

func (uc ArmorUseCase) Add(armor entities.ArmorDTO) error {
	if err := uc.checkArmor(armor); err != nil {
		return err
	}

	return uc.repo.Add(armor)
}

func (uc ArmorUseCase) checkArmor(armor entities.ArmorDTO) error {
	val := validator.New()
	val.
		Check(armor.Protection < 0, "защита брони не может быть меньше нуля").
		Check(armor.Weight < 0, "вес брони не может быть меньше нуля").
		Check(len(armor.Name) < 1, "название брони должно содержать хотя бы один символ")

	return val.Verdict()
}
