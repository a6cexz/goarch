package main

import (
	"fmt"
	"os"
	"reflectflags/rflags"
)

func main() {
	f := Flags{}
	fmt.Println(rflags.ParseFlags(&f, os.Args[1:]))
	fmt.Println(f)
}

type Flags struct {
	Source string `rflag:"source,s,src"`
	Debug  bool   `rflag:"debug,d"`
	Output string
	Count  int `rflag:"count"`
}
