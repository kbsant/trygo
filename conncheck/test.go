package conncheck

import (
	"fmt"
	"github.com/kbsant/trygo/conncheck/v1"
	"github.com/kbsant/trygo/conncheck/v2"
)

func testConn(desc string, nodes Nodes) {
	fmt.Println(desc)

	fmt.Println("start with ", nodes)
	fmt.Printf("Are %v and %v connected? %v\n", 2, 5, nodes.Connected(2, 5))
	fmt.Printf("Are %v and %v connected? %v\n", 6, 4, nodes.Connected(6, 4))
	fmt.Printf("Are %v and %v connected? %v\n", 5, 7, nodes.Connected(5, 7))
	fmt.Printf("Are %v and %v connected? %v\n", 5, 6, nodes.Connected(5, 6))
	fmt.Printf("Are %v and %v connected? %v\n", 2, 7, nodes.Connected(2, 7))
	fmt.Printf("Are %v and %v connected? %v\n", 2, 4, nodes.Connected(2, 4))

	fmt.Println("Merge 5 and 7.")
	fmt.Println("Merge 6 and 4.")
	nodes.Merge(5, 7)
	nodes.Merge(6, 4)
	fmt.Printf("Are %v and %v connected? %v\n", 2, 5, nodes.Connected(2, 5))
	fmt.Printf("Are %v and %v connected? %v\n", 6, 4, nodes.Connected(6, 4))
	fmt.Printf("Are %v and %v connected? %v\n", 5, 7, nodes.Connected(5, 7))
	fmt.Printf("Are %v and %v connected? %v\n", 5, 6, nodes.Connected(5, 6))
	fmt.Printf("Are %v and %v connected? %v\n", 2, 7, nodes.Connected(2, 7))
	fmt.Printf("Are %v and %v connected? %v\n", 2, 4, nodes.Connected(2, 4))

	fmt.Println("Merge 5 and 2.")
	nodes.Merge(5, 2)
	fmt.Printf("Are %v and %v connected? %v\n", 2, 5, nodes.Connected(2, 5))
	fmt.Printf("Are %v and %v connected? %v\n", 6, 4, nodes.Connected(6, 4))
	fmt.Printf("Are %v and %v connected? %v\n", 5, 7, nodes.Connected(5, 7))
	fmt.Printf("Are %v and %v connected? %v\n", 5, 6, nodes.Connected(5, 6))
	fmt.Printf("Are %v and %v connected? %v\n", 2, 7, nodes.Connected(2, 7))
	fmt.Printf("Are %v and %v connected? %v\n", 2, 4, nodes.Connected(2, 4))

	fmt.Println(desc, "- ends with ", nodes)
}

func Test() {
	testConn("Conncheck v1", v1.Init(10))
	testConn("Conncheck v2", v2.Init(10))
}
