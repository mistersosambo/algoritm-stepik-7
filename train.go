package main

import "fmt"

func main() {
	var a int
	var m []int
	fmt.Scan(&a)
	for j := a; j != 0; j /= 10 {
		m = append(m, j%10)
	}
	for i := range m {
		fmt.Print(m[i])
	}
}
