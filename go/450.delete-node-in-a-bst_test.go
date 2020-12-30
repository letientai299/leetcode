package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_deleteNode(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		key  int
	}{
		{
			root: treeFromLevelOrder(50, 30, 70, NA, 40, 60, 80),
			key:  50,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := deleteNode(tt.root, tt.key)
			if !assert.True(t, isValidBST(got)) {
				t.Errorf("deleteNode() = %v is invalid BST", got)
			}
		})
	}
}
