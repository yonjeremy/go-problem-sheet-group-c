//Guessing game
//Author Jeremy Yon G00330435

package main

import (
    "fmt"
	"math"
)

func main(){
    
    var input float64
	var z_next float64
    fmt.Print("Enter Number to get SquareRoot: ")
    fmt.Scanf("%f\n", &input)


	for z := 1.0; z != z_next; z = calculateZNext(z,input){
		fmt.Printf("The next z is: %12.8f\n",z)
		z_next = z
	}

	fmt.Printf("The actual value of the square root is: %12.8f",math.Sqrt(input))


}

func calculateZNext(z float64,x float64) float64{
    
	z_next := z - ((z*z - x) / (2 * z))
	return z_next
}