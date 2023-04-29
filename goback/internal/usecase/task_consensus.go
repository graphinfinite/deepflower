package usecase

import (
	"context"
	"deepflower/internal/model"
	"fmt"
)

type TaskConsensus struct {
	Tranzactor         Tranzactor
	Bot                BotInterface
	ProjectStorage     ProjectStorageInterface
	UserStorage        UserStorageInterface
	TaskStorage        TaskStorageInterface
	TaskUsersStorage   TaskUsersStorageInterface
	TaskProcessStorage TaskProcessStorageInterface
	TaskConsensus      TaskConsensusInterface
}

func NewTaskConsensus(
	tx Tranzactor,
	bot BotInterface,
	ps ProjectStorageInterface,
	us UserStorageInterface,
	ts TaskStorageInterface,
	tus TaskUsersStorageInterface,
	tps TaskProcessStorageInterface) *TaskConsensus {
	return &TaskConsensus{
		Tranzactor:         tx,
		Bot:                bot,
		ProjectStorage:     ps,
		UserStorage:        us,
		TaskStorage:        ts,
		TaskUsersStorage:   tus,
		TaskProcessStorage: tps,
	}
}

func (c *TaskConsensus) StartTaskConsensus(ctx context.Context, processId string) error {
	process, err := c.TaskProcessStorage.GetTaskConsensusProcessById(ctx, processId)
	if err != nil {
		return err
	}
	users, err := c.TaskUsersStorage.GetTaskUsersByTaskId(ctx, process.ProjectId, process.NodeId)
	if err != nil {
		return err
	}

	var usersDests = make([]int64, 0, len(users))
	fmt.Println("users", users)
	for _, user := range users {
		usersDests = append(usersDests, user.TgChatId)
		fmt.Println(user.TgChatId, usersDests)
	}

	execuser, err := c.UserStorage.GetUserById(ctx, process.ExecUserId)
	if err != nil {
		return err
	}

	project, err := c.ProjectStorage.GetProjectById(ctx, process.ProjectId)
	if err != nil {
		return err
	}
	task, err := c.TaskStorage.GetTaskData(ctx, project.ID, process.NodeId)
	if err != nil {
		return err
	}

	msg := fmt.Sprintf(`Подтвердите выполнение задачи:
	Исполнитель - %s
	Энергия задачи - %d
	Число участников - %d
	Проект - %s
	Описание задачи: 
	%s
	`,
		execuser.Username,
		process.EnergyTotal,
		process.InspectorsTotal,
		project.Name,
		task.Description,
	)

	ev := fmt.Sprintf("pc/%s/confirm", process.ID)
	if err := c.Bot.SendMessagesWithCallbacks(ctx, usersDests, msg, map[string]string{"ok": ev}); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (c *TaskConsensus) ConsensusConfirmation(ctx context.Context, processId string) error {
	pc, err := c.TaskProcessStorage.AddInspectorConfirmed(ctx, processId)
	if err != nil {
		return err
	}
	if pc.InspectorsConfirmed != pc.InspectorsTotal {
		return nil
	}
	if err := c.Tranzactor.WithTx(ctx, func(ctx context.Context) error {
		execuser, err := c.UserStorage.GetUserById(ctx, pc.ExecUserId)
		if err != nil {
			return err
		}
		// TODO
		// статус процесса = завершен
		_, err = c.TaskProcessStorage.UpsertTaskProcess(ctx, pc.ProjectId, pc.NodeId, pc.ExecUserId, TaskStatus_complited, 0, 0)
		if err != nil {
			return err
		}
		// статус задачи = выполнен
		if err := c.TaskStorage.UpdateTaskStatus(ctx, pc.ProjectId, pc.NodeId, TaskStatus_complited); err != nil {
			return err
		}
		// передать энергию от задачи к исполнителю
		if err := c.UserStorage.AddEnergy(ctx, pc.ExecUserId, pc.EnergyTotal); err != nil {
			return err
		}
		if err := c.TaskStorage.SubtractEnergy(ctx, pc.ProjectId, pc.NodeId, pc.EnergyTotal); err != nil {

			return err
		}
		// отправка уведомления исполнителю
		c.Bot.SendMessage(ctx, execuser.TgChatId, fmt.Sprintf("Confirmation complited for process %s", pc.ID))
		return nil

	}); err != nil {
		return err
	}

	return nil

}

func (c *TaskConsensus) GetAllUserTaskProcess(ctx context.Context, userId string) ([]model.ProcessTask, error) {
	process, err := c.TaskProcessStorage.GetTaskConsensusByExecUserId(ctx, userId)
	if err != nil {
		return []model.ProcessTask{}, err
	}
	return process, nil
}

func (c *TaskConsensus) SearchUserTaskProcesses(ctx context.Context, userId string,
	limit uint64, offset uint64, onlyActive bool, onlyForUser bool, order string, searchTerm string,
	sort string) ([]model.ProcessTask, int, error) {
	processes, cnt, err := c.TaskProcessStorage.SearchProcesses(ctx, userId,
		limit, offset, onlyActive, onlyForUser, order, searchTerm, sort)
	if err != nil {
		return []model.ProcessTask{}, 0, err
	}
	return processes, cnt, nil
}
