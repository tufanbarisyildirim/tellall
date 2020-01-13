package dht

import "sync"
import jump "github.com/lithammer/go-jump-consistent-hash"

type HashTable struct {
	Bucket []chan interface{}
	m      *sync.Mutex
}

func NewHashTable() *HashTable {
	return &HashTable{
		Bucket: make([]chan interface{}, 0),
		m:      &sync.Mutex{},
	}
}

func (h *HashTable) Add(ch chan interface{}) {
	h.m.Lock()
	h.Bucket = append(h.Bucket, ch)
	h.m.Unlock()
}

func (h *HashTable) Remove(ch chan interface{}) {
	h.m.Lock()
	i := 0
	for _, c := range h.Bucket {
		if c != ch {
			h.Bucket[i] = c
			i++
		}
	}
	h.Bucket = h.Bucket[:i]
	h.m.Unlock()
}

func (h *HashTable) Find(num uint64) chan interface{} {
	return h.Bucket[jump.Hash(num, int32(len(h.Bucket)))]
}

func (h *HashTable) FindKey(key string) chan interface{} {
	return h.Bucket[jump.HashString(key, int32(len(h.Bucket)), jump.NewCRC64())]
}

func (h *HashTable) Push(num uint64, data interface{}) {
	h.Find(num) <- data
}

func (h *HashTable) PushKey(key string, data interface{}) {
	h.FindKey(key) <- data
}
