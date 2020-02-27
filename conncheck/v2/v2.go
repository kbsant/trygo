package v2

import (
	"fmt"
)

type Nodes []int

// Initialize the connection map.
// A node is represented by a nonzero integer, and its entry in the nodes array
// is the node it is connected to.
// Zero represents the unconnected node, so initially, every entry is zero
func Init(max int) Nodes {
	nodes := make([]int, max+1)
	return nodes
}

// Check whether 2 nodes are connected.
// Group id 0 is unconnected, so skip those.
// If nodes or their ancestors have the same group, they are connected.
// If nodes are connected by their ancestors, connect them directly.
// TODO speed this up, since some logic is duplicated between merge and connected.
func (nodes Nodes) Connected(a int, b int) bool {
	i, j := sort2(a, b)
	for {
		if i == 0 {
			return false
		}
		if i == j {
			nodes.Merge(a, b)
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
	}
	return b, a
}

// Link the second node to the first, or its ancestor if present.
func (nodes Nodes) Merge(a int, b int) {
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
