package bloomfilter_test

import (
	"blog/pkg/bloomfilter"
	"testing"
)

func TestNewBloomFilter(t *testing.T) {

	data := [][]byte{
		[]byte("1"),
		[]byte("2"),
		[]byte("3"),
		[]byte("4"),
	}

	filter := bloomfilter.NewBloomFilter(500, 5)
	for i := range data {
		filter.Add(data[i])
	}

	if filter.Contains([]byte("5")) {
		t.Error("filter verify err")
	}
	t.Log("filter not contains `5`")

	if !filter.Contains([]byte("1")) {
		t.Error("filter verify err")
	}
	t.Log("filter contains `1`")

}

func BenchmarkBloomFilter_Contains(b *testing.B) {
	data := [][]byte{
		[]byte("1"),
		[]byte("2"),
		[]byte("3"),
		[]byte("4"),
	}

	filter := bloomfilter.NewBloomFilter(500, 5)
	for i := range data {
		filter.Add(data[i])
	}

	for i := 0; i < b.N; i++ {
		//filter.Contains([]byte("5"))
		filter.Contains([]byte("1"))
	}
}
