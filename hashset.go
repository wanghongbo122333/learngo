package main

// initialize hashset without inner hashset
import "fmt"

type MyHashSet struct {
	array [1000001]bool
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	myhashset := new(MyHashSet)

	return *myhashset
}

func (this *MyHashSet) Add(key int) {
	this.array[key] = true
}

func (this *MyHashSet) Remove(key int) {
	this.array[key] = false
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	return this.array[key]
}

func main() {
	key := 1
	obj := Constructor()
	obj.Add(key)
	param_3 := obj.Contains(key)
	fmt.Print(param_3)
	obj.Remove(key)
	param_3 = obj.Contains(key)
	fmt.Print(param_3)
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
