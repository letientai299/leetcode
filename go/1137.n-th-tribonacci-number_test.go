package main

import (
	"log"
	"testing"
)

func Test_tribonacci(t *testing.T) {
	for i := 0; i < 10; i++ {
		log.Println(tribonacci(i))
	}
}
