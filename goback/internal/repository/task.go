package repository

import (
	"context"
	"deepflower/internal/model"
	"deepflower/pkg/postgres"
	"encoding/json"
	"fmt"
)

type TaskStorage struct {
	Db *postgres.PG
}

func NewTaskStorage(db *postgres.PG) *TaskStorage {
	return &TaskStorage{Db: db}
}

func (s *TaskStorage) GetTaskData(ctx context.Context, projectId, nodeId string) (model.CellData, error) {
	tx := s.Db.ExtractTx(ctx)
	query2 := `SELECT graph FROM project WHERE id=$1;`
	var graphStr string
	err := tx.GetContext(ctx, &graphStr, query2, projectId)
	if err != nil {
		return model.CellData{}, err
	}
	var graph model.Graph
	err = json.Unmarshal([]byte(graphStr), &graph)
	if err != nil {
		return model.CellData{}, err
	}
	for _, c := range graph.Cells {
		if c.Shape == "slow-model" || c.Shape == "fast-model" {
			if c.Id == nodeId {
				return c.Data, nil
			}
		}
	}

	return model.CellData{}, fmt.Errorf("task not found")
}

func (s *TaskStorage) AddEnergyToTask(ctx context.Context, projectId, nodeId string, energy uint64) error {
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

func (s *TaskStorage) SubtractEnergy(ctx context.Context, projectId, nodeId string, energy uint64) error {
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
				graph.Cells[n].Data.Energy = c.Data.Energy - energy
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

// CODE PROTOTYPE !!!!!!!!!!!!!!
// / TODO :POSTGRES-> JSONB OR DECOMPOSE NODES DATA OR OTHER DB ?????
func (s *TaskStorage) UpdateTaskStatus(ctx context.Context, projectId, nodeId, newStatus string) error {
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
		return err
	}
	var query3 = `UPDATE project SET graph=$1 WHERE id=$2;`
	_, err = tx.ExecContext(ctx, query3, string(updatedGraphByte), projectId)
	if err != nil {
		return err
	}
	return nil
}
