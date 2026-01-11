package homework12

import (
	"fmt"
)

type Task struct {
	identifier int
	priority   int
}

func (t *Task) GetIdentifier() int {
	return t.identifier
}

func (t *Task) GetPriority() int {
	return t.priority
}

func (t *Task) SetIdentifier(val int) {
	t.identifier = val
}

func (t *Task) SetPriority(val int) {
	t.priority = val
}

type Scheduler struct {
	BinaryHeap
}

func NewScheduler() Scheduler {
	return Scheduler{
		BinaryHeap: NewBinaryHeap(),
	}
}

func (s *Scheduler) AddTask(task *Task) {
	s.Add(task)
}

func (s *Scheduler) ChangeTaskPriority(taskID int, newPriority int) {
	task, err := s.Get(taskID)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	newTask := task.(*Task)
	newTask.SetPriority(newPriority)
	s.Add(newTask)
}

func (s *Scheduler) GetTask() *Task {
	task, err := s.GetMax()
	if err != nil {
		fmt.Println(err.Error())
		return &Task{}
	}
	return task.(*Task)
}
