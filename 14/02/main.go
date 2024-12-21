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
	m, n := 103, 101
	// m, n := 7, 11
	grid := make([][]int, m)
	for i := range m {
		grid[i] = make([]int, n)
	}
	vis := make([][]int, m)
	for i := range m {
		vis[i] = make([]int, n)
	}

	robots := readInput()
outer:
	for seconds := range 10000 {
		for _, robot := range robots {
			robot.pos.X = (((robot.pos.X + robot.vel.X*seconds) % m) + m) % m
			robot.pos.Y = (((robot.pos.Y + robot.vel.Y*seconds) % n) + n) % n
			grid[robot.pos.X][robot.pos.Y] = seconds
			if vis[robot.pos.X][robot.pos.Y] == seconds {
				continue outer
			}
			vis[robot.pos.X][robot.pos.Y] = seconds
		}

		fmt.Println(seconds)
		for i := range grid {
			for j := range grid[i] {
				if grid[i][j] == seconds {
					fmt.Print("#")
					continue
				}
				fmt.Print(".")
			}
			fmt.Println()
		}
	}
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
