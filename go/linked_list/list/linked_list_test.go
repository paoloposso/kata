package list_test

import (
	"linkedlist/list"
	"testing"
)

func TestAppend(t *testing.T) {
	linkedList := &list.LinkedList[string]{}
	linkedList.Append("5")
	v, err := linkedList.Get(0)
	if err != nil {
		t.Error(err)
	}
	if *v != "5" {
		t.Fail()
	}
	linkedList.Append("10")
	v, err = linkedList.Get(0)
	if err != nil {
		t.Error(err)
	}
	if *v != "5" {
		t.Fail()
	}
}
