package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	var previous *NodeL
	for node := l.Head; node != nil; node = node.Next {
		if node.Data == data_ref {
			if previous == nil {
				l.Head = node.Next
			} else {
				previous.Next = node.Next
			}
		} else {
			previous = node
		}
	}
}
