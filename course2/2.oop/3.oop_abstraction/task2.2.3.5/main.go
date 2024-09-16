package main

import (
	"encoding/binary"
	"fmt"
	"hash"
	"hash/crc32"
	"hash/crc64"
)

type HashOption func(hm *HashMap)

type HashMap struct {
	Hasher hash.Hash
	keys   []*Data
}

type Data struct {
	key   string
	value interface{}
	next  *Data
}

type HashMaper interface {
	Set(key string, value interface{}) error
	Get(key string) (interface{}, bool)
}

func WithHashCRC64() HashOption {
	return func(hm *HashMap) {
		table := crc64.MakeTable(crc64.ISO)
		hm.Hasher = crc64.New(table)
	}
}

func WithHashCRC32() HashOption {
	return func(hm *HashMap) {
		hm.Hasher = crc32.NewIEEE()
	}
}

func NewHashMap(cap int, option HashOption) *HashMap {
	data := make([]*Data, cap)
	hm := &HashMap{keys: data}
	option(hm)
	return hm
}

func (hm HashMap) Set(key string, value interface{}) error {
	hashValue := CalculateHash(hm.Hasher, key)
	keysCapacity := uint64(cap(hm.keys))

	if keysCapacity == 0 {
		return fmt.Errorf("zero capacity of keys slice")
	}

	hashValue = (hashValue / 100000) % keysCapacity
	newData := &Data{key: key, value: value, next: nil}

	if hm.keys[hashValue] == nil {
		hm.keys[hashValue] = newData
		return nil
	}

	checkData := hm.keys[hashValue]
	for {
		if checkData.key == key {
			*checkData = *newData
			return nil
		}

		if checkData.next == nil {
			(*checkData).next = newData
			return nil
		}

		checkData = checkData.next
	}
}

func (hm HashMap) Get(key string) (interface{}, bool) {
	hashValue := CalculateHash(hm.Hasher, key)
	keysCapacity := uint64(cap(hm.keys))

	if keysCapacity == 0 {
		return nil, false
	}

	hashValue = (hashValue / 100000) % keysCapacity

	if hm.keys[hashValue] == nil {
		return nil, false
	}

	checkData := hm.keys[hashValue]
	for {
		if checkData.key == key {
			return checkData.value, true
		}

		if checkData.next == nil {
			return nil, false
		}

		checkData = checkData.next
	}
}

func CalculateHash(h hash.Hash, input string) uint64 {
	h.Reset()
	h.Write([]byte(input))
	return uint64(binary.BigEndian.Uint32(h.Sum(nil)))
}

func main() {

}
