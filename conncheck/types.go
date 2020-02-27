package conncheck

// Since there are multiple implementations of the connection checker,
// define an interface
type Nodes interface {
	Connected(a, b int) bool
	Merge(a, b int)
}
