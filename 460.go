// https://leetcode-cn.com/problems/lfu-cache/



package main

type LFUCache struct {
    Keys [100]int
    Values [100]int
    Count [100]int
}


func Constructor(capacity int) LFUCache {
    lfu := LFUCache{}

    return lfu
}


func (this *LFUCache) Get(key int) int {
    if this.Values[key] == 0 {
        return 0
    }

    this.Values[key] += 1
    return this.Values[key]
}


func (this *LFUCache) Put(key int, value int){
    this.Values[key] = value
}


/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */