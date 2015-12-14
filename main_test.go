package main

import "testing"

func TestPrime(t *testing.T) {
	cases := []struct{
		number int
		want Factors
	}{
		{2, Factors{[]Factor{Factor{1, 1}, Factor{2, 1}}}},
		{3, Factors{[]Factor{Factor{1, 1}, Factor{3, 1}}}},
		{4, Factors{[]Factor{Factor{1, 1}, Factor{2, 2}}}},
		{6, Factors{[]Factor{Factor{1, 1}, Factor{2, 1}, Factor{3, 1}}}},
		{8, Factors{[]Factor{Factor{1, 1}, Factor{2, 3}}}},
		{9, Factors{[]Factor{Factor{1, 1}, Factor{3, 2}}}},
		{12, Factors{[]Factor{Factor{1, 1}, Factor{2, 2}, Factor{3, 1}}}},
		{30, Factors{[]Factor{Factor{1, 1}, Factor{2, 1}, Factor{3, 1}, Factor{5, 1}}}},
	}

	for _, c := range cases {
		got := GetPrimeFactors(c.number)
		if !got.Equals(c.want) {
			t.Errorf("GetPrimeFactors(%d) == %v, want %v", c.number, got, c.want)
		}
	}
}