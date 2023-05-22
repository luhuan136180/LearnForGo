package main

import (
	"fmt"
	"strings"
	"time"
)

type Node struct {
	nodeId       string
	lastSeenTime time.Time
}

type Bucket struct {
	nodes []*Node
}

func (bucket *Bucket) insertNode(node *Node) {
	if len(bucket.nodes) < 3 {
		bucket.nodes = append(bucket.nodes, node)
	} else {
		if !strings.HasPrefix(bucket.nodes[0].nodeId, node.nodeId) {
			return
		}

		// Split bucket into two buckets of equal size
		splitIndex := len(bucket.nodes) / 2
		leftBucket := &Bucket{nodes: bucket.nodes[:splitIndex]}
		rightBucket := &Bucket{nodes: bucket.nodes[splitIndex:]}

		if strings.HasPrefix(leftBucket.nodes[0].nodeId, node.nodeId) {
			leftBucket.insertNode(node)
			bucket.nodes = leftBucket.nodes
			rightBucket = bucket
		} else {
			rightBucket.insertNode(node)
			bucket.nodes = rightBucket.nodes
			leftBucket = bucket
		}

		// Reassign nodes to the appropriate buckets
		for _, oldNode := range bucket.nodes {
			if strings.HasPrefix(leftBucket.nodes[0].nodeId, oldNode.nodeId) {
				leftBucket.insertNode(oldNode)
			} else {
				rightBucket.insertNode(oldNode)
			}
		}
	}
}

func (bucket *Bucket) removeNode(nodeId string) {
	for i, node := range bucket.nodes {
		if node.nodeId == nodeId {
			bucket.nodes = append(bucket.nodes[:i], bucket.nodes[i+1:]...)
			break
		}
	}
}

func (bucket *Bucket) updateNode(nodeId string) {
	for _, node := range bucket.nodes {
		if node.nodeId == nodeId {
			node.lastSeenTime = time.Now()
			break
		}
	}
}

type KBucket struct {
	buckets [20]*Bucket
}

func NewKBucket() *KBucket {
	kBucket := &KBucket{}
	for i := 0; i < len(kBucket.buckets); i++ {
		kBucket.buckets[i] = &Bucket{}
	}
	return kBucket
}

func (kBucket *KBucket) getBucketIndex(nodeId string) int {
	distance := 0
	for i := 0; i < len(nodeId) && i < len(kBucket.buckets); i++ {
		if nodeId[i] != kBucket.buckets[i].nodes[0].nodeId[i] {
			break
		}
		distance++
	}
	return distance
}

func (kBucket *KBucket) insertNode(nodeId string) {
	distance := kBucket.getBucketIndex(nodeId)
	kBucket.buckets[distance].insertNode(&Node{nodeId: nodeId, lastSeenTime: time.Now()})
}

func (kBucket *KBucket) printBucketContents() {
	for i, bucket := range kBucket.buckets {
		fmt.Printf("Bucket %d:\n", i)
		for _, node := range bucket.nodes {
			fmt.Println(node.nodeId)
		}
		fmt.Println()
	}
}

func main() {
	kBucket := NewKBucket()

	// 初始化插入6个节点
	kBucket.insertNode("0001")
	kBucket.insertNode("0002")
	kBucket.insertNode("0003")
	kBucket.insertNode("0010")
	kBucket.insertNode("0011")
	kBucket.insertNode("0012")

	// 打印所有桶的内容
	kBucket.printBucketContents()

	// 更新节点0010的时间为当前时间
	kBucket.buckets[3].updateNode("0010")

	// 删除节点0012
	kBucket.buckets[3].removeNode("0012")

	// 打印更新和删除后的桶的内容
	kBucket.printBucketContents()

	// 新增节点0013，触发桶分裂
	kBucket.insertNode("0013")

	// 打印分裂后的桶的内容
	kBucket.printBucketContents()

	fmt.Println("Done")
}
