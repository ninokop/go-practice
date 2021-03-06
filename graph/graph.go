package graph

type Graph interface {
	AddNode(Node) error
	AddEdge(src, dst string) error
	DFS()
	BFS()

	//For test
	Edges() interface{}
}

type Node interface {
	ID() string
	Run()
}
