package usecase

import "deepflower/internal/model"

type ConsensusProcess struct {
	Rep ProjectStorageInterface
}

func NewConsensusProcess(r ProjectStorageInterface) *ConsensusProcess {
	return &ConsensusProcess{Rep: r}
}

func (c *ConsensusProcess) StartConsensusProcess(processId string) error {

	return nil
}
func (c *ConsensusProcess) GetConsensusProcessById(processId string) (process model.ProcessTask, err error) {
	return model.ProcessTask{}, nil
}
