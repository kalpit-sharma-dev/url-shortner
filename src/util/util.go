package util

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func PadLeft(str, pad string, length int) string {
	for len(str) < length {
		str = pad + str
	}
	return str
}
