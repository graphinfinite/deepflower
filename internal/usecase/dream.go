package usecase

import "deepflower/internal/model"

type DreamStorageInterface interface {
	CreateDream(name, info, location, creater string) (model.Dream, error)
}

type DreamUsecase struct {
	Rep DreamStorageInterface
}

func (d *DreamUsecase) CreateDream(name, info, location, creater string) (model.Dream, error) {
	dream, err := d.Rep.CreateDream(name, info, location, creater)
	if err != nil {
		return model.Dream{}, err
	}
	return dream, nil

}
func (d *DreamUsecase) GetUserDreamById() {

}
func (d *DreamUsecase) UpdateUserDreamById() {

}
func (d *DreamUsecase) DeleteUserDreamById() {

}

func (d *DreamUsecase) PushUserDreamById() {

}
