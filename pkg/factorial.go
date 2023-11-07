// Package pkg Package factorial provides a concurrent method for calculating the factorial of a given number.
package pkg

import (
	"math/big"
)

func CalculateFactorialSequentially(num int) *big.Int {
	result := big.NewInt(1)
	for i := 2; i <= num; i++ {
		result.Mul(result, big.NewInt(int64(i)))
	}
	return result
}

// CalculateFactorialConcurrently calculates the factorial of a given integer concurrently
// utilizing multiple goroutines. It distributes the work among CPU cores and returns the result as a *big.Int.
//
// The function takes two arguments:
// - num: The integer for which the factorial is to be calculated.
// - cpuCount: The number of CPU cores to utilize for concurrent calculation.
//
// The function breaks down the factorial calculation into smaller ranges and utilizes goroutines to calculate
// different segments concurrently. It returns the final factorial result as a *big.Int.
func CalculateFactorialConcurrently(num int, cpuCount int) *big.Int {
	iteration := calculateIteration(num, cpuCount)
	ch, counter := launchConcurrentMultiplying(num, iteration)

	result := big.NewInt(1)
	for i := 0; i < counter; i++ {
		result.Mul(result, <-ch)
	}
	return result
}

// launchConcurrentMultiplying launches goroutines to calculate factorial parts concurrently in smaller ranges.
// It returns a channel for results and a counter for the number of goroutines launched.
func launchConcurrentMultiplying(num int, iteration int) (chan *big.Int, int) {
	ch := make(chan *big.Int)
	counter := 0
	for i := 1; ; i += iteration {
		counter += 1
		if i+iteration > num {
			go multiplier(ch, i, num)
			break
		}
		go multiplier(ch, i, i+iteration)
	}
	return ch, counter
}

// multiplier calculates the factorial in the range (x; y].
// It sends the result to the provided channel.
func multiplier(ch chan *big.Int, x int, y int) {
	result := big.NewInt(1)
	for i := x + 1; i <= y; i++ {
		result.Mul(result, big.NewInt(int64(i)))
	}
	ch <- result
}

// calculateIteration determines the range for concurrent calculation based on the input num and cpuCount.
func calculateIteration(num int, cpuCount int) int {
	if num < cpuCount {
		return 1
	} else {
		return num / cpuCount
	}
}
