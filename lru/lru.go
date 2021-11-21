package lru

import (
	"container/list"
	"sync"
)

type LRU struct {
	maxByte int64                    //最多存储数据的字节个数，超过此数量便会触发数据的淘汰
	curByte int64                    //目前存储的字节个数
	list    *list.List               //使用go语言内置的双向链表存储节点
	cache   map[string]*list.Element //通过节点的key快速定位到属于哪个节点，不需遍历双向链表
	mu      sync.RWMutex             //读写锁，保证并发安全
}

type Entry struct {
	Key   string //每个节点的唯一标识，作为key储存到lru的cache里
	Value []byte //携带的数据
}

// New Constructor of LRU
func New(maxByte int64) *LRU {
	return &LRU{
		maxByte: maxByte,
		curByte: 0,
		list:    list.New(),
		cache:   make(map[string]*list.Element),
	}
}

// Get look up a key`s value
func (l *LRU) Get(key string) ([]byte, bool) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	if ele, exist := l.cache[key]; exist {
		l.list.MoveToFront(ele) //此元素被访问，移动到链表最前面
		if entry, ok := ele.Value.(*Entry); ok {
			return entry.Value, true
		}
	}

	return nil, false
}

// RemoveOldest remove entry from linklist back
func (l *LRU) RemoveOldest() {
	l.mu.Lock()
	defer l.mu.Unlock()

	ele := l.list.Back() //取出链表尾部节点
	if ele != nil {
		l.list.Remove(ele) //删除节点
		if entry, ok := ele.Value.(*Entry); ok {
			delete(l.cache, entry.Key)                                   //哈希表删除该节点key
			l.curByte -= int64(len(entry.Key)) + int64(len(entry.Value)) // 调整已使用字节数
		}
	}
}

// Add a value to lru
func (l *LRU) Add(key string, value []byte) {
	l.mu.Lock()
	defer l.mu.Unlock()

	// 两种情况
	if ele, ok := l.cache[key]; ok {
		l.list.MoveToFront(ele)
		if entry, ok := ele.Value.(*Entry); ok {
			l.curByte += int64(len(value)) - int64(len(entry.Value)) // value改变的话需要改变已用字节数
			entry.Value = value
		}
	} else {
		ele := l.list.PushFront(&Entry{Key: key, Value: value})
		l.cache[key] = ele
		l.curByte = int64(len(key)) + int64(len(value))
	}

	// 已使用字节数大于最大字节数时，需要移除链表尾部节点，知到已使用字节数小于最大字节数
	for l.maxByte > 0 && l.maxByte < l.curByte {
		l.RemoveOldest()
	}
}

// Len the number of cache entries
func (l *LRU) Len() int {
	return l.list.Len()
}
