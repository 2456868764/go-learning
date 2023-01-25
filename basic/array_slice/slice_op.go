package main

import (
	"errors"
	"fmt"
)

func main() {
	s := []int{1, 2, 4, 7}
	fmt.Printf("a = %v\n", s)
	// 结果应该是 5, 1, 2, 4, 7
	s, _ = Add(s, 0, 5)
	fmt.Printf("a = %v\n", s)
	// 结果应该是5, 9, 1, 2, 4, 7
	s, _ = Add(s, 1, 9)
	fmt.Printf("a = %v\n", s)
	// 结果应该是5, 9, 1, 2, 4, 7, 13
	s, _ = Add(s, 6, 13)
	fmt.Printf("a = %v\n", s)
	// 结果应该是5, 9, 2, 4, 7, 13
	s, _ = Remove(s, 2)
	fmt.Printf("a = %v\n", s)
	// 结果应该是9, 2, 4, 7, 13
	s, _ = Remove(s, 0)
	fmt.Printf("a = %v\n", s)
	// 结果应该是9, 2, 4, 7
	s, _ = Remove(s, 4)
	fmt.Printf("a = %v\n", s)
}

func Add(arr []int, index int, value int) ([]int, error) {
	if index > len(arr) || index < 0 {
		return arr, errors.New("index exceed slice length")
	}

	if index == 0 {
		arr = append([]int{value}, arr[:]...)
		return arr, nil
	}
	if index == len(arr)-1 {
		return append(arr, value), nil
	}

	temp := append(arr[:index], value)
	arr = append(temp, arr[index:]...)
	return arr, nil

	//s = append(s, zero_value)
	//copy(s[i+1:], s[i:])
	//s[i] = x
}

func Remove(arr []int, index int) ([]int, error) {
	if index >= len(arr) || index < 0 {
		return nil, errors.New("index exceed slice length")
	}

	if index == len(arr)-1 {
		return arr[:index-1], nil
	}
	if index == 0 {
		return arr[1:], nil
	}
	return append(arr[:index], arr[index+1:]...), nil
}
