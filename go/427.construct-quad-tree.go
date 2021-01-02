package main

/*
 * @lc app=leetcode id=427 lang=golang
 *
 * [427] Construct Quad Tree
 *
 * https://leetcode.com/problems/construct-quad-tree/description/
 *
 * algorithms
 * Medium (59.29%)
 * Total Accepted:    25.6K
 * Total Submissions: 41.1K
 * Testcase Example:  '[[0,1],[1,0]]'
 *
 * Given a n * n matrix grid of 0's and 1's only. We want to represent the grid
 * with a Quad-Tree.
 *
 * Return the root of the Quad-Tree representing the grid.
 *
 * Notice that you can assign the value of a node to True or False when isLeaf
 * is False, and both are accepted in the answer.
 *
 * A Quad-Tree is a tree data structure in which each internal node has exactly
 * four children. Besides, each node has two attributes:
 *
 *
 * val: True if the node represents a grid of 1's or False if the node
 * represents a grid of 0's. 
 * isLeaf: True if the node is leaf node on the tree or False if the node has
 * the four children.
 *
 *
 *
 * class Node {
 * ⁠   public boolean val;
 * public boolean isLeaf;
 * public Node topLeft;
 * public Node topRight;
 * public Node bottomLeft;
 * public Node bottomRight;
 * }
 *
 * We can construct a Quad-Tree from a two-dimensional area using the following
 * steps:
 *
 *
 * If the current grid has the same value (i.e all 1's or all 0's) set isLeaf
 * True and set val to the value of the grid and set the four children to Null
 * and stop.
 * If the current grid has different values, set isLeaf to False and set val to
 * any value and divide the current grid into four sub-grids as shown in the
 * photo.
 * Recurse for each of the children with the proper sub-grid.
 *
 *
 * If you want to know more about the Quad-Tree, you can refer to the wiki.
 *
 * Quad-Tree format:
 *
 * The output represents the serialized format of a Quad-Tree using level order
 * traversal, where null signifies a path terminator where no node exists
 * below.
 *
 * It is very similar to the serialization of the binary tree. The only
 * difference is that the node is represented as a list [isLeaf, val].
 *
 * If the value of isLeaf or val is True we represent it as 1 in the list
 * [isLeaf, val] and if the value of isLeaf or val is False we represent it as
 * 0.
 *
 *
 * Example 1:
 *
 *
 * Input: grid = [[0,1],[1,0]]
 * Output: [[0,1],[1,0],[1,1],[1,1],[1,0]]
 * Explanation: The explanation of this example is shown below:
 * Notice that 0 represnts False and 1 represents True in the photo
 * representing the Quad-Tree.
 *
 *
 *
 * Example 2:
 *
 *
 *
 *
 * Input: grid =
 * [[1,1,1,1,0,0,0,0],[1,1,1,1,0,0,0,0],[1,1,1,1,1,1,1,1],[1,1,1,1,1,1,1,1],[1,1,1,1,0,0,0,0],[1,1,1,1,0,0,0,0],[1,1,1,1,0,0,0,0],[1,1,1,1,0,0,0,0]]
 * Output:
 * [[0,1],[1,1],[0,1],[1,1],[1,0],null,null,null,null,[1,0],[1,0],[1,1],[1,1]]
 * Explanation: All values in the grid are not the same. We divide the grid
 * into four sub-grids.
 * The topLeft, bottomLeft and bottomRight each has the same value.
 * The topRight have different values so we divide it into 4 sub-grids where
 * each has the same value.
 * Explanation is shown in the photo below:
 *
 *
 *
 * Example 3:
 *
 *
 * Input: grid = [[1,1],[1,1]]
 * Output: [[1,1]]
 *
 *
 * Example 4:
 *
 *
 * Input: grid = [[0]]
 * Output: [[1,0]]
 *
 *
 * Example 5:
 *
 *
 * Input: grid = [[1,1,0,0],[1,1,0,0],[0,0,1,1],[0,0,1,1]]
 * Output: [[0,1],[1,1],[1,0],[1,0],[1,1]]
 *
 *
 *
 * Constraints:
 *
 *
 * n == grid.length == grid[i].length
 * n == 2^x where 0 <= x <= 6
 *
 *
 */
/**
 * Definition for a QuadTree node.
 */
type QuadNode struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *QuadNode
	TopRight    *QuadNode
	BottomLeft  *QuadNode
	BottomRight *QuadNode
}

func construct(grid [][]int) *QuadNode {
	n := len(grid)
	var build func(x, y, w, h int) *QuadNode
	build = func(x, y, w, h int) *QuadNode {
		if w == 0 || h == 0 {
			return nil
		}

		if w == 1 && h == 1 {
			return &QuadNode{
				Val:    grid[y][x] == 1,
				IsLeaf: true,
			}
		}

		// This for-loop is much faster than the recursive method
		// This results in 4ms (100%), while recursive then check all 4 subtrees
		// got 12ms (52%)
		v := grid[y][x]
		for j := y; j < y+h; j++ {
			for i := x; i < x+w; i++ {
				if v != grid[j][i] {
					goto notLeaf
				}
			}
		}
		return &QuadNode{
			Val:    grid[y][x] == 1,
			IsLeaf: true,
		}

	notLeaf:
		mw := w / 2
		mh := h / 2

		t1 := build(x, y, mw, mh)
		t2 := build(x+mw, y, w-mw, mh)
		t3 := build(x, y+mh, mw, h-mh)
		t4 := build(x+mw, y+mh, w-mw, h-mh)

		return &QuadNode{
			IsLeaf:      false,
			TopLeft:     t1,
			TopRight:    t2,
			BottomLeft:  t3,
			BottomRight: t4,
		}
	}

	return build(0, 0, n, n)
}
