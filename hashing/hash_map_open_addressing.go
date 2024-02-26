package hashing

import (
	"fmt"
	"strconv"
)

type hashMapOpenAddressing struct {
	size        int
	capacity    int
	loadThres   float64
	extendRatio int
	buckets     []pair
	removed     pair
}

func newHashMapOpenAddressing() *hashMapOpenAddressing {
	buckets := make([]pair, 4)
	return &hashMapOpenAddressing{
		size:        0,
		capacity:    4,
		loadThres:   2.0 / 3.0,
		extendRatio: 2,
		buckets:     buckets,
		removed: pair{
			key:   -1,
			value: "-1",
		},
	}
}

func (m *hashMapOpenAddressing) hashFunc(key int) int {
	return key % m.capacity
}

func (m *hashMapOpenAddressing) loadFactor() float64 {
	return float64(m.size) / float64(m.capacity)
}

func (m *hashMapOpenAddressing) get(key int) string {
	idx := m.hashFunc(key)
	for i := 0; i < m.capacity; i++ {
		j := (idx + i) % m.capacity
		if m.buckets[j] == (pair{}) {
			return ""
		}

		if m.buckets[j].key == key && m.buckets[j] != m.removed {
			return m.buckets[j].value
		}
	}
	return ""
}

func (m *hashMapOpenAddressing) put(key int, val string) {
	if m.loadFactor() > m.loadThres {
		m.extend()
	}
	idx := m.hashFunc(key)
	for i := 0; i < m.capacity; i++ {
		j := (idx + i) % m.capacity
		if m.buckets[j] == (pair{}) || m.buckets[j] == m.removed {
			m.buckets[j] = pair{key, val}
			m.size++
			return
		}
		if m.buckets[j].key == key {
			m.buckets[j].value = val
			return
		}
	}
}

func (m *hashMapOpenAddressing) remove(key int) {
	idx := m.hashFunc(key)
	for i := 0; i < m.capacity; i++ {
		j := (idx + i) % m.capacity
		if m.buckets[j] == (pair{}) {
			return
		}
		if m.buckets[j].key == key {
			m.buckets[j] = m.removed
			m.size++
		}
	}
}

func (m *hashMapOpenAddressing) extend() {
	tmpBuckets := make([]pair, m.capacity)
	copy(tmpBuckets, m.buckets)

	m.capacity *= m.extendRatio
	m.buckets = make([]pair, m.capacity)
	m.size = 0
	for _, p := range tmpBuckets {
		if p != (pair{}) && p != m.removed {
			m.put(p.key, p.value)
		}
	}
}

func (m *hashMapOpenAddressing) print() {
	for _, p := range m.buckets {
		if p != (pair{}) {
			fmt.Println(strconv.Itoa(p.key) + "->" + p.value)
		} else {
			fmt.Println("nil")
		}
	}
}
