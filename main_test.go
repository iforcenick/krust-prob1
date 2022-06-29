package main

import (
	"fmt"
	"testing"
)

type Pair struct {
	input  int
	output bool
}

func TestIsPrime(t *testing.T) {
	arr := []Pair{
		{2, true},
		{3, true},
		{4, false},
		{5, true},
		{6, false},
		{7, true},
		{8, false},
		{9, false},
		{10, false},
	}
	fmt.Print(arr)
	for _, pair := range arr {
		test_name := fmt.Sprintf("%d primetest : ", pair.input)
		t.Run(test_name, func(t *testing.T) {
			ans := isPrime(pair.input)
			if ans != pair.output {
				t.Errorf("got %t, want %t", ans, pair.output)
			}
		})
	}
}
