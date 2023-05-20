package geecache

import (
	"fmt"
	"log"
	"sync"
)

/*
如果缓存不存在，应从数据源（文件，数据库等）获取数据并添加到缓存中。
GeeCache 是否应该支持多种数据源的配置呢？不应该，一是数据源的种类太多，没办法一一实现；二是扩展性不好。如何从源头获取数据，应该是用户决定的事情，我们就把这件事交给用户好了。因此，我们设计了一个回调函数(callback)，
在缓存不存在时，调用这个函数，得到源数据。
*/
type Getter interface {
	Get(key string) ([]byte, error)
}

/*
定义一个函数类型 F，并且实现接口 A 的方法，然后在这个方法中调用自己。
这是 Go 语言中将其他函数（参数返回值定义与 F 一致）转换为接口 A 的常用技巧。
*/

// A GetterFunc implements Getter with a function.
//定义函数类型 GetterFunc，并实现 Getter 接口的 Get 方法。
type GetterFunc func(key string) ([]byte, error)

// Get implements Getter interface function
func (f GetterFunc) Get(key string) ([]byte, error) {
	return f(key)
}

type Group struct {
	name string //每个 Group 拥有一个唯一的名称 name
	//比如可以创建三个 Group，缓存学生的成绩命名为 scores，缓存学生信息的命名为 info，缓存学生课程的命名为 courses。
	getter    Getter //即缓存未命中时获取源数据的回调(callback)---用于在group中没有所需数据时，能过通过回调函数获取值
	mainCache cache  //实现的并发缓存
}

var (
	mu     sync.RWMutex
	groups = make(map[string]*Group) //全局变量
)

func NewGroup(name string, cacheBytes int64, getter Getter) *Group {
	if getter == nil {
		panic("nil Getter")
	}
	mu.Lock()
	defer mu.Unlock()

	g := &Group{
		name:      name,
		getter:    getter,
		mainCache: cache{cacheBytes: cacheBytes},
	}
	groups[name] = g
	return g
}

//获取对应名字的group
func GetGroup(name string) *Group {
	mu.RLock()
	g := groups[name]
	mu.Unlock()

	return g

}

func (g *Group) Get(key string) (ByteView, error) {
	if key == "" {
		return ByteView{}, fmt.Errorf("key is required")
	}

	if v, ok := g.mainCache.get(key); ok { //存在缓存值
		log.Println("[Geecache] hit")
		return v, nil
	}
	//缓存不存在,
	return g.load(key)
}

//流程 ⑶ ：缓存不存在，则调用 load 方法，load 调用 getLocally（分布式场景下会调用 getFromPeer 从其他节点获取），
//getLocally 调用用户回调函数 g.getter.Get() 获取源数据，并且将源数据添加到缓存 mainCache 中（通过 populateCache 方法）
func (g *Group) load(key string) (value ByteView, err error) {
	return g.getlocally(key)
}

func (g *Group) getlocally(key string) (ByteView, error) {
	bytes, err := g.getter.Get(key)
	if err != nil {
		return ByteView{}, err
	}

	value := ByteView{b: cloneBytes(bytes)}
	g.populateCache(key, value)
	return value, nil
}

func (g *Group) populateCache(key string, value ByteView) {
	g.mainCache.add(key, value)
}
