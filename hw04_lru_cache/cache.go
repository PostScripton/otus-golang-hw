package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type cacheItem struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	listItem, ok := c.items[key]
	cached := cacheItem{key: key, value: value}

	if ok {
		listItem.Value = cached
		c.queue.MoveToFront(listItem)
	} else {
		if c.queue.Len() == c.capacity {
			old := c.queue.Back()
			c.queue.Remove(old)
			delete(c.items, old.Value.(cacheItem).key)
		}
		listItem = c.queue.PushFront(cached)
	}
	c.items[key] = listItem

	return ok
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	listItem, ok := c.items[key]

	if ok {
		c.queue.MoveToFront(listItem)
		return listItem.Value.(cacheItem).value, true
	}

	return nil, false
}

func (c *lruCache) Clear() {
	c.queue = NewList()
	c.items = map[Key]*ListItem{}
}
