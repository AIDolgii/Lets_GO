package main

import "fmt"

func countPrimes(n int) []int {
	sieve := make([]bool, n + 1)
	primes := []int{}
	for i := 2; i <= n; i++ {
		if sieve[i] == false {
			primes = append(primes, i)
			for j := i * i; j <= n; j += i {
				sieve[j] = true
			}
		}
	}
	return primes
}

func main() {
	var limit int
	fmt.Scan(&limit)
	for _, el := range countPrimes(limit) {
		fmt.Println(el)
	}
}