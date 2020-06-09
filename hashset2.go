type MyHashSet struct {
	// 当前大小
	size int
	// 容量
	cap int
	// 负载因子
	growFactor   float64
	shrinkFactor float64
	buf          []*Node
}
type Node struct {
	Val  int
	Next *Node
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	return MyHashSet{
		size:         0,
		growFactor:   0.75,
		shrinkFactor: 0.25,
		cap:          16,
		buf:          make([]*Node, 16),
	}
}

func (this *MyHashSet) Add(key int) {
	if this.Contains(key) {
		return
	}
	this.size++
	if this.needGrow() {
		this.grow()
	}

	hash := this.hashcode(key)
	node := &Node{Val: key}
	if this.buf[hash] == nil {
		this.buf[hash] = node
	} else {
		node.Next = this.buf[hash]
		this.buf[hash] = node
	}

}
func (this *MyHashSet) hashcode(key int) int {
	return key % this.cap
}
func (this *MyHashSet) needGrow() bool {
	return this.size >= int(float64(this.cap)*this.growFactor)
}
func (this *MyHashSet) needShrink() bool {
	return this.size <= int(float64(this.cap)*this.shrinkFactor) && this.cap > 16
}
func (this *MyHashSet) grow() {
	this.cap = this.cap * 2
	this.rehash()
}
func (this *MyHashSet) shrink() {
	this.cap = this.cap / 2
	this.rehash()
}
func (this *MyHashSet) rehash() {
	// lock
	newBuf := make([]*Node, this.cap)
	for i := 0; i < len(this.buf); i++ {
		head := this.buf[i]
		for head != nil {
			hash := this.hashcode(head.Val)
			node := &Node{Val: head.Val}
			if newBuf[hash] == nil {
				newBuf[hash] = node
			} else {
				// 插入到新链表头部
				node.Next = newBuf[hash]
				newBuf[hash] = node
			}
			head = head.Next
		}
	}
	this.buf = newBuf

}

func (this *MyHashSet) Remove(key int) {
	// 自动缩容
	hash := this.hashcode(key)
	if this.buf[hash] == nil {
		return
	} else {
		dummy := &Node{Val: 1<<31 - 1}
		dummy.Next = this.buf[hash]
		head := dummy
		for head != nil && head.Next != nil {
			if head.Next.Val == key {
				t := head.Next
				head.Next = head.Next.Next
				t.Next = nil
				this.size--
				break
			}
			head = head.Next
		}
		this.buf[hash] = dummy.Next
	}
	if this.needShrink() {
		this.shrink()
	}
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	hash := this.hashcode(key)
	if this.buf[hash] == nil {
		return false
	} else {
		head := this.buf[hash]
		for head != nil {
			if head.Val == key {
				return true
			}
			head = head.Next
		}
	}
	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */

// 作者：greyireland
// 链接：https://leetcode-cn.com/problems/design-hashset/solution/dan-lian-biao-he-zi-dong-kuo-rong-suo-rong-by-grey/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。