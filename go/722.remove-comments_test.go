package main

import (
	"reflect"
	"testing"
)

func Test_removeComments(t *testing.T) {
	tests := []struct {
		name   string
		source []string
		want   []string
	}{
		{
			source: []string{
				"a/*/b//*c",
				"blank",
				"d/*/e*//f",
			},
			want: []string{
				"ae*",
			},
		},


		{
			source: []string{
				"a//*b//*c",
				"blank",
				"d/*/e*//f",
			},
			want: []string{
				"a",
				"blank",
				"d/f",
			},
		},

		{
			source: []string{
				"void func(int k) {",
				"// this function does nothing /*",
				"   k = k*2/4;",
				"   k = k/2;*/",
				"}",
			},
			want: []string{
				"void func(int k) {",
				"   k = k*2/4;",
				"   k = k/2;*/",
				"}",
			},
		},

		{
			source: []string{"a/*comment", "line", "more_comment*/b"},
			want:   []string{"ab"},
		},
		{
			source: []string{
				"/*Test program */",
				"int main()",
				"{ ",
				"  // variable declaration ",
				"int a, b, c;",
				"/* This is a test",
				"   multiline  ",
				"   comment for ",
				"   testing */",
				"a = b + c;", "}",
			},
			want: []string{
				"int main()",
				"{ ",
				"  ", "int a, b, c;",
				"a = b + c;",
				"}",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeComments(tt.source); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeComments() = %v, want %v", got, tt.want)
				for _, s := range got {
					t.Log(s)
				}
			}
		})
	}
}
