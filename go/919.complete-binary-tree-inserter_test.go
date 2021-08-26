package main

import (
	"testing"
)

func Test_CBTInserter(t *testing.T) {
	cb := Constructor919(treeFromLevelOrder(1,2,3,4,5,6))
	cb.Insert(7)
	cb.Insert(8)
	cb.Get_root().print()
}
