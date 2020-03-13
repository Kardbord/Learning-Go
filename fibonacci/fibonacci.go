package fibonacci

// Computes the n'th fibonacci number
func Fibonacci(n uint32) uint64 {
	if n == 0 || n == 1 { return uint64(n) }
	var val1, val2 uint64 = 0, 1
	
	for i := uint32(2); i <= n; i++ {
		val1, val2 = val2, val1+val2
	}
	return val2
}
