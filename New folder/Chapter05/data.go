package main

import "time"

const (
	dateFormat = "02 Jan 06 15:04"

	lowPriority  = 0
	midPriority  = 1
	highPriority = 2
)

func dateValidator(text string) error {
	_, err := time.Parse(dateFormat, text)
	return err
}

type task struct {
	title, description string
	done               bool
	category           string
	priority           int
	due                *time.Time
	completion         float64
}

type taskList struct {
	tasks []*task
}

func (l *taskList) add(t *task) {
	l.tasks = append([]*task{t}, l.tasks...)
}

func (l *taskList) remaining() []*task {
	var items []*task
	for _, task := range l.tasks {
		if !task.done {
			items = append(items, task)
		}
	}

	return items
}

func (l *taskList) done() []*task {
	var items []*task
	for _, task := range l.tasks {
		if task.done {
			items = append(items, task)
		}
	}

	return items
}

func dummyData() *taskList {
	return &taskList{
		tasks: []*task{
			{title: "Nearly done",
				description: `You can tick my checkbox
and I will be marked as
done and disappear`},
			{title: "Functions",
				description: `Tap the plus icon above to add
a new task, or tap the minus
icon to remove this one`},
		}}
}
