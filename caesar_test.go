package go_encrypt

import (
	"testing"
)

func TestEncrypt(t *testing.T) {
	want := "helloworld"
	if result := Encrypt(want, KeyGen(3)); result != "khoorzruog" {
		t.Fatalf("want: %s. but result: %s", want, result)
	}
}

func TestDecrypt(t *testing.T) {
	want := "khoorzruog"
	if result := Decrypt(want, KeyGen(3)); result != "helloworld" {
		t.Fatalf("want: %s. but result: %s", want, result)
	}
}
