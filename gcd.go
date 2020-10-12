package main

import (
	"fmt"
)

func gcd(m, n int) int {
	if m % n == 0 {
		return n
	}

	return gcd(m, m%n)
}

func main(){
	ans := gcd(51, 15)
	fmt.Println(ans)
}
