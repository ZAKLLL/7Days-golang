package consistenthash

import (
	"hash/crc32"
	"sort"
	"strconv"
)

// Hash maps bytes to uint32
type Hash func(data []byte) uint32

// Map constains all hashed vNodes
type Map struct {
	hash     Hash
	replicas int   //虚拟节点复制数量
	vNodes   []int // Sorted
	hashMap  map[int]string
}

// New creates a Map instance
func New(replicas int, fn Hash) *Map {
	m := &Map{
		replicas: replicas,
		hash:     fn,
		hashMap:  make(map[int]string),
	}
	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE
	}
	return m
}

// Add adds some vNodes to the hash.
func (m *Map) Add(nodes ...string) {
	for _, node := range nodes {
		for i := 0; i < m.replicas; i++ {
			hash := int(m.hash([]byte(strconv.Itoa(i) + node)))
			m.vNodes = append(m.vNodes, hash)
			m.hashMap[hash] = node
		}
	}
	sort.Ints(m.vNodes)
}

// Get gets the closest item in the hash to the provided key.
func (m *Map) Get(key string) string {
	if len(m.vNodes) == 0 {
		return ""
	}

	hash := int(m.hash([]byte(key)))
	// Binary search for appropriate replica.
	//二分查找第一个>=hash 值的key
	idx := sort.Search(len(m.vNodes), func(i int) bool {
		return m.vNodes[i] >= hash
	})

	//这里之所以要idx%len(m.vNodes)
	//是因为 二分搜索的情况下 idx 可能== len(m.vNodes),因为当前key对应的hash可能是比现有的hash值都大
	//防止越界
	return m.hashMap[m.vNodes[idx%len(m.vNodes)]]
}
