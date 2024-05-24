package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countDaysTogether(t *testing.T) {
	type args struct {
		arriveAlice string
		leaveAlice  string
		arriveBob   string
		leaveBob    string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				arriveAlice: "03-05",
				leaveAlice:  "07-14",
				arriveBob:   "04-14",
				leaveBob:    "09-21",
			},
			want: 92,
		},

		{
			args: args{
				arriveAlice: "09-01",
				leaveAlice:  "10-19",
				arriveBob:   "06-19",
				leaveBob:    "10-20",
			},
			want: 49,
		},

		{
			args: args{
				arriveAlice: "08-06",
				leaveAlice:  "12-08",
				arriveBob:   "02-04",
				leaveBob:    "09-01",
			},
			want: 27,
		},

		{
			args: args{
				arriveAlice: "08-15",
				leaveAlice:  "08-18",
				arriveBob:   "08-16",
				leaveBob:    "08-19",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, countDaysTogether(tt.args.arriveAlice, tt.args.leaveAlice, tt.args.arriveBob, tt.args.leaveBob), "countDaysTogether(%v, %v, %v, %v)", tt.args.arriveAlice, tt.args.leaveAlice, tt.args.arriveBob, tt.args.leaveBob)
		})
	}
}
