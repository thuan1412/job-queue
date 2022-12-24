package main

import "fmt"

type I interface {
	f()
}

type s struct {
}

func (a s) f() {

}

func hello() {

}

func main() {
	fmt.Println("=============")
	hello()
}
