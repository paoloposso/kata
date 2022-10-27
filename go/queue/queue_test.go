package queue_test

import (
	"queue"
	"testing"
)

func TestPush(t *testing.T) {
	q := queue.NewQueue[string]()
	q.Push("aaaa")
	q.Push("bbbb")
	q.Push("cccc")
	q.Push("dddd")
	q.Push("eeee")

	v := q.Pop()
	if v.Data != "aaaa" {
		t.Fail()
	}

	v = q.Pop()
	if v.Data != "bbbb" {
		t.Fail()
	}
}
