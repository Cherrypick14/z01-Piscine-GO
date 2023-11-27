package piscine

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := BTreeLevelCount(root.Left)
	right := BTreeLevelCount(root.Right)

	if right > left {
		return right + 1
	} else {
		return left + 1
	}
}
