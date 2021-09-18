package main

// Faulty Sensor
//
// Easy
//
// https://leetcode.com/problems/faulty-sensor/
func badSensor(sensor1 []int, sensor2 []int) int {
	n := len(sensor1)
	i := 0

	for i < n && sensor1[i] == sensor2[i] {
		i++
	}

	if i >= n-1 {
		return -1 // either none faults or error occur at last data point
	}

	s1, s2 := true, true
	for ; i+1 < n; i++ {
		s1 = s1 && sensor1[i] == sensor2[i+1]
		s2 = s2 && sensor2[i] == sensor1[i+1]
	}

	if s1 && s2 {
		return -1
	}

	if s1 {
		return 1
	}

	return 2
}
