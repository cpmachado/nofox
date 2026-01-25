package nofox

type AST []Node

type NodeType int

const (
	NodeTypeError NodeType = iota
	NodeTypeMove
	NodeTypeIncrement
	NodeTypeLoop
	NodeTypeRead
	NodeTypePrint
)

type Node interface {
	Type() NodeType
}

type NodeMove struct {
	Value int
}

func (n *NodeMove) Type() NodeType {
	return NodeTypeMove
}

type NodeIncrement struct {
	Value int
}

func (n *NodeIncrement) Type() NodeType {
	return NodeTypeIncrement
}

type NodeLoop struct {
	Nodes []Node
}

func (n *NodeLoop) Type() NodeType {
	return NodeTypeLoop
}

type NodeRead struct{}

func (n *NodeRead) Type() NodeType {
	return NodeTypeRead
}

type NodePrint struct{}

func (n *NodePrint) Type() NodeType {
	return NodeTypePrint
}
