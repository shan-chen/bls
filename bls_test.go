package bls

import (
	"fmt"
	"testing"
)

func TestBls(t *testing.T) {
	priv, pub := GenKeyPair()
	fmt.Println("sk:", priv)
	fmt.Println("pk:", pub.String())
	data := "abcdjsskfklsjf"
	bytesData := []byte(data)
	sign := Signature(bytesData, priv)
	fmt.Println("sign:", sign)
	fmt.Println(Verify(bytesData, sign, pub))
}
