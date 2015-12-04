package main

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
