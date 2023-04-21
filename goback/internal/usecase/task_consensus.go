package usecase

import (
	"context"
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

	msg := fmt.Sprintf(`Подтвердите выполнение задачи:
	process.ID - %s
	process.ExecUserId %s
	process.ProjectId %s
	process.NodeId %s
	`, process.ID, process.ExecUserId, process.ProjectId, process.NodeId)

	print(msg)
	print(usersDests)

	if err := c.Bot.SendMessagesWithCallbacks(ctx, usersDests, msg, map[string]string{"ok": "ok"}); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
