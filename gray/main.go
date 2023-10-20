package main

import (
	"fmt"
	"strconv"
)

func convertBinaryToInt(bin string) int {
	i, _ := strconv.ParseInt(bin, 2, 64)
	return int(i)
}

func grayCode(n int) []int {
	if n <= 0 {
		return []int{}
	}

	result := make([]int, 0)
	arr := make([]string, 0)

	arr = append(arr, "0")
	arr = append(arr, "1")

	var i, j int
	for i = 2; i < (1 << n); i = i << 1 {
		for j = i - 1; j >= 0; j-- {
			arr = append(arr, arr[j])
		}

		for j = 0; j < i; j++ {
			arr[j] = fmt.Sprintf("0%v", arr[j])
		}

		for j = i; j < 2*i; j++ {
			arr[j] = fmt.Sprintf("1%v", arr[j])
		}
	}

	for _, bin := range arr {
		result = append(result, convertBinaryToInt(bin))
	}

	return result
}

func main() {
	for i := 1; i <= 16; i++ {
		fmt.Println(grayCode(i))
	}
}
