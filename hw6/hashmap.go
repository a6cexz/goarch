package main

import "fmt"

const scopeSize = 2

type Cell struct {
	originalKey int
	value       interface{}
}

type list []Cell

type HashMap struct {
	scope      []list
	collisions int
	length     int
}

func NewHashMap() *HashMap {
	return newHashMap(scopeSize)
}

func newHashMap(size int) *HashMap {
	h := HashMap{}

	h.scope = make([]list, size)
	for i := range h.scope {
		h.scope[i] = make(list, 0)
	}

	return &h
}

func (h *HashMap) extend() {
	nh := newHashMap(len(h.scope) * len(h.scope))

	for _, list := range h.scope {
		for _, item := range list {
			nh.Set(item.originalKey, item.value)
		}
	}

	h.scope = nh.scope
	h.length = nh.length
	h.collisions = nh.collisions
}

func (h *HashMap) hash(key int) int {
	return key % len(h.scope)
}

func (h *HashMap) Set(key int, a interface{}) {
	hash := h.hash(key)
	list := h.scope[hash]

	for i := range list {
		if list[i].originalKey == key {
			list[i].value = a
			h.scope[hash] = list
			return
		}
	}

	if len(list) != 0 {
		h.collisions++
	}

	list = append(list, Cell{
		originalKey: key,
		value:       a,
	})
	h.scope[hash] = list

	h.length++

	fmt.Printf("LOG: LENGTH: %d, COLLISIONS: %d\n", h.length, h.collisions)
	if h.length/2 > len(h.scope) || h.collisions > h.length/2 {
		fmt.Println("LOG: EXTENDED")
		h.extend()
		fmt.Println("LOG: ", h.scope)
	}
}

func (h *HashMap) Get(key int) (interface{}, bool) {
	hash := h.hash(key)
	list := h.scope[hash]

	fmt.Println(key, list)

	for i := range list {
		if list[i].originalKey == key {
			return list[i].value, true
		}
	}

	return nil, false
}

func (h *HashMap) Delete(key int) bool {
	hash := h.hash(key)
	list := h.scope[hash]

	newList := []Cell{}
	deleted := false
	for i := range list {
		if list[i].originalKey == key {
			deleted = true
		} else {
			newList = append(newList, list[i])
		}
	}
	if deleted {
		h.scope[hash] = newList
	}
	return deleted
}
