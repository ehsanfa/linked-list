package linked_list

import (
	"testing"
)

func TestLinkedList(t *testing.T) {
	ll := NewLinkedList()
	ll.Append("node1")
	ll.Append("node2")
	ll.Append("node3")
	ll.Append("node4")
	ll.Append("tail")
	ll.Append(struct {Name string}{"mammad"})
	if ll.Count() != 6 {
		t.Error("miscount")
	}
	if !ll.Contains("node3") {
		t.Error("contains doesn't work as expected")
	}
	ll.RemoveValue("node3")
	if ll.Contains("node3") {
		t.Error("remove doesn't work as expected")
	}
	if ll.Count() != 5 {
		t.Error("miscount")
	}
}