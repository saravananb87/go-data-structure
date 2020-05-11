package main

import (
    "fmt"
)

func doBubbleSort(items []int){

	runRequired := true	

	for runRequired {
		
		isSwap :=false

		for i:=0; i < len(items)-1; i++{

			if( items[i+1] < items[i]){
				items[i] , items[i+1] = items[i+1], items[i]
				isSwap= true
			}
		} 
		if(!isSwap){
			runRequired = false;
		}
	}

	fmt.Println(items)

}

func main() {
	items := []int{10,11,17,15,11,17,19,20,9,5,1,2,5,6,4,3,8,1,3,5,7}
    doBubbleSort(items)
}