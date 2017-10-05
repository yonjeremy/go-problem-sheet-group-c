//Factorial digit sum
//Author Jeremy Yon G00330435

package main

import (
    "fmt"
    "math/big"
    "strconv"
    "strings"
)


func main(){
	bigInt := factorial(100)
    bigStr := bigInt.String()


    a := strings.Split(bigStr, "")
    totalSum := 0
    for _, digit := range a {
        sum, _ := strconv.Atoi(digit)
        totalSum += sum
    }

    
    fmt.Println(totalSum);
}

func factorial(n int64) *big.Int{
    // num := big.NewInt(1)
    // i := big.NewInt(1)
    // for i = 1; i <= factorial; i++{
    //     num.Mul(i,num)
    // }

    // sum := 0;
    // for (num > 0) {
    //     sum = sum + num % 10;
    //     num = num / 10;
    // }
    // fmt.Println(sum);

    // return sum

    if n < 0{
        return big.NewInt(1)
    }
    if n == 0{
        return big.NewInt(1)
    }
    bigN:= big.NewInt(n)
    return bigN.Mul(bigN,factorial(n-1))
}