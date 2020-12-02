package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
	"time"
)

func Test_longestSubsequence(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		diff int
		want int
	}{
		{arr: []int{1, 5, 7, 8, 5, 3, 4, 2, 1}, diff: -2, want: 4},
		{arr: []int{1, 2, 3, 4}, diff: 0, want: 1},
		{arr: []int{1, 2, 3, 4}, diff: 1, want: 4},
		{arr: []int{1, 2, 3, 4}, diff: 2, want: 2},
		{arr: []int{1, 1, 1, 1}, diff: 2, want: 1},
		{arr: []int{1, 1, 1, 1}, diff: 0, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestSubsequence(tt.arr, tt.diff); got != tt.want {
				t.Errorf("longestSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_longestSubsequence_big(t *testing.T) {
	arr, k := parseInput1218()
	start := time.Now()
	defer func() {
		fmt.Println("time: ", time.Now().Sub(start))
	}()

	n := longestSubsequence(arr, k)
	fmt.Println(n)
}

func parseInput1218() ([]int, int) {
	const input = "./testdata/1218-big.txt"
	wd, _ := os.Getwd()
	f, err := os.Open(filepath.Join(wd, input))
	if err != nil {
		log.Printf("err=%v", err)
		return nil, 0
	}

	info, err := f.Stat()
	if err != nil {
		log.Printf("err=%v", err)
		return nil, 0
	}

	sc := bufio.NewScanner(f)
	sc.Buffer(make([]byte, info.Size()), int(info.Size()))

	sc.Scan()
	line := sc.Text()
	line = strings.TrimPrefix(line, "[")
	line = strings.TrimSuffix(line, "]")
	arr := make([]int, 0, 1000)
	for _, s := range strings.Split(line, ",") {
		x, _ := strconv.Atoi(s)
		arr = append(arr, x)
	}

	sc.Scan()
	line = sc.Text()
	k, _ := strconv.Atoi(line)
	return arr, k
}
