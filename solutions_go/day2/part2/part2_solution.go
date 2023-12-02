package part2

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
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

		req := make(map[string]float64)

		for _, gameDesc := range games {

			gameData := strings.Split(gameDesc, ", ")
			for _, value := range gameData {
				values := strings.Split(value, " ")
				cnt, err := strconv.Atoi(values[0])

				if err != nil {
					fmt.Println("Error: ", err)
					break
				}

				req[values[1]] = math.Max(req[values[1]], float64(cnt))

			}
		}

		var power int = 1
		for _, value := range req {
			power *= int(value)
		}

		ans += power

	}

	fmt.Println(ans)
}