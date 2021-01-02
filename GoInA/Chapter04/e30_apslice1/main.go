// 使用append向切片增加元素
package main

import "fmt"

func main() {
	// 创建一个整型切片，其长度和容量都是5个元素
	slice := []int{10, 20, 30, 40, 50}

	// 创建一个新切片，其长度为2个元素，容量为4个元素
	newSlice := slice[1:3]
	fmt.Printf("slice: %d\n", slice)
	fmt.Printf("newSlice: %d\n", newSlice)

	// 使用原有的容量来分配一个新元素，将新元素赋值为60
	newSlice = append(newSlice, 60)
	fmt.Printf("slice: %d\n", slice)
	fmt.Printf("newSlice: %d\n", newSlice)
}

/*
>>> Execution Result:
slice: [10 20 30 40 50]
newSlice: [20 30]
slice: [10 20 30 60 50]
newSlice: [20 30 60]
*/
