package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"unicode"
	"unicode/utf8"
)

const (
	STATE_M     = iota
	STATE_U     = iota
	STATE_L     = iota
	STATE_OB    = iota // ( open brace
	STATE_NUM1  = iota
	STATE_COMMA = iota
	STATE_NUM2  = iota
	STATE_CB    = iota // ) closing brace
	STATE_OK    = iota
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanRunes)

	result := 0
	num1 := new(bytes.Buffer)
	num2 := new(bytes.Buffer)

	state := STATE_M
	hasNext := s.Scan()
	for hasNext {
		switch state {
		case STATE_M:
			if s.Text() == "m" {
				state = STATE_U
			}

			hasNext = s.Scan()
		case STATE_U:
			if s.Text() != "u" {
				state = STATE_M
				continue
			}

			state = STATE_L
			hasNext = s.Scan()
		case STATE_L:
			if s.Text() != "l" {
				state = STATE_M
				continue
			}

			state = STATE_OB
			hasNext = s.Scan()
		case STATE_OB:
			if s.Text() != "(" {
				state = STATE_M
				continue
			}

			state = STATE_NUM1
			hasNext = s.Scan()
		case STATE_NUM1:
			num1.Reset()
			for {
				if !unicode.IsDigit([]rune(s.Text())[0]) {
					break
				}

				num1.WriteString(s.Text())
				hasNext = s.Scan()
			}

			if utf8.RuneCount(num1.Bytes()) > 3 {
				state = STATE_M
				continue
			}

			state = STATE_COMMA
		case STATE_COMMA:
			if s.Text() != "," {
				state = STATE_M
				continue
			}

			state = STATE_NUM2
			hasNext = s.Scan()
		case STATE_NUM2:
			num2.Reset()
			for hasNext {
				if !unicode.IsDigit([]rune(s.Text())[0]) {
					break
				}

				num2.WriteString(s.Text())
				hasNext = s.Scan()
			}

			if utf8.RuneCount(num2.Bytes()) > 3 {
				state = STATE_M
				continue
			}

			state = STATE_CB
		case STATE_CB:
			if s.Text() != ")" {
				state = STATE_M
				continue
			}

			state = STATE_OK
			hasNext = s.Scan()
		case STATE_OK:
			a, err := strconv.Atoi(num1.String())
			if err != nil {
				panic("not a number")
			}

			b, err := strconv.Atoi(num2.String())
			if err != nil {
				panic("not a number")
			}

			result += a * b
			state = STATE_M
		}
	}

	fmt.Println(result)
}
