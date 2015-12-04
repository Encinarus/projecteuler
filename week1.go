package main

import "fmt"

// Day1 calculates the sum of the numbers between 0 and 1000
func Day1() int {
	sum := 0
	for i := 0; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}

// Day2 calculates the sum of the even fibonacci numbers below 4 million
func Day2() int {
	lastTerm := 0
	currentTerm := 1
	nextTerm := 1

	sum := 0
	for nextTerm <= 4000000 {
		if nextTerm%2 == 0 {
			sum += nextTerm
		}
		lastTerm = currentTerm
		currentTerm = nextTerm
		nextTerm = lastTerm + currentTerm
	}

	return sum
}

// Day3 returns the largest prime factor of 600851475143
func Day3() int {
	targetValue := 600851475143

	knownPrimes := make(map[int]bool)
	knownPrimes[2] = true

	factors := make(map[int]bool)
	for targetValue%2 == 0 {
		factors[2] = true
		targetValue /= 2
	}

Sieve:
	for primePotential := 3; primePotential <= targetValue; primePotential += 2 {
		for knownPrime := range knownPrimes {
			if primePotential%knownPrime == 0 {
				// Not a prime, continue
				continue Sieve
			}
		}
		// primePotential is prime! woo!
		for targetValue%primePotential == 0 {
			factors[primePotential] = true
			targetValue /= primePotential
		}
	}

	maxFactor := 0
	for factor := range factors {
		if factor > maxFactor {
			maxFactor = factor
		}
	}

	return maxFactor
}

func reverse(s string) string {
	// Get Unicode code points.
	n := 0
	rune := make([]rune, len(s))
	for _, r := range s {
		rune[n] = r
		n++
	}
	rune = rune[0:n]
	// Reverse
	for i := 0; i < n/2; i++ {
		rune[i], rune[n-1-i] = rune[n-1-i], rune[i]
	}
	// Convert back to UTF-8.
	return string(rune)
}

func Day4() int {
	largestPalindrome := 0
	for i := 999; i > 99; i-- {
		for j := i; j > 99; j-- {
			product := i * j
			if product > largestPalindrome {
				productString := fmt.Sprintf("%v", product)
				reversedString := reverse(productString)

				if productString == reversedString {
					largestPalindrome = product
				}
			} else {
				// Only getting smaller from here, don't even bother continuing
				// since we can't have a larger palindrome.
				break
			}
		}
	}
	return largestPalindrome
}

func Day5() int {
	return -1
}

func Day6() int {
	return -1
}

func Day7() int {
	return -1
}
