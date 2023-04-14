package usecase

import (
	"context"
	"fmt"
)

type ConsensusProcess struct {
	Bot        BotInterface
	RepProject ProjectStorageInterface
}

func NewConsensusProcess(rpj ProjectStorageInterface, bot BotInterface) *ConsensusProcess {
	return &ConsensusProcess{RepProject: rpj, Bot: bot}
}

func (c *ConsensusProcess) StartTaskConsensusProcess(ctx context.Context, processId string) error {
	process, err := c.RepProject.GetTaskConsensusProcessById(ctx, processId)
	if err != nil {
		return err
	}
	users, err := c.RepProject.SelectTaskUsers(ctx, process.ProjectId, process.NodeId)
	if err != nil {
		return err
	}

	var usersDests = make([]int64, 0, len(users))
	for _, user := range users {
		usersDests = append(usersDests, user.TgChatId)
	}

	msg := fmt.Sprintf(`Подтвердите выполнение задачи:
	process.ID - %s
	process.ExecUserId %s
	process.ProjectId %s
	process.NodeId %s
	`, process.ID, process.ExecUserId, process.ProjectId, process.NodeId)

	if err := c.Bot.SendMessagesWithOkButton(ctx, usersDests, msg); err != nil {
		return err
	}

	return nil
}

// func (c *ConsensusProcess) GetTaskConsensusProcessById(processId string) (process model.ProcessTask, err error) {
// 	return model.ProcessTask{}, nil
// }
