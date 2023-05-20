package consistenthash

import (
	"hash/crc32"
	"sort"
	"strconv"
)

// Hash maps bytes to uint32
//定义了函数类型 Hash，采取依赖注入的方式，允许用于替换成自定义的 Hash 函数，也方便测试时替换，默认为 crc32.ChecksumIEEE 算法。
type Hash func(data []byte) uint32

// Map constains all hashed keys
type Map struct {
	hash     Hash           //
	replicas int            //虚拟节点的倍数
	keys     []int          //哈希环 Sorted
	hashMap  map[int]string //虚拟节点 -> 真实节点的映射表
}

// New creates a Map instance
func New(replicas int, fn Hash) *Map {
	m := &Map{
		replicas: replicas,
		hash:     fn,
		hashMap:  make(map[int]string),
	}
	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE //默认算法
	}
	return m
}

// 添加真实节点，传入真实节点的名称（个数不限），
func (m *Map) Add(keys ...string) {
	for _, key := range keys {
		for i := 0; i < m.replicas; i++ {
			//针对每一个真实节点 key ，对应创建 m.replicas 个虚拟节点，虚拟节点的名称是： strconv.Ttoa(i)+key ,即通过添加编号的方式区分不同的虚拟节点
			hash := int(m.hash([]byte(strconv.Itoa(i) + key))) //m.hash ：计算出每一个虚拟节点的名称的哈希
			m.keys = append(m.keys, hash)                      //**将该虚拟节点的哈希值填入哈希环中
			m.hashMap[hash] = key                              //将虚拟节点的哈希值和对应的真实节点绑定，key：虚拟节点hash值，value：真实节点的名称
		}
	}
	sort.Ints(m.keys) //排序
}

// Get gets the closest item in the hash to the provided key.
// Get获取哈希中最接近所提供键的项。----根据传入的key选择适合的节点存储
func (m *Map) Get(key string) string {
	if len(m.keys) == 0 {
		return ""
	}

	hash := int(m.hash([]byte(key)))
	// Binary search for appropriate replica.
	idx := sort.Search(len(m.keys), func(i int) bool {
		return m.keys[i] >= hash
	})

	return m.hashMap[m.keys[idx%len(m.keys)]] //通过映射获取真实的节点名称
}
