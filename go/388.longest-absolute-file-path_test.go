package main

import "testing"

func Test_lengthLongestPath(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			input: "a",
			want:  0,
		},

		{
			input: "file1.txt\nfile2.txt\nlongfile.txt",
			want:  12,
		},

		{
			input: "dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext",
			want:  32,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthLongestPath(tt.input); got != tt.want {
				t.Errorf("lengthLongestPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
