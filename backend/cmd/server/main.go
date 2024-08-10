package main

import (
	"log/slog"

	"github.com/ayahiro1729/onpu/api/controller"
)

func main() {
	s, err := controller.NewServer()
	if err != nil {
		slog.Error(err.Error())
	}

	if err := s.Run(":8080"); err != nil {
		slog.Error(err.Error())
	}
}
