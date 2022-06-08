package easy

import (
	"fmt"
)

func RunCase21() {

	fmt.Println("Case21")

	list1 := ListNode{}
	list1.Val = 1
	list1.Next = &ListNode{}
	list1.Next.Val = 2
	list1.Next.Next = &ListNode{}
	list1.Next.Next.Val = 5
	list1.Next.Next.Next = &ListNode{}
	list1.Next.Next.Next.Val = 9

	list2 := ListNode{}
	list2.Val = 1
	list2.Next = &ListNode{}
	list2.Next.Val = 3
	list2.Next.Next = &ListNode{}
	list2.Next.Next.Val = 4
	list2.Next.Next.Next = &ListNode{}
	list2.Next.Next.Next.Val = 7
	list2.Next.Next.Next.Next = &ListNode{}
	list2.Next.Next.Next.Next.Val = 8

	res := mergeTwoLists(&list1, &list2)
	printNode(res)
}

func printNode(list *ListNode) {
	lst := list
	for lst != nil {
		fmt.Println(lst.Val)
		lst = lst.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// Faster than 100%
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	node := ListNode{}

	if list1 == nil && list2 != nil {
		return list2
	} else if list2 == nil && list1 != nil {
		return list1
	} else if list1 == nil && list2 == nil {
		return nil
	}

	if list1.Val <= list2.Val {
		res := mergeTwoLists(list1.Next, list2)
		node.Next = res
		node.Val = list1.Val
	} else {
		res := mergeTwoLists(list1, list2.Next)
		node.Next = res
		node.Val = list2.Val
	}

	return &node
}
