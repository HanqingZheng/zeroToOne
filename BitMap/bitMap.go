package bitMap

import (
	"unsafe"
)

type BM int64

type BitMap struct {
	BitSlice []BM
	BitNum uint
}

var bmBitNum = uint(unsafe.Sizeof(BM(1)) * 8)

func NewBitMap(n int) *BitMap {
	bitNum := uint(n) / bmBitNum + 1
	return &BitMap {
		BitSlice : make([]BM,bitNum,bitNum),
		BitNum : uint(n),
	}
}

func (bm *BitMap) Set (n uint) {
	if n > bm.BitNum {
		return
	}
	byteIndex := n / bmBitNum
	bitIndex := n % bmBitNum
	bm.BitSlice[byteIndex] |= BM(uint(1) << bitIndex)
}

func (bm *BitMap) Get (n uint) bool{
	if n > bm.BitNum {
		return false
	}
	byteIndex := n / bmBitNum
	bitIndex := n % bmBitNum
	return (bm.BitSlice[byteIndex] & BM(uint(1) << bitIndex)) != 0
}