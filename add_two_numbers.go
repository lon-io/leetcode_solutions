package main

import (
    "fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	resNode := &ListNode{}
	curResNode := resNode
	curNode1 := l1
	curNode2 := l2
	carry := 0

	for {
		var d1, d2 int
		if curNode1 != nil {
			d1 = curNode1.Val
			curNode1 = curNode1.Next
		}

		if curNode2 != nil {
			d2 = curNode2.Val
			curNode2 = curNode2.Next
		}

		dSum := d1 + d2 + carry
		carry = dSum / 10
		dSum = dSum % 10

		fmt.Printf("\nNode1 %d + Node2 %d ==> %d;; carry: %d", d1, d2, dSum, carry)

		curResNode.Val = dSum

		if curNode1 == nil && curNode2 == nil {
            if carry > 0 {
				carryNode := ListNode{
					Val: carry,
					Next: nil,
				}
				curResNode.Next = &carryNode
			}

			break
		}

		nextNode := ListNode{}
		curResNode.Next = &nextNode
		curResNode = curResNode.Next
	}

    return resNode
}

// populateLinkedList: Helper function to build LinkedList for testCases
func populateLinkedList(ln *ListNode, acc []int, ci int) {
    iSum := acc[ci]
    ln.Val = int(iSum)

    if ci < len(acc) - 1 {
        next := ListNode{}
        ln.Next = &next
        populateLinkedList(ln.Next, acc, ci + 1)
    }
}

// walk: Helper function to walk a Linkedlist and populate build a
// string representation for it
func walk (l *ListNode, a *string) {
    v := l.Val
    *a =  *a + fmt.Sprintf("%d", v)

    if l.Next != nil {
        walk(l.Next, a)
    }
}

func main() {
	// Generated Test Case
	test1 := []int{1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1}
	test2 := []int{5,6,4}

	var listNode1, listNode2 ListNode
	populateLinkedList(&listNode1, test1, 0)
	populateLinkedList(&listNode2, test2, 0)

	resNode := addTwoNumbers(&listNode1, &listNode2)
	resStr := ""
	walk(resNode, &resStr)

	fmt.Printf("\nRESULT::: %s \n", resStr)
}
