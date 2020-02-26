package conncheck2

import (
	"fmt"
)

// Initialize the connection map.
// Node zero is a virtual node, representing leaf nodes (non-membership).
// Node max + 1 is a virtual node, representing all nodes.
// Initially, every node has its own group since it isn't connected.
// So every node's group is zero.
func initialize(max int) []int {
	nodes := make([]int, max+2)
	nodes[max+1] = max + 1
	return nodes
}

// Check whether 2 nodes are connected.
// Group id 0 is unconnected, so skip those.
// If nodes or their ancestors have the same group, they are connected.
// If nodes are connected by their ancestors, connect them directly.
func connected(nodes []int, a int, b int) bool {
	i, j := a, b
	for {
		left, right := nodes[i], nodes[j]
		if left == 0 || right == 0 {
			return false
		}
		if left == right {
			return true
		}
		upleft, upright := nodes[left], nodes[right]
		nodes[i] = upleft
		i = upleft
		nodes[j] = upright
		j = upright
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

// Merge 2 nodes by assigning them to the same group.
func merge(nodes []int, a int, b int) {
	lo, hi := sort2(a, b)
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
	fmt.Printf("Are %v and %v connected? %v\n", 5, 7, nodesConnected(5, 7))
	fmt.Printf("Are %v and %v connected? %v\n", 2, 7, nodesConnected(2, 7))

	fmt.Println("Merge 5 and 7.")
	merge(nodes, 5, 7)
	fmt.Printf("Are %v and %v connected? %v\n", 2, 5, nodesConnected(2, 5))
	fmt.Printf("Are %v and %v connected? %v\n", 5, 7, nodesConnected(5, 7))
	fmt.Printf("Are %v and %v connected? %v\n", 2, 7, nodesConnected(2, 7))

	fmt.Println("Merge 5 and 2.")
	merge(nodes, 5, 2)
	fmt.Printf("Are %v and %v connected? %v\n", 2, 5, nodesConnected(2, 5))
	fmt.Printf("Are %v and %v connected? %v\n", 5, 7, nodesConnected(5, 7))
	fmt.Printf("Are %v and %v connected? %v\n", 2, 7, nodesConnected(2, 7))

}
