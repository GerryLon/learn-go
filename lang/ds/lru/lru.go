// Least Recently Used
// list & map implements lru algorithm
// list保存数据
// map实现检测key是否存在
package lru

import (
	"container/list"
	"errors"
)

type CacheEntry struct {
	Key, Value interface{}
}

type CacheDB struct {
	Capacity int
	list     *list.List
	cacheMap map[interface{}]*list.Element
}

// 创建一个缓存数据库
func New(cap int) *CacheDB {
	return &CacheDB{
		Capacity: cap,
		list:     list.New(),                          // 初始化链表
		cacheMap: make(map[interface{}]*list.Element), // 初始化内部map
	}
}

// 目前缓存的元素个数
func (c *CacheDB) size() int {
	return c.list.Len()
}

func (c *CacheDB) Set(k, v interface{}) error {
	if c.list == nil {
		return errors.New("cache db is not initialized")
	}

	// 命中缓存
	if ele, ok := c.cacheMap[k]; ok {
		c.list.MoveToFront(ele)           // 刚访问的元素置换到list开始位置
		ele.Value.(*CacheEntry).Value = v // 修改链表元素的值
		return nil
	}

	// 未命中时, 创建新节点, 并插入到链表开始, 表示这是最近访问的
	newEle := c.list.PushFront(&CacheEntry{k, v})
	c.cacheMap[k] = newEle

	// 如果缓存已满, 移除最后一个元素(也就是最近最少使用的)
	if c.list.Len() > c.Capacity {
		lastEle := c.list.Back() // 链表最后一个元素

		// 移除最后一个元素, list.Remove方法会返回节点的value
		cacheEntry := c.list.Remove(lastEle).(*CacheEntry)
		delete(c.cacheMap, cacheEntry.Key) // 移除相应的key
	}
	return nil
}

func (c *CacheDB) Get(key interface{}) (v interface{}, success bool, err error) {
	if c.list == nil {
		return nil, false, errors.New("cache db is not initialized")
	}

	// get操作命中缓存
	if ele, ok := c.cacheMap[key]; ok {
		c.list.MoveToFront(ele) // 将对应的元素置换到链表头, 表示最近访问过了
		return ele.Value.(*CacheEntry).Value, true, nil
	}

	return nil, false, nil
}

func (c *CacheDB) Remove(key interface{}) bool {
	if c.list == nil {
		return false
	}

	// 从链表的map中删除相应的元素
	if ele, ok := c.cacheMap[key]; ok {
		cacheEntry := c.list.Remove(ele)
		delete(c.cacheMap, cacheEntry.(*CacheEntry).Key)
		return true
	}

	return false
}

// 辅助函数, 遍历缓存的值
func (c *CacheDB) Traverse(f func(k, v interface{}) bool) {
	var cacheEntry *CacheEntry
	for e := c.list.Front(); e != nil; e = e.Next() {
		cacheEntry = e.Value.(*CacheEntry)
		if !f(cacheEntry.Key, cacheEntry.Value) {
			break
		}
	}
}
