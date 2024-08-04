package main

import (
	"github.com/ayahiro1729/onpu/api/controller"
)

func main() {
	s, err := controller.NewServer()
	if err != nil {
		panic(err)
	}
	s.Run()
}
