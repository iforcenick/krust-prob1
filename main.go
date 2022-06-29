package main

import (
	"fmt"
)

func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
func main() {
	cnt := 0
	for i := 2; ; i++ {
		if isPrime(i) {
			cnt++
			if cnt == 10001 {
				fmt.Println(i)
				return
			}
		}
	}
}
