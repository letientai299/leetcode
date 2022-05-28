package main

import "testing"

func Test_countPairs1530(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		distance int
		want     int
	}{
		{
			root: treeFromLevelOrder(
				61, 46, 16, 61, 49, 93, 12, 92, 44, 4, 69, 56, 25, 92, 6, 10, 39, 59, 93, 39, 5, 72, 12, 16, 31, 5, 54, 75, 62, NA, NA, 31, 79, 10, NA, 81, 93, NA, NA, 2, 47, 100, 65, 68, NA, 97, NA, 60, 87, 80, NA, 36, NA, NA, NA, 21, 11, NA, 55, 98, 50, 1, 51, 87, NA, 51, 73, 68, NA, 15, 38, 10, NA, 100, 61, 52, NA, NA, NA, NA, NA, 52, 6, 29, NA, 63, NA, 37, 11, NA, 56, NA, NA, NA, NA, 95, 83, NA, NA, 37, 29, NA, NA, NA, NA, 68, 45, 91, NA, 24, NA, 36, NA, 28, NA, NA, NA, NA, NA, 84, NA, NA, NA, 60, 21, NA, NA, 64, 54, 38, NA, 45, NA, NA, NA, NA, NA, 92, NA, 51, 45, 100, NA, NA, NA, 81, NA, NA, 37, 34, NA, NA, 97, 74, NA, NA, NA, NA, NA, 6, NA, 9, NA, NA, 50, NA, NA, NA, 50, NA, 36, 80, 7, NA, NA, NA, NA, NA, NA, NA, NA, NA, NA, NA, NA, NA, NA, NA, NA, NA, 1, NA, NA, NA, NA, NA, NA, NA, NA, 32, 71, NA, NA, NA, NA, 37, 7, NA, 28,
			),
			distance: 5,
			want:     38,
		},

		{
			root: treeFromLevelOrder(
				35, 96, 44, NA, 57, 22, 73, 2, 24, NA, 88, 50, 70, NA, NA, NA, NA, 89, 90, 92, 91, 96, 40, 16, 34, NA, 57, NA, 36, NA, 65, 87, 52, 100, 57, NA, NA, NA, NA, 39, NA, NA, NA, NA, 12, NA, NA, NA, NA, NA, 8, 57, 94, NA, NA, 71, 37, NA, 56, NA, 61, 7, 23, NA, NA, NA, NA, NA, NA, NA, 83, 95, 9, 66, 71, 92, 22, NA, NA, NA, 97, NA, 71, 89, 18, NA, NA, NA, 36, NA, 50, NA, 47, NA, 87, 97, 100, NA, NA, NA, 3, NA, 56, NA, 60, NA, NA, 15, 41, NA, 23, NA, NA, 47, 67, NA, NA, 85, 64, NA, NA, NA, NA, NA, NA, NA, 32, 28, 71, NA, 65, NA, NA, NA, 54, NA, 61, 10, 60, NA, NA, NA, NA, NA, 13, 2, 66, NA, NA, NA, 40,
			),
			distance: 5,
			want:     15,
		},

		{
			root: treeFromLevelOrder(
				15, 66, 55, 97, 60, 12, 56, NA, 54, NA, 49, NA, 9, NA, NA, NA, NA, NA, 90,
			),
			distance: 5,
			want:     3,
		},

		{
			root:     treeFromLevelOrder(1, 2, 3, 4, 5, 6, 7),
			distance: 3,
			want:     2,
		},

		{
			root:     treeFromLevelOrder(1, 1, 1),
			distance: 2,
			want:     1,
		},

		{
			root:     treeFromLevelOrder(1, 2, 3, NA, 4),
			distance: 3,
			want:     1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPairs1530(tt.root, tt.distance); got != tt.want {
				t.Errorf("countPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
