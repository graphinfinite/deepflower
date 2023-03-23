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

func (d *DreamUsecase) SearchDreams(ctx context.Context, userId string,
	limit uint64, offset uint64, onlyMyDreams bool, order string, searchTerm string,
	sort string) ([]model.Dream, error) {

	// search
	dreams, err := d.Rep.SearchDreams(ctx, userId,
		limit, offset, onlyMyDreams, order, searchTerm, sort)
	if err != nil {
		return []model.Dream{}, err
	}
	return dreams, nil
}

const EnergyForPublish uint64 = 1

func (d *DreamUsecase) PublishDream(ctx context.Context, userId, dreamId string) error {
	dream, err := d.Rep.GetDreamById(ctx, dreamId)
	if err != nil {
		return err
	}

	if dream.Creater != userId {
		return fmt.Errorf("error: not available for user: %s", userId)
	}
	if dream.Published {
		return fmt.Errorf("error: dream has already been published")
	}
	if err := d.Rep.EnergyTxUserToDream(ctx, userId, dreamId, EnergyForPublish); err != nil {
		return err
	}
	if _, err := d.Rep.UpdateUserDream(ctx, dreamId, map[string]interface{}{"Published": true}); err != nil {
		return err
	}
	return nil
}

func (d *DreamUsecase) AddEnergyToDream(ctx context.Context, userId, dreamId string, energy uint64) error {
	dream, err := d.Rep.GetDreamById(ctx, dreamId)
	if err != nil {
		return err
	}
	if !dream.Published {
		return fmt.Errorf("error: article not published")
	}
	if err := d.Rep.EnergyTxUserToDream(ctx, userId, dreamId, energy); err != nil {
		return err
	}
	return nil
}

func (d *DreamUsecase) UpdateUserDream(ctx context.Context, userId, dreamId string, patchDream map[string]interface{}) (model.Dream, error) {
	dream, err := d.Rep.GetDreamById(ctx, dreamId)
	if err != nil {
		return model.Dream{}, err
	}
	if dream.Creater != userId {
		return model.Dream{}, fmt.Errorf("error: not available for user: %s", userId)
	}
	if dream.Published {
		return model.Dream{}, fmt.Errorf("you can't edit a published dream")
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
