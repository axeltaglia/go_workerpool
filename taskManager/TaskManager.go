package taskManager

import "fmt"

// Task interface with an ID and a process function
type Task interface {
	Process()      // The function that does the actual task work
	GetID() string // A method to get the task ID
}

// TaskManager manages a channel of tasks
type TaskManager struct {
	taskQueue chan Task // Buffered channel for tasks
}

// NewTaskManager initializes the TaskManager with a buffer size
func NewTaskManager(bufferSize int) *TaskManager {
	return &TaskManager{
		taskQueue: make(chan Task, bufferSize),
	}
}

// AddTask adds a new task to the task queue
func (tm *TaskManager) AddTask(task Task) {
	tm.taskQueue <- task
	fmt.Printf("Task with ID: %s added to the queue\n", task.GetID())
}

// Start runs a worker that processes tasks from the task queue
func (tm *TaskManager) Start() {
	go func() {
		for task := range tm.taskQueue {
			task := task
			go func() {
				task.Process() // Process each task
			}()
		}
	}()
}
