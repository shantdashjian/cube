package task

import (
	"github.com/google/uuid"
	"github.com/docker/go-connections/nat"
)


type State int

const (
	Pending State = iota
	Scheduled
	Running
	Completed
	Failed
)

type Task struct (
	ID				uuid.UUID
	Name			string
	State			State
	Image			string
	Memory			int
	Disk			int
	ExposedPorts	nat.PortSet
	PortBindings	map[string]string
	RestartPolicy	string	
	StartTime		time.Time
	FinishTime		time.Time
)

type Task struct (
	ID			uuid.UUID
	State		State
	Timestamp	time.Time
	Task		Task
)