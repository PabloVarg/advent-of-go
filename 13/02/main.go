package main

import (
	"fmt"
)

func main() {
	in := readInput()

	res := 0
	for in != nil {
		in[2][0] += 10000000000000
		in[2][1] += 10000000000000

		t1, t2, found := check(in)
		if found {
			res += 3*t1 + t2
		}

		fmt.Scanln()
		in = readInput()
	}

	fmt.Println(res)
}

// Solved equations
// t2 = (X - ax*t1) / bx, t1 = (Y*bx - by*X) / (bx*ay - by*ax)
func check(in [][]int) (int, int, bool) {
	d1 := in[1][0]*in[0][1] - in[1][1]*in[0][0]
	d2 := in[1][0]

	if d1 == 0 || d2 == 0 {
		return 0, 0, false
	}

	n1 := in[2][1]*in[1][0] - in[1][1]*in[2][0]
	if n1%d1 != 0 {
		return 0, 0, false
	}

	t1 := n1 / d1
	n2 := in[2][0] - in[0][0]*t1
	if n2%d2 != 0 {
		return 0, 0, false
	}

	t2 := n2 / d2
	return t1, t2, true
}

func readInput() [][]int {
	res := make([][]int, 0, 3)
	c := '0'
	x, y := 0, 0

	_, err := fmt.Scanf("Button %c: X+%d, Y+%d\n", &c, &x, &y)
	if err != nil {
		return nil
	}
	res = append(res, []int{x, y})

	_, err = fmt.Scanf("Button %c: X+%d, Y+%d\n", &c, &x, &y)
	if err != nil {
		return nil
	}
	res = append(res, []int{x, y})

	_, err = fmt.Scanf("Prize: X=%d, Y=%d\n", &x, &y)
	if err != nil {
		return nil
	}
	res = append(res, []int{x, y})

	return res
}
