//Merge List and Sort
//Author Jeremy Yon G00330435

package main

import (
    "fmt"
)

func main(){
    
    list1 := []int{1,3,5,3443}
    list2 := []int{2,4,6}

    mergeAndSortList(list1,list2)





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

func mergeAndSortList(list1 []int, list2 []int) {

    var mergedList = make([]int, len(list1) + len(list2))
    var i = 0;
    var j = 0;

    for i<len(list1) && j<len(list2) {
        if(list1[i] < list2[j]){
            mergedList[i+j] = list1[i]
            i++
        }else{
            mergedList[i+j] = list2[j]
            j++
        }
    }

    for i < len(list1){
        mergedList[i+j] = list1[i]
        i++
    }

    for j < len(list2){
        mergedList[i+j] = list2[j]
        j++
    } 

    for i=0;i<len(mergedList);i++{
        fmt.Println(mergedList[i])
    }   


}