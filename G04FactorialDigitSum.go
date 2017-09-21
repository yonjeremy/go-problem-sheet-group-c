//Factorial digit sum
//Author Jeremy Yon G00330435

package main

import (
    "fmt"
    "math/big"
)


func main(){
	fmt.Println(getSumOfFactorial(20))
}

func getSumOfFactorial(factorial int) int{
    num := big.NewInt(1)
    i := big.NewInt(1)
    for i = 1; i <= factorial; i++{
        num.Mul(i,num)
    }

    sum := 0;
    for (num > 0) {
        sum = sum + num % 10;
        num = num / 10;
    }
    fmt.Println(sum);

    return sum
}