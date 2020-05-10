package main

import "fmt"

func main() {

	items := []int{1, 2, 4, 5, 6, 8, 9, 10, 20, 30, 50, 33, 44}
	fmt.Println(IsExist(items, 11))
	fmt.Println(IsExistOptimised(items, 44))

}

//IsExist check item exist in array using linear
func IsExist(items []int, value int) bool {

	for _, v := range items {
		if v == value {
			return true
		}
	}
	return false
}

//IsExistOptimised linear loop optimization test
func IsExistOptimised(items []int, value int) bool {

	var middle int = (len(items) / 2) + 1
	for i := 1; i <= middle; i++ {
		if items[i-1] == value || items[len(items)-i] == value {
			return true
		}
	}
	return false
}
