package main

import (
	"fmt"
)

//https://www.coursera.org/learn/golang-functions-methods/peer/UxoIW/module-1-activity-bubble-sort-program

func BubbleSort(my_slice []int) {
	for j:=0; j<len(my_slice); j++ {
		for pos:=0; pos<(len(my_slice)-1); pos++ {
			Swap(my_slice, pos)
		}
	}
}

func Swap(my_slice []int, pos int) {
	if my_slice[pos] > my_slice[pos+1] {
		my_slice[pos], my_slice[pos+1] = my_slice[pos+1], my_slice[pos]
	}
}

func main(){
	length := 0
	fmt.Println("Please enter the length of numbers, and then numbers you'd like to sort: ")
	fmt.Scanln(&length)
	numbers := make([]int, length)
	for i:=0; i<length; i++ {
		fmt.Scanln(&numbers[i])
	}
	BubbleSort(numbers)
	for j:=0; j<len(numbers); j++ {
		fmt.Print(numbers[j], " ")
	}
}