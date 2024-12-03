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

		increasing := true
		if safe(reports, increasing) || safe(reports, !increasing) {
			safeReports++
		}
	}

	fmt.Println(safeReports)
}

func safe(reports []int, increasing bool) bool {
	if len(reports) == 1 {
		return true
	}

	firstBad := -1
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
			firstBad = i - 1
			break
		}

		if curr-prev > 3 {
			firstBad = i - 1
			break
		}
	}

	return testExcluding(reports, increasing, firstBad) ||
		testExcluding(reports, increasing, firstBad+1)
}

func testExcluding(reports []int, increasing bool, firstBad int) bool {
	if firstBad == -1 {
		return true
	}

	for i, report := range reports {
		if i == 0 || i == firstBad {
			continue
		}

		prevIdx := i - 1
		if i-1 == firstBad {
			prevIdx--
		}
		if prevIdx < 0 {
			continue
		}

		prev := reports[prevIdx]
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
