# BruteForceCaesarCipher
Simple function to brute force all possible combinations of standard Caesar cipher


## Usage Example:
- Replace "***exxegoexsrgi***" with your search string and check answer :+1:
```Go
for id, answer := range BruteForceCaesarCipher("exxegoexsrgi") {
	// Check answer here if not printing whole array
	if answer == "ATTACKATONCE" {
		fmt.Println(id, "SUCCESS: ", answer)
	}
}
```
