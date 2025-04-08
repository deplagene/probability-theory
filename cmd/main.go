package main

import (
	"homework/probability/cmd/gui"
	"homework/probability/services/probability"
)

func main() {
	s := probability.NewService()
	g := gui.NewGui(s)
	g.Run()
}