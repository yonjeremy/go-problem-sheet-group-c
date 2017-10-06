// Get Largest and Smallest value in an array
//Author Jeremy Yon G00330435

package main

import (
    "fmt"
)

func main(){
    
    // Hardcoded user array (only works with integers)
    numbers := []int{1,2,3,4,5}

    // call function to get smallest and largest number
    getSmallestandLargest(numbers)

}

func getSmallestandLargest(numbers [] int) {

    // hardcode default smallest and largest number
    var smallest = 999999
    var largest = -999999

    // loop over array
    for i:=0;i<len(numbers);i++{

        // assign new smallest and largest number
        if numbers[i]<smallest{
            smallest = numbers[i]
        }
        if numbers[i]>largest{
            largest = numbers[i]
        }
    }

    // Print results
    fmt.Printf("The smallest number is %d\n",smallest)
    fmt.Printf("The largest number is %d",largest)

}