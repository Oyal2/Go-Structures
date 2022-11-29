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
