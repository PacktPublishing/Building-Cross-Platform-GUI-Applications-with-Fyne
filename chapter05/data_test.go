package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTaskList_Remaining(t *testing.T) {
	item := &task{title: "Remain"}
	list := &taskList{tasks: []*task{item}}

	remain := list.remaining()
	assert.Equal(t, 1, len(remain))
	done := list.done()
	assert.Equal(t, 0, len(done))
}

func TestTaskList_Done(t *testing.T) {
	item := &task{title: "Done", done: true}
	list := &taskList{tasks: []*task{item}}

	remain := list.remaining()
	assert.Equal(t, 0, len(remain))
	done := list.done()
	assert.Equal(t, 1, len(done))
}
