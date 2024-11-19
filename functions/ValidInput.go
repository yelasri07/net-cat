package functions

func ValidInput(s []byte) bool {
	if len(s) == 0 {
		return false
	}

	for _, ele := range s {
		if ele < 32 && ele != 10 && ele != 9 {
			return false
		}
	}
	return true
}

func CheckSpaceName(s string) bool {
	if len(s) == 0 || len(s) > 15 {
		return false
	}

	for _, char := range s {
		if char == ' ' {
			return false
		}
	}

	return true
}
