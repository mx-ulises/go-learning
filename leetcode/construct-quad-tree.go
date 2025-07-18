package leetcode

/**
 * Definition for a QuadTree node.
 * type Node struct {
 *     Val bool
 *     IsLeaf bool
 *     TopLeft *Node
 *     TopRight *Node
 *     BottomLeft *Node
 *     BottomRight *Node
 * }
 */

func buildQuadTree(grid [][]int, xStart, xEnd, yStart, yEnd int) *Node {
	val := grid[xStart][yStart] == 1
	if (xEnd - xStart) == 1 {
		return &Node{Val: val, IsLeaf: true}
	}
	xMid, yMid := (xStart+xEnd)/2, (yStart+yEnd)/2
	topLeft := buildQuadTree(grid, xStart, xMid, yStart, yMid)
	topRight := buildQuadTree(grid, xStart, xMid, yMid, yEnd)
	bottomLeft := buildQuadTree(grid, xMid, xEnd, yStart, yMid)
	bottomRight := buildQuadTree(grid, xMid, xEnd, yMid, yEnd)
	sameValue := topLeft.Val == topRight.Val && topLeft.Val == bottomLeft.Val && topLeft.Val == bottomRight.Val
	isLeaf := sameValue && topLeft.IsLeaf && topRight.IsLeaf && bottomLeft.IsLeaf && bottomRight.IsLeaf
	if isLeaf {
		topLeft, topRight, bottomLeft, bottomRight = nil, nil, nil, nil
	}
	return &Node{
		Val:         val,
		IsLeaf:      isLeaf,
		TopLeft:     topLeft,
		TopRight:    topRight,
		BottomLeft:  bottomLeft,
		BottomRight: bottomRight,
	}
}

func construct(grid [][]int) *Node {
	n := len(grid)
	return buildQuadTree(grid, 0, n, 0, n)
}
