// Package mathutil provides primitives that should have been part 
// of golang in the first place.
package mathutil

// Max accpets two integers and returns the maximum of the two
func Max(a int, b int) (mx int) {
	if a > b {
		mx = a
	} else {
		mx = b
	}
	return
}

// Min accepts two integers and returns the minimum of the two
func Min(a int, b int) (mn int) {
	if a < b {
		mn = a
	} else {
		mn = b
	}
	return
}

// MaxV accepts any number of integer arguments and returns the maximum
func MaxV(numbers ...int) (mx int) {				
	switch len(numbers) {
		case 0: return mx
		case 1: return numbers[0]
	default:
		return Max(numbers[0], MaxV(numbers[1:]...))		
	}	
}
