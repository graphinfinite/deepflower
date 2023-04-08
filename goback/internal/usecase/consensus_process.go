package usecase

import "deepflower/internal/model"

type ConsensusProcess struct {
}

func NewConsensusProcess() *ConsensusProcess {
	return &ConsensusProcess{}
}

func (c *ConsensusProcess) StartConsensusProcess(processId string) error {

	return nil
}
func (c *ConsensusProcess) GetConsensusProcessById(processId string) (process model.ProcessTask, err error) {

	return model.ProcessTask{}, nil
}
