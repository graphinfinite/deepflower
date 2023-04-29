package repository

import (
	"context"
	"deepflower/internal/model"
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

func (s *TaskProcessStorage) GetTaskConsensusProcessById(ctx context.Context, processId string) (model.ProcessTask, error) {
	tx := s.Db.ExtractTx(ctx)
	query := `SELECT * FROM task_process WHERE id=$1;`
	var process model.ProcessTask
	if err := tx.GetContext(ctx, &process, query, processId); err != nil {
		return model.ProcessTask{}, err
	}
	return process, nil
}

func (s *TaskProcessStorage) UpsertTaskProcess(ctx context.Context, projectId, nodeId, userId, status string, taskEnerge, taskLeadTime uint64) (model.ProcessTask, error) {
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
	RETURNING *
	;`
	var process model.ProcessTask
	fmt.Println(projectId, nodeId, projectId, nodeId, userId, 0, taskEnerge, taskLeadTime, status, status, time.Now())

	err := tx.GetContext(ctx, &process, query4, projectId, nodeId, userId, projectId, nodeId, 0, taskEnerge, taskLeadTime, status, status, time.Now())
	if err != nil {
		return model.ProcessTask{}, err
	}
	return process, nil
}

func (s *TaskProcessStorage) GetTaskConsensusByExecUserId(ctx context.Context, userId string) ([]model.ProcessTask, error) {
	tx := s.Db.ExtractTx(ctx)
	q := `SELECT * FROM "task_process" WHERE exec_userid=$1;`
	var processes []model.ProcessTask

	if err := tx.SelectContext(ctx, processes, q, userId); err != nil {
		return []model.ProcessTask{}, err
	}
	return processes, nil
}

func (s *TaskProcessStorage) AddInspectorConfirmed(ctx context.Context, processId string) (model.ProcessTask, error) {
	tx := s.Db.ExtractTx(ctx)
	q := `UPDATE "task_process" SET inspectors_confirmed=inspectors_confirmed+1 WHERE id=$1 RETURNING *;`

	var process model.ProcessTask
	err := tx.GetContext(ctx, &process, q, processId)
	if err != nil {
		return model.ProcessTask{}, err
	}
	return process, nil
}

func (s *TaskProcessStorage) SearchProcesses(ctx context.Context, userId string,
	limit uint64, offset uint64, onlyActive bool, onlyForUser bool, order string, searchTerm string,
	sort string) ([]model.ProcessTask, int, error) {
	tx := s.Db.ExtractTx(ctx)
	var processes []model.ProcessTask
	var args []interface{}
	var query string
	var queryCnt string
	var count int

	queryCnt = `SELECT count(id) FROM task_process WHERE LOWER(status) LIKE CONCAT('%%',$1::text,'%%')`
	query = `SELECT * FROM "task_process" WHERE LOWER(status) LIKE CONCAT('%%',$1::text,'%%')`

	if onlyForUser {
		queryCnt += " AND exec_userid=$2"
		query += " AND exec_userid=$2"
		args = append(args, searchTerm, userId)
	}
	if onlyActive {
		queryCnt += " AND completed=$3"
		query += " AND completed=$3"
		args = append(args, !onlyActive)
	}

	filter := fmt.Sprintf(` ORDER BY %s %s LIMIT %d OFFSET %d;`, order, sort, limit, offset)
	q := query + filter

	if err := tx.SelectContext(ctx, &processes, q, args...); err != nil {
		return []model.ProcessTask{}, 0, err
	}
	tx.GetContext(ctx, &count, queryCnt, args...)
	return processes, count, nil
}
