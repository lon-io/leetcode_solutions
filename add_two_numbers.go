package main

import (
    "fmt"
    "strconv"
    "strings"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func walk (l *ListNode, a *string) {
    v := l.Val
    *a =  *a + fmt.Sprintf("%d", v)

    if l.Next != nil {
        walk(l.Next, a)
    }
}

func extendLinkedList(ln *ListNode, acc []string, ci int) {
    cmp := len(acc) - ci - 1
    iSum, _ := strconv.ParseInt(acc[cmp], 10, 32)
    ln.Val = int(iSum)

    if cmp != 0 {
        next := ListNode{}
        ln.Next = &next
        extendLinkedList(ln.Next, acc, ci + 1)
    }
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    acc1 := ""
    acc2 := ""

    walk(l1, &acc1)
    walk(l2, &acc2)

    lng := acc1
    srt := acc2
    if len(acc1) < len(acc2) {
        lng = acc2
        srt = acc1
    }

    sumStr := ""
    carry := 0
    for i := 0; i <= len(lng); i ++ {
        if i == len(lng) {
            if carry > 0 {
                sumStr = fmt.Sprintf("%d", carry) + sumStr
            }
        } else {
            var vs, vl string

            if i < len(srt)  {
                vs = string(srt[i])
            }

            vl = string(lng[i])
            ds, _ := strconv.Atoi(vs)
            dl, _ := strconv.Atoi(vl)

            sI := ds + dl + carry
            carry = sI / 10
            sI = sI % 10

            fmt.Printf("\n Index: %d ::: short %d + long %d ==> %d;; carry: %d", i, ds, dl, sI, carry)
            sumStr = fmt.Sprintf("%d", sI) + sumStr
        }
    }

    fmt.Printf("\n FINAL SUM::: %s", sumStr)
    sumSl := strings.Split(sumStr, "")

    lnSum := ListNode{}
    extendLinkedList(&lnSum, sumSl, 0)

    return &lnSum
}

func extendTestCaseLinkedList(ln *ListNode, acc []int, ci int) {
    iSum := acc[ci]
    ln.Val = int(iSum)

    if ci < len(acc) - 1 {
        next := ListNode{}
        ln.Next = &next
        extendTestCaseLinkedList(ln.Next, acc, ci + 1)
    }
}

func main() {
	// Generated Test Case
	test1 := []int{1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1}
	test2 := []int{5,6,4}

	var listNode1, listNode2 ListNode
	extendTestCaseLinkedList(&listNode1, test1, 0)
	extendTestCaseLinkedList(&listNode2, test2, 0)

	addTwoNumbers(&listNode1, &listNode2)
}
