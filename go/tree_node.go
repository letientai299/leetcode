package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

const NA = math.MaxInt32

type TreeNode struct {
	Val   int       `json:"val"`
	Left  *TreeNode `json:"left"`
	Right *TreeNode `json:"right"`
}

func (t *TreeNode) String() string {
	s := fmt.Sprint(t.levelOrder())
	return strings.ReplaceAll(s, strconv.Itoa(NA), "_")
}

func (t *TreeNode) inOrder() []int {
	if t == nil {
		return nil
	}

	result := append(t.Left.inOrder(), t.Val)
	return append(result, t.Right.inOrder()...)
}

func (t *TreeNode) levels() [][]int {
	if t == nil {
		return nil
	}
	res := make([][]int, 0, 16)
	q := []*TreeNode{t}
	res = append(res, []int{t.Val})
	for len(q) > 0 {
		var level []int
		var next []*TreeNode
		for _, r := range q {
			if r.Left != nil {
				level = append(level, r.Left.Val)
				next = append(next, r.Left)
			} else {
				level = append(level, NA)
			}

			if r.Right != nil {
				level = append(level, r.Right.Val)
				next = append(next, r.Right)
			} else {
				level = append(level, NA)
			}
		}
		q = next
		res = append(res, level)
	}
	return res[:len(res)-1]
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

func (t *TreeNode) Clone() *TreeNode {
	if t == nil {
		return nil
	}

	clone := new(TreeNode)
	clone.Val = t.Val

	clone.Left = t.Left.Clone()
	clone.Right = t.Right.Clone()
	return clone
}

func (t *TreeNode) Pretty() {
	levels := t.levels()
	levelStrings := make([][]string, 0, len(levels))
	for _, vs := range levels {
		ss := make([]string, 0, len(vs))
		for i, x := range vs {
			s := "_"
			if x != NA {
				s = strconv.Itoa(x)
			}

			if len(s)%2 == 0 {
				if i%2 == 0 {
					s += " "
				} else {
					s = " " + s
				}
			}

			ss = append(ss, s)
		}
		levelStrings = append(levelStrings, ss)
	}

	n := len(levelStrings)
	visuals := make([]string, 0, n)

	re := strings.NewReplacer("_", " ")

	var sb strings.Builder
	for _, s := range levelStrings[n-1] {
		sb.WriteString(re.Replace(s))
		sb.WriteString(strings.Repeat(" ", len(s)*(n-1)))
	}
	visuals = append(visuals, sb.String())

	for i := n - 2; i >= 0; i-- {
		prefixSpace := strings.Repeat(" ", 2*(n-i-1)-1)
		sb.Reset()

		sb.WriteString(prefixSpace)
		for j, s := range levelStrings[i+1] {
			var slash = "/"
			if j%2 == 1 {
				slash = "\\"
			}
			if strings.Contains(s, "_") {
				slash = " "
			}
			sb.WriteString(slash)
			sb.WriteString(strings.Repeat(" ", len(s)*(i+1)))
		}
		visuals = append(visuals, sb.String())

		sb.Reset()
		sb.WriteString(prefixSpace)
		sb.WriteString(" ")
		for _, s := range levelStrings[i] {
			sb.WriteString(re.Replace(s))
			sb.WriteString(strings.Repeat(" ", len(s)*i))
		}
		visuals = append(visuals, sb.String())
	}

	for i := len(visuals) - 1; i >= 0; i-- {
		fmt.Println(visuals[i])
	}
}

