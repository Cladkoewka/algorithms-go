package tbankcontestmay2025

import (
	"bufio"
	"fmt"
	"os"
)


func Test2Solve() {
	solve2()
}

func solve2() {
	writer := bufio.NewWriter(os.Stdout)
	reader := bufio.NewReader(os.Stdin)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	a := make([]int64, n)
	b := make([]int64, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i], &b[i])
	}

	var q int
	fmt.Fscan(reader, &q)

	for i := 0; i < q; i++ {
		var t int
		var d int64
		fmt.Fscan(reader, &t, &d)
		t--

		diff := d - a[t]
		if diff <= 0 {
			fmt.Fprintln(writer, a[t])
		} else {
			add := ceilDiv(diff, b[t]) * b[t]
			fmt.Fprintln(writer, a[t]+add)
		}
	}
}


func ceilDiv(a, b int64) int64 {
	return (a + b - 1) / b
}
