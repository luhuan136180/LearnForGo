package main

import (
	"fmt"
	"math/big"
)

const (
	AddressLength  = 20
	NodesPerBucket = 3
)

type Node struct {
	ID string
}

type Bucket struct {
	Nodes []Node
	Left  *Bucket
	Right *Bucket
}

func (b *Bucket) InsertNode(node Node) {
	if len(b.Nodes) < NodesPerBucket {
		b.Nodes = append(b.Nodes, node)
		return
	}

	// Calculate the distance between the node ID and the IDs in the bucket
	nodeID := node.ID
	for _, n := range b.Nodes {
		distance := xorDistance(nodeID, n.ID)
		if distance.Cmp(big.NewInt(0)) == 0 {
			return // Node already exists in the bucket
		}
	}

	// Determine whether to insert the node in the left or right subtree
	leftDistance := xorDistance(nodeID, b.Nodes[0].ID)
	if leftDistance.Cmp(big.NewInt(int64(AddressLength*8/2))) <= 0 {
		if b.Left == nil {
			b.Left = &Bucket{}
		}
		b.Left.InsertNode(node)
	} else {
		if b.Right == nil {
			b.Right = &Bucket{}
		}
		b.Right.InsertNode(node)
	}
}

func (b *Bucket) PrintBucketContents() {
	for _, node := range b.Nodes {
		fmt.Println(node.ID)
	}

	if b.Left != nil {
		b.Left.PrintBucketContents()
	}
	if b.Right != nil {
		b.Right.PrintBucketContents()
	}
}

func xorDistance(a, b string) *big.Int {
	xor := new(big.Int)
	xor.SetString(a, 16)
	temp := new(big.Int)
	temp.SetString(b, 16)
	xor.Xor(xor, temp)

	return xor
}

func main() {
	nodeList := []Node{
		{ID: "node1"},
		{ID: "node2"},
		{ID: "node3"},
		{ID: "node4"},
		{ID: "node5"},
	}

	rootBucket := &Bucket{}

	for _, node := range nodeList {
		rootBucket.InsertNode(node)
	}

	rootBucket.PrintBucketContents()
}
