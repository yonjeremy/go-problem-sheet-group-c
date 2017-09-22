//FizzBuzz code in GO language
//Author Jeremy Yon G00330435

package main

import (
    "fmt"
)

func main(){
	for i := 1; i <= 100; i++ {
        if ( (i % 5 == 0) && (i%3 == 0)){
        fmt.Println("FizzBuzz")
        } else if ( i % 3 == 0){
            fmt.Println("Fizz")
        } else if ( i % 5 == 0){
            fmt.Println("Buzz")
        }  
	    fmt.Println(i)

    }
}