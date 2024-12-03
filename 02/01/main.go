package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/PabloVarg/advent-of-go/input"
)

func main() {
	solve()
}

func solve() {
	s := bufio.NewScanner(os.Stdin)

	safeReports := 0
	for s.Scan() {
		reports := input.ReadIntSlice(strings.NewReader(s.Text()))

		if safe(reports) {
			safeReports++
		}
	}

	fmt.Println(safeReports)
}

func safe(reports []int) bool {
	if len(reports) == 1 {
		return true
	}

	if reports[0] == reports[1] {
		return false
	}

	increasing := reports[0] < reports[1]
	for i, report := range reports {
		if i == 0 {
			continue
		}

		prev := reports[i-1]
		curr := report
		if !increasing {
			prev, curr = curr, prev
		}

		if prev >= curr {
			return false
		}

		if curr-prev > 3 {
			return false
		}
	}

	return true
}
