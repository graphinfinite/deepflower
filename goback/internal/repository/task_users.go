package repository

import (
	"context"
	"deepflower/internal/model"
	"deepflower/pkg/postgres"
	"time"
)

type TaskUsersStorage struct {
	Db *postgres.PG
}

func NewTaskUsersStorage(db *postgres.PG) *TaskUsersStorage {
	return &TaskUsersStorage{Db: db}
}

func (s *TaskUsersStorage) AddTaskUser(ctx context.Context, userId, projectId, nodeId string, energy uint64) error {
	tx := s.Db.ExtractTx(ctx)

	query4 := `INSERT INTO "task_users" (projectid, nodeid, userid, updatedAt, energy, confirmed)
	VALUES($1,$2,$3,$4,$5,$6) 
	ON CONFLICT (nodeid, userid)
	DO 
	   UPDATE SET energy="task_users".energy + $7, updatedat=$8;`

	_, err := tx.ExecContext(ctx, query4, projectId, nodeId, userId, time.Now(), energy, false, energy, time.Now())
	if err != nil {
		return err
	}
	return nil
}

func (s *TaskUsersStorage) GetTaskUsersByTaskId(ctx context.Context, projectId, nodeId string) ([]model.User, error) {
	tx := s.Db.ExtractTx(ctx)
	query := `SELECT * FROM "users" WHERE id IN (SELECT userid FROM "task_users" WHERE projectid=$1 AND nodeid=$2);`
	var users []model.User
	if err := tx.SelectContext(ctx, &users, query, projectId, nodeId); err != nil {
		return []model.User{}, err
	}
	return users, nil
}
