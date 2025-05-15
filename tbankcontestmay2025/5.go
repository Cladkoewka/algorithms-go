package tbankcontestmay2025

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Test5Solve() {
	solve5()
}


func solve5() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	parts := strings.Fields(line)

	n, _ := strconv.Atoi(parts[0])
	a, _ := strconv.Atoi(parts[1])
	b, _ := strconv.Atoi(parts[2])

	sLine, _ := reader.ReadString('\n')
	s := []rune(strings.TrimSpace(sLine))

	fmt.Println(solve(n, a, b, s))
}

func solve(n, a, b int, s []rune) int {
	if len(s) == 0 {
		return 0
	}

	lenS := 2 * n
	totalBrackets := 0
	totalCloseBrackets := 0
	totalOpenBrackets := 0
	bracketsCloseOpenBrackets := 0

	for totalBrackets < lenS {
		if totalBrackets != lenS-1 &&
			s[totalBrackets] == '(' &&
			s[totalBrackets+1] == ')' &&
			(totalBrackets%2) == 0 {
			totalBrackets += 2
		} else if totalBrackets != lenS-1 &&
			s[totalBrackets] == ')' &&
			s[totalBrackets+1] == '(' &&
			(totalBrackets%2) == 0 {
			bracketsCloseOpenBrackets++
			totalBrackets += 2
		} else if s[totalBrackets] == '(' {
			totalOpenBrackets++
			totalBrackets++
		} else if s[totalBrackets] == ')' {
			totalCloseBrackets++
			totalBrackets++
		}
	}

	cost1 := 0
	cost2 := 0
	sumOpenCloseBrackets := totalCloseBrackets + totalOpenBrackets

	if bracketsCloseOpenBrackets > 0 {
		cost1 = min(a, 2*b)
		cost2 = min(a, 2*b)
	}

	if totalCloseBrackets == totalOpenBrackets {
		cost1 += int(math.Ceil(float64(totalCloseBrackets+totalOpenBrackets) / 4.0)) * a
	} else if totalCloseBrackets == 0 || totalOpenBrackets == 0 {
		cost1 += int(math.Ceil(float64(totalCloseBrackets+totalOpenBrackets) / 2.0)) * b
	} else {
		diff := int(math.Abs(float64(totalCloseBrackets - totalOpenBrackets)))
		cost1 += (diff / 2 * a) + ((sumOpenCloseBrackets - diff*2) / 2 * b)
	}

	cost2 += int(math.Ceil(float64(totalCloseBrackets+totalOpenBrackets) / 2.0)) * b

	return min(cost1, cost2)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}