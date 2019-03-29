package main

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func Test_findNthDigit(t *testing.T) {
	sb := strings.Builder{}
	sb.WriteString("0")
	limit := 100
	// 9 + 90*2 + 3
	for i := 1; i <= limit; i++ {
		sb.WriteString(strconv.Itoa(i))
	}
	baseString := sb.String()
	fmt.Println(baseString)

	// 1234567891011
	//for n := len(baseString) - 1; n > 0; n-- {
	for n := 0; n < len(baseString); n++ {
		testName := fmt.Sprintf("%v", n)
		t.Run(testName, func(t *testing.T) {
			got := findNthDigit(n)
			want := int(baseString[n] - '0')
			if got != want {
				t.Errorf("findNthDigit(%v) = %v, want %v", n, got, want)
			}
		})
	}
}
