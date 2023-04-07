package usecase

import (
	"context"
	"deepflower/internal/model"
	"fmt"
)

type ProjectUsecase struct {
	Rep ProjectStorageInterface
}

func NewProjectUsecase(s ProjectStorageInterface) *ProjectUsecase {
	return &ProjectUsecase{Rep: s}
}

// TODO проверка что мечта опубликована
func (d *ProjectUsecase) CreateProject(ctx context.Context, name, info, graph, dreamName, creater string) (model.Project, error) {
	project, err := d.Rep.CreateProject(ctx, name, info, graph, dreamName, creater)
	if err != nil {
		return model.Project{}, err
	}
	return project, nil

}

func (d *ProjectUsecase) SearchProjects(ctx context.Context, userId string,
	limit uint64, offset uint64, onlyMyProjects bool, order string, searchTerm string,
	sort string) ([]model.Project, int, error) {
	// search
	projects, cnt, err := d.Rep.SearchProjects(ctx, userId,
		limit, offset, onlyMyProjects, order, searchTerm, sort)
	if err != nil {
		return []model.Project{}, 0, err
	}
	return projects, cnt, nil
}

func (d *ProjectUsecase) PublishProject(ctx context.Context, userId, projectId string) error {
	project, err := d.Rep.GetProjectById(ctx, projectId)
	if err != nil {
		return err
	}
	if project.Creater != userId {
		return fmt.Errorf("error: not available for user: %s", userId)
	}
	if project.Published {
		return fmt.Errorf("error: project has already been published")
	}

	// объединить?
	if err := d.Rep.EnergyTxUserToProject(ctx, userId, projectId, EnergyForPublish); err != nil {
		return err
	}
	if _, err := d.Rep.UpdateUserProject(ctx, projectId, map[string]interface{}{"Published": true}); err != nil {
		return err
	}
	return nil
}

func (d *ProjectUsecase) AddEnergyToProject(ctx context.Context, userId, projectId string, energy uint64) error {
	project, err := d.Rep.GetProjectById(ctx, projectId)
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
func (d *ProjectUsecase) DeleteUserProject(ctx context.Context, userId, projectId string) error {
	project, err := d.Rep.GetProjectById(ctx, projectId)
	if err != nil {
		return err
	}
	if project.Creater != userId || project.Published {
		return fmt.Errorf("not available")
	}
	if err := d.Rep.DeleteUserProject(ctx, projectId); err != nil {
		return err
	}
	return nil
}

func (d *ProjectUsecase) AddEnergyToTask(ctx context.Context, userId, projectId, nodeId string, energy uint64) error {
	project, err := d.Rep.GetProjectById(ctx, projectId)
	if err != nil {
		return err
	}
	if !project.Published {
		return fmt.Errorf("not available")
	}
	// TODO список участников полного консенсуса
	if err := d.Rep.EnergyTxUserToTask(ctx, userId, projectId, nodeId, energy); err != nil {
		return err
	}
	return nil
}

func (d *ProjectUsecase) ToWorkTask(ctx context.Context, userId, projectId, nodeId string) error {
	project, err := d.Rep.GetProjectById(ctx, projectId)
	if err != nil {
		return err
	}
	if !project.Published {
		return fmt.Errorf("not available for no published project")
	}
	if err := d.Rep.UpdateTaskStatus(ctx, projectId, nodeId, "inwork"); err != nil {
		return err
	}
	return nil
}

func (d *ProjectUsecase) CloseTask(ctx context.Context, userId, projectId, nodeId string) error {
	project, err := d.Rep.GetProjectById(ctx, projectId)
	if err != nil {
		return err
	}
	if !project.Published {
		return fmt.Errorf("not available for no published project")
	}
	// check status task
	// change status tast to 'confirmation'
	if err := d.Rep.UpdateTaskStatus(ctx, projectId, nodeId, "confirmation"); err != nil {
		return err
	}
	// TODO
	// start consensus process
	// выбор в настройках тип конценсуса и инструмент консенсуса
	// если консенсус полный запустить процесс по отправке уведомлений в тг/другой инструмент для подтверждения
	//const consensus_tool = "telegram"
	//const consensus_type = "total"

	// if err := GoConsensusProcessToNode(userId, projectId, nodeId); err != nil {
	// 	// откат состояния задачи до created
	// 	return err

	// }
	//
	return nil
}
