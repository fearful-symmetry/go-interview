package math

import (
	"reflect"
	"testing"
)

func TestBasicPrimeCheck(t *testing.T) {

	testVals := []float64{10, 17, 19, 25, 37}
	testOut := []bool{false, true, true, false, true}

	for index, toCheck := range testVals {
		isPrime := BasicPrimeCheck(toCheck)
		if isPrime != testOut[index] {
			t.Fatalf("Bad result checking for prime. %f is %v. Got %v", toCheck, testOut[index], isPrime)
		}
	}

}

func TestSieveOfEratosthenes(t *testing.T) {

	primes := SieveOfEratosthenes(50)

	knownPrimes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 49}

	if !reflect.DeepEqual(primes, knownPrimes) {
		t.Fatalf("Primes not equal. Got %v", primes)
	}

}
