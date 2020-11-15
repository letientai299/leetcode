package main

import "testing"

func Test_lemonadeChange(t *testing.T) {
	tests := []struct {
		bills []int
		want  bool
	}{
		{
			bills: []int{5, 5, 5, 5, 20, 20, 5, 5, 5},
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := lemonadeChange(tt.bills); got != tt.want {
				t.Errorf("lemonadeChange() = %v, want %v", got, tt.want)
			}
		})
	}
}
