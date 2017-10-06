//Test if user input String is a Pallindrome
//Author Jeremy Yon G00330435

package main

import (
    "fmt"
)

func main(){
    
    // Prompt for user input String
    var input string
    fmt.Println("Enter String to test for Palindrome:")
    fmt.Scanf("%s\n", &input)

    // Calls the test Pallindrome function
    fmt.Printf("Test input for Palindrome: %t",testPalindrome(input))

}

func testPalindrome(input string) bool{
    last := len(input) - 1
   
    // compares first char in array with last char, 2nd with 2nd to last char and so forth
    for i :=0; i < (len(input))/2;i++{
        if input[i] != input[last - i]{
            return false
        }
    }

    // prints true if all characters matched in order
    return true

}