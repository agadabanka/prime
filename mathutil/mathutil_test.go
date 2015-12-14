package mathutil

import "testing"

// TestMax tests the correctness of the Max utility
func TestMax(t *testing.T) {
	cases := []struct {
		a, b, want int
	}{
		{10, 20, 20},
		{-1, -10, -1},
		{10, 10, 10},
	}
	for _, c := range cases {
		got := Max(c.a, c.b)
		if got != c.want {
			t.Errorf("Max(%d, %d) == %d, want %d", c.a, c.b, got, c.want)
		}
	}
}

// TestMin tests the correctness of the Max utility
func TestMin(t *testing.T) {
	cases := []struct {
		a, b, want int
	}{
		{10, 20, 10},
		{-1, -10, -10},
		{10, 10, 10},
	}
	for _, c := range cases {
		got := Min(c.a, c.b)
		if got != c.want {
			t.Errorf("Min(%d, %d) == %d, want %d", c.a, c.b, got, c.want)
		}
	}
}

// TestMaxVR tests the correctness of the Max utility
func TestMaxVR(t *testing.T) {	
	testMaxV(t, "MaxVR")
}

// TestMaxVI tests the correctness of the Max utility
func TestMaxVI(t *testing.T) {		
	testMaxV(t, "MaxVI")
}

func testMaxV(t *testing.T, fname string) {
	cases_pass := []struct {
		a []int
		want int
	}{
		{[]int{10, 20, 30}, 30},
		{[]int{-1, -10, -5}, -1},
		{[]int{10, 10, 10}, 10},		
		{[]int{8, 9, 10, 100, 200, 300, -1, 10, -10}, 300},		
	}

	for _, c := range cases_pass {
		got := MaxV(c.a...)
		if got != c.want {
			t.Errorf("%s(%v) == %d, want %d", fname, c.a, got, c.want)
		}
	}
}


