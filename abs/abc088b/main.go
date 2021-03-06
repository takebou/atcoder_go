package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func main() {
	var Alice, Bob int
	N := readI()
	A := make([]int, N)

	for i := 0; i < N; i++ {
		A[i] = readI()
	}

	sort.Sort(sort.Reverse(sort.IntSlice(A)))

	for i := 0; i < N; i++ {
		if i%2 == 0 {
			Alice += A[i]
		} else {
			Bob += A[i]
		}
	}
	fmt.Println(Alice - Bob)
}
