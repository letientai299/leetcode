package main

import (
	"encoding/json"
	"log"
	"math"
	"strings"
)

const NA = math.MaxInt32

type TreeNode struct {
	Val   int       `json:"val"`
	Left  *TreeNode `json:"left"`
	Right *TreeNode `json:"right"`
}

func (t *TreeNode) inOrder() []int {
	if t == nil {
		return nil
	}

	result := append(t.Left.inOrder(), t.Val)
	return append(result, t.Right.inOrder()...)
}

func (t *TreeNode) levelOrder() []int {
	if t == nil {
		return []int{NA}
	}

	q := make([]*TreeNode, 1)
	var result []int = nil
	q[0] = t
	for len(q) != 0 {
		nextQ := make([]*TreeNode, 0)
		for _, n := range q {
			if n == nil {
				result = append(result, NA)
				continue
			}
			result = append(result, n.Val)
			nextQ = append(nextQ, n.Left)
			nextQ = append(nextQ, n.Right)
		}
		q = nextQ
	}

	i := len(result) - 1
	for result[i] == NA {
		i--
	}

	return result[:i+1]
}

func treeFromJson(str string) *TreeNode {
	decoder := json.NewDecoder(strings.NewReader(str))
	tr := new(TreeNode)
	if err := decoder.Decode(tr); err != nil {
		log.Println(err)
		return nil
	}

	return tr
}

func treeFromLevelOrder(values ...int) *TreeNode {
	if len(values) == 0 {
		return nil
	}

	nodes := make([]*TreeNode, 1)
	root := &TreeNode{
		Val: values[0],
	}
	nodes[0] = root
	i := 1
	for len(nodes) != 0 && i < len(values) {
		// get the first node of the queue
		node := nodes[0]

		// make new nodes and append to the queue
		if i < len(values) && values[i] != NA {
			// this node has value
			node.Left = &TreeNode{
				Val: values[i],
			}
			nodes = append(nodes, node.Left)
		}
		i++

		if i < len(values) && values[i] != NA {
			// this node has value
			node.Right = &TreeNode{
				Val: values[i],
			}
			nodes = append(nodes, node.Right)
		}
		i++
		nodes = nodes[1:]
	}

	return root
}
