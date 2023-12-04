package part1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Solve() {
	var ans int
	scanner := bufio.NewScanner(os.Stdin)
	for {
		var inp string
		scanner.Scan()
		inp = scanner.Text()
		if inp == "" {
			break
		}
		var data string
		for i := 0; i < len(inp); i++ {
			if inp[i] == ':' {
				data = inp[i+2:]
				break
			}
		}
		dataArr := strings.Split(data, " | ")
		winingCards := strings.Fields(dataArr[0])
		isWining := make(map[int]bool)
		for _, card := range winingCards {
			cardNo, err := strconv.Atoi(strings.Trim(card, "[]"))
			if err != nil {
				fmt.Println("Error: ", err)
				break
			}
			isWining[cardNo] = true
		}
		myCards := strings.Fields(dataArr[1])
		var value int
		for _, card := range myCards {
			cardNo, err := strconv.Atoi(strings.Trim(card, "[]"))
			if err != nil {
				fmt.Println("Error: ", err)
				break
			}
			if isWining[cardNo] {
				if value == 0 {
					value = 1
				} else {
					value *= 2
				}
			}
		}
		ans += value
	}
	fmt.Println(ans)
}