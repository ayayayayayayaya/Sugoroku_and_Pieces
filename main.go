package main

import (
		"fmt"
		"bufio"
		"os"
)

var scanner = bufio.NewScanner(os.Stdin)

func contains(s []int, e int) bool {
	for _, v := range s{
		if e == v {
			return true
		}
	}
	return false
}

func main() {
	var n,m int
	
	fmt.Scan(&n)
	var x = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&x[i])
	}

	fmt.Scan(&m)
	var a = make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&a[i])
	}
	
	for i := 0; i < m; i++ {
		if !contains(x, x[a[i] - 1] + 1) && x[a[i] - 1] != 2019 {
			x[a[i] - 1]++
		}
	}

	for i := 0; i < n; i++ {
		fmt.Println(x[i])
	}
		
}
