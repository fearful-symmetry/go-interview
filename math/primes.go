package math

import (
	"math"
)

//Prime number-related problems. I'm terrible at these, but I love them

//BasicPrimeCheck is the most basic possible prime search
//Primality tests are a huge *thing* in itself,
// and I could devote an entire project on different algos
func BasicPrimeCheck(n float64) bool {

	if n < 2 {
		return false
	}

	//search up to the square root of the number
	sqrt := math.Sqrt(n)
	for index := float64(2); index <= sqrt; index++ {
		if math.Mod(n, index) == 0 {
			return false
		}
	}
	return true

}

//SieveOfEratosthenes generates prime numbers up to i.
//The real interview question is pronouncing 'Eratosthenes'
//I've seen some crazy optimizations for this.
func SieveOfEratosthenes(i int) []int {
	//array of bool values for primes
	isPrimes := make([]bool, i)
	for index := 2; index < i; index++ {
		isPrimes[index] = true
	}

	for index := 2; index < int(math.Floor(math.Sqrt(float64(i)))); index++ {
		if isPrimes[index] {
			for indexInner := index * index; indexInner < len(isPrimes); indexInner += index {
				//fmt.Printf("Index: %d, indexInner: %d\n", index, indexInner)
				isPrimes[indexInner] = false
			}
		}
	}

	var primes []int
	//Yes, this is the lazy way to do this
	for i := range isPrimes {
		if isPrimes[i] {
			primes = append(primes, i)
		}
	}

	return primes
}
