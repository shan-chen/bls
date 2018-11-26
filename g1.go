package bls

import (
	"crypto/sha1"
	"math/big"
)

// y^2=x^3+ax+b
// y^2=x^3+2x+1

type G1 struct {
	A int64
	B int64
}

type Point struct {
	X *big.Int
	Y *big.Int
}

// 杂凑将明文映射到椭圆曲线上一点
func (g *G1) MapToGroup(data []byte) *Point {
	hash := sha1.Sum(data)
	hashData := hash[:]
	bigData := new(big.Int)
	bigData.SetBytes(hashData)

	p := GetBigPrime(PrimeRanges)
	k := new(big.Int)
	k.Quo(p, bigData.Add(bigData, One))
	base := new(big.Int)
	base.Mul(bigData, k)

	for i := big.NewInt(0); i.Cmp(k) == -1; i.Add(i, One) {
		base.Add(base, i)
		if y := g.isOnCurve(base, p); y != nil {
			return &Point{
				X: base,
				Y: y,
			}
		}
	}
	return nil
}

// 判断是否是曲线上的点
// 如果是，返回对应y；否则返回空
func (g *G1) isOnCurve(x, p *big.Int) *big.Int {
	a, b := g.GetCurveParams()
	y := new(big.Int)
	y.Mul(x, x)
	y.Mul(x, y)
	tmp := new(big.Int)
	tmp.Mul(a, x)
	y.Add(y, tmp)
	y.Add(y, b)
	if isQuadraticResidue(y, p) {
		return Sqrt(y)
	}
	return nil
}

// 判断y是否是模p平方剩余
func isQuadraticResidue(y, p *big.Int) bool {
	mod := new(big.Int)
	mod.Mod(y, p)
	if mod.Cmp(Zero) == 0 {
		return true
	}
	index := new(big.Int)
	index.Sub(p, One)
	index.Div(index, Two)
	mod.Exp(y, index, p)
	if mod == One {
		return true
	}
	return false
}

// 计算kx
func (g *G1) Mul(x *Point, k *big.Int) *Point {
	//TODO
	return &Point{}
}

// 获取生成元
func (g *G1) GetGenerator() *Point {
	//TODO
	return &Point{}
}

// 获取曲线参数
func (g *G1) GetCurveParams() (*big.Int, *big.Int) {
	var a, b *big.Int
	a.SetInt64(g.A)
	b.SetInt64(g.B)
	return a, b
}

// 双线性映射
func (g *G1) BilinearMap(g1 *Point, g2 *Point) *Point {
	return &Point{}
}

func (g *G1) doubleJacobian(x, y, z *big.Int) (*big.Int, *big.Int, *big.Int) {
	p := GetBigPrime(PrimeRanges)

	delta := new(big.Int).Mul(z, z)
	delta.Mod(delta, p)
	gamma := new(big.Int).Mul(y, y)
	gamma.Mod(gamma, p)
	alpha := new(big.Int).Sub(x, delta)
	if alpha.Sign() == -1 {
		alpha.Add(alpha, p)
	}
	alpha2 := new(big.Int).Add(x, delta)
	alpha.Mul(alpha, alpha2)
	alpha2.Set(alpha)
	alpha.Lsh(alpha, 1)
	alpha.Add(alpha, alpha2)

	beta := alpha2.Mul(x, gamma)

	x3 := new(big.Int).Mul(alpha, alpha)
	beta8 := new(big.Int).Lsh(beta, 3)
	x3.Sub(x3, beta8)
	for x3.Sign() == -1 {
		x3.Add(x3, p)
	}
	x3.Mod(x3, p)

	z3 := new(big.Int).Add(y, z)
	z3.Mul(z3, z3)
	z3.Sub(z3, gamma)
	if z3.Sign() == -1 {
		z3.Add(z3, p)
	}
	z3.Sub(z3, delta)
	if z3.Sign() == -1 {
		z3.Add(z3, p)
	}
	z3.Mod(z3, p)

	beta.Lsh(beta, 2)
	beta.Sub(beta, x3)
	if beta.Sign() == -1 {
		beta.Add(beta, p)
	}
	y3 := alpha.Mul(alpha, beta)

	gamma.Mul(gamma, gamma)
	gamma.Lsh(gamma, 3)
	gamma.Mod(gamma, p)

	y3.Sub(y3, gamma)
	if y3.Sign() == -1 {
		y3.Add(y3, p)
	}
	y3.Mod(y3, p)

	return x3, y3, z3
}
