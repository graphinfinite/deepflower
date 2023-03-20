package usecase

import (
	"deepflower/internal/model"
	"fmt"
)

type DreamUsecase struct {
	Rep DreamStorageInterface
}

func NewDreamUsecase(s DreamStorageInterface) DreamUsecase {
	return DreamUsecase{Rep: s}
}

func (d *DreamUsecase) CreateDream(name, info, location, creater string) (model.Dream, error) {
	dream, err := d.Rep.CreateDream(name, info, location, creater)
	if err != nil {
		return model.Dream{}, err
	}
	return dream, nil

}

func (d *DreamUsecase) GetAllUserDreams(userId string) ([]model.Dream, error) {
	dreams, err := d.Rep.GetAllUserDreams(userId)
	if err != nil {
		return []model.Dream{}, err
	}
	return dreams, nil
}

func (d *DreamUsecase) UpdateUserDream(userId, dreamId string, patchDream map[string]interface{}) (model.Dream, error) {
	dream, err := d.Rep.GetDreamById(dreamId)
	if err != nil {
		return model.Dream{}, err
	}
	if dream.Creater != userId || dream.Publised {
		return model.Dream{}, fmt.Errorf("not available")
	}

	// Энергию нельзя забрать у мечты
	energyNew, ok := patchDream["Energy"]
	if ok {
		energyNew, _ := energyNew.(uint64)
		if energyNew-dream.Energy <= 0 {
			return model.Dream{},
				fmt.Errorf("the new energy must be greater than the original")
		}
	}

	dreamUpdated, err := d.Rep.UpdateUserDream(dreamId, patchDream)
	if err != nil {
		return model.Dream{}, err
	}
	return dreamUpdated, nil
}
func (d *DreamUsecase) DeleteUserDream(userId, dreamId string) error {
	dream, err := d.Rep.GetDreamById(dreamId)
	if err != nil {
		return err
	}
	if dream.Creater != userId || dream.Publised {
		// TODO
		return fmt.Errorf("not available")
	}
	if err := d.Rep.DeleteUserDream(dreamId); err != nil {
		return err
	}
	return nil
}
