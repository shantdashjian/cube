package worker

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/golang-collections/collections/queue"

	"cube/task"
)

type Worker struct {
	Name		string
	Queue 		queue.Queue
	Db 			map[uuid.UUID]*task.Task
	TaskCount	int
}

func (w *Worker) CollectStats() {
	fmt.Println("I will collect stats")
}

func (w *Worker) RunTask() {
	fmt.Println("I will start or stop a task")
}

func (w *Worker) StartTask() {
	fmt.Println("I will start a task")
}

func (w *Worker) StopTask() {
	fmt.Println("I will stop a task")
}
