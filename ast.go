package nofox

import "errors"

var (
	ErrLoopEnd        = errors.New("received loop end")
	ErrMissingLoopEnd = errors.New("missing loop end")
)

func Parse(token chan Token) (AST, error) {
	var base AST

	for tok := <-token; tok != TokenEOF; tok = <-token {
		var curr Node

		switch tok {
		case TokenMoveRight:
			curr = &NodeMove{Value: 1}
		case TokenMoveLeft:
			curr = &NodeMove{Value: -1}
		case TokenIncrement:
			curr = &NodeIncrement{Value: 1}
		case TokenDecrement:
			curr = &NodeIncrement{Value: -1}
		case TokenPrint:
			curr = &NodePrint{}
		case TokenRead:
			curr = &NodeRead{}
		case TokenLoopStart:
			nodes, err := Parse(token)
			if err != ErrLoopEnd {
				return nil, ErrMissingLoopEnd
			}
			curr = &NodeLoop{Nodes: nodes}
		case TokenLoopEnd:
			return base, ErrLoopEnd
		}
		base = append(base, curr)
	}

	return base, nil
}

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
