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
	result := int64(0)

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		total, nums := readLine(strings.NewReader(s.Text()))

		if combinations(total, nums, 0, 0) {
			result += total
		}
	}

	fmt.Println(result)
}

func combinations(total int64, nums []int64, acc, i int64) bool {
	if i == int64(len(nums)) {
		return acc == total
	}

	if acc > total {
		return false
	}

	if i == 0 {
		return combinations(total, nums, nums[i], i+1)
	}

	return combinations(total, nums, acc+nums[i], i+1) ||
		combinations(total, nums, acc*nums[i], i+1) ||
		combinations(total, nums, concatenate(acc, nums[i]), i+1)
}

func concatenate(a int64, b int64) int64 {
	res := b + a*(pow(10, numDigits(b)))
	return res
}

func pow(b, exp int64) int64 {
	result := int64(1)
	for exp > 0 {
		if exp&1 == 1 {
			result *= b
		}
		b *= b
		exp >>= 1
	}

	return result
}

func numDigits(num int64) int64 {
	digits := int64(0)
	for num > 0 {
		num /= 10
		digits++
	}

	return digits
}

func readLine(r io.Reader) (int64, []int64) {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanWords)

	s.Scan()
	total, err := strconv.ParseInt(strings.TrimSuffix(s.Text(), ":"), 10, 64)
	if err != nil {
		panic("not a number")
	}

	nums := make([]int64, 0)
	for s.Scan() {
		num, err := strconv.ParseInt(s.Text(), 10, 64)
		if err != nil {
			panic("not a number")
		}

		nums = append(nums, num)
	}

	return total, nums
}
