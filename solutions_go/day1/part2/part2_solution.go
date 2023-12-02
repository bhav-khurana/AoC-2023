package part2

import (
	"fmt"
	"math"
	"unicode"
)

func Solve() {

	var ans int
	var digits map[string]int
	digits = make(map[string]int)

	digits["one"] = 1
	digits["two"] = 2
	digits["three"] = 3
	digits["four"] = 4
	digits["five"] = 5
	digits["six"] = 6
	digits["seven"] = 7
	digits["eight"] = 8
	digits["nine"] = 9

	for true {

		var inp string
		fmt.Scanln(&inp)
		if inp == "" {
			break
		}
		number := 0
		l := len(inp)

		for i := 0; i < l; i++ {

			sub3 := inp[i:int(math.Min(float64(i+3), float64(l)))]
			sub4 := inp[i:int(math.Min(float64(i+4), float64(l)))]
			sub5 := inp[i:int(math.Min(float64(i+5), float64(l)))]

			v, ok := digits[sub3]
			if ok {
				number += v * 10
				break
			}

			v, ok = digits[sub4]
			if ok {
				number += v * 10
				break
			}

			v, ok = digits[sub5]
			if ok {
				number += v * 10
				break
			}

			if unicode.IsDigit(rune(inp[i])) {
				number += (int(inp[i]) - '0') * 10
				break
			}

		}

		for i := l - 1; i >= 0; i-- {

			sub3 := inp[int(math.Max(float64(i-2), float64(0))) : i+1]
			sub4 := inp[int(math.Max(float64(i-3), float64(0))) : i+1]
			sub5 := inp[int(math.Max(float64(i-4), float64(0))) : i+1]

			v, ok := digits[sub3]
			if ok {
				number += v
				break
			}

			v, ok = digits[sub4]
			if ok {
				number += v
				break
			}

			v, ok = digits[sub5]
			if ok {
				number += v
				break
			}

			if unicode.IsDigit(rune(inp[i])) {
				number += (int(inp[i]) - '0')
				break
			}

		}

		ans += number
	}

	fmt.Println(ans)
}
