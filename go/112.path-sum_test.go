package main

import "testing"

func Test_hasPathSum(t *testing.T) {
	tests := []struct {
		root *TreeNode
		sum  int
		want bool
	}{
		{
			root: treeFromJson(`
{
  "val": 5,
  "left": {
    "val": 4,
    "left": {
      "val": 11,
      "left": { "val": 7 },
      "right": { "val": 2 }
    }
  },
  "right": {
    "val": 8,
    "left": { "val": 13 },
    "right": { "val": 4, "right": { "val": 1 } }
  }
}
`),
			sum:  22,
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := hasPathSum(tt.root, tt.sum); got != tt.want {
				t.Errorf("hasPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
