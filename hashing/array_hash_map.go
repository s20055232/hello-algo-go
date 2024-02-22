package hashing

import "fmt"

type pair struct {
	key   int
	value string
}

type arrayHashMap struct {
	buckets []*pair
}

func newArrayHashMap() *arrayHashMap {
	return &arrayHashMap{
		buckets: make([]*pair, 100),
	}
}

func (a *arrayHashMap) hashFunc(key int) int {
	index := key % 100
	return index
}

func (a *arrayHashMap) put(key int, value string) {
	idx := a.hashFunc(key)
	a.buckets[idx] = &pair{key, value}
}

func (a *arrayHashMap) get(key int) string {
	idx := a.hashFunc(key)
	tmp := a.buckets[idx]
	if tmp != nil {
		return tmp.value
	}
	return "Not Found"
}

func (a *arrayHashMap) remove(key int) {
	idx := a.hashFunc(key)
	a.buckets[idx] = nil
}

func (a *arrayHashMap) keySet() []int {
	var keys []int
	for _, pair := range a.buckets {
		if pair != nil {
			keys = append(keys, pair.key)
		}
	}
	return keys
}

func (a *arrayHashMap) valueSet() []string {
	var values []string
	for _, pair := range a.buckets {
		if pair != nil {
			values = append(values, pair.value)
		}
	}
	return values
}

func (a *arrayHashMap) print() {
	for _, pair := range a.buckets {
		if pair != nil {
			fmt.Println(pair.key, "->", pair.value)
		}
	}

}

func (a *arrayHashMap) pairSet() []*pair {
	var result []*pair
	for _, pair := range a.buckets {
		if pair != nil {
			result = append(result, pair)
		}
	}
	return result
}
