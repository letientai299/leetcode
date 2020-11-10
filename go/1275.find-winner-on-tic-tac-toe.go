package main

func tictactoe(moves [][]int) string {
	b := [3][3]int{}
	for i, m := range moves {
		b[m[0]][m[1]] = i%2 + 1
	}

	winner := func(x int) string {
		if x == 1 {
			return "A"
		}
		return "B"
	}

	moreMove := false

	if b[0][0] == b[0][1] && b[0][1] == b[0][2] {
		if b[0][0] == 0 {
			moreMove = true
		} else {
			return winner(b[0][0])
		}
	}

	if b[1][0] == b[1][1] && b[1][1] == b[1][2] {
		if b[1][0] == 0 {
			moreMove = true
		} else {
			return winner(b[1][0])
		}
	}

	if b[2][0] == b[2][1] && b[2][1] == b[2][2] {
		if b[2][0] == 0 {
			moreMove = true
		} else {
			return winner(b[2][0])
		}
	}

	if b[0][0] == b[1][1] && b[1][1] == b[2][2] {
		if b[0][0] == 0 {
			moreMove = true
		} else {
			return winner(b[0][0])
		}
	}

	if b[2][0] == b[1][1] && b[1][1] == b[0][2] {
		if b[2][0] == 0 {
			moreMove = true
		} else {
			return winner(b[2][0])
		}
	}

	if b[0][0] == b[1][0] && b[1][0] == b[2][0] {
		if b[0][0] == 0 {
			moreMove = true
		} else {
			return winner(b[0][0])
		}
	}

	if b[0][1] == b[1][1] && b[1][1] == b[2][1] {
		if b[0][1] == 0 {
			moreMove = true
		} else {
			return winner(b[0][1])
		}
	}

	if b[0][2] == b[1][2] && b[1][2] == b[2][2] {
		if b[0][2] == 0 {
			moreMove = true
		} else {
			return winner(b[0][2])
		}
	}

	if moreMove || len(moves) < 9 {
		return "Pending"
	}

	return "Draw"
}
