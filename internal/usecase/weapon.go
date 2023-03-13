package usecase

import (
	"github.com/vmoltaemcrkonrgcechd/pocu/internal/entities"
	"github.com/vmoltaemcrkonrgcechd/pocu/internal/usecase/repo"
	"github.com/vmoltaemcrkonrgcechd/pocu/pkg/validator"
)

type WeaponUseCase struct {
	repo repo.WeaponRepo
}

func NewWeaponUseCase(repo repo.WeaponRepo) WeaponUseCase {
	return WeaponUseCase{repo}
}

func (uc WeaponUseCase) Add(weapon entities.WeaponDTO) error {
	if err := uc.checkWeapon(weapon); err != nil {
		return err
	}
	
	return uc.repo.Add(weapon)
}

func (uc WeaponUseCase) All(qp repo.AllWeaponsQP) (entities.AllWeaponsDTO, error) {
	weapons, err := uc.repo.All(qp)
	if err != nil {
		return weapons, err
	}

	if weapons.Weapons == nil {
		weapons.Weapons = make([]entities.Weapon, 0, 0)
	}

	return weapons, nil
}

func (uc WeaponUseCase) Edit(weapon entities.WeaponDTO, id uint16) error {
	if err := uc.checkWeapon(weapon); err != nil {
		return err
	}

	return uc.repo.Edit(weapon, id)
}

func (uc WeaponUseCase) Delete(id uint16) error {
	return uc.repo.Delete(id)
}

func (uc WeaponUseCase) checkWeapon(weapon entities.WeaponDTO) error {
	val := validator.New()
	val.
		Check(weapon.Attack < 0, "урон оружия не может быть меньше нуля").
		Check(weapon.Weight < 0, "вес оружия не может быть меньше нуля").
		Check(len(weapon.Name) < 1, "название оружия должно содержать хотя бы один символ")

	return val.Verdict()
}
