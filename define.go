package bls

import "math/big"

var Zero = big.NewInt(0)
var One = big.NewInt(1)
var Two = big.NewInt(2)

func GetBigPrime(prime int) *big.Int {
	bigPrime := new(big.Int)
	bigPrime.SetInt64(int64(prime))
	return bigPrime
}

// 计算平方根
func Sqrt(a *big.Int) *big.Int {
	// TODO
	return nil
}
