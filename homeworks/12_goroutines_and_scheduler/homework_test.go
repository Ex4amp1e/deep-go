package homework12

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrace(t *testing.T) {
	task1 := &Task{identifier: 1, priority: 10}
	task2 := &Task{identifier: 2, priority: 20}
	task3 := &Task{identifier: 3, priority: 30}
	task4 := &Task{identifier: 4, priority: 40}
	task5 := &Task{identifier: 5, priority: 50}
	taskNil := &Task{}

	scheduler := NewScheduler()

	task := scheduler.GetTask()
	assert.Equal(t, taskNil, task)

	scheduler.AddTask(task1)
	scheduler.AddTask(task2)
	scheduler.AddTask(task3)
	scheduler.AddTask(task4)
	scheduler.AddTask(task5)

	task = scheduler.GetTask()
	assert.Equal(t, task5, task)

	task = scheduler.GetTask()
	assert.Equal(t, task4, task)

	scheduler.ChangeTaskPriority(1, 100)

	task = scheduler.GetTask()
	assert.Equal(t, task1, task)

	task = scheduler.GetTask()

	assert.Equal(t, task3, task)
}
