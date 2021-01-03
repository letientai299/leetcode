package main

// medium
// https://leetcode.com/problems/logical-or-of-two-binary-grids-represented-as-quad-trees/

func intersect(a *QuadNode, b *QuadNode) *QuadNode {
	if a.IsLeaf && b.IsLeaf {
		return &QuadNode{
			IsLeaf: true,
			Val:    a.Val || b.Val,
		}
	}

	clean := func(q *QuadNode) *QuadNode {
		if q.IsLeaf {
			return q
		}

		if q.TopLeft.IsLeaf &&
			q.TopRight.IsLeaf &&
			q.BottomLeft.IsLeaf &&
			q.BottomRight.IsLeaf &&
			q.TopLeft.Val == q.TopRight.Val &&
			q.TopRight.Val == q.BottomLeft.Val &&
			q.BottomRight.Val == q.BottomLeft.Val {
			q.IsLeaf = true
			q.Val = q.TopLeft.Val
			q.TopLeft = nil
			q.TopRight = nil
			q.BottomLeft = nil
			q.BottomRight = nil
		}

		return q
	}

	if !a.IsLeaf && !b.IsLeaf {
		return clean(&QuadNode{
			IsLeaf:      false,
			TopLeft:     intersect(a.TopLeft, b.TopLeft),
			TopRight:    intersect(a.TopRight, b.TopRight),
			BottomLeft:  intersect(a.BottomLeft, b.BottomLeft),
			BottomRight: intersect(a.BottomRight, b.BottomRight),
		})
	}

here:
	if a.IsLeaf {
		return clean(&QuadNode{
			IsLeaf:      false,
			Val:         b.Val,
			TopLeft:     intersect(a, b.TopLeft),
			TopRight:    intersect(a, b.TopRight),
			BottomLeft:  intersect(a, b.BottomLeft),
			BottomRight: intersect(a, b.BottomRight),
		})
	}

	a, b = b, a
	goto here
}
