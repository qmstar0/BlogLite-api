package idtools

import "hash/fnv"

func NewHashID(s string) uint32 {
	hash := fnv.New32()
	_, _ = hash.Write([]byte(s))
	return hash.Sum32()
}
