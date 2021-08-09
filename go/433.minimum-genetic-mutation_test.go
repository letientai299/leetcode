package main

import "testing"

func Test_minMutation(t *testing.T) {
	type args struct {
		start string
		end   string
		bank  []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				bank:  []string{"AACCGGTA","AACCGCTA","AAACGGTA"},
				start: "AACCGGTT",
				end:   "AAACGGTA",
			},
			want: 2,
		},


		{
			args: args{
				bank:  []string{"AACCGGTA"},
				start: "AACCGGTT",
				end:   "AACCGGTA",
			},
			want: 1,
		},

		{
			args: args{
				bank:  []string{"AAAACCCC", "AAACCCCC", "AACCCCCC"},
				start: "AAAAACCC",
				end:   "AACCCCCC",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minMutation(tt.args.start, tt.args.end, tt.args.bank); got != tt.want {
				t.Errorf("minMutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
