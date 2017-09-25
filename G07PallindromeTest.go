//Guessing game
//Author Jeremy Yon G00330435

package main

import (
    "fmt"
)

func main(){
    
    var input string
    fmt.Println("Enter String to test for Palindrome:")
    fmt.Scanf("%s\n", &input)

    fmt.Printf("Test input for Palindrome: %t",testPalindrome(input))

}

func testPalindrome(input string) bool{
    last := len(input) - 1
   
    for i :=0; i < (len(input))/2;i++{
        if input[i] != input[last - i]{
            return false
        }
    }

    return true

}