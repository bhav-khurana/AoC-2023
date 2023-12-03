package part1

import (
	"fmt"
	"strconv"
)

func isSymbol(c rune) bool {
	return (c < '0' || c > '9') && c != '.'
}

var moves [][]int = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}, {1, 1}, {-1, -1}, {-1, 1}, {1, -1}}

func isInside(i, j, n, m int) bool {
	return i >= 0 && i < n && j >= 0 && j < m
}

func Solve() {

	var ans int
	var inp []string

	for {
		var s string
		_, err := fmt.Scanln(&s)
		if s == "" {
			break
		}

		if err != nil {
			fmt.Println("Error: ", err)
			break
		}

		if s == "" {
			break
		}

		inp = append(inp, s)
	}

	for i := 0; i < len(inp); i++ {
		s := inp[i]
		l := len(s)
		var idx []int
		idx = append(idx, -1)
		for i, c := range s {
			if c < '0' || c > '9' {
				idx = append(idx, i)
			}
		}
		idx = append(idx, l)
		for k := 1; k < len(idx); k++ {
			st := idx[k-1]+1
			end := idx[k]-1
			if st > end {
				continue
			}
			num, err := strconv.Atoi(s[st:end+1])
			if err != nil {
				fmt.Println("Error: ", err)
				break
			}
		OuterLoop:
			for j := st; j <= end; j++ {
				for _, move := range moves {
					ni := i + move[0]
					nj := j + move[1]
					if isInside(ni, nj, len(inp), len(inp[0])) && isSymbol(rune(inp[ni][nj])) {
						ans += num
						break OuterLoop
					}
				}
			}
		}
	}

	fmt.Println(ans)

}