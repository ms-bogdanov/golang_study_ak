package main

import (
	"hash/crc32"
	"strconv"
	"testing"
)

type testCase struct {
	key   string
	value int
}

func TestSetAndGetKeyValue(t *testing.T) {
	m := NewHashMap(16, WithHashCRC64())
	err := m.Set("key", "values")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	value, ok := m.Get("key")
	if !ok {
		t.Fatalf("expected key to be found")
	}

	if value != "values" {
		t.Fatalf("expected value to be 'values', got %v", value)
	}
}

func TestGetNonExistentKey(t *testing.T) {
	m := NewHashMap(16, WithHashCRC32())
	_, ok := m.Get("nonexistent")
	if ok {
		t.Fatalf("expected key to not be found")
	}
}

func TestCalculateHashForStringInput(t *testing.T) {
	hasher := crc32.NewIEEE()
	input := "test"
	expected := uint64(3632233996)

	result := CalculateHash(hasher, input)

	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func MassTest(t *testing.T) {
	testCases := []testCase{}

	for i := 0; i < 100; i++ {
		tc := testCase{key: strconv.Itoa(i), value: i}
		testCases = append(testCases, tc)
	}

	m := NewHashMap(100, WithHashCRC64())

	for _, tc := range testCases {
		err := m.Set(tc.key, tc.value)
		if err != nil {
			t.Errorf("Dont expected error: %v", err)
		}
	}

	for _, tc := range testCases {
		result, ok := m.Get(tc.key)

		if !ok {
			t.Errorf("Didnt find value, that is in map")
		}

		if result != tc.value {
			t.Errorf("Expected %d, got %d", tc.value, result)
		}
	}

}
func Benchmark64(b *testing.B) {
	m := NewHashMap(1000, WithHashCRC64())
	for i := 0; i < 1000; i++ {
		m.Set(strconv.Itoa(i), i)
	}

	for i := 0; i < 1000; i++ {
		m.Get(strconv.Itoa(i))
	}
}

func Benchmark32(b *testing.B) {
	m := NewHashMap(1000, WithHashCRC32())
	for i := 0; i < 1000; i++ {
		m.Set(strconv.Itoa(i), i)
	}

	for i := 0; i < 1000; i++ {
		m.Get(strconv.Itoa(i))
	}
}
