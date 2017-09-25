//Guessing game
//Author Jeremy Yon G00330435

package main

import (
    "fmt"
)

func main(){
    
    numbers := []int{1,2,3,4,5}

    getSmallestandLargest(numbers)





    // var input int

    // for input != -9999{
    //     var input int
    //     fmt.Print("\nEnter Integer: ")
    //     _, err := fmt.Scan(&input)
    //     if err != nil {
    //         fmt.Println(err)
    //         return
    //     }

    //     numbers = append(numbers, input)

    // }

   



}

func getSmallestandLargest(numbers []int) {

    var smallest = 9999
    var largest = -9999

    for i:=0;i<len(numbers);i++{
        if numbers[i]<smallest{
            smallest = numbers[i]
        }
        if numbers[i]>largest{
            largest = numbers[i]
        }
    }

    fmt.Println(smallest)
    fmt.Println(largest)

}