package util

type SearchList struct {
	front *DoublyLinkedList
	back  *DoublyLinkedList

	count uint
}

type linkedListIterator struct { // Iterates from front to back
	lIdx    *DoublyLinkedList
	reverse bool
}

func (l *linkedListIterator) Next() bool {
	if l.lIdx == nil {
		return false
	}

	if !l.reverse {
		l.lIdx = l.lIdx.Last
	} else {
		l.lIdx = l.lIdx.Next
	}

	return l.lIdx != nil
}

func (l *linkedListIterator) Value() interface{} {
	return l.lIdx.Data
}

type LinkedSearchStatus uint

const (
	SearchMatch              LinkedSearchStatus = iota
	SearchMatchNoContinue                       // SearchMatchNoContinue informs the searcher to NOT continue searching after this match.
	SearchContinue                              // SearchContinue will not return the last element. Reaching the back cancels an insert entirely.
	SearchContinueAcceptBack                    // SearchContinueAcceptBack will return or insert at the back of the array.
	SearchNoMatch                               // SearchNoMatch indicates that there will be no future match. This cancels an insert entirely.
)

// LinkedSearchFunc specifies a function that will be used to select the insert point for the linked list element
type LinkedSearchFunc func(next interface{}) LinkedSearchStatus

func LinkedSearchFront(next interface{}) LinkedSearchStatus { return SearchMatchNoContinue }    // Find front of array
func LinkedSearchBack(next interface{}) LinkedSearchStatus  { return SearchContinueAcceptBack } // Find back of array
// CurryLinkedSearchIndex : Find what's at an index in an array. Should be supplied non-match continuation policy.
// NOTE: Using SearchNoMatch is invalid and will NOT return anything. only Continue and ContinueAcceptBack should be used.
// These functions should _not_ be re-used.
func CurryLinkedSearchIndex(idx uint, continuationPolicy LinkedSearchStatus) LinkedSearchFunc {
	o := idx

	return func(next interface{}) LinkedSearchStatus {
		if o == 0 {
			return SearchMatchNoContinue
		}

		o--
		return continuationPolicy
	}
}
func CurryLinkedSearchEq(eqData interface{}, matchPolicy LinkedSearchStatus) LinkedSearchFunc { // Find one or more matches (specify SearchMatchNoContinue for a single match)
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
		l.back = &DoublyLinkedList{Data: data}
		l.front = l.back
		l.count++

		return
	}

	l.count++

	for {
		if resp := sortFunc(lIdx.Data); resp == SearchMatch || resp == SearchMatchNoContinue { // We're inserting here.
			lIdx.InsertNext(&DoublyLinkedList{
				Data: data,
			})
			if lIdx == l.front {
				l.front = lIdx.Next
			}

			return
		} else if resp == SearchNoMatch || (resp == SearchContinue && lIdx.Last == nil) {
			l.count-- // Cancel the insertion.
			return
		}

		if lIdx.Last == nil { // We're definitely inserting at the back of the array.
			lIdx.Last = &DoublyLinkedList{
				Data: data,
				Next: lIdx,
			}
			l.back = lIdx.Last

			return
		}

		lIdx = lIdx.Last // step towards the back of the array
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

		resp := searchFunc(lIdx.Data)

		if resp == SearchMatchNoContinue {
			return []interface{}{lIdx.Data}, []uint{idx}
		} else if resp == SearchMatch || (resp == SearchContinueAcceptBack && lIdx.Last == nil) { // does it match with the search query or status?
			out = append(out, lIdx.Data)
			outIdx = append(outIdx, idx)
		} else if resp == SearchNoMatch { // We should exit immediately, because we're not going to find anything else.
			return out, outIdx
		}

		lIdx = lIdx.Last
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

		resp := searchFunc(lIdx.Data)

		tmpIdx := lIdx
		lIdx = lIdx.Last

		if resp == SearchMatch || resp == SearchMatchNoContinue {
			tmpIdx.Remove()
			count++
			l.count--

			if resp == SearchMatchNoContinue {
				return
			}
		} else if resp == SearchContinueAcceptBack && lIdx == nil {
			tmpIdx.Remove()
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

	return l.front.Data
}

func (l *SearchList) Back() interface{} {
	if l.back == nil {
		return nil
	}

	return l.back.Data
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
