package hashing

import (
	"fmt"
	"testing"
)

func TestArrayHashMap(t *testing.T) {
	hmap := newArrayHashMap()
	hmap.put(12836, "A")
	hmap.put(15937, "B")
	hmap.put(16750, "C")
	hmap.put(13276, "D")
	hmap.put(10583, "E")
	hmap.print()

	name := hmap.get(15937)
	fmt.Println(name)
	if name != "B" {
		t.Error("expected: B, got: ", name)
	}

	hmap.remove(10583)
	fmt.Println("remove ok")
	name = hmap.get(10583)
	fmt.Println(name)
	if name != "Not Found" {
		t.Error("expected: Not Found, got: ", name)
	}

	for _, kv := range hmap.pairSet() {
		fmt.Println(kv.key, "->", kv.value)
	}

	for _, k := range hmap.keySet() {
		fmt.Println(k)
	}

	for _, v := range hmap.valueSet() {
		fmt.Println(v)
	}
}
