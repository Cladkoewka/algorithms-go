package tbankcontestmay2025

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func Test3Solve() {
	solve3()
}

func solve3() {
	writer := bufio.NewWriter(os.Stdout)
	reader := bufio.NewReader(os.Stdin)
	defer writer.Flush()

	used := make(map[int]struct{})

	var n int
	fmt.Fscan(reader, &n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	sort.Slice(a, func(i, j int) bool {
		return a[i] > a[j]
	})

	for _, x := range a {
		for x >= 0 {
			if _, exists := used[x]; !exists {
				used[x] = struct{}{}
				break
			}
			if x == 0 {
				break
			}
			x /= 2
		}
	}

	fmt.Fprintln(writer, len(used))
}