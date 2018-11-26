package bls

const (
	PrimeRanges = 11
)

/*
var g1 = &G1{
	A: 2,
	B: 1,
}


// BLS数字签名
func Signature(data []byte) ([]byte, *Point) {
	privKey, pubKey := genKeyPair()
	return Sign(data, privKey), pubKey
}

// 生成密钥
func genKeyPair() (*big.Int, *Point) {
	// 强随机
	randNum, err := rand.Int(rand.Reader, big.NewInt(PrimeRanges))
	if err != nil {
		return nil, nil
	}
	gen := g1.GetGenerator()
	pubKey := g1.Mul(gen, randNum)
	return randNum, pubKey
}

// 计算签名
func Sign(data []byte, privKey *big.Int) []byte {
	hashPoint := g1.MapToGroup(data)
	cipherPoint := g1.Mul(hashPoint, privKey)

	//TODO 序列化
	signature, err := json.Marshal(cipherPoint)
	if err != nil {
		return nil
	}
	return signature
}

// 验证签名
// e(P,S)=e(A,M)
func Verify(data []byte, signature []byte, pubKey *Point) bool {
	var cipherPoint *Point
	err := json.Unmarshal(signature, cipherPoint)
	if err != nil {
		return false
	}
	hashPoint := g1.MapToGroup(data)
	return g1.BilinearMap(g1.GetGenerator(), cipherPoint) == g1.BilinearMap(pubKey, hashPoint)
}
*/
