package util

type DoublyLinkedList struct {
	Data       interface{}
	Next, Last *DoublyLinkedList
}

func (d *DoublyLinkedList) InsertNext(newNext *DoublyLinkedList) {
	tmpNext := d.Next

	// Splice it in front
	d.Next = newNext
	newNext.Last = d

	// if there WAS a last next element, splice that in.
	if tmpNext != nil {
		tmpNext.Last = newNext
		newNext.Next = tmpNext
	}
}

func (d *DoublyLinkedList) Remove() (onlyElement bool) {
	// Stitch together the sides if needbe.
	if d.Next != nil {
		d.Next.Last = d.Last
	}
	if d.Last != nil {
		d.Last.Next = d.Next
	}

	// Are we the only element in the array?
	onlyElement = d.Next == nil && d.Last == nil

	// Clear our connections.
	d.Next = nil
	d.Last = nil

	return onlyElement
}
