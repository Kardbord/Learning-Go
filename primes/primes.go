package primes

// Returns a closure that will generate sequential primes beginning with 2.
// For all those times when you want to generate primes, but only one at a time. :)
func SequentialPrime() func() uint64 {
	n := uint64(2)
	primesToN := SieveOfEratosthenes(n)
	lastNumPrimes := len(primesToN)
	
	return func() uint64 {
		defer func(){
			for len(primesToN) == lastNumPrimes {
				primesToN = SieveOfEratosthenes(n)
				n++
			}
			lastNumPrimes = len(primesToN)
		}()
		
		return primesToN[len(primesToN)-1]
	}
}

// Returns a list of primes up to n
func SieveOfEratosthenes(n uint64) []uint64 {
	A := make([]bool, n+1, n+1)
	for i := 2; i < len(A); i++ { A[i] = true }
	
	for i := uint64(2); i*i <= n; i++ {
		if A[i] {
			for j := i*2; j <= n; j += i {
				A[j] = false
			}
		}
	}
	
	var primes []uint64
	
	for i := uint64(2); i <= n; i++ {
		if A[i] {
			primes = append(primes, i)
		}
	}
	
	return primes
}


