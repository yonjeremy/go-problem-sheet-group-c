//Factorial digit sum
//Author Jeremy Yon G00330435
//Part of code extracted from https://stackoverflow.com/questions/11270547/go-big-int-factorial-with-recursion

package main

import (
    "fmt"
    "math/big"
    "strconv"
    "strings"
)


func main(){
    // Harcoded user input of factorial to be summed ie:100
    // Calls the medthod factorial() recursively
	bigInt := factorial(100)

    // Convert result into a string
    bigStr := bigInt.String()

    // split the string into an array
    array := strings.Split(bigStr, "")
    totalSum := 0

    // loop through the array and add each digit to totalSum
    for _, digit := range array {
        sum, _ := strconv.Atoi(digit)
        totalSum += sum
    }

    // Print results to console
    fmt.Println(totalSum);
}

func factorial(n int64) *big.Int{
    // function that calls itself recursively until n is equal or less than 0
    if n < 0{
        return big.NewInt(1)
    }
    if n == 0{
        return big.NewInt(1)
    }
    bigN:= big.NewInt(n)
    return bigN.Mul(bigN,factorial(n-1))
}