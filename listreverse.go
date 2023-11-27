package piscine

func ListReverse(l *List) {
	var prev *NodeL
	current := l.Head

	for current != nil {
		// Store next node
		next := current.Next

		// Reverse pointer
		current.Next = prev

		// Advance pointers
		prev = current
		current = next
	}

	// Swap head and tail
	l.Head = l.Tail
	l.Tail = prev
}
