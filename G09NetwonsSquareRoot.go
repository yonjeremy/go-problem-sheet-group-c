//Program that uses Netwons Square root rule 
//Author Jeremy Yon G00330435

package main

import (
    "fmt"
	"math"
)

func main(){
    // get user input (float64)
    var input float64
	var z_next float64
	var z float64
    fmt.Print("Enter Number to get SquareRoot: ")
    fmt.Scanf("%f\n", &input)

	// loop til the z_next number doesnt change
	for z = 1.0; z != (z_next); z = calculateZNext(z,input){
		fmt.Printf("The next z is: %12.8f\n\n",z)
		z_next = z
	}

	// Compare actual value and estimated value of the square root of input
	fmt.Printf("The actual value of the square root is: %12.8f\n",math.Sqrt(input))

}

// function to calculate the next z
func calculateZNext(z float64,x float64) float64{
    // use Newton's formula to calculate next z number
	z_next := z - ((z*z - x) / (2 * z))
	return (toFixed(z_next,8))
}

// function to round the number up or down
func round(num float64) int {
    return int(num + math.Copysign(0.5, num))
}
// function to fix the z number to 8 decimal places
func toFixed(num float64, precision int) float64 {
    output := math.Pow(10, float64(precision))
    return float64(round(num * output)) / output
}