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
	A := readI()
	B := readI()
	C := readI()
	X := readI()
	var cnt int = 0

	for i := 0; i <= A; i++ {
		for k := 0; k <= B; k++ {
			for j := 0; j <= C; j++ {
				if i*500+k*100+j*50 == X {
					cnt++
				}
			}
		}
	}
	fmt.Println(cnt)
}
