// ring is a simple consistent hashing package
package ring

import "hash/crc32"
import "fmt"

// Node is a frame that contain the real-node value and weight
type Node struct {
	Value	string
	Weight 	int
}

// Vnode is frame that contains the real-node "hash" and its new virtual hash
type Vnode struct {
	Real	uint32
	Virtual	uint32
}

// Ring is our main ring that contains the real/virtual nodes
type Ring struct {
	Nodes		map[uint32]Node
	Ring 		[]Vnode
}

// Create a new ring
func NewRing() *Ring {
	this := new(Ring)
	this.Nodes = map[uint32]Node{}
	this.Ring = []Vnode{}
	return this
}

// Return the hash of the specified key
func (this *Ring) Hash(key string) uint32 {
	return crc32.ChecksumIEEE([]byte(key))
}

// Add new node with a {weight "how many replicas ?"}
func (this *Ring) Add(node string, weight int) *Ring {
	hash := this.Hash(node)
	this.Nodes[hash] = Node{node, weight}
	for i := 0; i < weight; i ++ {
		this.Ring = append(this.Ring, Vnode{hash, this.Hash(fmt.Sprintf("%s:%d", node, i))})
	}
	return this
}

// Return the key's node
func (this *Ring) Get(key string) string {
	hash := this.Hash(key)
	for _, v := range(this.Ring) {
		if (hash <= v.Virtual) && (this.Nodes[v.Real].Value != "") {
			fmt.Println(this.Nodes[v.Real])
			return this.Nodes[v.Real].Value
		}
	}
	return this.Nodes[this.Ring[0].Real].Value
}

// Remove a node from the ring
func (this *Ring) Remove(node string) *Ring {
	hash := this.Hash(node)
	for k,v := range(this.Ring) {
		if hash == v.Real {
			this.Ring = append(this.Ring[:k], this.Ring[k:]...)
		}
	}
	delete(this.Nodes, hash)
	return this
}
