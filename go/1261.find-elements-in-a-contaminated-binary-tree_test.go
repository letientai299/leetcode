package main

import "testing"

func TestFindElements(t *testing.T) {
	tests := []struct {
		name    string
		root    *TreeNode
		targets []int
		want    []bool
	}{
		{
			root:    treeFromLevelOrder(-1, NA, -1),
			targets: []int{1, 2},
			want:    []bool{false, true},
		},

		{
			root:    treeFromLevelOrder(-1, NA, -1, -1, NA, -1),
			targets: []int{2, 3, 4, 5},
			want:    []bool{true, false, false, true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fe := &FindElements{root: tt.root}
			for i, v := range tt.targets {
				if got := fe.Find(v); got != tt.want[i] {
					t.Errorf("Find(%d) = %v, want %v", v, got, tt.want[i])
				}
			}
		})
	}
}
