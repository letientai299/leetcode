package main

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	n := 0
	for i, s := range startTime {
		if queryTime >= s && queryTime <= endTime[i] {
			n++
		}
	}
	return n
}
