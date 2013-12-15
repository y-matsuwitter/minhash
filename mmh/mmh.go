package mmh

import (
	"bytes"
	"encoding/binary"
)

var mask32 = uint32(0xffffffff)

func rotl(x, r uint32) uint32 {
	return ((x << r) | (x >> (32 - r))) & mask32
}

func mmix(h uint32) uint32 {
	h &= mask32
	h ^= h >> 16
	h = (h * 0x85ebca6b) & mask32
	h ^= h >> 13
	h = (h * 0xc2b2ae35) & mask32
	return h ^ (h >> 16)
}

func Murmurhash3_32(key string, seed uint32) uint32 {
	var h1 uint32 = seed
	var k uint32
	var c1, c2 uint32 = 0xcc9e2d51, 0x1b873593
	buffer := bytes.NewBufferString(key)
	keyBytes := []byte(key)
	length := buffer.Len()
	if length == 0 {
		return 0
	}

	nblocks := length / 4
	for i := 0; i < nblocks; i++ {
		binary.Read(buffer, binary.LittleEndian, &k)
		k *= c1
		k = rotl(k, 15)
		k *= c2

		h1 ^= k
		h1 = rotl(h1, 13)
		h1 = h1*5 + 0xe6546b64
	}

	// 4バイトブロック外の最後のデータを反映する
	var k1 uint32 = 0
	tail := nblocks * 4
	switch length & 3 {
	case 3:
		k1 ^= uint32(keyBytes[tail+2]) << 16
		fallthrough
	case 2:
		k1 ^= uint32(keyBytes[tail+1]) << 8
		fallthrough
	case 1:
		k1 ^= uint32(keyBytes[tail])
		k *= c1
		k = rotl(k, 15)
		k *= c2
		h1 ^= k
	}
	// finalize
	h1 ^= uint32(length)
	return mmix(h1)
}
