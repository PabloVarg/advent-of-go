package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	result := 0

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		total, nums := readLine(strings.NewReader(s.Text()))

		if combinations(total, nums, 0, 0) {
			result += total
		}
	}

	fmt.Println(result)
}

func combinations(total int, nums []int, acc, i int) bool {
	if i == len(nums) {
		return acc == total
	}

	if i == 0 {
		combinations(total, nums, nums[i], i+1)
	}

	return combinations(total, nums, acc+nums[i], i+1) ||
		combinations(total, nums, acc*nums[i], i+1)
}

func readLine(r io.Reader) (int, []int) {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanWords)

	s.Scan()
	total, err := strconv.Atoi(strings.TrimSuffix(s.Text(), ":"))
	if err != nil {
		panic("not a number")
	}

	nums := make([]int, 0)
	for s.Scan() {
		num, err := strconv.Atoi(s.Text())
		if err != nil {
			panic("not a number")
		}

		nums = append(nums, num)
	}

	return total, nums
}
