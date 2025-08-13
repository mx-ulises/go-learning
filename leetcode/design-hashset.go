package leetcode

import "math"

type MyHashSet struct {
	Mem [15874]int
}

func Constructor() MyHashSet {
	return MyHashSet{Mem: [15874]int{}}
}

func (this *MyHashSet) Add(key int) {
	bucket := key / 63
	flag := 1 << (key % 63)
	this.Mem[bucket] |= flag
}

func (this *MyHashSet) Remove(key int) {
	bucket := key / 63
	flag := 1 << (key % 63)
	this.Mem[bucket] &= math.MaxInt ^ flag
}

func (this *MyHashSet) Contains(key int) bool {
	bucket := key / 63
	flag := 1 << (key % 63)
	return this.Mem[bucket]&flag != 0
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
