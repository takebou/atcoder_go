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

func main() {
	N := readI()
	A := make([]int, N)
	for i := 0; i < N; i++ {
		A[i] = readI()
	}

	for cnt := 0; ; cnt++ {
		for i := 0; i < N; i++ {
			if A[i]%2 == 0 {
				A[i] = A[i] / 2
			} else {
				fmt.Println(cnt)
				os.Exit(0)
			}
		}
	}
}
