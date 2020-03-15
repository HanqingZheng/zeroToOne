package dataStruct

type segmentTree struct {
	data []int
	tree []int
}

func NewSegmentTree(num []int) *segmentTree {
	countNum := len(num)
	data := make([]int, countNum)
	for k, v := range num {
		data[k] = v
	}
	tree := make([]int, 4*countNum)
	if countNum > 0 {
		var buildTree func(int, int, int)
		buildTree = func(index, left, right int) {
			if left == right {
				tree[index] = num[left]
				return
			}
			leftChild := leftChild(index)
			rightChild := rightChild(index)
			mid := left + ((right - left) >> 1)
			buildTree(leftChild, left, mid)
			buildTree(rightChild, mid+1, right)
			tree[index] = tree[leftChild] + tree[rightChild]
		}
		buildTree(0, 0, countNum-1)
	}
	return &segmentTree{data, tree}
}

func (st *segmentTree) SumRange(start, end int) int {
	var sum func(int, int, int, int, int) int
	sum = func(index, left, right, start, end int) int {
		if left == start && right == end {
			return st.tree[index]
		}
		leftChild := leftChild(index)
		rightChild := rightChild(index)
		mid := left + ((right - left) >> 1)
		if start >= mid+1 {
			return sum(rightChild, mid+1, right, start, end)
		} else if end <= mid {
			return sum(leftChild, left, mid, start, end)
		}
		return sum(leftChild, left, mid, start, mid) + sum(rightChild, mid+1, right, mid+1, end)
	}
	return sum(0, 0, len(st.data)-1, start, end)
}

func (st *segmentTree) Update(i int, value int) {
	countNum := len(st.data)
	if i >= len(st.data) {
		return
	}
	st.data[i] = value
	var up func(int, int, int)
	up = func(index, left, right int) {
		if left == right {
			st.tree[index] = value
			return
		}
		leftChild := leftChild(index)
		rightChild := rightChild(index)
		mid := left + ((right - left) >> 1)
		if i >= mid+1 {
			up(rightChild, mid+1, right)
		} else if i <= mid {
			up(leftChild, left, mid)
		}
		st.tree[index] = st.tree[leftChild] + st.tree[rightChild]
	}
	up(0, 0, countNum-1)
}

func leftChild(i int) int {
	return (i << 1) + 1
}
func rightChild(i int) int {
	return (i << 1) + 2
}
