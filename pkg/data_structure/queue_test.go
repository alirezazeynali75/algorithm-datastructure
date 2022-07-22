package datastructure

import (
	"testing"
)

func TestEnqueueInQueue(t *testing.T) {
	q := CreateNewQueue()
	q.Enqueue("hello")
	newLength := q.GetLength()
	if newLength != 1 {
		t.Errorf("enqueue process failed")
	}
}

func TestDequeue(t *testing.T) {
	q := CreateNewQueue()
	q.Enqueue("hello")
	got, status := q.Dequeue()
	if !status {
		t.Errorf("dequeue false")
	}
	if got != "hello" {
		t.Errorf("some problem on dequeue")
	}
}