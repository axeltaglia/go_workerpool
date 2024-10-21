package taskManager

import (
	"fmt"
	"time"
)

// ConcreteTask is a simple implementation of the Task interface
type ConcreteTask struct {
	ID       string
	Workload int // Simulates the amount of work (sleep duration)
}

func (t *ConcreteTask) Process() {
	for i := 0; i <= t.Workload; i++ {
		time.Sleep(1 * time.Second) // Simulate work
		progress := int(float64(i) / float64(t.Workload) * 100)
		// Print the progress on the same line for this task
		fmt.Printf("Task %s. Progress: %d%%\n", t.ID, progress)
	}
}

// GetID returns the ID of the task
func (t *ConcreteTask) GetID() string {
	return t.ID
}
