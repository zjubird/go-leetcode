package main

import "fmt"
import "math"
func reverse(x int) int{
	if x > math.MaxInt32 || x < math.MinInt32{
		return 0
	}
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

	if signed * res > math.MaxInt32 || signed * res < math.MinInt32{
		return 0
	}
	return signed*res
}

func main(){
	fmt.Println(reverse(120))
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(1534236469))
	fmt.Println(reverse(-2147483648))
}
