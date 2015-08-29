package stats

import (
	"fmt"
	"sort"

	"github.com/wiless/vlib"
)

func CDF(xInput vlib.VectorF) (xOut, yOut vlib.VectorF) {
	n := xInput.Size()
	str := fmt.Sprintf("%f:%f:1", 1/float64(n), 1/float64(n))
	fmt.Println(str, n)
	yOut = vlib.ToVectorF(str)
	if yOut[yOut.Size()-1] < 1 {
		yOut = yOut.Insert(yOut.Size()-1, 1)
	}
	fmt.Println(yOut[0], yOut[yOut.Size()-1], yOut.Size())
	sort.Float64s([]float64(xInput))
	xOut = xInput
	// xx := vlib.NewVectorI(n)
	for i := 0; i < n-1; i++ {
		if xOut[i] == xOut[i+1] {
			xOut.Delete(i)
			yOut.Delete(i)
			i--
			n--
		}
	}
	yOut = yOut.Insert(0, 0)
	fmt.Println(yOut[0], yOut[yOut.Size()-1], yOut.Size())
	xOut1 := vlib.NewVectorF(2 * n)
	yOut1 := vlib.NewVectorF(2 * n)
	// fmt.Printf("\n %v \n", yOut)
	for i := 0; i < n; i++ {
		xOut1[2*i] = xOut[i]
		xOut1[2*i+1] = xOut[i]
		yOut1[2*i] = yOut[i+1]
		yOut1[2*i+1] = yOut[i+1]
	}

	return xOut1, yOut1
}
