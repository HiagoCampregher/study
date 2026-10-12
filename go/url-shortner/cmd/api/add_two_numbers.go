package main

// ListNode representa um dígito da lista encadeada.
type ListNode struct {
	Val  int
	Next *ListNode
}

// addTwoNumbers soma dois números cujos dígitos estão em ordem reversa.
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		tail.Next = &ListNode{
			Val: sum % 10,
		}
		tail = tail.Next
		carry = sum / 10
	}

	return dummy.Next
}

// lengthOfLongestSubstring retorna o tamanho da maior substring sem caracteres repetidos.
func lengthOfLongestSubstring(s string) int {
	var lastSeen [256]int
	left, longest := 0, 0

	for right := 0; right < len(s); right++ {
		char := s[right]
		if lastSeen[char] > left {
			left = lastSeen[char]
		}

		lastSeen[char] = right + 1
		currentLength := right - left + 1
		if currentLength > longest {
			longest = currentLength
		}
	}

	return longest
}
