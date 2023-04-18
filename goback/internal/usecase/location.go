package usecase

import (
	"context"
	"deepflower/internal/model"
	"fmt"
)

type LocationUsecase struct {
	Tranzactor      Tranzactor
	LocationStorage LocationStorageInterface
	UserStorage     UserStorageInterface
}

func NewLocationUsecase(ls LocationStorageInterface, us UserStorageInterface, tx Tranzactor) *LocationUsecase {
	return &LocationUsecase{LocationStorage: ls, UserStorage: us, Tranzactor: tx}
}

func (d *LocationUsecase) CreateLocation(ctx context.Context, creater string, Name string, info string, geolocation string, radius uint64, height uint64) (model.Location, error) {
	location, err := d.LocationStorage.CreateLocation(ctx, creater, Name, info, geolocation, radius, height)
	if err != nil {
		return model.Location{}, err
	}
	return location, nil

}

func (d *LocationUsecase) SearchLocations(ctx context.Context, userId string,
	limit uint64, offset uint64, onlyMyLocations bool,
	order string, searchTerm string,
	sort string) ([]model.Location, int, error) {
	locations, cnt, err := d.LocationStorage.SearchLocations(ctx, userId,
		limit, offset, onlyMyLocations,
		order, searchTerm,
		sort)
	if err != nil {
		return []model.Location{}, 0, err
	}
	return locations, cnt, nil
}

func (d *LocationUsecase) GetLocationDreams(ctx context.Context, locationId string) ([]model.Dream, error) {
	dreams, err := d.LocationStorage.GetLocationDreams(ctx, locationId)
	if err != nil {
		return []model.Dream{}, err
	}
	return dreams, nil
}

/*
	func (d *LocationUsecase) UpdateUserLocation(ctx context.Context, userId, locationId string, patchLocation map[string]interface{}) (model.Location, error) {
		location, err := d.LocationStorage.GetLocationById(ctx, locationId)
		if err != nil {
			return model.Location{}, err
		}
		if location.Creater != userId {
			return model.Location{}, fmt.Errorf("error:update not available for user: %s", userId)
		}
		locationUpdated, err := d.LocationStorage.UpdateUserLocation(ctx, locationId, patchLocation)
		if err != nil {
			return model.Location{}, err
		}
		return locationUpdated, nil
	}
*/
func (d *LocationUsecase) DeleteUserLocation(ctx context.Context, userId string, locationId string) error {
	location, err := d.LocationStorage.GetLocationById(ctx, locationId)
	if err != nil {
		return err
	}
	if location.Creater != userId {
		return fmt.Errorf("error: delete not available for user: %s", userId)
	}
	if err := d.LocationStorage.DeleteUserLocation(ctx, locationId); err != nil {
		return err
	}
	return nil
}

func (d *LocationUsecase) EnergyTxUserToLocation(ctx context.Context, userId, locationId string, energy uint64) error {
	err := d.Tranzactor.WithTx(ctx, func(ctx context.Context) error {
		if err := d.UserStorage.SubtractEnergy(ctx, userId, energy); err != nil {
			return err
		}
		if err := d.LocationStorage.AddEnergy(ctx, locationId, energy); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
