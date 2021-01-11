package main

import "testing"

type numIslandTestCase struct {
	name string
	grid []string
	want int
}

func getNumIslandTests() []numIslandTestCase {
	return []numIslandTestCase{
		{
			grid: []string{
				"10110010111101011110",
				"01001010111111011011",
				"10010101011011100110",
				"01100110111100100011",
				"11010010001010111011",
				"00001011001001011110",
				"10111101101101110010",
				"01100010010111001101",
				"00001101001101001010",
				"00111010101110111110",
				"10101110111010101011",
				"00111101110100011101",
				"11100000110111011110",
				"00111001001111110110",
				"00011000011010011111",
				"01110100111110111001",
				"00001111000010000110",
				"11111111110110111111",
				"01001001111110101111",
				"00111110001111110110",
			},
			want: 23,
		},

		{
			grid: []string{
				"111111111",
				"100000001",
				"101010101",
				"101111101",
				"101010101",
				"100000001",
				"111111111",
			},

			want: 2,
		},

		{
			grid: []string{
				"10111",
				"10101",
				"11101",
			},
			want: 1,
		},

		{
			grid: []string{
				"10001",
				"01000",
				"00101",
				"00010",
			},
			want: 6,
		},

		{
			grid: []string{
				"11000",
				"11000",
				"00100",
				"00011",
			},
			want: 3,
		},

		{
			grid: []string{
				"11000110000",
				"01101100000",
				"00111000000",
				"01000000000",
				"01100000000",
				"01100001000",
			},
			want: 3,
		},
	}
}

func Test_numIslands(t *testing.T) {
	tests := getNumIslandTests()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bs := make([][]byte, 0, len(tt.grid))
			for _, s := range tt.grid {
				bs = append(bs, []byte(s))
			}

			if got := numIslands(bs); got != tt.want {
				t.Errorf("numIslands() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_numIslands(b *testing.B) {
	tests := getNumIslandTests()

	bss := make([][][]byte, 0, len(tests))
	for _, ss := range tests {
		bs := make([][]byte, 0, len(ss.grid))
		for _, s := range ss.grid {
			bs = append(bs, []byte(s))
		}
		bss = append(bss, bs)
	}

	benchs := []struct {
		fn   func(grid [][]byte) int
		name string
	}{
		{name: "mine", fn: numIslands},
		{name: "other", fn: numIslands_1},
	}

	for _, ben := range benchs {
		b.Run(ben.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for _, bs := range bss {
					ben.fn(bs)
				}
			}
		})
	}
}

func dfs_numIslands_1(grid [][]byte, row, col int) bool {
	if grid[row][col] == '1' {
		grid[row][col] = '0'
		if row+1 < len(grid) {
			dfs_numIslands_1(grid, row+1, col)
		}
		if row-1 >= 0 {
			dfs_numIslands_1(grid, row-1, col)
		}
		if col+1 < len(grid[0]) {
			dfs_numIslands_1(grid, row, col+1)
		}
		if col-1 >= 0 {
			dfs_numIslands_1(grid, row, col-1)
		}
		return true
	}
	return false
}

func numIslands_1(grid [][]byte) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if dfs_numIslands_1(grid, i, j) {
				count++
			}
		}
	}
	return count
}
