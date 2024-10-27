package main

import (
	"os"
	"time"

	"github.com/63square/cm2go2/block"
)

func main() {
	creation := block.NewCreation()

	start := time.Now()

	var x float64
	var y float64
	var z float64

	for x = 0; x < 100; x++ {
		for y = 0; y < 100; y++ {
			for z = 0; z < 100; z++ {
				creation.Add(block.AND, x, y, z, []float64{})
			}
		}
	}

	t := time.Now()
	elapsed := t.Sub(start)
	println(elapsed)

	creation.Compile(os.Stdout)
}
