// golang squareroot calculation, z in the loop can be changed, as well as the amount of loop repetitions to change accuracy

package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
z := 1.0
for i := 0;i<10 ; i++ {
z -= (z*z - x) / (2*z)
}
return z
}

func main() {
	fmt.Println(Sqrt(2))
}
