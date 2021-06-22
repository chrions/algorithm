package main

import "DataStructureAndAlgorithm/LinkedList/Link"

func main() {

	node1 := getNode1()
	node2 := getNode2()
	node3 := Link.MergeSortedLists(node1, node2)
	Link.Traverse(node3)
}

func getNode1() *Link.Node {
	node := &Link.Node{}
	nodeArray := make([]Link.Element, 0)
	for i := 0; i < 5; i++ {
		nodeArray = append(nodeArray, Link.Element(i+i*2))
		Link.Add(node, nodeArray[i])
	}
	return node
}

func getNode2() *Link.Node {
	node := &Link.Node{}
	nodeArray := make([]Link.Element, 0)
	for i := 0; i < 5; i++ {
		nodeArray = append(nodeArray, Link.Element(i+i*1))
		Link.Add(node, nodeArray[i])
	}
	return node
}
