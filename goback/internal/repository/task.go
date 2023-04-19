package repository

import (
	"context"
	"deepflower/internal/model"
	"deepflower/pkg/postgres"
	"encoding/json"
	"fmt"
	"time"
)

type TaskStorage struct {
	Db *postgres.PG
}

func NewTaskStorage(db *postgres.PG) *TaskStorage {
	return &TaskStorage{Db: db}
}

func (s *TaskStorage) AddEnergyToTask(ctx context.Context, userId, projectId, nodeId string, energy uint64) error {
	tx := s.Db.ExtractTx(ctx)
	query2 := `SELECT graph FROM project WHERE id=$1;`
	var graphStr string
	err := tx.GetContext(ctx, &graphStr, query2, projectId)
	if err != nil {
		return err
	}
	var graph model.Graph
	err = json.Unmarshal([]byte(graphStr), &graph)
	if err != nil {
		return err
	}
	for n, c := range graph.Cells {
		if c.Shape == "slow-model" || c.Shape == "fast-model" {
			if c.Id == nodeId {
				graph.Cells[n].Data.Energy = c.Data.Energy + energy
			}
		}
	}
	updatedGraphByte, err := json.Marshal(graph)
	if err != nil {
		return err
	}

	var query3 = `UPDATE project SET graph=$1 WHERE id=$2;`
	_, err = tx.ExecContext(ctx, query3, string(updatedGraphByte), projectId)
	if err != nil {
		return err
	}
	return nil
}

func (s *TaskStorage) GetTaskConsensusProcessById(ctx context.Context, processId string) (model.ProcessTask, error) {
	tx := s.Db.ExtractTx(ctx)
	query := `SELECT * FROM task_process WHERE id=$1;`
	var process model.ProcessTask
	if err := tx.GetContext(ctx, &process, query, processId); err != nil {
		return model.ProcessTask{}, err
	}
	return process, nil
}

// CODE PROTOTYPE !!!!!!!!!!!!!!
// / TODO :POSTGRES-> JSONB OR DECOMPOSE NODES DATA OR OTHER DB ?????
func (s *TaskStorage) UpdateTaskStatus(ctx context.Context, projectId, nodeId, userId, newStatus string) (processId string, err error) {
	tx := s.Db.ExtractTx(ctx)
	query2 := `SELECT graph FROM project WHERE id=$1;`
	var graphStr string
	err = tx.GetContext(ctx, &graphStr, query2, projectId)
	if err != nil {
		return "", err
	}
	var graph model.Graph
	err = json.Unmarshal([]byte(graphStr), &graph)
	if err != nil {
		return "", err
	}
	var celldata model.CellData
	for n, c := range graph.Cells {
		if c.Id == nodeId {
			graph.Cells[n].Data.Status = newStatus
			celldata.LeadTime = c.Data.LeadTime
			celldata.Energy = c.Data.Energy
		}

	}
	updatedGraphByte, err := json.Marshal(graph)
	if err != nil {
		return "", err
	}
	var query3 = `UPDATE project SET graph=$1 WHERE id=$2;`
	_, err = tx.ExecContext(ctx, query3, string(updatedGraphByte), projectId)
	if err != nil {
		return "", err
	}

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
	fmt.Println(projectId, nodeId, projectId, nodeId, userId, 0, celldata.Energy, celldata.LeadTime, newStatus, newStatus, time.Now())

	err = tx.GetContext(ctx, &insId, query4, projectId, nodeId, userId, projectId, nodeId, 0, celldata.Energy, celldata.LeadTime, newStatus, newStatus, time.Now())
	if err != nil {
		return "", err
	}
	return insId, nil
}
