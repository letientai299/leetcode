package main

import "testing"

func Test_findReplaceString(t *testing.T) {
	type args struct {
		s       string
		indices []int
		sources []string
		targets []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				s:       "vmokgggqzp",
				indices: []int{3, 5, 1},
				sources: []string{"kg", "ggq", "mo"},
				targets: []string{"s", "so", "bfr"},
			},
			want: "vbfrssozp",
		},

		{
			args: args{
				s:       "abcd",
				indices: []int{0, 2},
				sources: []string{"ab", "ec"},
				targets: []string{"eee", "ffff"},
			},
			want: "eeecd",
		},

		{
			args: args{
				s:       "xxabcdyy",
				indices: []int{2, 4},
				sources: []string{"a", "cd"},
				targets: []string{"eee", "ffff"},
			},
			want: "xxeeebffffyy",
		},

		{
			args: args{
				s:       "abcd",
				indices: []int{0, 2},
				sources: []string{"a", "cd"},
				targets: []string{"eee", "ffff"},
			},
			want: "eeebffff",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findReplaceString(tt.args.s, tt.args.indices, tt.args.sources, tt.args.targets); got != tt.want {
				t.Errorf("findReplaceString() = %v, want %v", got, tt.want)
			}
		})
	}
}
