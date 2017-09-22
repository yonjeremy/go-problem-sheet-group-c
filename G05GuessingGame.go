//Guessing game
//Author Jeremy Yon G00330435

package main

import (
    "time"
    "fmt"
    "math/rand"
)

func main(){
    rand.Seed(time.Now().UTC().UnixNano())
	secretNum := rand.Intn(100)

    var numOfTries = 1
    var input int
    fmt.Print("\nEnter some text and press enter: ")
    // using fmt.Scan, we can read single words in ascii string
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println(err)
		return
	}

    for (input != secretNum){
        if input > secretNum{
            fmt.Println("Your input is too high!")
        }else{
            fmt.Println("Your input is too low!")
        }

        fmt.Print("\nEnter some text and press enter: ")
        // using fmt.Scan, we can read single words in ascii string
        _, err := fmt.Scan(&input)
        if err != nil {
            fmt.Println(err)
            return
        }
        numOfTries++
    }
	fmt.Println(numOfTries)	
    fmt.Println(secretNum)

}