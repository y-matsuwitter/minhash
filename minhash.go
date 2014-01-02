package minhash

import (
	"math"
	"math/big"
	"math/rand"
)

var bitMask = uint32(0x1)
var hashFuncNum = 128

func minKey(l map[string]uint32) (string, uint32) {
	var result string
	m := uint32(math.MaxUint32)
	for k := range l {
		if m > l[k] {
			m = l[k]
			result = k
		}
	}
	return result, m
}

func minHash(data []string, seed uint32) uint32 {
	vector := make(map[string]uint32)
	for k := range data {
		vector[data[k]] = Murmurhash3_32(data[k], seed)
	}
	_, value := minKey(vector)
	return value
}

func signature(data []string) uint32 {
	rand.Seed(1)
	sig := uint32(0)
	for i := 0; i < hashFuncNum; i++ {
		sig += (minHash(data, rand.Uint32()) & bitMask) << uint32(i)
	}
	return sig
}

func signatureBig(data []string) *big.Int {
	rand.Seed(1)
	sigBig := big.NewInt(0)
	for i := 0; i < hashFuncNum; i++ {
		sigBig.SetBit(sigBig, i, uint(minHash(data, rand.Uint32())&bitMask))
	}
	return sigBig
}

func popCount(bits uint32) uint32 {
	bits = (bits & 0x55555555) + (bits >> 1 & 0x55555555)
	bits = (bits & 0x33333333) + (bits >> 2 & 0x33333333)
	bits = (bits & 0x0f0f0f0f) + (bits >> 4 & 0x0f0f0f0f)
	bits = (bits & 0x00ff00ff) + (bits >> 8 & 0x00ff00ff)
	return (bits & 0x0000ffff) + (bits >> 16 & 0x0000ffff)
}

func popCountBig(bits *big.Int) int {
	result := 0
	for _, v := range bits.Bytes() {
		result += int(popCount(uint32(v)))
	}
	return result
}

func Minhash(v1, v2 []string) float32 {
	commonBig := big.NewInt(0)
	commonBig.Xor(signatureBig(v1), signatureBig(v2))
	return 2.0 * (float32((hashFuncNum-popCountBig(commonBig)))/float32(hashFuncNum) - 0.5)
}
