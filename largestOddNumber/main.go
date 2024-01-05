package main

import "fmt"

func main() {
	num := "52"

	fmt.Println((num[0] & 0x1) > 0)

	if (num[0] & 0x1) > 0 {
		fmt.Println(num)
	}

	for i := len(num) - 2; i >= 0; i-- {
		if (num[i] & 0x1) > 0 {
			fmt.Println(string(num[0 : i+1]))
		}
	}
}
