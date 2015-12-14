package diagnostics

import "fmt"

// AssertI will panic in case of a mismatch between the expected value
// and the actual output. Assert is used only for enforcing invariants
// during development and should not be used in production code.
// This form of Assert takes two integer values
func AssertI(want, got int, message string) {
	if (want != got) {
		panic(fmt.Sprintf("want= %d, got = %d, message: %s", want, got, message))
	}
}

// AssertT will panic in case of a mismatch between the expected value
// and the actual output. Assert is used only for enforcing invariants
// during development and should not be used in production code.
// This form of Assert takes one boolean value
func AssertT(assertion bool, message string) {
	if (!assertion) {
		panic(fmt.Sprintf("assertion failed: message: %s", message))
	}
}

