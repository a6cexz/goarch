package main

import (
	"fmt"
)

func main() {

	h := NewHashMap()

	h.Set(5, "test5")
	h.Set(15, "test15")
	h.Set(3, "test3")
	h.Set(3, "test3")
	h.Set(3, "test3")
	h.Set(3, "test3")
	h.Set(3, "test3")
	h.Set(3, "test30")
	h.Set(6, "test6")
	h.Set(7, "test7")
	h.Set(8, "test8")
	h.Set(9, "test9")
	h.Set(10, "test10")
	h.Set(11, "test11")
	h.Set(12, "test12")
	h.Set(13, "test13")

	fmt.Println(h.Get(13))
	fmt.Println(h.Delete(13))
	fmt.Println(h.Get(13))
}
