package main

// medium

var romanSymbols = []byte{'I', 'V', 'X', 'L', 'C', 'D', 'M'}
var romanSymbolValues = []int{1, 5, 10, 50, 100, 500, 1000}

func intToRoman(num int) string {
	var buf []byte
	for i := len(romanSymbolValues) - 1; i >= 0 && num > 0; i-- {
		var p byte
		var adjust int
		switch romanSymbolValues[i] {
		case 5, 10:
			adjust = 1
			p = 'I'
		case 50, 100:
			p = 'X'
			adjust = 10
		case 500, 1000:
			p = 'C'
			adjust = 100
		}

		if romanSymbolValues[i] > num+adjust {
			continue
		}

		if romanSymbolValues[i] <= num+adjust && romanSymbolValues[i] > num {
			if p != 0 {
				buf = append(buf, p, romanSymbols[i])
			} else {
				buf = append(buf, romanSymbols[i])
			}
			num = num + adjust - romanSymbolValues[i]
			continue
		}

		a := num / romanSymbolValues[i]
		if a <= 3 {
			num -= romanSymbolValues[i] * a
			for x := 0; x < a; x++ {
				buf = append(buf, romanSymbols[i])
			}
			i++
			continue
		}

	}
	return string(buf)
}
