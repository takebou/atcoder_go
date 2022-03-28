package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func init() {
	if len(os.Args) >= 2 {
		if os.Args[1] == "debug" {
			debug()
		}
	}
	sc.Split(bufio.ScanWords)
}

func debug() {
	testFile, err := os.Open("./tests/sample-1.in")
	if err != nil {
		fmt.Fprintln(os.Stderr, "There is no testfile.")
		os.Exit(1)
	}
	sc = bufio.NewScanner(testFile)
}

func readS() string {
	sc.Scan()
	return sc.Text()
}

func readI() int {
	i, _ := strconv.Atoi(readS())
	return i
}

func readR() []rune {
	return []rune(readS())
}

func reverseS(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	return string(r)
}

func digit(i int, list []int) []int {
	if i > 0 {
		return digit(i/10, append(list, i%10))
	}
	return list
}

func main() {
	sc.Buffer(make([]byte, 64*1024), 100001)
	_ = readI()
	T := readS()
	Rcnt := 0
	Xcnt := 0
	Ycnt := 0

	arrT := strings.Split(T, "")

	for _, t := range arrT {
		if t == "R" {
			Rcnt++
		} else {
			switch Rcnt % 4 {
			case 0:
				Xcnt++
			case 1:
				Ycnt--
			case 2:
				Xcnt--
			case 3:
				Ycnt++
			}
		}
	}
	fmt.Println(Xcnt, Ycnt)
}
