package worker

import (
	"fmt"
	"orchestrator/task"
	"sync"

	"github.com/google/uuid"
)

//Worker’s requirements are:
//1. Run tasks as Docker containers.
//2. Accept tasks to run from a manager.
//3. Provide relevant statistics to the manager for the purpose of scheduling
//tasks.
//4. Keep track of its tasks and their state.
////////////////////////////////////////////////////////

// WORKER
// Name is the name of the worker.
// Queue holds tasks in a FIFO order.
// Db is a map that associates task UUIDs with their corresponding tasks.
// TaskCount tracks the number of tasks the worker is currently managing.
// mu is a mutex to ensure thread-safe operations on the worker's fields.
type Worker struct {
	Name      string
	Queue     uuid.UUIDs
	Db        map[uuid.UUID]task.Task
	TaskCount int
	Mu        sync.Mutex
}

// some basic utils
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
