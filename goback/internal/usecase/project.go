package usecase

import (
	"context"
	"deepflower/internal/model"
	"fmt"
)

type ProjectUsecase struct {
	Tranzactor     Tranzactor
	ProjectStorage ProjectStorageInterface
	UserStorage    UserStorageInterface
}

func NewProjectUsecase(ps ProjectStorageInterface, us UserStorageInterface, tx Tranzactor) *ProjectUsecase {
	return &ProjectUsecase{ProjectStorage: ps, UserStorage: us, Tranzactor: tx}
}

// TODO проверка что мечта опубликована
func (d *ProjectUsecase) CreateProject(ctx context.Context, name, info, graph, dreamName, creater string) (model.Project, error) {
	project, err := d.ProjectStorage.CreateProject(ctx, name, info, graph, dreamName, creater)
	if err != nil {
		return model.Project{}, err
	}
	return project, nil

}

func (d *ProjectUsecase) SearchProjects(ctx context.Context, userId string,
	limit uint64, offset uint64, onlyMyProjects bool, order string, searchTerm string,
	sort string) ([]model.Project, int, error) {
	// search
	projects, cnt, err := d.ProjectStorage.SearchProjects(ctx, userId,
		limit, offset, onlyMyProjects, order, searchTerm, sort)
	if err != nil {
		return []model.Project{}, 0, err
	}
	return projects, cnt, nil
}

func (d *ProjectUsecase) PublishProject(ctx context.Context, userId, projectId string) error {
	project, err := d.ProjectStorage.GetProjectById(ctx, projectId)
	if err != nil {
		return err
	}
	if project.Creater != userId {
		return fmt.Errorf("error: not available for user: %s", userId)
	}
	if project.Published {
		return fmt.Errorf("error: project has already been published")
	}

	if err := d.Tranzactor.WithTx(ctx, func(ctx context.Context) error {
		if err = d.UserStorage.SubtractEnergy(ctx, userId, EnergyForPublish); err != nil {
			return err
		}

		if err = d.ProjectStorage.AddEnergyToProject(ctx, projectId, EnergyForPublish); err != nil {
			return err
		}

		if err = d.ProjectStorage.UpdateProjectToPublished(ctx, projectId); err != nil {
			return err
		}

		return nil

	}); err != nil {
		return err
	}
	return nil
}

/*

func (d *ProjectUsecase) AddEnergyToProject(ctx context.Context, userId, projectId string, energy uint64) error {
	project, err := d.ProjectStorage.GetProjectById(ctx, projectId)
	if err != nil {
		return err
	}
	if !project.Published {
		return fmt.Errorf("error: project not published")
	}
	if err := d.Rep.EnergyTxUserToProject(ctx, userId, projectId, energy); err != nil {
		return err
	}
	return nil
}

*/

/*

// TODO no realize
func (d *ProjectUsecase) UpdateUserProject(ctx context.Context, userId, projectId string, patchProject map[string]interface{}) (model.Project, error) {
	project, err := d.Rep.GetProjectById(ctx, projectId)
	if err != nil {
		return model.Project{}, err
	}
	if project.Creater != userId {
		return model.Project{}, fmt.Errorf("error: not available for user: %s", userId)
	}
	if project.Published {
		return model.Project{}, fmt.Errorf("you can't edit a published project")
	}
	projectUpdated, err := d.Rep.UpdateUserProject(ctx, projectId, patchProject)
	if err != nil {
		return model.Project{}, err
	}
	return projectUpdated, nil
}
*/

func (d *ProjectUsecase) DeleteUserProject(ctx context.Context, userId, projectId string) error {
	project, err := d.ProjectStorage.GetProjectById(ctx, projectId)
	if err != nil {
		return err
	}
	if project.Creater != userId || project.Published {
		return fmt.Errorf("not available")
	}
	if err := d.ProjectStorage.DeleteUserProject(ctx, projectId); err != nil {
		return err
	}
	return nil
}
