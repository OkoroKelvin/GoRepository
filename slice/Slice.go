package main

import "fmt"

func main() {
	slice := []int{1, 2, 3}
	fmt.Println(slice)
	slice = append(slice, 4, 34, 98)
	fmt.Println(slice)
	s2 := slice[1:]
	s3 := slice[:2]
	s4 := slice[1:2]
	s2[0] = 1000

	fmt.Printf("slice %v %T \n", s2, s2)
	fmt.Println(s2, s3, s4)

	arr := [4]int{1, 2, 3, 4}

	sArr := arr[:]
	sArr[1] = 20
	fmt.Println(arr)
	fmt.Println(sArr)
	fmt.Printf("%T\n", arr)
	fmt.Printf("%T\n", sArr)
	sArr = append(sArr, 2)
	fmt.Printf("%T\n", sArr)

	fmt.Println(sArr)
	fmt.Println(arr)

}
