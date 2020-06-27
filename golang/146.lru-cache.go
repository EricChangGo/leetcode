/*
 * @lc app=leetcode id=146 lang=golang
 *
 * [146] LRU Cache
 */

// @lc code=start
type Node struct {
    key int // use to delete key in map
    val int
    prev *Node
    next *Node
}

type LRUCache struct {
    capacity int
    leastRecentUsed  *Node // Record the last element in orderChain
    orderChain *Node // New -> Old
    cache map[int]*Node // Use map to gurantee operations excute in O(1) 
}

func Constructor(capacity int) LRUCache {
    return LRUCache{
		capacity: capacity,
		cache: make(map[int]*Node),
	}
}

func (this *LRUCache) addToChain(tNode *Node) {
    /** add the node into list*/
    tNode.prev = nil 
	if this.orderChain != nil {
        // cache - head is not empty
        this.orderChain.prev = tNode
        tNode.next = this.orderChain
	}
	this.orderChain = tNode
    
    if this.leastRecentUsed == nil {
        this.leastRecentUsed = tNode
    }
}

func (this *LRUCache) removeFromChain(tNode *Node) {
    /** remove targetNode from the list*/
    if this.leastRecentUsed == tNode {
        this.leastRecentUsed = this.leastRecentUsed.prev
    }
    if this.orderChain == tNode {
        this.orderChain = this.orderChain.next
    }
    
    if tNode.next != nil {
        tNode.next.prev = tNode.prev
    }
    if tNode.prev != nil {
        tNode.prev.next = tNode.next
    }
}

func (this *LRUCache) Get(key int) int {
    /** get value of the key */
    has_value := -1
    tNode, ok := this.cache[key]
    if ok {
        // get key in Cache
        has_value = this.cache[key].val
        this.removeFromChain(tNode)
        this.addToChain(tNode)
    }
    return has_value
}

func (this *LRUCache) Put(key int, value int)  {
    /** 
        put (key, value), 
        if reached the capacity limitation -> pop the recently used
    */
    tNode, ok := this.cache[key]
    if ok {
        tNode.val = value
		this.removeFromChain(tNode)
    } else {
        if len(this.cache) == this.capacity {
            // pop the recentUsed node
            removeNode := this.leastRecentUsed
            this.removeFromChain(this.leastRecentUsed)
            delete(this.cache, removeNode.key)
        }
        // create a new node space for (key, value)
        newNode := &Node{key: key, val: value}
		this.cache[key] = newNode
        tNode = newNode
    }
    this.addToChain(tNode)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
