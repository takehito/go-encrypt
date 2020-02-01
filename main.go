package main

import (
	"fmt"
	"github.com/takehito/go-encrypt/caesar"
)

func main() {
	encrypt := "kzdvwczvjczbvreriifn"
	for i := 1; i < 26; i++ {
		key := caesar.KeyGen(i)
		fmt.Printf("%d: %s\n", i, caesar.Encrypt(encrypt, key))
	}
}
