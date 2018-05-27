package main
import "fmt"

func twoSum(s []int, target int) (int, int){
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++{
			if s[i] + s[j] == target{
				return i,j
			}
		}
	}
	return 0,0
}

func main(){
	data := []int{1,2,3,4,5}
	fmt.Println(twoSum(data, 9))
}
