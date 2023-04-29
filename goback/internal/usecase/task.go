package usecase

import (
	"context"
	"deepflower/internal/model"
	"fmt"
)

type TaskUsecase struct {
	Tranzactor         Tranzactor
	ProjectStorage     ProjectStorageInterface
	UserStorage        UserStorageInterface
	TaskStorage        TaskStorageInterface
	TaskUsersStorage   TaskUsersStorageInterface
	TaskProcessStorage TaskProcessStorageInterface
	TaskConsensus      TaskConsensusInterface
}

func NewTaskUsecase(
	tx Tranzactor,
	ps ProjectStorageInterface,
	us UserStorageInterface,
	ts TaskStorageInterface,
	tus TaskUsersStorageInterface,
	tps TaskProcessStorageInterface,
	tc TaskConsensusInterface) *TaskUsecase {
	return &TaskUsecase{
		Tranzactor:         tx,
		ProjectStorage:     ps,
		UserStorage:        us,
		TaskStorage:        ts,
		TaskUsersStorage:   tus,
		TaskProcessStorage: tps,
		TaskConsensus:      tc,
	}
}

func (s *TaskUsecase) AddEnergyToTask(ctx context.Context, userId, projectId, nodeId string, energy uint64) error {
	project, err := s.ProjectStorage.GetProjectById(ctx, projectId)
	if err != nil {
		return err
	}
	if !project.Published {
		return fmt.Errorf("not available")
	}
	if err := s.Tranzactor.WithTx(ctx, func(ctx context.Context) error {
		if err = s.UserStorage.SubtractEnergy(ctx, userId, energy); err != nil {
			return err
		}
		if err = s.TaskStorage.AddEnergyToTask(ctx, projectId, nodeId, energy); err != nil {
			return err
		}
		s.TaskUsersStorage.AddTaskUser(ctx, userId, projectId, nodeId, energy)
		return nil

	}); err != nil {
		return err
	}
	return nil
}

func (s *TaskUsecase) ToWorkTask(ctx context.Context, userId, projectId, nodeId string) error {
	project, err := s.ProjectStorage.GetProjectById(ctx, projectId)
	if err != nil {
		return err
	}
	if !project.Published {
		return fmt.Errorf("not available for no published project")
	}
	if err := s.Tranzactor.WithTx(ctx, func(ctx context.Context) error {
		if err = s.TaskStorage.UpdateTaskStatus(ctx, projectId, nodeId, TaskStatus_inwork); err != nil {
			return err
		}
		task, err := s.TaskStorage.GetTaskData(ctx, projectId, nodeId)
		if err != nil {
			return err
		}
		_, err = s.TaskProcessStorage.UpsertTaskProcess(ctx, projectId, nodeId, userId, TaskStatus_inwork, task.Energy, task.LeadTime)

		if err != nil {
			return err
		}

		return nil

	}); err != nil {
		return err
	}
	return nil
}

func (s *TaskUsecase) CloseTask(ctx context.Context, userId, projectId, nodeId string) error {
	project, err := s.ProjectStorage.GetProjectById(ctx, projectId)
	if err != nil {
		return err
	}
	if !project.Published {
		return fmt.Errorf("not available for no published project")
	}

	var pc model.ProcessTask

	if err := s.Tranzactor.WithTx(ctx, func(ctx context.Context) error {
		if err = s.TaskStorage.UpdateTaskStatus(ctx, projectId, nodeId, TaskStatus_confirmation); err != nil {
			return err
		}
		task, err := s.TaskStorage.GetTaskData(ctx, projectId, nodeId)
		if err != nil {
			return err
		}
		pc, err = s.TaskProcessStorage.UpsertTaskProcess(ctx, projectId, nodeId, userId, TaskStatus_confirmation, task.Energy, task.LeadTime)

		if err != nil {
			return err
		}

		return nil

	}); err != nil {
		return err
	}

	//TODO
	fmt.Printf("START ---> PROCESS: %#v  ...", pc)
	// start consensus process
	if errProcess := s.TaskConsensus.StartTaskConsensus(ctx, pc.ID); err != nil {
		fmt.Println(errProcess)
		//// что-то откатываем
		return errProcess
	}
	return nil
}
