package main

import (
	"encoding/json"
	"log"
	"strings"
)

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

func treeFromJson(str string) *TreeNode {
	decoder := json.NewDecoder(strings.NewReader(str))
	tr := new(TreeNode)
	if err := decoder.Decode(tr); err != nil {
		log.Println(err)
		return nil
	}

	return tr
}
