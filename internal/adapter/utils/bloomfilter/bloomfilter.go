package bloomfilter

import (
	"hash"
	"hash/fnv"
	"sync"
)

// BloomFilter 表示一个布隆过滤器数据结构。
type BloomFilter struct {
	bitset       []byte
	bitsetLength uint32
	numHash      int
	mu           sync.RWMutex
	hasher       hash.Hash32
}

// NewBloomFilter 创建一个新的布隆过滤器，指定大小。
func NewBloomFilter(size int, numHash int) *BloomFilter {

	// 计算需要的字节数
	byteSize := (size + 7) / 8

	return &BloomFilter{
		bitset:       make([]byte, byteSize),
		numHash:      numHash,
		mu:           sync.RWMutex{},
		hasher:       fnv.New32a(),
		bitsetLength: uint32(byteSize * 8),
	}
}

// Add 将元素添加到布隆过滤器中。
func (bf *BloomFilter) Add(element []byte) {
	bf.mu.Lock()
	defer bf.mu.Unlock()

	for _, hashValue := range bf.hash(element) {
		v := hashValue % bf.bitsetLength
		byteIndex := v / 8
		bitIndex := v % 8
		bf.bitset[byteIndex] |= 1 << bitIndex
	}
}

// Contains 检查元素是否可能在布隆过滤器中。
func (bf *BloomFilter) Contains(element []byte) bool {
	bf.mu.RLock()
	defer bf.mu.RUnlock()
	for _, hashValue := range bf.hash(element) {
		v := hashValue % bf.bitsetLength
		byteIndex := v / 8
		bitIndex := v % 8
		if (bf.bitset[byteIndex] & (1 << bitIndex)) == 0 {
			return false
		}
	}
	return true
}

func (bf *BloomFilter) hash(b []byte) []uint32 {
	bf.hasher.Reset()
	_, _ = bf.hasher.Write(b)
	var result = make([]uint32, bf.numHash)
	for i := range 5 {
		_, _ = bf.hasher.Write([]byte{uint8(i)})
		result[i] = bf.hasher.Sum32()
	}
	return result
}
