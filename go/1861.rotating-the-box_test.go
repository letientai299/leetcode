package main

import (
	"reflect"
	"testing"
)

func Test_rotateTheBox(t *testing.T) {
	tests := []struct {
		name string
		box  []string
		want []string
	}{
		{
			box: []string{
				"##*.*.",
				"###*..",
				"###.#.",
			},
			want: []string{
				".##",
				".##",
				"##*",
				"#*.",
				"#.*",
				"#..",
			},
		},

		{
			box:  []string{"#.#"},
			want: []string{".", "#", "#"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var box [][]byte
			for _, b := range tt.box {
				box = append(box, []byte(b))
			}

			var want [][]byte
			for _, b := range tt.want {
				want = append(want, []byte(b))
			}

			if got := rotateTheBox(box); !reflect.DeepEqual(got, want) {
				t.Errorf("rotateTheBox() = %v, want %v", got, want)
				t.Logf("got: ")
				for _, bs := range got {
					t.Logf(string(bs))
				}

				t.Logf("want: ")
				for _, bs := range want {
					t.Logf(string(bs))
				}
			}
		})
	}
}
