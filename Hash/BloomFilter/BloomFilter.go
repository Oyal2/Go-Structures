package BloomFilter

import (
	"math"
	"math/big"

	"github.com/oyal2/Go-Structures/utils"
)

type BloomFilter[T utils.Ordered] struct {
	hashFuncs []func(element T, size uint64) uint64
	bits      *big.Int
	m         uint64
	size      uint64
}

func New[T utils.Ordered](hashFunc []func(element T, size uint64) uint64, m uint64) *BloomFilter[T] {
	return &BloomFilter[T]{
		bits:      big.NewInt(0),
		m:         m,
		hashFuncs: hashFunc,
	}
}

func NewWithFalsePositiveRate[T utils.Ordered](hashFunc []func(element T, size uint64) uint64, falsePositive float64, n uint64) *BloomFilter[T] {
	return &BloomFilter[T]{
		bits:      big.NewInt(0),
		m:         uint64(math.Ceil((float64(n) * math.Log(falsePositive)) / (math.Log(1 / math.Pow(2, math.Log(2)))))),
		hashFuncs: hashFunc,
	}
}

func (BloomFilter *BloomFilter[T]) Add(element T) {
	BloomFilter.size++
	val := new(big.Int)
	val.SetUint64(BloomFilter.hash(element))
	val = val.Lsh(big.NewInt(1), uint(val.Uint64()))
	BloomFilter.bits = BloomFilter.bits.Or(BloomFilter.bits, val)
}

func (BloomFilter *BloomFilter[T]) Contains(element T) bool {
	val := new(big.Int)
	val.SetUint64(BloomFilter.hash(element))
	val = val.Lsh(big.NewInt(1), uint(val.Uint64()))
	val = val.And(BloomFilter.bits, val)

	return val.String() != "0"
}

func (BloomFilter *BloomFilter[T]) Size() uint64 {
	return BloomFilter.size
}

func (BloomFilter *BloomFilter[T]) Clear() {
	BloomFilter.bits = big.NewInt(0)
	BloomFilter.size = 0
}

func (BloomFilter *BloomFilter[T]) hash(element T) uint64 {
	value := BloomFilter.hashFuncs[0](element, BloomFilter.m)
	for i := 1; i < len(BloomFilter.hashFuncs); i++ {
		value = BloomFilter.hashFuncs[i](T(rune(value)), BloomFilter.m)
	}
	return value
}
