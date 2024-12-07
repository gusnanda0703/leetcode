package main

import (
	"fmt"
	"leetcode/solutions"
)

// func encryptor(a0 *[]uint32, a1 *string) {
// 	var v0, v1, v2 uint32
// 	// var v3 byte
// 	// var v5 uintptr

// 	// v5 = uintptr(unsafe.Pointer(&v3))
// 	v1 = uint32(len(*a1))
// 	for v2 = 0; v2 < v1-7; v2++ {
// 		v0 = uint32((*a1)[7+v2]) ^ uint32((*a1)[v2]) ^ uint32((*a1)[1+v2])
// 		(*a0)[v2] = v0
// 	}
// }

// func checker(a0 *[]uint32) int {
// 	var v0 uint32

// 	for v0 = 0; v0 <= 55; v0++ {
// 		if (*a0)[v0] != (*a0)[70+v0] {
// 			return 0
// 		}
// 	}
// 	return 1
// }

func main() {
	fmt.Println(solutions.DeleteAndEarn([]int{2, 2, 3, 3, 3, 4}))
}

// func dynamicArray(n int32, queries [][]int32) []int32 {
// 	var lastAnswer int32 = 0
// 	arr := make([][]int32, n)
// 	var result []int32

// 	for _, v := range queries {
// 		idx := (v[1] ^ lastAnswer) % n
// 		if v[0] == 1 {
// 			arr[idx] = append(arr[idx], v[2])
// 		} else {
// 			i := int(v[2]) % len(arr[idx])
// 			lastAnswer = arr[idx][i]
// 			result = append(result, lastAnswer)
// 		}
// 	}

// 	return result
// }

// func rotateLeft(d int32, arr []int32) []int32 {
// 	idx := int(d) % len(arr)
// 	return append(arr[idx:], arr[:idx]...)
// }

// func matchingStrings(stringList []string, queries []string) []int32 {
// 	counts := make(map[string]int32)
// 	for _, str := range stringList {
// 		counts[str]++
// 	}

// 	result := make([]int32, len(queries))
// 	for i, query := range queries {
// 		result[i] = counts[query]
// 	}

// 	return result
// }

// func arrayManipulation(n int32, queries [][]int32) int64 {
// 	arr := make([]int64, n+1)
// 	for _, v := range queries {
// 		a, b, k := v[0], v[1], v[2]
// 		arr[a-1] += int64(k)
// 		arr[b] -= int64(k)
// 	}

// 	var max int64
// 	var sum int64
// 	for _, v := range arr {
// 		sum += v
// 		if sum > max {
// 			max = sum
// 		}
// 	}

// 	return max

// }

// func removeNthFromEnd(head *ListNode, n int) *ListNode {
// 	dummy := &ListNode{0, head}
// 	first, second := dummy, dummy

// 	// Advances first pointer so that the gap between first and second is n nodes apart
// 	for i := 1; i <= n+1; i++ {
// 		first = first.Next
// 	}

// 	// Move first to the end, maintaining the gap
// 	for first != nil {
// 		first = first.Next
// 		second = second.Next
// 	}
// 	second.Next = second.Next.Next
// 	return dummy.Next
// }
