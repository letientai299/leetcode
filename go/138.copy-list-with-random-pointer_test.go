package main

import (
	"reflect"
	"testing"
)

func Test_copyRandomList(t *testing.T) {
	tests := []struct {
		name string
		head *Node
		want *Node
	}{
		{head: newNodes(1, 2, 3), want: newNodes(1, 2, 3)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := copyRandomList(tt.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("copyRandomList() = %v, want %v", got, tt.want)
			}
		})
	}
}
