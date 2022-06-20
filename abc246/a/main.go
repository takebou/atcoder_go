package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func intSliceUnique(list []int) []int {
	m := make(map[int]int)
	slim := make([]int, 0)
	for _, e := range list {
		if _, ok := m[e]; !ok {
			m[e] = 0
			slim = append(slim, e)
		}
	}
	return slim
}

func abs(i int) int {
	if i < 0 {
		i = i * -1
	}
	return i
}

func main() {
	sc.Buffer(make([]byte, 64*1024), 100001)
	x := 0
	y := 0
	xy1 := make([]int, 2)
	xy2 := make([]int, 2)
	xy3 := make([]int, 2)

	for i := 0; i < 2; i++ {
		xy1[i] = readI()
	}
	for i := 0; i < 2; i++ {
		xy2[i] = readI()
	}
	for i := 0; i < 2; i++ {
		xy3[i] = readI()
	}

	if xy1[0] == xy2[0] {
		x = xy3[0]
	} else if xy1[0] == xy3[0] {
		x = xy2[0]
	} else {
		x = xy1[0]
	}

	if xy1[1] == xy2[1] {
		y = xy3[1]
	} else if xy1[1] == xy3[1] {
		y = xy2[1]
	} else {
		y = xy1[1]
	}

	fmt.Println(x, y)
}
