//Guessing game
//Author Jeremy Yon G00330435

package main

import (
    "fmt"
)

func main(){
    
    var input string
    fmt.Println("Enter String to reverse it:")
    fmt.Scanf("%s\n", &input)

    fmt.Printf("The reversed string is: %s",reverse(input))

}

// func reverseString(input string) string{
//     last := len(input) - 1
// 	var temp byte
   
//     for i :=0; i < (len(input))/2;i++{
//         temp = input[i]
// 		input[i] = input[last-i]
// 		input[last-i] = temp
//     }

//     return input

// }

func reverse(s string) string {
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}