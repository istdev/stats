package main

import (
	"fmt"
	"time"

	"github.com/istdev/stats"
	"github.com/wiless/vlib"
)

func main() {
	t := time.Now()
	var matlab *vlib.Matlab
	x := vlib.RandNFVec(701, 1.0)
	// x = x.Scale(2)
	xx, yy := stats.CDF(x)
	matlab = vlib.NewMatlab("xx.m")
	matlab.Export("x", xx)
	matlab.Export("y", yy)
	matlab.Command("plot([-Inf x Inf],[0 0 y])")
	matlab.Close()
	fmt.Printf("time = %v", time.Since(t))
}
