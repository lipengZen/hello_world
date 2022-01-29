package leetcode

import "testing"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	nodeMap := make(map[*ListNode]bool, 0)

	node := head
	for node != nil {
		if nodeMap[node] {
			return true
		}
		nodeMap[node] = true
		node = node.Next
	}
	return false
}

var (
	node1 = &ListNode{Next: node2, Val: 1}
	node2 = &ListNode{Next: node3, Val: 2}
	node3 = &ListNode{Next: node4, Val: 3}
	node4 = &ListNode{Next: node5, Val: 4}
	node5 = &ListNode{Val: 5}
)

func reverseKGroup(head *ListNode, k int) *ListNode {

	currentNode := head
	var lastNode *ListNode

	var flag, notEnough bool
	ansNode := head

	pre := &ListNode{}

	for currentNode != nil {

		i := 1
		nextNode := currentNode // 下一个k个节点的第一个,改成:这k个节点的最后一个
		notEnough = false
		for i < k {
			i++
			nextNode = nextNode.Next
			if nextNode == nil {
				notEnough = true
				break
			}
		}

		if notEnough {
			break
		}

		// 反转k个节点

		// 记录这k个节点的第一个
		first := currentNode

		// lastNode = currentNode
		// currentNode = currentNode.Next
		lastNode = nextNode.Next

		nextK := nextNode.Next
		for currentNode != nextK { // nextNode.Next {  当currentNode到nextNode时, nextNode指针会变, 空指针访问异常

			node := currentNode.Next
			currentNode.Next = lastNode
			lastNode = currentNode
			currentNode = node
		}

		// 记录下一串初始的 lastNode
		// 错:lastNode = first, 应该是第一个节点的next指向 nextNode.Next

		pre.Next = nextNode
		// pre = nextNode
		pre = first

		if !flag {
			ansNode = nextNode
			flag = true
		}
	}

	return ansNode

}

func TestHasCircle(t *testing.T) {
	reverseKGroup(node1, 2)
	for node := node2; node != nil; node = node.Next {
		t.Log(node)
	}
	return
}
