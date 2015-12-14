package main

import (
	"bytes"
	"fmt"

	dg "github.com/agadabanka/tic-tac-go/diagnostics"
)

type Factor struct {
	base  int
	radix int
}

type Factors struct {
	list []Factor
}

// Equals on the factor compares two Factor slices for sematical equality
// The Factors are expected to be in the same order in both slices
func (this Factors) Equals(that Factors) bool {
	if len(this.list) != len(that.list) {
		return false
	}
	for i, thatFactor := range that.list {
		thisFactor := this.list[i]
		if thisFactor != thatFactor {
			return false
		}
	}
	return true
}

func (this Factor) String() string {
	if this.radix > 1 {
		return fmt.Sprintf("%d^%d", this.base, this.radix)
	} else {
		return fmt.Sprintf("%d", this.base)
	}
}

func (this Factors) String() string {
	var buffer bytes.Buffer
	var p string
	prefix := " * "
	
	for i, factor := range this.list {
		if p != prefix && i > 0 {
			p = prefix
		}
		buffer.WriteString(fmt.Sprintf("%s%s", p, factor.String()))
	}
	return buffer.String()
}

func GetPrimeFactors(n int) Factors {
	dg.AssertT(n > 1, "Input number should be greater than 1")
	var factors Factors = Factors{[]Factor{Factor{1, 1}}} // 1 is always a factor
	for i := 2; i <= n; i++ {
		var f Factor = Factor{base: i}
		for n%i == 0 {
			n /= i
			f.radix += 1
		}
		if f.radix > 0 {
			factors.list = append(factors.list, f)
		}
	}
	dg.AssertT(len(factors.list) >= 2, "All numbers should have atleast two prime factors")
	return factors
}
