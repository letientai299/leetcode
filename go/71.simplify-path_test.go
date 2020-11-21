package main

import "testing"

func Test_simplifyPath(t *testing.T) {
	tests := []struct {
		path string
		want string
	}{
		{
			path: "/a/./b/../../c/",
			want: "/c",
		},

		{
			path: "/a///b/../..",
			want: "/",
		},
		{
			path: "/a///b/../.",
			want: "/a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			if got := simplifyPath(tt.path); got != tt.want {
				t.Errorf("simplifyPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
