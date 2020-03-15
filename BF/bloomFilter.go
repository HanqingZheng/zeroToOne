package BF

import (
	"zeroToOne/BitMap"
)

var cap = []uint{7, 11, 13}
var mod = []uint{31, 37, 101}

type BloomFilter struct {
	BitMap   *bitMap.BitMap
}

func NewBloomFilter(n int) *BloomFilter {
	return &BloomFilter {
		BitMap:bitMap.NewBitMap(n),
	}
}

func (bf BloomFilter) Set(value string) {
	for i := 0; i < len(cap); i++ {
		bf.BitMap.Set(hash(value,i))
	}
}

func (bf BloomFilter) Exist(value string) bool {
	for i := 0; i < len(cap); i++ {
		if !bf.BitMap.Get(hash(value,i)) {
			return false
		}
	}
	return true
}

func hash(s string,index int) uint {
	bit := uint(1)
	for i := 0; i < len(s); i++ {
		bit = (bit * cap[index] + (uint(s[i] - 'a') + uint(1))) % mod[index]
	}
	return bit
}