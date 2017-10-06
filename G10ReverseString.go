//Reverse String
//Author Jeremy Yon G00330435
//Code extracted from https://stackoverflow.com/questions/1752414/how-to-reverse-a-string-in-go

package main

import (
    "fmt"
)

func main(){
    // Prompt for user input String
    var input string
    fmt.Println("Enter String to reverse it:")
    fmt.Scanf("%s\n", &input)

    // call function to reverse string
    fmt.Printf("The reversed string is: %s",reverse(input))

}

func reverse(s string) string {
    // use rune to store characters and switch places with each other
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}

    // return reversed string
	return string(chars)
}