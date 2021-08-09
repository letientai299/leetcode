package main

import (
	"fmt"
	"testing"
)

func Test_rand10(t *testing.T) {
	n := 1000
	m := make(map[int]int)
	for i := 0; i < n; i++ {
		m[rand10()]++
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
}
