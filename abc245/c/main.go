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
	testFile, err := os.Open("./tests/sample-2.in")
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
	N := readI()
	K := readI()
	A := make([]int, N)
	B := make([]int, N)
	dpa := make([]bool, N)
	dpb := make([]bool, N)

	dpa[0], dpb[0] = true, true

	for i := 0; i < N; i++ {
		A[i] = readI()
	}
	for i := 0; i < N; i++ {
		B[i] = readI()
	}

	for k := 1; k < N; k++ {
		if dpa[k-1] && abs(A[k-1]-A[k]) <= K {
			dpa[k] = true
		}
		if dpa[k-1] && abs(A[k-1]-B[k]) <= K {
			dpb[k] = true
		}
		if dpb[k-1] && abs(B[k-1]-B[k]) <= K {
			dpb[k] = true
		}
		if dpb[k-1] && abs(B[k-1]-A[k]) <= K {
			dpa[k] = true
		}
	}

	if dpa[N-1] || dpb[N-1] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
