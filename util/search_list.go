package util

type SearchList struct {
	front *doublyLinkedElement
	back *doublyLinkedElement
	
	count uint
}

type doublyLinkedElement struct {
	data interface{}
	next, last *doublyLinkedElement
}

type linkedListIterator struct { // Iterates from front to back
	lIdx *doublyLinkedElement
	reverse bool
}

func (l *linkedListIterator) Next() bool {
	if l.lIdx == nil {
		return false
	}
	
	if !l.reverse {
		l.lIdx = l.lIdx.last
	} else {
		l.lIdx = l.lIdx.next
	}
	
	return l.lIdx != nil
}

func (l *linkedListIterator) Value() interface{} {
	return l.lIdx.data
}

func (d *doublyLinkedElement) insertNext(newNext *doublyLinkedElement) {
	tmpNext := d.next
	
	// Splice it in front
	d.next = newNext
	newNext.last = d
	
	// if there WAS a last next element, splice that in.
	if tmpNext != nil {
		tmpNext.last = newNext
		newNext.next = tmpNext
	}
}

func (d *doublyLinkedElement) remove() (onlyElement bool) {
	// Stitch together the sides if needbe.
	if d.next != nil {
		d.next.last = d.last
	}
	if d.last != nil {
		d.last.next = d.next
	}
	
	// Are we the only element in the array?
	onlyElement = d.next == nil && d.last == nil
	
	// Clear our connections.
	d.next = nil
	d.last = nil
	
	return onlyElement
}

type LinkedSearchStatus uint

const (
	SearchMatch LinkedSearchStatus = iota
	SearchMatchNoContinue // SearchMatchNoContinue informs the searcher to NOT continue searching after this match.
	SearchContinue // SearchContinue will not return the last element. Reaching the back cancels an insert entirely.
	SearchContinueAcceptBack // SearchContinueAcceptBack will return or insert at the back of the array.
	SearchNoMatch // SearchNoMatch indicates that there will be no future match. This cancels an insert entirely.
)

// LinkedSearchFunc specifies a function that will be used to select the insert point for the linked list element
type LinkedSearchFunc func(next interface{}) LinkedSearchStatus

func LinkedSearchFront(next interface{}) LinkedSearchStatus { return SearchMatchNoContinue } // Find front of array
func LinkedSearchBack(next interface{}) LinkedSearchStatus  { return SearchContinueAcceptBack } // Find back of array
// CurryLinkedSearchIndex : Find what's at an index in an array. Should be supplied non-match continuation policy.
// NOTE: Using SearchNoMatch is invalid and will NOT return anything. only Continue and ContinueAcceptBack should be used.
// These functions should _not_ be re-used.
func CurryLinkedSearchIndex (idx uint, continuationPolicy LinkedSearchStatus) LinkedSearchFunc {
	o := idx
	
	return func(next interface{}) LinkedSearchStatus {
		if o == 0 {
			return SearchMatchNoContinue
		}
		
		o--;
		return continuationPolicy
	}
}
func CurryLinkedSearchEq (eqData interface{}, matchPolicy LinkedSearchStatus) LinkedSearchFunc { // Find one or more matches (specify SearchMatchNoContinue for a single match)
	return func(next interface{}) LinkedSearchStatus {
		if next == eqData {
			return matchPolicy
		}
		
		return SearchContinue
	}
}

// Insert inserts data at the sort point specified. Searches from the front back.
func (l *SearchList) Insert(data interface{}, sortFunc LinkedSearchFunc) {
	lIdx := l.front
	
	if lIdx == nil { // insert will always work here
		l.back = &doublyLinkedElement{data: data}
		l.front = l.back
		l.count++
		
		return
	}
	
	l.count++
	
	for {
		if resp := sortFunc(lIdx.data); resp == SearchMatch || resp == SearchMatchNoContinue { // We're inserting here.
			lIdx.insertNext(&doublyLinkedElement{
				data: data,
			})
			if lIdx == l.front {
				l.front = lIdx.next
			}
			
			return
		} else if resp == SearchNoMatch || (resp == SearchContinue && lIdx.last == nil) {
			l.count-- // Cancel the insertion.
			return
		}
		
		if lIdx.last == nil { // We're definitely inserting at the back of the array.
			lIdx.last = &doublyLinkedElement{
				data: data,
				next: lIdx,
			}
			l.back = lIdx.last
			
			return
		}
		
		lIdx = lIdx.last // step towards the back of the array
	}
}

// Search finds multiple entities unless SearchMatchNoContinue is returned.
func (l *SearchList) Search(searchFunc LinkedSearchFunc) ([]interface{}, []uint) {
	lIdx := l.front
	out := make([]interface{}, 0)
	outIdx := make([]uint, 0)
	
	idx := uint(0)
	
	for {
		if lIdx == nil {
			return out, outIdx
		}
		
		resp := searchFunc(lIdx.data)
		
		if resp == SearchMatchNoContinue {
			return []interface{}{lIdx.data}, []uint{idx}
		} else if resp == SearchMatch || (resp == SearchContinueAcceptBack && lIdx.last == nil) { // does it match with the search query or status?
			out = append(out, lIdx.data)
			outIdx = append(outIdx, idx)
		} else if resp == SearchNoMatch { // We should exit immediately, because we're not going to find anything else.
			return out, outIdx
		}
		
		lIdx = lIdx.last
		idx++
	}
}

// Remove removes all matching entities unless SearchMatchNoContinue is returned.
func (l *SearchList) Remove(searchFunc LinkedSearchFunc) (count uint) {
	lIdx := l.front
	
	for {
		if lIdx == nil {
			return
		}
		
		resp := searchFunc(lIdx.data)
		
		tmpIdx := lIdx
		lIdx = lIdx.last
		
		if resp == SearchMatch || resp == SearchMatchNoContinue {
			tmpIdx.remove()
			count++
			l.count--
			
			if resp == SearchMatchNoContinue {
				return
			}
		} else if resp == SearchContinueAcceptBack && lIdx == nil {
			tmpIdx.remove()
			count++
			l.count--
			return
		} else if resp == SearchNoMatch {
			return
		}
	}
}

func (l *SearchList) Front() interface{} {
	if l.front == nil {
		return nil
	}
	
	return l.front.data
}

func (l *SearchList) Back() interface{} {
	if l.back == nil {
		return nil
	}
	
	return l.back.data
}

// Iterator Gets an Iterator from front to back.
func (l *SearchList) Iterator(reverse bool) Iterator {
	idx := l.front
	if reverse {
		idx = l.back
	}
	return &linkedListIterator{lIdx: idx, reverse: reverse}
}

func (l *SearchList) Len() uint {
	return l.count
}

func (l *SearchList) Clear() {
	l.front = nil
	l.back = nil
}