package functions

import "fmt"

func isValidInput(str []byte, n int) bool {
	n = n - 1
	if n < 1 {
		return false
	}

	fmt.Println(string(str[:n]))

	return true
}
