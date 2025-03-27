package manager

import (
	"cube/task"
	"fmt"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

type Manager struct {
	Pending			queue.Queue
	TaskDb 			map[uuid.UUID]*task.Task
	EventDb 		map[uuid.UUID]*task.TaskEvent
	Workers			[]string
	WorkerTaskMap	map[string][]uuid.UUID
	TasWorkerkMap	map[uuid.UUID]string
}

func (m *Manager) SelectWorker() {
	fmt.Println("I will select an appropriate worker")
}

func (m *Manager) UpdateTasks() {
	fmt.Println("I will update tasks")
}

func (m *Manager) SendWork() {
	fmt.Println("I will send work to workers")
}