package BloomFilter

import (
	"hash/fnv"
	"testing"
)

func TestInsert1(t *testing.T) {
	hash := func(element string, size uint64) uint64 {
		h := fnv.New32a()
		h.Write([]byte(element))
		return uint64(h.Sum32()) % uint64(size)
	}

	bloom := New([]func(element string, size uint64) uint64{hash}, 20)

	bloom.Add("TAATCCC")
	bloom.Add("AATCCCT")
	bloom.Add("ATCCCTT")

	if actualValue := bloom.Contains("TAATCCC"); !actualValue {
		t.Errorf("Got %v expected %v", actualValue, true)
	}

	if actualValue := bloom.Contains("AATCCCT"); !actualValue {
		t.Errorf("Got %v expected %v", actualValue, true)
	}

	if actualValue := bloom.Contains("ATCCCTT"); !actualValue {
		t.Errorf("Got %v expected %v", actualValue, true)
	}

	if actualValue := bloom.Contains("TGCATGC"); actualValue {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
}

func TestInsert2(t *testing.T) {
	hash := func(element int, size uint64) uint64 {
		return uint64(element) % uint64(size)
	}

	bloom := New([]func(element int, size uint64) uint64{hash}, 10)

	for i := 0; i < 9; i++ {
		bloom.Add(i)
	}

	for i := 0; i < 9; i++ {
		if actualValue := bloom.Contains(i); !actualValue {
			t.Errorf("Got %v expected %v", actualValue, true)
		}
	}

	if actualValue := bloom.Contains(9); actualValue {
		t.Errorf("Got %v expected %v", actualValue, false)
	}

	if actualValue := bloom.Contains(10); !actualValue {
		t.Errorf("Got %v expected %v", actualValue, true)
	}

	if actualValue := bloom.Contains(11); !actualValue {
		t.Errorf("Got %v expected %v", actualValue, true)
	}

}

func TestInsertNewFP(t *testing.T) {
	hash := func(element string, size uint64) uint64 {
		h := fnv.New32a()
		h.Write([]byte(element))
		return uint64(h.Sum32()) % uint64(size)
	}

	bloom := NewWithFalsePositiveRate([]func(element string, size uint64) uint64{hash}, 0.01, 3)

	bloom.Add("TAATCCC")
	bloom.Add("AATCCCT")
	bloom.Add("ATCCCTT")

	if actualValue := bloom.Contains("TAATCCC"); !actualValue {
		t.Errorf("Got %v expected %v", actualValue, true)
	}

	if actualValue := bloom.Contains("AATCCCT"); !actualValue {
		t.Errorf("Got %v expected %v", actualValue, true)
	}

	if actualValue := bloom.Contains("ATCCCTT"); !actualValue {
		t.Errorf("Got %v expected %v", actualValue, true)
	}

	if actualValue := bloom.Contains("TGCATGC"); actualValue {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
}

func TestInsertNewFP2(t *testing.T) {
	hash := func(element int, size uint64) uint64 {
		return uint64(element) % uint64(size)
	}

	bloom := NewWithFalsePositiveRate([]func(element int, size uint64) uint64{hash}, 0.01, 10)

	for i := 0; i < 9; i++ {
		bloom.Add(i)
	}

	for i := 0; i < 9; i++ {
		if actualValue := bloom.Contains(i); !actualValue {
			t.Errorf("Got %v expected %v", actualValue, true)
		}
	}

	if actualValue := bloom.Contains(9); actualValue {
		t.Errorf("Got %v expected %v", actualValue, false)
	}

	if actualValue := bloom.Contains(10); actualValue {
		t.Errorf("Got %v expected %v", actualValue, true)
	}

	if actualValue := bloom.Contains(11); actualValue {
		t.Errorf("Got %v expected %v", actualValue, true)
	}

}

func TestClear(t *testing.T) {
	hash := func(element int, size uint64) uint64 {
		return uint64(element) % uint64(size)
	}

	bloom := NewWithFalsePositiveRate([]func(element int, size uint64) uint64{hash}, 0.01, 10)

	for i := 0; i < 9; i++ {
		bloom.Add(i)
	}

	for i := 0; i < 9; i++ {
		if actualValue := bloom.Contains(i); !actualValue {
			t.Errorf("Got %v expected %v", actualValue, true)
		}
	}

	if actualValue := bloom.Contains(9); actualValue {
		t.Errorf("Got %v expected %v", actualValue, false)
	}

	if actualValue := bloom.Contains(10); actualValue {
		t.Errorf("Got %v expected %v", actualValue, true)
	}

	if actualValue := bloom.Contains(11); actualValue {
		t.Errorf("Got %v expected %v", actualValue, true)
	}

	if actualValue := bloom.bits; actualValue.String() == "0" {
		t.Errorf("Got %v expected %v", actualValue, 0)
	}

	if actualValue := bloom.size; actualValue == 0 {
		t.Errorf("Got %v not expected %v", actualValue, 0)
	}

	bloom.Clear()

	if actualValue := bloom.bits; actualValue.String() != "0" {
		t.Errorf("Got %v expected %v", actualValue, "0")
	}

	if actualValue := bloom.size; actualValue != 0 {
		t.Errorf("Got %v expected %v", actualValue, 0)
	}

}

func TestSize(t *testing.T) {
	hash := func(element int, size uint64) uint64 {
		return uint64(element) % uint64(size)
	}

	bloom := NewWithFalsePositiveRate([]func(element int, size uint64) uint64{hash}, 0.01, 10)

	for i := 0; i < 9; i++ {
		bloom.Add(i)
	}

	for i := 0; i < 9; i++ {
		if actualValue := bloom.Contains(i); !actualValue {
			t.Errorf("Got %v expected %v", actualValue, true)
		}
	}

	if actualValue := bloom.size; actualValue != 9 {
		t.Errorf("Got %v expected %v", actualValue, 9)
	}
}

func TestNew(t *testing.T) {
	hashFuncs := []func(element int, size uint64) uint64{
		func(element int, size uint64) uint64 { return uint64(element) },
		func(element int, size uint64) uint64 { return uint64(element) * 2 },
	}
	bf := New(hashFuncs, 100)

	if bf.bits.String() != "0" {
		t.Errorf("Expected bits to be 0, got %s", bf.bits)
	}
	if bf.m != 100 {
		t.Errorf("Expected m to be 100, got %d", bf.m)
	}
	if len(bf.hashFuncs) != 2 {
		t.Errorf("Expected 2 hash functions, got %d", len(bf.hashFuncs))
	}
}

func TestNewWithFalsePositiveRate(t *testing.T) {
	hashFuncs := []func(element int, size uint64) uint64{
		func(element int, size uint64) uint64 { return uint64(element) },
		func(element int, size uint64) uint64 { return uint64(element) * 2 },
	}
	bf := NewWithFalsePositiveRate(hashFuncs, 0.01, 1000)

	if bf.bits.String() != "0" {
		t.Errorf("Expected bits to be 0, got %s", bf.bits)
	}
	if bf.m == 0 {
		t.Errorf("Expected m to be non-zero, got %d", bf.m)
	}
	if len(bf.hashFuncs) != 2 {
		t.Errorf("Expected 2 hash functions, got %d", len(bf.hashFuncs))
	}
}

func TestContains(t *testing.T) {
	bf := New([]func(element int, size uint64) uint64{
		func(element int, size uint64) uint64 { return uint64(element) },
	}, 100)
	bf.Add(1)

	if !bf.Contains(1) {
		t.Error("Expected Contains to return true for 1")
	}
	if bf.Contains(2) {
		t.Error("Expected Contains to return false for 2")
	}
}

func TestSize2(t *testing.T) {
	bf := New([]func(element int, size uint64) uint64{
		func(element int, size uint64) uint64 { return uint64(element) },
	}, 100)
	bf.Add(1)
	bf.Add(2)

	if bf.Size() != 2 {
		t.Errorf("Expected size to be 2, got %d", bf.Size())
	}
}

func TestClear2(t *testing.T) {
	bf := New([]func(element int, size uint64) uint64{
		func(element int, size uint64) uint64 { return uint64(element) },
	}, 100)
	bf.Add(1)
	bf.Add(2)
	bf.Clear()

	if bf.bits.String() != "0" {
		t.Errorf("Expected bits to be 0, got %s", bf.bits)
	}
	if bf.size != 0 {
		t.Errorf("Expected size to be 0, got %d", bf.size)
	}
}
