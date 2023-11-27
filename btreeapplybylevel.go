package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	q := make([]*TreeNode, 0)
	q = append(q, root)

	for len(q) > 0 {
		Root := q[0]
		q = q[1:]

		f(Root.Data)

		if Root.Left != nil {
			q = append(q, Root.Left)
		}
		if Root.Right != nil {
			q = append(q, Root.Right)
		}
	}
}
