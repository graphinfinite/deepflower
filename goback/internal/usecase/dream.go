package usecase

import (
	"context"
	"deepflower/internal/model"
	"fmt"
)

type DreamUsecase struct {
	Rep DreamStorageInterface
}

func NewDreamUsecase(s DreamStorageInterface) DreamUsecase {
	return DreamUsecase{Rep: s}
}

func (d *DreamUsecase) CreateDream(ctx context.Context, name, info, location, creater string) (model.Dream, error) {
	dream, err := d.Rep.CreateDream(ctx, name, info, location, creater)
	if err != nil {
		return model.Dream{}, err
	}
	return dream, nil

}

func (d *DreamUsecase) GetAllUserDreams(ctx context.Context, userId string) ([]model.Dream, error) {
	dreams, err := d.Rep.GetAllUserDreams(ctx, userId)
	if err != nil {
		return []model.Dream{}, err
	}
	return dreams, nil
}

func (d *DreamUsecase) UpdateUserDream(ctx context.Context, userId, dreamId string, patchDream map[string]interface{}) (model.Dream, error) {
	dream, err := d.Rep.GetDreamById(ctx, dreamId)
	if err != nil {
		return model.Dream{}, err
	}
	if dream.Creater != userId || dream.Published {
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

	dreamUpdated, err := d.Rep.UpdateUserDream(ctx, dreamId, patchDream)
	if err != nil {
		return model.Dream{}, err
	}
	return dreamUpdated, nil
}
func (d *DreamUsecase) DeleteUserDream(ctx context.Context, userId, dreamId string) error {
	dream, err := d.Rep.GetDreamById(ctx, dreamId)
	if err != nil {
		return err
	}
	if dream.Creater != userId || dream.Published {
		// TODO
		return fmt.Errorf("not available")
	}
	if err := d.Rep.DeleteUserDream(ctx, dreamId); err != nil {
		return err
	}
	return nil
}
