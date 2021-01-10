package main

import (
	"reflect"
	"testing"
)

func Test_mergeInBetween(t *testing.T) {
	type args struct {
		x *ListNode
		a int
		b int
		y *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			args: args{
				x: newList(0, 1, 2, 3, 4, 5, 6),
				y: newList(100, 101, 102),
				a: 0,
				b: 6,
			},
			want: newList(100, 101, 102),
		},

		{
			args: args{
				x: newList(0, 1, 2, 3, 4, 5, 6),
				y: newList(100, 101, 102),
				a: 0,
				b: 3,
			},
			want: newList(100, 101, 102, 4, 5, 6),
		},

		{
			args: args{
				x: newList(0, 1, 2, 3, 4, 5, 6),
				y: newList(100, 101, 102),
				a: 1,
				b: 3,
			},
			want: newList(0, 100, 101, 102, 4, 5, 6),
		},

		{
			args: args{
				x: newList(0, 1, 2, 3, 4, 5, 6),
				y: newList(100, 101, 102),
				a: 3,
				b: 4,
			},
			want: newList(0, 1, 2, 100, 101, 102, 5, 6),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeInBetween(tt.args.x, tt.args.a, tt.args.b, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeInBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}
