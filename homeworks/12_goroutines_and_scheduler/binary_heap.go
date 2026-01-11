package homework12

import "errors"

var ErrNotFound = errors.New("not found")

type Prioritisable interface {
	GetIdentifier() int
	GetPriority() int
}

type BinaryHeap struct {
	list []Prioritisable
}

func NewBinaryHeap() BinaryHeap {
	return BinaryHeap{
		list: make([]Prioritisable, 0),
	}
}

func (bh *BinaryHeap) Size() int {
	return len(bh.list)
}

func (bh *BinaryHeap) Add(elem Prioritisable) {
	bh.list = append(bh.list, elem)

	i := bh.Size() - 1
	parent := (i - 1) / 2

	for bh.list[parent].GetPriority() < bh.list[i].GetPriority() {
		bh.list[i], bh.list[parent] = bh.list[parent], bh.list[i]
		i = parent
		parent = (i - 1) / 2
	}
}

func (bh *BinaryHeap) GetMax() (result Prioritisable, err error) {
	if bh.Size() == 0 {
		return nil, ErrNotFound
	}

	result = bh.list[0]

	bh.list[0], bh.list[bh.Size()-1] = bh.list[bh.Size()-1], bh.list[0]
	bh.list = bh.list[0 : bh.Size()-1]
	bh.hapify(0)

	return result, nil
}

func (bh *BinaryHeap) Get(identifier int) (result Prioritisable, err error) {
	for i, v := range bh.list {
		if v.GetIdentifier() == identifier {
			result = bh.list[i]
			bh.list[i], bh.list[bh.Size()-1] = bh.list[bh.Size()-1], bh.list[i]
			bh.list = bh.list[:bh.Size()-1]
			bh.hapify(i)
			return result, nil
		}
	}
	return nil, ErrNotFound
}

func (bh *BinaryHeap) hapify(i int) {
	if i < 0 {
		return
	}

	var leftChildIdx, rightChildIdx, LargestChildIdx int

	for {
		leftChildIdx = 2*i + 1
		rightChildIdx = 2*i + 2
		LargestChildIdx = i

		if leftChildIdx < bh.Size() && bh.list[leftChildIdx].GetPriority() > bh.list[LargestChildIdx].GetPriority() {
			LargestChildIdx = leftChildIdx
		}
		if rightChildIdx < bh.Size() && bh.list[rightChildIdx].GetPriority() > bh.list[LargestChildIdx].GetPriority() {
			LargestChildIdx = rightChildIdx
		}
		if LargestChildIdx == i {
			return
		}
		bh.list[LargestChildIdx], bh.list[i] = bh.list[i], bh.list[LargestChildIdx]
		i = LargestChildIdx
	}
}
