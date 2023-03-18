package usecase

import "deepflower/internal/model"

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

func (d *DreamUsecase) GetUserDreamById() {

}
func (d *DreamUsecase) UpdateUserDreamById() {

}
func (d *DreamUsecase) DeleteUserDreamById() {

}

func (d *DreamUsecase) PushUserDreamById() {

}
