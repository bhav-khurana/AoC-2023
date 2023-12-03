package part2

import (
	"fmt"
	"strconv"
	"unicode"
)

var moves [][]int = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}, {1, 1}, {-1, -1}, {-1, 1}, {1, -1}}

func isInside(i, j, n, m int) bool {
	return i >= 0 && i < n && j >= 0 && j < m
}

type Pair struct {
    First  int
    Second int
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
		for j := 0; j < len(inp[i]); j++ {
			if inp[i][j] == '*' {
				var taken map[Pair]bool = make(map[Pair]bool)
				var cnt int
				var gearRatio int = 1
				for _, move := range moves {
					x := i + move[0] 
					y := j + move[1]
					if isInside(x, y, len(inp), len(inp[i])) && unicode.IsDigit(rune(inp[x][y])) && !taken[Pair{x, y}] {
						cnt++
						var s string
						var nx, ny int = x, y
						taken[Pair{x, y}] = true
						s += string(inp[nx][ny])
						ny--
						for ny >= 0 && unicode.IsDigit(rune(inp[nx][ny])) {
							s = string(inp[nx][ny]) + s
							taken[Pair{nx, ny}] = true
							ny--
						}
						ny = y+1
						for ny < len(inp[i]) && unicode.IsDigit(rune(inp[nx][ny])) {
							s += string(inp[nx][ny])
							taken[Pair{nx, ny}] = true
							ny++
						}
						val, err := strconv.Atoi(s)
						if err != nil {
							fmt.Println("Error: ", err)
							break
						}
						gearRatio *= val
					}
				}
				if cnt == 2 {
					ans += gearRatio
				}
			}
		}
	}

	fmt.Println(ans)

}