package trees

func NewRadixTrees() *RadixTrees {
	return &RadixTrees{}
}

type Node struct {
	Value    interface{}
	ID       int
	Children []*Node
}

type Edge struct {
	Label     interface{}
	StartNode Node
	EndNode   Node
}

type RadixTrees struct {
}

func (r *RadixTrees) Insert(value interface{}) error {
	panic("impls")
}

func (r *RadixTrees) Delete(value interface{}) error {
	panic("impls")
}

func (r *RadixTrees) FindPrefix(prefix interface{}) interface{} {
	panic("impls")
}

func (r *RadixTrees) FindLongestPrefix(prefix interface{}) interface{} {
	panic("impls")
}
