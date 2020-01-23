package go_encrypt

func KeyGen(n int) int {
	return n
}

func Encrypt(m string, n int) string {
	result := make([]byte, len(m))
	for i, b := range []byte(m) {
		result[i] = b + byte(n)
	}
	return string(result)
}
