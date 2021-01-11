package main

import "testing"

func Test_isValidSerialization(t *testing.T) {
	tests := []struct {
		preorder string
		want     bool
	}{
		{
			preorder: "9,3,4,#,#,1,#,#,2,#,6,#",
			want:     false,
		},

		{
			preorder: "9,3,4,#,#,1,#,#,2,#,6,#,#,#",
			want:     false,
		},

		{
			preorder: "9,3,4,#,#,1,#,#,2,#,6,#,#",
			want:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.preorder, func(t *testing.T) {
			if got := isValidSerialization(tt.preorder); got != tt.want {
				t.Errorf("isValidSerialization() = %v, want %v", got, tt.want)
			}
		})
	}
}
