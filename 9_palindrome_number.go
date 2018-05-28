package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0{
		return false
	}

	s := x
	reverse := 0
	for s != 0{
		digit := s%10
		s /= 10
		reverse = 10 *reverse + digit
	}

	if reverse == x{
		return true
	}else{
		return false
	}
}


func main(){
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(10))
}
