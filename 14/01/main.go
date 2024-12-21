package main

import (
	"fmt"
)

type Vector struct {
	X int
	Y int
}

type Robot struct {
	pos Vector
	vel Vector
}

func main() {
	seconds := 100
	m, n := 103, 101
	// m, n := 7, 11

	robots := readInput()
	quadrants := make([]int, 5)
	for _, robot := range robots {
		robot.pos.X = (((robot.pos.X + robot.vel.X*seconds) % m) + m) % m
		robot.pos.Y = (((robot.pos.Y + robot.vel.Y*seconds) % n) + n) % n
		quadrants[quadrant(robot.pos.X, robot.pos.Y, m, n)] += 1
	}

	res := 1
	for i := range 4 {
		if quadrants[i] == 0 {
			continue
		}

		res *= quadrants[i]
	}
	fmt.Println(res)
}

func quadrant(x, y, m, n int) int {
	res := 0

	if x == m/2 || y == n/2 {
		return 4
	}

	if x > m/2 {
		res += 1
	}
	if y > n/2 {
		res += 2
	}
	return res
}

func readInput() []Robot {
	res := make([]Robot, 0)

	for {
		var robot Robot
		_, err := fmt.Scanf(
			"p=%d,%d v=%d,%d\n",
			&robot.pos.Y,
			&robot.pos.X,
			&robot.vel.Y,
			&robot.vel.X,
		)
		if err != nil {
			break
		}

		res = append(res, robot)
	}

	return res
}
