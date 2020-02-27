package conncheck

type Nodes interface {
	Connected(a, b int) bool
	Merge(a, b int)
}
