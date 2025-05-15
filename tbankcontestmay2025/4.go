package tbankcontestmay2025

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func Test4Solve() {
	solve4()
}

const INF = int(1e9 + 7) 

func getNext(a, b, idxA int, st [][]int) int {
	c := 2*b - a
	if c < 1 || c > 10 {
		return INF
	}

	idxB := sort.Search(len(st[b]), func(i int) bool {
		return st[b][i] > idxA
	})
	if idxB == len(st[b]) {
		return INF
	}
	posB := st[b][idxB]

	idxC := sort.Search(len(st[c]), func(i int) bool {
		return st[c][i] > posB
	})
	if idxC == len(st[c]) {
		return INF
	}
	return st[c][idxC]
}

func solve4() {
	writer := bufio.NewWriter(os.Stdout)
	reader := bufio.NewReader(os.Stdin)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	v := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &v[i])
	}

	st := make([][]int, 20) 
	for i := 0; i < n; i++ {
		st[v[i]] = append(st[v[i]], i)
	}

	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = INF
		a := v[i]
		for b := 1; b <= 10; b++ {
			next := getNext(a, b, i, st)
			if next < dp[i] {
				dp[i] = next
			}
		}
	}

	ans := 0
	minDp := INF
	for i := n - 1; i >= 0; i-- {
		if dp[i] < minDp {
			minDp = dp[i]
		}
		add := n - minDp
		if add > 0 {
			ans += add
		}
	}

	fmt.Fprintln(writer, ans)
}