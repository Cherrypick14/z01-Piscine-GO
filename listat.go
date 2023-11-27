package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	for i := 0; l != nil; i++ {
		if i == pos {
			return l
		}
		l = l.Next
	}
	return nil
}
