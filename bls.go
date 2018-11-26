package bls

import (
	"bn256"
	"crypto/rand"
	"log"
	"math/big"
)

func GenKeyPair() (*big.Int, *bn256.G2) {
	priv, pub, err := bn256.RandomG2(rand.Reader)
	if err != nil {
		log.Println("genKeyPair error", err)
		return nil, nil
	}
	return priv, pub
}

func Signature(data []byte, priv *big.Int) []byte {
	hashPoint := mapToG1(data)
	cipherPoint := new(bn256.G1).ScalarMult(hashPoint, priv)
	signature := cipherPoint.Marshal()
	return signature
}

func Verify(data, signature []byte, pubKey *bn256.G2) bool {
	cipherPoint := new(bn256.G1)
	hashPoint := mapToG1(data)
	_, err := cipherPoint.Unmarshal(signature)
	if err != nil {
		log.Println("unmarshal error", err)
		return false
	}
	twistGen := new(bn256.G2).ScalarBaseMult(big.NewInt(1))
	ate1 := bn256.Pair(cipherPoint, twistGen)
	ate2 := bn256.Pair(hashPoint, pubKey)
	//ate1 := bn256.Miller(cipherPoint, twistGen).Finalize()
	//ate2 := bn256.Miller(hashPoint, pubKey).Finalize()
	return ate1.String() == ate2.String()
}

func mapToG1(data []byte) *bn256.G1 {
	hashData := new(big.Int).SetBytes(data)
	hashPoint := new(bn256.G1).ScalarBaseMult(hashData)
	return hashPoint
}
