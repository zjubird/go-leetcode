package main

import "fmt"

type ListNode struct{
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode{
	carrier := 0
	var head *ListNode  = nil
	var prev *ListNode = nil
	s1 := l1
	s2 := l2
	for s1 != nil && s2 != nil{
		newNode := ListNode{}
		newNode.Val = (s1.Val + s2.Val + carrier) % 10
		carrier = (s1.Val + s2.Val + carrier) / 10
		newNode.Next = nil
		if head == nil{
			head = &newNode
			prev = &newNode
		}else{
			prev.Next = &newNode
			prev = &newNode
		}

		s1 = s1.Next
		s2 = s2.Next

	}
	var remainNode *ListNode = nil
	if s1 != nil{
		remainNode = s1
	}else if s2 != nil{
		remainNode = s2
	}

	for remainNode != nil{
		newNode := ListNode{}
		newNode.Val = (remainNode.Val + carrier) % 10
		carrier = (remainNode.Val + carrier) / 10
		remainNode = remainNode.Next

		prev.Next = &newNode
		prev = &newNode
	}
	if carrier != 0{
		newNode := ListNode{1, nil}
		prev.Next = &newNode
	}

	return head
}

func main(){
	node1 := ListNode{2,nil}
	node2 := ListNode{4,nil}
	node3 := ListNode{3,nil}
	node1.Next = &node2
	node2.Next = &node3

	node4 := ListNode{5, nil}
	node5 := ListNode{6, nil}
	node6 := ListNode{4, nil}
	node4.Next = &node5
	node5.Next = &node6

	res := addTwoNumbers(&node1, &node4)

	for res != nil{
		fmt.Println(res.Val)
		res = res.Next
	}

}
