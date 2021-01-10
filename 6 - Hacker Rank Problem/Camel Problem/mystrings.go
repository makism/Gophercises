package main

func startsWithUppercase(b []byte) bool {
	return isUppercase(b[0])
}

func startsWithLowercase(b []byte) bool {
	return isLowercase(b[0])
}

func isLowercase(b byte) bool {
	return b >= 97 && b <= 122
}

func isUppercase(b byte) bool {
	return b >= 65 && b <= 90
}
