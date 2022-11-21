package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Brute forces & returns all possible combinations of standard Caesar cipher
func BruteForceCaesarCipher(targetStr string) []string {

	fmt.Println("Starting Caesar brute force: ", targetStr)

	alphabet := [26]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	// Fix all input uppercase
	targetStr = strings.ToUpper(targetStr)

	// Variable buffers
	target, answers, cipher := []string{}, []string{}, []string{}

	// Append characters to target buffer
	for _, r := range targetStr {
		target = append(target, string(r))
	}

	// Reverse loop alphabet
	for v := len(alphabet) - 1; v >= 0; v-- {

		aBufList := alphabet[0 : v+1]
		for r := len(alphabet) - 1; r >= 0; r-- {
			if r == v {
				cipher = alphabet[v+1:]
				cipher = append(cipher, aBufList...)
			}
		}

		// Gets sequence by comparing index between cipher & alphabet (Normal caesar cipher checking)
		sequence := []string{}
		for _, x := range target {

			// Check for spaces
			if x == " " {
				sequence = append(sequence, " ")
				continue
			}

			// If not a space we idiotically assume character is in the defined alphabet

			// Check characters
			for i, letter := range alphabet {
				if letter == x {
					sequence = append(sequence, strconv.Itoa(i))
				}
			}
		}

		// Get possible answer from current sequence
		pa := ""
		for _, idx := range sequence {

			// For space support
			if idx == " " {
				pa += " "
				continue
			}

			id, _ := strconv.Atoi(idx)
			pa = pa + cipher[id]
		}

		answers = append(answers, pa)
	}

	return answers
}
