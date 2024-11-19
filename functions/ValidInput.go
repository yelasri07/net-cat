package functions

func ValidInput(s []byte) bool {
	if len(s) == 0 {
		return false
	}
	
	for _, ele := range s {
		if (ele < 32 && ele != 10)  {
			return false
		}
	}
	return true
}
