package model

import "time"

type TaskUsers struct {
	ID        string    `db:"id" json:"ID,omitempty"`
	UpdatedAt time.Time `db:"updatedat" json:"UpdatedAt,omitempty"`
	ProjectId string    `db:"projectid" json:"ProjectId,omitempty"`
	NodeId    string    `db:"nodeid" json:"NodeId,omitempty"`
	UserId    string    `db:"userid" json:"UserId,omitempty"`
	Energy    uint64    `db:"energy" json:"Energy,omitempty"`
	Confirmed bool      `db:"confirmed" json:"Confirmed"`
}

type ProcessTask struct {
	ID                  string    `db:"id" json:"ID,omitempty"`
	CreatedAt           time.Time `db:"createdat" json:"CreatedAt,omitempty"`
	UpdatedAt           time.Time `db:"updatedat" json:"UpdatedAt,omitempty"`
	ProjectId           string    `db:"projectid" json:"ProjectId,omitempty"`
	NodeId              string    `db:"nodeid" json:"NodeId,omitempty"`
	ExecUserId          string    `db:"exec_userid" json:"ExecUserId,omitempty"`
	InspectorTotal      uint64    `db:"inspectors_total" json:"InspectorTotal,omitempty"`
	InspectorsConfirmed uint64    `db:"inspectors_confirmed" json:"InspectorsConfirmed,omitempty"`
	EnergyTotal         uint64    `db:"energy_total" json:"EnergyTotal,omitempty"`
	Completed           bool      `db:"complited" json:"Complited"`
}
