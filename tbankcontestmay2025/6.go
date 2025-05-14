package tbankcontestmay2025

import (
	"fmt"
	"sort"
	"bufio"
	"os"
)

func Test6Solve() {
	solve6()
}

func solve6() {
	writer := bufio.NewWriter(os.Stdout)
	reader := bufio.NewReader(os.Stdin)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	v := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &v[i])
	}

	sort.Slice(v, func(i, j int) bool {
		return v[i] > v[j]
	})

	sum := 0 

	lenV := len(v)
	for i := 0; i < lenV / 2; i++ {
		sum += v[i] -  v[lenV - i - 1]
	}
		

	fmt.Fprintln(writer, sum)
}