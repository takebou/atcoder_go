package main

import (
	"bufio"
	"fmt"
	"math"
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

func main() {
	sc.Buffer(make([]byte, 64*1024), 100001)
	var t int
	var x, y float64
	t, x, y = 0, 0, 0

	N := readI()
	for i := 0; i < N; i++ {
		A := make([]int, 3)
		for k := 0; k < 3; k++ {
			A[k] = readI()
		}
		t = A[0] - t
		x = math.Abs(float64(A[1]) - x)
		y = math.Abs(float64(A[2]) - y)

		if t < int(x+y) {
			fmt.Println("No")
			os.Exit(0)
		}
		if (t-int(x+y))%2 != 0 {
			fmt.Println("No")
			os.Exit(0)
		}
	}
	fmt.Println("Yes")
}
