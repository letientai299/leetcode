package main

import "testing"

func Test_flipEquiv(t *testing.T) {
	tests := []struct {
		name string
		a    *TreeNode
		b    *TreeNode
		want bool
	}{
		{
			a:    treeFromLevelOrder(0, 1, 2, 5, NA, 3, 6, NA, NA, NA, 4),
			b:    treeFromLevelOrder(0, 1, 2, NA, 5, 3, 6, NA, NA, 4),
			want: true,
		},

		{
			a:    treeFromLevelOrder(1, 2, 3, 4, 5, 6, NA, NA, NA, 7, 8),
			b:    treeFromLevelOrder(1, 3, 2, NA, 6, 4, 5, NA, NA, NA, NA, 8, 7),
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flipEquiv(tt.a, tt.b); got != tt.want {
				t.Errorf("flipEquiv() = %v, want %v", got, tt.want)
				t.Log("a:")
				tt.a.print()
				t.Log("b:")
				tt.b.print()
			}
		})
	}
}
