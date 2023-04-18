package usecase

import (
	"context"
	"deepflower/internal/model"
	"fmt"
)

type DreamUsecase struct {
	Tranzactor   Tranzactor
	DreamStorage DreamStorageInterface
	UserStorage  UserStorageInterface
}

func NewDreamUsecase(ds DreamStorageInterface, us UserStorageInterface, tx Tranzactor) *DreamUsecase {
	return &DreamUsecase{DreamStorage: ds, UserStorage: us, Tranzactor: tx}
}

func (d *DreamUsecase) CreateDream(ctx context.Context, name, info, location, creater string) (model.Dream, error) {
	dream, err := d.DreamStorage.CreateDream(ctx, name, info, location, creater)
	if err != nil {
		return model.Dream{}, err
	}
	return dream, nil

}

// не используется
func (d *DreamUsecase) GetAllUserDreams(ctx context.Context, userId string) ([]model.Dream, error) {
	dreams, err := d.DreamStorage.GetAllUserDreams(ctx, userId)
	if err != nil {
		return []model.Dream{}, err
	}
	return dreams, nil
}

func (d *DreamUsecase) SearchDreams(ctx context.Context, userId string,
	limit uint64, offset uint64, onlyMyDreams bool, order string, searchTerm string,
	sort string) ([]model.Dream, int, error) {
	dreams, cnt, err := d.DreamStorage.SearchDreams(ctx, userId,
		limit, offset, onlyMyDreams, order, searchTerm, sort)
	if err != nil {
		return []model.Dream{}, 0, err
	}
	return dreams, cnt, nil
}

func (d *DreamUsecase) PublishDream(ctx context.Context, userId, dreamId string) error {
	dream, err := d.DreamStorage.GetDreamById(ctx, dreamId)
	if err != nil {
		return err
	}
	if dream.Creater != userId {
		return fmt.Errorf("error: not available for user: %s", userId)
	}
	if dream.Published {
		return fmt.Errorf("error: dream has already been published")
	}
	if err = d.Tranzactor.WithTx(ctx, func(ctx context.Context) error {
		if err = d.UserStorage.SubtractEnergy(ctx, userId, EnergyForPublish); err != nil {
			return err
		}

		if err = d.DreamStorage.AddEnergy(ctx, dreamId, EnergyForPublish); err != nil {
			return err
		}
		if err := d.DreamStorage.UpdateDreamPublished(ctx, dreamId); err != nil {

			return err
		}
		return nil

	}); err != nil {
		return err
	}
	return nil
}

func (d *DreamUsecase) AddEnergyToDream(ctx context.Context, userId, dreamId string, energy uint64) error {
	dream, err := d.DreamStorage.GetDreamById(ctx, dreamId)
	if err != nil {
		return err
	}
	if !dream.Published {
		return fmt.Errorf("error: article not published")
	}

	err = d.Tranzactor.WithTx(ctx, func(ctx context.Context) error {
		if err = d.UserStorage.SubtractEnergy(ctx, userId, EnergyForPublish); err != nil {
			return err
		}

		if err = d.DreamStorage.AddEnergy(ctx, dreamId, EnergyForPublish); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

/*
func (d *DreamUsecase) UpdateUserDream(ctx context.Context, userId, dreamId string, patchDream map[string]interface{}) (model.Dream, error) {
	dream, err := d.DreamStorage.GetDreamById(ctx, dreamId)
	if err != nil {
		return model.Dream{}, err
	}
	if dream.Creater != userId {
		return model.Dream{}, fmt.Errorf("error: not available for user: %s", userId)
	}
	if dream.Published {
		return model.Dream{}, fmt.Errorf("you can't edit a published dream")
	}
	dreamUpdated, err := d.DreamStorage.UpdateUserDream(ctx, dreamId, patchDream)
	if err != nil {
		return model.Dream{}, err
	}
	return dreamUpdated, nil
}
*/

func (d *DreamUsecase) DeleteUserDream(ctx context.Context, userId, dreamId string) error {
	dream, err := d.DreamStorage.GetDreamById(ctx, dreamId)
	if err != nil {
		return err
	}
	if dream.Creater != userId || dream.Published {
		return fmt.Errorf("not available")
	}
	if err := d.DreamStorage.DeleteUserDream(ctx, dreamId); err != nil {
		return err
	}
	return nil
}
