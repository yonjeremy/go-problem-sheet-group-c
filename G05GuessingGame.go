//Guessing game
//Author Jeremy Yon G00330435

package main

import (
    "time"
    "fmt"
    "math/rand"
)

func main(){
    // seeds the random number with currect time
    rand.Seed(time.Now().UTC().UnixNano())
    //generate secret number
	secretNum := rand.Intn(100)

    // gets the number of tries taken for the user to guess the secret number
    var numOfTries = 1
    // variable for user input
    var input int

    // prompt and get user input
    fmt.Print("\nEnter guess (integer) and press enter: ")
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println(err)
		return
	}

    // keep looping until the secret number is guessed
    for (input != secretNum){
        // checks if user input is higher or lower than secret number
        if input > secretNum{
            fmt.Println("Your input is too high!")
        }else{
            fmt.Println("Your input is too low!")
        }

    // keep looping
    fmt.Print("\nEnter guess (integer) and press enter: ")
        _, err := fmt.Scan(&input)
        if err != nil {
            fmt.Println(err)
            return
        }
        // adds the number of tries taken
        numOfTries++
    }

    // print results
	fmt.Printf("You took %d tries to guess the secret number\n",numOfTries)	
    fmt.Printf("The secret number is %d",secretNum)

}