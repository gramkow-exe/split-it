package main

import (
	"github.com/gramkow-dev/split-it/app"
)

func main() {
	err := app.SetupAndRunApp()
	if err != nil {
		panic(err)
	}
}
