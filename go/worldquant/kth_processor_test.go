package worldquant

import (
	"fmt"
	"testing"
)

func TestKthProcessor(t *testing.T) {
	tests := []struct {
		name string
		m    int
		n    int
		k    int
		want int
	}{

		{m: 20, n: 5, k: 2, want: 5},
		{m: 11, n: 5, k: 1, want: 4},
		{m: 11, n: 5, k: 5, want: 4 /* 1 1 2 3 4 */},
		{m: 11, n: 5, k: 4, want: 3},
		{m: 5, n: 5, k: 5, want: 1},
		{m: 6, n: 5, k: 5, want: 2},
		{m: 7, n: 5, k: 5, want: 2},

		{m: 11, n: 5, k: 3, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KthProcessor(tt.m, tt.n, tt.k); got != tt.want {
				t.Errorf("KthProcessor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_a(t *testing.T) {
	for m := 5; m <= 30; m++ {
		fmt.Println("m=", m, "res=", KthProcessor(m, 5, 2))
		fmt.Println()
	}
}
