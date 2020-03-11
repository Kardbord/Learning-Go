package fibonacci

// Computes the n'th fibonacci number
func Fibonacci(n uint32) uint64 {
	if n == 0 || n == 1 { return uint64(n) }
	var val1 uint64 = 0
	var val2 uint64 = 1
	
	var i uint32 = 2
	for i <= n {
		if (i % 2) == 0 {
			val1 += val2
		} else {
			val2 += val1
		}
		i++
	}
	if (i % 2) == 0 { return val2 }
	return val1
}
