package main

import "time"

type task struct {
	title, description string
	done               bool
	category           string
	priority           int
	due                time.Time
	completion         float32
}

type taskList struct {
	tasks []*task
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
