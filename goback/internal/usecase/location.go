// CreateLocation(ctx context.Context, creater string, Name string, info string, geolocation string, radius uint64, height uint64, idFiles string) (model.Location, error)
// UpdateUserLocation(ctx context.Context, userId, locationId string, locationUpdate map[string]interface{}) (model.Location, error)
// DeleteUserLocation(ctx context.Context, userId string, locationId string) error
// AddEnergyToLocation(ctx context.Context, userId, locationId string, energy uint64) error
// SearchLocations(ctx context.Context, userId string,
// 	limit uint64, offset uint64, onlyMyLocations bool,
// 	order string, searchTerm string,
// 	sort string) ([]model.Location, int, error)

package usecase

import (
	"context"
	"deepflower/internal/model"
	"fmt"
)

type LocationUsecase struct {
	Rep LocationStorageInterface
}

func NewLocationUsecase(s LocationStorageInterface) LocationUsecase {
	return LocationUsecase{Rep: s}
}

func (d *LocationUsecase) CreateLocation(ctx context.Context, creater string, Name string, info string, geolocation string, radius uint64, height uint64, idFiles string) (model.Location, error) {
	location, err := d.Rep.CreateLocation(ctx, creater, Name, info, geolocation, radius, height, idFiles)
	if err != nil {
		return model.Location{}, err
	}
	return location, nil

}

func (d *LocationUsecase) SearchLocations(ctx context.Context, userId string,
	limit uint64, offset uint64, onlyMyLocations bool,
	order string, searchTerm string,
	sort string) ([]model.Location, int, error) {
	// search
	locations, cnt, err := d.Rep.SearchLocations(ctx, userId,
		limit, offset, onlyMyLocations,
		order, searchTerm,
		sort)
	if err != nil {
		return []model.Location{}, 0, err
	}
	return locations, cnt, nil
}

func (d *LocationUsecase) AddEnergyToLocation(ctx context.Context, userId, locationId string, energy uint64) error {
	//location, err := d.Rep.GetLocationById(ctx, locationId)
	// if err != nil {
	// 	return err
	// }
	if err := d.Rep.EnergyTxUserToLocation(ctx, userId, locationId, energy); err != nil {
		return err
	}
	return nil
}

func (d *LocationUsecase) UpdateUserLocation(ctx context.Context, userId, locationId string, patchLocation map[string]interface{}) (model.Location, error) {
	location, err := d.Rep.GetLocationById(ctx, locationId)
	if err != nil {
		return model.Location{}, err
	}
	if location.Creater != userId {
		return model.Location{}, fmt.Errorf("error: not available for user: %s", userId)
	}
	locationUpdated, err := d.Rep.UpdateUserLocation(ctx, locationId, patchLocation)
	if err != nil {
		return model.Location{}, err
	}
	return locationUpdated, nil
}
func (d *LocationUsecase) DeleteUserLocation(ctx context.Context, userId string, locationId string) error {
	location, err := d.Rep.GetLocationById(ctx, locationId)
	if err != nil {
		return err
	}
	if location.Creater != userId {
		// TODO
		return fmt.Errorf("not available")
	}
	if err := d.Rep.DeleteUserLocation(ctx, locationId); err != nil {
		return err
	}
	return nil
}
