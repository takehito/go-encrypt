package caesar

var alphabets = "abcdefghijklmnopqrstuvwxyz"

func KeyGen(n int) int {
	return n
}

func Encrypt(m string, n int) string {
	result := make([]byte, len(m))
	for i, c := range m {
		for j, alphabet := range alphabets {
			if c != alphabet {
				continue
			}

			shift := j + n
			if shift > 25 {
				result[i] = alphabets[shift-26]
				continue
			}

			result[i] = alphabets[shift]
		}
	}

	return string(result)
}

func Decrypt(c string, n int) string {
	shift := 26 - n
	return Encrypt(c, shift)
}
