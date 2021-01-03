package main

import (
	"reflect"
	"testing"
)

func Test_exclusiveTime(t *testing.T) {
	tests := []struct {
		name string
		n    int
		logs []string
		want []int
	}{
		{
			n: 3,
			logs: []string{
				"0:start:0",
				"1:start:1",
				"1:end:2",
				"2:start:2",
				"2:end:4",
				"0:end:10",
			},
			want: []int{6, 2, 3},
		},
		/*
			"0:start:0",
				"1:start:5",
					"2:start:6",
						"3:start:9",
							"4:start:11",
								"5:start:12",
									"6:start:14",
										"7:start:15",
											"1:start:24",
											"1:end:29",
										"7:end:34",
									"6:end:37",
								"5:end:39",
							"4:end:40",
						"3:end:45",
						"0:start:49",
						"0:end:54",
						"5:start:55",
						"5:end:59",
						"4:start:63",
						"4:end:66",
						"2:start:69",
						"2:end:70",
						"2:start:74",
							"6:start:78",
								"0:start:79",
								"0:end:80",
							"6:end:85",
							"1:start:89",
							"1:end:93",
						"2:end:96",
					"2:end:100",
				"1:end:102",
				"2:start:105",
				"2:end:109",
			"0:end:114",
		*/
		{
			n: 8,
			logs: []string{"0:start:0", "1:start:5", "2:start:6", "3:start:9",
				"4:start:11", "5:start:12", "6:start:14", "7:start:15", "1:start:24",
				"1:end:29", "7:end:34", "6:end:37", "5:end:39", "4:end:40", "3:end:45",
				"0:start:49", "0:end:54", "5:start:55", "5:end:59", "4:start:63",
				"4:end:66", "2:start:69", "2:end:70", "2:start:74", "6:start:78",
				"0:start:79", "0:end:80", "6:end:85", "1:start:89", "1:end:93",
				"2:end:96", "2:end:100", "1:end:102", "2:start:105", "2:end:109",
				"0:end:114",
			},
			want: []int{20, 14, 35, 7, 6, 9, 10, 14},
		},

		{
			n:    2,
			logs: []string{"0:start:0", "0:start:2", "0:end:5", "1:start:7", "1:end:7", "0:end:8"},
			want: []int{8, 1},
		},

		{
			n:    1,
			logs: []string{"0:start:0", "0:end:7", "0:start:8", "0:end:8"},
			want: []int{9},
		},

		{
			n:    2,
			logs: []string{"0:start:0", "1:start:7", "1:end:7", "0:end:8"},
			want: []int{8, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := exclusiveTime(tt.n, tt.logs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("exclusiveTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
