package usecase

import (
	"github.com/vmoltaemcrkonrgcechd/pocu/internal/entities"
	"github.com/vmoltaemcrkonrgcechd/pocu/internal/usecase/repo"
)

type WeaponUseCase struct {
	repo repo.WeaponRepo
}

func NewWeaponUseCase(repo repo.WeaponRepo) WeaponUseCase {
	return WeaponUseCase{repo}
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
