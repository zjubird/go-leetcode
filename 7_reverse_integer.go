package main

import "fmt"

func reverse(x int) int{
	signed := 1
	if x < 0{
		signed = -1
		x = -x
	}
	res := 0
	for x != 0{
		res = 10 * res + x % 10
		x = x/10
	}
	return signed*res
}

func main(){
	fmt.Println(reverse(120))
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
}
