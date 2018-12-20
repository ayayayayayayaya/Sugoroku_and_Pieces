package main

import (
		"fmt"
		"bufio"
		"os"
		"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func nextInt() int {
	scanner.Scan()
	i, e := strconv.Atoi(scanner.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func contains(s []int, e int) bool {
	for _, v := range s{
		if e == v {
			return true
		}
	}
	return false
}

func main() {
	var N, M int
	var X, A []int
	scanner.Split(bufio.ScanWords)
	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		X = append(X, nextInt())
	}

	fmt.Scan(&M)
	
	fmt.Println(N,X,M,A)
	for i := 0; i < M; i++ {
		A = append(A, nextInt())
	}
	
	for i := 0; i < M; i++ {
		if !contains(X, X[A[i] - 1] + 1) && X[A[i] - 1] != 2019 {
			X[A[i] - 1]++
		}
	}

	for i := 0; i < N; i++ {
		fmt.Println(X[i])
	}
		
}
