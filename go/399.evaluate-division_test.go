package main

import (
	"reflect"
	"testing"
)

func Test_calcEquation(t *testing.T) {
	type args struct {
		equations [][]string
		values    []float64
		queries   [][]string
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			args: args{
				equations: [][]string{{"a", "b"}, {"c", "d"}},
				values:    []float64{1.0, 1.0},
				queries:   [][]string{{"a", "c"}, {"b", "d"}, {"b", "a"}, {"d", "c"}},
			},
			want: []float64{-1, -1, 1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcEquation(tt.args.equations, tt.args.values, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calcEquation() = %v, want %v", got, tt.want)
			}
		})
	}
}
