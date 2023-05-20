package geecache

import (
	"fmt"
	"geecache6/singleflight"
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
	name      string
	getter    Getter
	mainCache cache

	peers PeerPicker

	loader *singleflight.Group
}

var (
	mu     sync.RWMutex
	groups = make(map[string]*Group)
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
		loader:    &singleflight.Group{}, //给一个空的
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

//
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

//将实现了 PeerPicker 接口的HTTPPool 注入到Group中
// RegisterPeers registers a PeerPicker for choosing remote peer
func (g *Group) RegisterPeers(peers PeerPicker) {
	if g.peers != nil {
		panic("RegisterPeerPicker called more than once")
	}
	g.peers = peers
}

func (g *Group) load(key string) (value ByteView, err error) {
	//原来的 load 的逻辑，使用 g.loader.Do 包裹起来即可，这样确保了并发场景下针对相同的 key，load 过程只会调用一次。
	viewi, err := g.loader.Do(key, func() (interface{}, error) {
		if g.peers != nil {
			//根据g.peers 的具体结构体的方法
			if peer, ok := g.peers.PickPeer(key); ok {
				if value, err = g.getFromPeer(peer, key); err == nil {
					return value, nil
				}
				log.Println("[GeeCache] Failed to get from peer", err)
			}
		}
		//失败，回退函数
		return g.getlocally(key)
	})

	if err == nil {
		return viewi.(ByteView), nil
	}
	return
}

//使用实现了PeerGetter接口的httpGetter从访问远程节点，获取缓存值
func (g *Group) getFromPeer(peer PeerGetter, key string) (ByteView, error) {
	bytes, err := peer.Get(g.name, key)
	if err != nil {
		return ByteView{}, err
	}
	return ByteView{b: bytes}, nil
}

//singleflight的使用
