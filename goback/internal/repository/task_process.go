package repository

import (
	"context"
	"deepflower/pkg/postgres"
	"fmt"
	"time"
)

type TaskProcessStorage struct {
	Db *postgres.PG
}

func NewTaskProcessStorage(db *postgres.PG) *TaskProcessStorage {
	return &TaskProcessStorage{Db: db}
}

func (s *TaskProcessStorage) UpsertTaskProcess(ctx context.Context, projectId, nodeId, userId, status string, taskEnerge, taskLeadTime uint64) (processId string, err error) {
	tx := s.Db.ExtractTx(ctx)
	query4 := `
	INSERT INTO "task_process" 
	(projectid, nodeid,exec_userid,
	inspectors_total,inspectors_confirmed,energy_total,
	leadtime,status) VALUES (
	$1,$2,$3, (SELECT count(*) FROM "task_users" WHERE projectid=$4 AND nodeid=$5), $6, $7, $8, $9)
	ON CONFLICT (nodeid) 
	DO 
	   UPDATE SET status=$10, updatedat=$11
	RETURNING id;`
	var insId string
	fmt.Println(projectId, nodeId, projectId, nodeId, userId, 0, taskEnerge, taskLeadTime, status, status, time.Now())
	err = tx.GetContext(ctx, &insId, query4, projectId, nodeId, userId, projectId, nodeId, 0, taskEnerge, taskLeadTime, status, status, time.Now())
	if err != nil {
		return "", err
	}
	return insId, nil
}
