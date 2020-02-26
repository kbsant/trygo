package main

import (
	"fmt"
)

// Initialize the connection map.
// A node is represented by a nonzero integer, and its entry in the nodes array
// is the node it is connected to.
// Zero represents the unconnected node, so initially, every entry is zero
func initialize(max int) []int {
	nodes := make([]int, max+1)
	return nodes
}

// Check whether 2 nodes are connected.
// Group id 0 is unconnected, so skip those.
// If nodes or their ancestors have the same group, they are connected.
// If nodes are connected by their ancestors, connect them directly.
// TODO speed this up, since some logic is duplicated between merge and connected.
func connected(nodes []int, a int, b int) bool {
	i, j := sort2(a, b)
	for {
		if i == 0 {
			return false
		}
		if i == j {
			merge(nodes, a, b)
			return true
		}
		i, j = sort2(i, nodes[j])
	}
	fmt.Errorf("Error checking connection between %v and %v\n", a, b)
	return false
}

func sort2(a, b int) (int, int) {
	if a < b {
		return a, b
	} else {
		return b, a
	}
}

// Link the second node to the first, or its ancestor if present.
func merge(nodes []int, a int, b int) {
	lo, hi := sort2(a, b)
	for {
		parent := nodes[lo]
		if parent == 0 {
			break
		}
		lo = parent
	}
	nodes[hi] = lo
}

func main() {
	fmt.Println("hello, this is the connection checker v2.")

	nodes := initialize(10)
	fmt.Println("start with ", nodes)
	nodesConnected := func(a, b int) bool {
		return connected(nodes, a, b)
	}
	fmt.Printf("Are %v and %v connected? %v\n", 2, 5, nodesConnected(2, 5))
	fmt.Printf("Are %v and %v connected? %v\n", 6, 4, nodesConnected(6, 4))
	fmt.Printf("Are %v and %v connected? %v\n", 5, 7, nodesConnected(5, 7))
	fmt.Printf("Are %v and %v connected? %v\n", 5, 6, nodesConnected(5, 6))
	fmt.Printf("Are %v and %v connected? %v\n", 2, 7, nodesConnected(2, 7))
	fmt.Printf("Are %v and %v connected? %v\n", 2, 4, nodesConnected(2, 4))

	fmt.Println("Merge 5 and 7.")
	fmt.Println("Merge 6 and 4.")
	merge(nodes, 5, 7)
	merge(nodes, 6, 4)
	fmt.Printf("Are %v and %v connected? %v\n", 2, 5, nodesConnected(2, 5))
	fmt.Printf("Are %v and %v connected? %v\n", 6, 4, nodesConnected(6, 4))
	fmt.Printf("Are %v and %v connected? %v\n", 5, 7, nodesConnected(5, 7))
	fmt.Printf("Are %v and %v connected? %v\n", 5, 6, nodesConnected(5, 6))
	fmt.Printf("Are %v and %v connected? %v\n", 2, 7, nodesConnected(2, 7))
	fmt.Printf("Are %v and %v connected? %v\n", 2, 4, nodesConnected(2, 4))

	fmt.Println("Merge 5 and 2.")
	merge(nodes, 5, 2)
	fmt.Printf("Are %v and %v connected? %v\n", 2, 5, nodesConnected(2, 5))
	fmt.Printf("Are %v and %v connected? %v\n", 6, 4, nodesConnected(6, 4))
	fmt.Printf("Are %v and %v connected? %v\n", 5, 7, nodesConnected(5, 7))
	fmt.Printf("Are %v and %v connected? %v\n", 5, 6, nodesConnected(5, 6))
	fmt.Printf("Are %v and %v connected? %v\n", 2, 7, nodesConnected(2, 7))
	fmt.Printf("Are %v and %v connected? %v\n", 2, 4, nodesConnected(2, 4))

	fmt.Println("end with ", nodes)
}
