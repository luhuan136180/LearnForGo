package lru

import "container/list"

//list包实现了双向链表。要遍历一个链表：
type Cache struct {
	maxBytes int64      //允许使用的最大内存
	nbytes   int64      //当前已使用内存
	ll       *list.List //List代表一个双向链表。List零值为一个空的、可用的链表。

	// 字典 键是字符串，值是双向链表中对应节点的指针。
	cache map[string]*list.Element //Element类型代表是双向链表的一个元素。
	// 某条记录被移除时的回调函数，可以为 nil。
	onEvicted func(key string, value Value)
}

//双向链表节点的数据结构
type entry struct {
	key   string
	value Value
}

type Value interface {
	// Value使用Len来计算需要多少字节
	Len() int
}

//cache 的实例化
func New(maxBytes int64, onEvicted func(string, Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		ll:        list.New(),                     //初始化
		cache:     make(map[string]*list.Element), //初始化,
		onEvicted: onEvicted,
	}
}

//查找有两步，第一步从字典中找到对应的双向链表的节点，第二步将该节点移动到队尾
func (c *Cache) Get(key string) (value Value, ok bool) {
	if ele, ok := c.cache[key]; ok {
		//键对应的链表节点存在，则将对应节点移动到队尾，并返回查找到的值。
		//c.ll.MoveToFront(ele)，即将链表中的节点 ele 移动到队尾,,约定 front 为队尾
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		return kv.value, true
	}
	return
}

//实际上是缓存淘汰。即移除最近最少访问的节点（队首）
func (c *Cache) RemoveOldest() {
	ele := c.ll.Back() //获取双向链表的队首元素

	if ele != nil { //有这个元素
		//func (l *List) Remove(e *Element) interface{}:Remove删除链表中的元素e，并返回e.Value。
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)
		delete(c.cache, kv.key) //将字典中key等于某值的映射关系
		//更新当前所用的内存 c.nbytes。
		c.nbytes -= int64(len(kv.key)) + int64(kv.value.Len())
		if c.onEvicted != nil {
			c.onEvicted(kv.key, kv.value)
		}
	}
}

//如果键存在，则更新对应节点的值，并将该节点移到队尾。
//不存在则是新增场景，首先队尾添加新节点 &entry{key, value}, 并字典中添加 key 和节点的映射关系。
//更新 c.nbytes，如果超过了设定的最大值 c.maxBytes，则移除最少访问的节点。
func (c *Cache) Add(key string, value Value) {
	if ele, ok := c.cache[key]; ok { //存在，更新
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		c.nbytes += int64(value.Len()) - int64(kv.value.Len()) //改变以用的内存大小数值
		kv.value = value
	} else {
		ele := c.ll.PushFront(&entry{key, value})
		c.cache[key] = ele
		c.nbytes += int64(len(key)) + int64(value.Len())
	}

	for c.maxBytes != 0 && c.maxBytes < c.nbytes {
		c.RemoveOldest()
	}
}

func (c *Cache) Len() int {
	return c.ll.Len()
}
