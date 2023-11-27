package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	if l == nil || l.Next == nil {
		return l
	}

	IsSorted := true
	for IsSorted {
		IsSorted = false
		current := l

		for current.Next != nil {
			if current.Data > current.Next.Data {
				temp := current.Data
				current.Data = current.Next.Data
				current.Next.Data = temp
				IsSorted = true
			}
			current = current.Next
		}
	}
	return l
}
