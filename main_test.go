package main

import (
	"fmt"
	"testing"
)

func TestBruteForceCaesarCipher(t *testing.T) {

	for i, v := range BruteForceCaesarCipher("Xvezlj nzkyflk vultrkzfe zj czbv jzcmvi ze kyv dzev") {
		fmt.Println(i, v)
	}

	for id, answer := range BruteForceCaesarCipher("exxegoexsrgi") {

		// Test
		if answer == "ATTACKATONCE" {
			fmt.Println(id, "SUCCESS: ", answer)
		}
	}
}
