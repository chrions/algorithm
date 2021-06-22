package heap

type head struct {
	data  []int //存储格式，数组，从下标1开始
	count int   //堆中已经存储的数据个数
	num   int   //堆中最大存储的个数
}

func NewHead(capacity int) *head {
	return &head{
		data:  make([]int, capacity+1),
		count: 0,
		num:   capacity,
	}
}

//新插入的数，放到最后一位，和他的父节点对比
func (this *head) InsertData(v int) {
	if this.num == this.count {
		return
	} //存满了
	this.count++
	this.data[this.count] = v
	i := this.count
	for i/2 > 0 && this.data[i] > this.data[i/2] {
		swap(this.data, i, i/2)
		i = i / 2
	}
}

func swap(data []int, i int, i2 int) {
	data[i], data[i2] = data[i2], data[i]
}

//把最后一位放在头结点，在和其子结点对比
func (this *head) RemoveMax() {
	if this.count == 0 {
		return
	} //堆内没有数据

	this.data[1] = this.data[this.count]
	this.count--
	heapify(this.data, this.count)
}

func heapify(data []int, count int) {
	for i := 1; i <= count/2; {
		maxIndex := i
		if data[i] < data[i*2] {
			maxIndex = i * 2
		}
		if i*2+1 <= count && data[maxIndex] < data[i*2+1] {
			maxIndex = i*2 + 1
		}
		if maxIndex == i {
			break
		}
		swap(data, i, maxIndex)
		i = maxIndex
	}
}

func BuildHeap(data []int, n int) {
	for i := n / 2; i > 0; i-- {
		buildHeap(data, i, n)
	}
}

func buildHeap(data []int, i, n int) {

}
