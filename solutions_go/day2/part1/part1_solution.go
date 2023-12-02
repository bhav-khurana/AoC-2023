package part1

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"os"
)

const (
	maxRed int = 12
	maxGreen int = 13
	maxBlue int = 14
)

func Solve() {
	
	var ans, total int
	scanner := bufio.NewScanner(os.Stdin)

	for {

		var inp, game, data string

		scanner.Scan()
    	inp = scanner.Text()

		if inp == "" {
			break
		}

		l := len(inp)

		for i := 0; i < len(inp); i++ {
			if inp[i] == ':' {
				game = inp[0:i]
				data = inp[i+2:l]
				break
			}
		}

		gameArr := strings.Split(game, " ")
		gameId, err := strconv.Atoi(gameArr[1])
		total += gameId

		if err != nil {
			fmt.Println("Error: ", err)
			break
		}

		games := strings.Split(data, "; ")

	OuterLoop:
		for _, gameDesc := range games {

			gameData := strings.Split(gameDesc, ", ")
			for _, value := range gameData {
				values := strings.Split(value, " ")
				cnt, err := strconv.Atoi(values[0])

				if err != nil {
					fmt.Println("Error: ", err)
					break
				}

				if values[1] == "red" && cnt > maxRed || values[1] == "green" && cnt > maxGreen || values[1] == "blue" && cnt > maxBlue {
					ans += gameId
					break OuterLoop
				}
			}
		}

	}

	fmt.Println(total-ans)
}