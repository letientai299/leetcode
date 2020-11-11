package main

func replaceElements(arr []int) []int {
	n := len(arr)
	if n == 0 {
		return arr
	}

	t := arr[n-1]
	arr[n-1] = -1
	for i := n - 2; i >= 0; i-- {
		if t < arr[i+1] {
			t = arr[i+1]
		}
		t, arr[i] = arr[i], t
	}
	return arr
}
