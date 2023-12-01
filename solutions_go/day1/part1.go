package main

import (
	"fmt"
	"unicode"
)

func main() {

	var ans int

	for true {

		var inp string
		fmt.Scanln(&inp)
		if inp == "" {
			break
		}
		number := 0
		for _, c := range inp {
			if unicode.IsDigit(c) {
				number += (int(c) - '0') * 10
				break
			}
		}
		l := len(inp)
		for i := l - 1; i >= 0; i-- {
			if unicode.IsDigit(rune(inp[i])) {
				number += (int(inp[i]) - '0')
				break
			}
		}

		ans += number
	}

	fmt.Println(ans)
}
