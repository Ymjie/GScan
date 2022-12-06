package ringbuffer

import (
	"errors"
	"fmt"
	"sync/atomic"
)

type T interface{}

var ErrIsEmpty = errors.New("ringbuffer is empty")

type cell struct {
	Data     []T   // 数据部分
	fullFlag bool  // cell满的标志
	next     *cell // 指向后一个cellBuffer
	pre      *cell // 指向前一个cellBuffer

	r int // 下一个要读的指针
	w int // 下一个要下的指针
}

type RingBuffer struct {
	cellSize  int   // cell大小
	cellCount int   // cell数量
	count     int32 // 有效元素个数

	readCell  *cell // 下一个要读的cell
	writeCell *cell // 下一个要写的cell
}

// NewRingBuffer 新建一个RingBuffer，包含两个cell
func NewRingBuffer(cellSize int) (buf *RingBuffer, err error) {
	if cellSize <= 0 || cellSize&(cellSize-1) != 0 {
		err = fmt.Errorf("初始大小必须是 2 的幂")
		return
	}

	rootCell := &cell{
		Data: make([]T, cellSize),
	}
	lastCell := &cell{
		Data: make([]T, cellSize),
	}
	rootCell.pre = lastCell
	lastCell.pre = rootCell
	rootCell.next = lastCell
	lastCell.next = rootCell

	buf = &RingBuffer{
		cellSize:  cellSize,
		cellCount: 2,
		count:     0,
		readCell:  rootCell,
		writeCell: rootCell,
	}

	return
}

// Read 读取数据
func (r *RingBuffer) Read() (data T, err error) {
	// 无数据
	if r.IsEmpty() {
		err = ErrIsEmpty
		return
	}

	// 读取数据，并将读指针向右移动一位
	data = r.readCell.Data[r.readCell.r]
	r.readCell.r++
	atomic.AddInt32(&r.count, -1)

	// 此cell已经读完
	if r.readCell.r == r.cellSize {
		// 读指针归零，并将该cell状态置为非满
		r.readCell.r = 0
		r.readCell.fullFlag = false
		// 将readCell指向下一个cell
		r.readCell = r.readCell.next
	}

	return
}

// Pop 读一个元素，读完后移动指针
func (r *RingBuffer) Pop() (data T) {
	data, err := r.Read()
	if errors.Is(err, ErrIsEmpty) {
		panic(ErrIsEmpty.Error())
	}
	return
}

// Peek 窥视 读一个元素，仅读但不移动指针
func (r *RingBuffer) Peek() (data T) {
	if r.IsEmpty() {
		panic(ErrIsEmpty.Error())
	}

	// 仅读
	data = r.readCell.Data[r.readCell.r]
	return
}

// Write 写入数据
func (r *RingBuffer) Write(value T) {
	// 在 r.writeCell.w 位置写入数据，指针向右移动一位
	r.writeCell.Data[r.writeCell.w] = value
	r.writeCell.w++
	atomic.AddInt32(&r.count, 1)

	// 当前cell写满了
	if r.writeCell.w == r.cellSize {
		// 指针置0，将该cell标记为已满，并指向下一个cell
		r.writeCell.w = 0
		r.writeCell.fullFlag = true
		r.writeCell = r.writeCell.next
	}

	// 下一个cell也已满，扩容
	if r.writeCell.fullFlag == true {
		r.grow()
	}
}

// grow 扩容
func (r *RingBuffer) grow() {
	// 新建一个cell
	newCell := &cell{
		Data: make([]T, r.cellSize),
	}

	// 总共三个cell，writeCell，preCell，newCell
	// 本来关系： preCell <===> writeCell
	// 现在将newcell插入：preCell <===> newCell <===> writeCell
	pre := r.writeCell.pre
	pre.next = newCell
	newCell.pre = pre
	newCell.next = r.writeCell
	r.writeCell.pre = newCell

	// 将writeCell指向新建的cell
	r.writeCell = r.writeCell.pre

	// cell 数量加一
	r.cellCount++
}

// IsEmpty 判断RingBuffer是否为空
func (r *RingBuffer) IsEmpty() bool {
	return r.Len() == 0
}

// Capacity RingBuffer容量
func (r *RingBuffer) Capacity() int {
	return r.cellCount * r.cellSize
}

// Len RingBuffer数据长度
func (r *RingBuffer) Len() (count int) {
	count = int(r.count)
	return
}

// Reset 重置为仅指向两个cell的ring
func (r *RingBuffer) Reset() {
	// 没有数据切cellCount只有两个时，无需重置
	if r.count == 0 && r.cellCount == 2 {
		return
	}

	lastCell := r.readCell.next

	lastCell.w = 0
	lastCell.r = 0
	r.readCell.r = 0
	r.readCell.w = 0
	r.cellCount = 2
	r.count = 0

	lastCell.next = r.readCell
}
