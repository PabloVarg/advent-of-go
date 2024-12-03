package input

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func ReadIntSlice(in io.Reader) []int {
	s := bufio.NewScanner(in)

	if !s.Scan() {
		return nil
	}

	line := s.Text()
	parts := strings.Split(line, " ")

	result := make([]int, 0)
	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			panic(err)
		}

		result = append(result, num)
	}

	return result
}
