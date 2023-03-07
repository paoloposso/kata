package preordersearch

import "trees/datastruct"

func walk(curr *datastruct.BynaryNode, path []int32) []int32 {
	//base case
	if curr == nil {
		return path
	}

	//pre
	path = append(path, curr.Value)

	//recurse
	path = walk(curr.Left, path)
	path = walk(curr.Right, path)

	//post
	return path
}

func PreOrderSearch(head datastruct.BynaryNode) []int32 {
	return walk(&head, []int32{})
}
