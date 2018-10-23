// golang squareroot calculation, z in the loop can be changed, as well as the amount of loop repetitions to change accuracy
// This method is called Newton's method, more about it can be found on wikipeida https://en.wikipedia.org/wiki/Newton%27s_method

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
