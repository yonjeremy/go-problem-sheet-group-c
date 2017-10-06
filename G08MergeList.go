//Merge List and Sort
//Author Jeremy Yon G00330435
//Algorithm from https://github.com/hugopeixoto/mergesort/blob/master/go/mergesort.go

package main

import (
    "fmt"
)

func main(){
    
    // Hardcoded slice array
    // User input goes in here
    // note that slice has to contain ints that are already sorted
    list1 := []int{1,3,5}
    list2 := []int{2,4,6}

    // Call function to merge slice
    mergeAndSortList(list1,list2)

}

func mergeAndSortList(list1 []int, list2 []int) {

    // initialise final merged list
    var mergedList = make([]int, len(list1) + len(list2))

    var i = 0;
    var j = 0;

    //Compare values from both lists, push the smaller value into the final merged list until one of the list is empty
    for i<len(list1) && j<len(list2) {
        if(list1[i] < list2[j]){
            mergedList[i+j] = list1[i]
            i++
        }else{
            mergedList[i+j] = list2[j]
            j++
        }
    }

    // push all leftover values from list 1 into merged list
    for i < len(list1){
        mergedList[i+j] = list1[i]
        i++
    }

    // push all leftover values from list 2 into merged list
    for j < len(list2){
        mergedList[i+j] = list2[j]
        j++
    } 

    // print out merged and sorted list
    for i=0;i<len(mergedList);i++{
        fmt.Println(mergedList[i])
    }   


}