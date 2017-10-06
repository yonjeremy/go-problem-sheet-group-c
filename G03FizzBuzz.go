//FizzBuzz code in GO language
//Author Jeremy Yon G00330435

package main

import (
    "fmt"
)

func main(){
    // Loops 100 times
	for i := 1; i <= 100; i++ {

        // print FIZZBUZZ if multiples of 3 and 5
        if ( (i % 5 == 0) && (i%3 == 0)){
        fmt.Println("FizzBuzz")
        } else if ( i % 3 == 0){ // print Fizz if multiples of 3
            fmt.Println("Fizz")
        } else if ( i % 5 == 0){ // print Buzz if multiples of 5
            fmt.Println("Buzz")
        }

        // print results to console  
	    fmt.Println(i)

    }
}