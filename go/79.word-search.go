package main

// medium

func exist(board [][]byte, word string) bool {
	r := len(board)
	c := len(board[0])
	buf := make([]byte, 0, len(word))
	var find func(x, y int) bool
	find = func(x, y int) bool {
		b := board[y][x]
		if b != word[len(buf)] {
			return false
		}

		buf = append(buf, b)
		board[y][x] = 0
		defer func() {
			board[y][x] = b
			buf = buf[:len(buf)-1]
		}()

		if len(buf) == len(word) {
			return true
		}

		// put this 2 check for same x before the 2 for same y make this code
		// 100% faster (8ms to 4ms). This is just due to cache line optimized.
		if x > 0 && board[y][x-1] != 0 && find(x-1, y) {
			return true
		}

		if x < c-1 && board[y][x+1] != 0 && find(x+1, y) {
			return true
		}

		if y > 0 && board[y-1][x] != 0 && find(x, y-1) {
			return true
		}

		if y < r-1 && board[y+1][x] != 0 && find(x, y+1) {
			return true
		}

		return false
	}

	for y := 0; y < r; y++ {
		for x := 0; x < c; x++ {
			if find(x, y) {
				return true
			}
		}
	}

	return false
}
