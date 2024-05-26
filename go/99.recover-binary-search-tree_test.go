package main

import "testing"

func Test_recoverTree(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
	}{
		{root: treeFromLevelOrder(1, 3, NA, NA, 2)},
		{root: treeFromLevelOrder(3, 1, 4, NA, NA, 2)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recoverTree(tt.root)
			if !isValidBST(tt.root) {
				t.Errorf("not valid bst")
				tt.root.print(func(s string) { t.Log(s) })
			}
		})
	}
}
