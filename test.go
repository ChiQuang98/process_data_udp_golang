package main

import "bytes"

func main() {
	s1 := []byte("tran chis quang")
	println(bytes.Contains(s1, []byte("chi")))
}
