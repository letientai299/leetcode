package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minNumberOfHours(t *testing.T) {
	type args struct {
		initialEnergy     int
		initialExperience int
		energy            []int
		experience        []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				initialEnergy:     1,
				initialExperience: 1,
				energy:            []int{1, 1, 1, 1},  // 4
				experience:        []int{1, 1, 1, 50}, // 1, 46
			},
			want: 51,
		},

		{
			args: args{
				initialEnergy:     5,
				initialExperience: 3,
				energy:            []int{1, 4},
				experience:        []int{2, 5},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, minNumberOfHours(tt.args.initialEnergy, tt.args.initialExperience, tt.args.energy, tt.args.experience), "minNumberOfHours(%v, %v, %v, %v)", tt.args.initialEnergy, tt.args.initialExperience, tt.args.energy, tt.args.experience)
		})
	}
}
