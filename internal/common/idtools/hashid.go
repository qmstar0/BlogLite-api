package idtools

import (
	"hash"
	"hash/fnv"
	"sync"
)

var fnvPool = sync.Pool{New: func() any {
	return fnv.New32()
}}

func NewHashID(s []byte) uint32 {
	hash32 := fnvPool.Get().(hash.Hash32)
	defer func() {
		hash32.Reset()
		fnvPool.Put(hash32)
	}()
	_, _ = hash32.Write(s)
	return hash32.Sum32()
}
