package main

import (
	"os"

	"github.com/ryan2/planner/internal/gart"
)

func main() {
	f, err := os.Create("test.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	gart.GeneratePNG(f, 1080, 1350)
}
