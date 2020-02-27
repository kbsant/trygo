package v1

type Nodes []int

// Initialize the connection map.
// Every node is denoted by a nonzero number.
// And every entry in the array is the node's group.
// Initially, every node has its own group since it isn't connected.
// Node zero is a guard element.
func Init(max int) Nodes {
	nodes := make([]int, max+1)
	for i := 0; i <= max; i++ {
		nodes[i] = i
	}
	return nodes
}

// Check whether 2 nodes are connected.
// If they have the same group, they are connected.
func (nodes Nodes) Connected(a int, b int) bool {
	return nodes[a] == nodes[b]
}

// Sort 2 nodes (helper function, used in merge)
func sort2(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

// Merge 2 nodes.
// Every node should end up with the same group id.
// To ensure transitivity, sort the ids before merging.
// By convention, always use the lower group id as the winner.
func (nodes Nodes) Merge(a int, b int) {
	lo, hi := sort2(a, b)
	for i := 0; i < len(nodes); i++ {
		if nodes[i] == hi {
			nodes[i] = lo
		}
	}
}
