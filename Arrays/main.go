package main

import "fmt"

func main() {
	//1> create & initializing arr
	var arr [5]string
	arr1 := [...]int{1, 2, 3}
	arr3 := [...]int{9, 4, 3}
	arr4 := [...]int{1, 2, 3}
	fmt.Printf("%+v,%+v\n", arr, arr1)
	//2> access the values of the arr using for range
	fmt.Println("Loop through the arr1")
	for _, v := range arr1 {
		fmt.Printf("%3d", v)
	}
	fmt.Println()
	//3> array is a value type
	//create new arr2 and initialize with arr1
	arr2 := arr1
	fmt.Println("New array before change:", arr2)
	arr2[0] = 4
	fmt.Println("New array after change:", arr2)
	fmt.Println("The original arr1 after arr2 change:", arr1)

	//4> compare two arrays
	fmt.Println(arr1 == arr2)
	fmt.Println(arr1 == arr3)
	fmt.Println(arr1 == arr4)

	//5> copy an array into another array
	//create a copy arr by value -> not reflect to the original arr
	arr5 := arr1
	fmt.Println("arr5 before change:", arr5)
	arr5[0] = 100
	fmt.Println("arr5 after change:", arr5)
	fmt.Println("arr1 after arr5 change:", arr1)

	//create a copy arr by reference -> reflect to the original arr
	arr6 := &arr3
	fmt.Println("arr6 before change:", *arr6)
	arr6[0] = 20
	fmt.Println("arr6 after change:", *arr6)
	fmt.Println("arr1 after arr5 change:", arr3)

	//6> pass an array to a function
	var totalSum = calculateSumArr(arr1)
	fmt.Println("Total Of List Number:", totalSum)
}

func calculateSumArr(nums [3]int) int {
	var sum int
	for _, v := range nums {
		sum += v
	}
	return sum
}
