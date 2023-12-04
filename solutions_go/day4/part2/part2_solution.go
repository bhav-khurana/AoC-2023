package part2

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
	var occurences [1000]int
	for i := 0; i < 1000; i++ {
		occurences[i] = 1
	}
	var curr int = 1
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
		var cnt int
		for _, card := range myCards {
			cardNo, err := strconv.Atoi(strings.Trim(card, "[]"))
			if err != nil {
				fmt.Println("Error: ", err)
				break
			}
			if isWining[cardNo] {
				cnt++
			}
		}
		for i := curr + 1; i <= curr + cnt; i++ {
			occurences[i] += occurences[curr];
		}
		curr++
	}
	for i := 1; i < curr; i++ {
		ans += occurences[i]
	}
	fmt.Println(ans)
}