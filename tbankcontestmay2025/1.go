package tbankcontestmay2025

import (
	"bufio"
	"fmt"
	"os"
)

func CheckIsPalindrome() {
	isPalindrome()
}

// Решение
func isPalindrome() {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	symbols := []rune(str)

	// x_dx, xd_x, x_xd, dx_x
	if symbols[0] == symbols[3] || symbols[0] == symbols[2] || symbols[1] == symbols[3] {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}