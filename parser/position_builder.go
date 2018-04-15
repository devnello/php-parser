package parser

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/scanner"
)

// Builder provide functions to constuct positions
type Builder struct {
	Positions *Positions
}

type startPos struct {
	startLine int
	startPos  int
}

type endPos struct {
	endLine int
	endPos  int
}

func (b *Builder) getListStartPos(l []node.Node) startPos {
	if l == nil {
		return startPos{-1, -1}
	}

	if len(l) == 0 {
		return startPos{-1, -1}
	}

	return b.getNodeStartPos(l[0])
}

func (b *Builder) getNodeStartPos(n node.Node) startPos {
	sl := -1
	sp := -1

	if n == nil {
		return startPos{-1, -1}
	}

	p := (*b.Positions)[n]
	if p != nil {
		sl = p.StartLine
		sp = p.StartPos
	}

	return startPos{sl, sp}
}

func (b *Builder) getListEndPos(l []node.Node) endPos {
	if l == nil {
		return endPos{-1, -1}
	}

	if len(l) == 0 {
		return endPos{-1, -1}
	}

	return b.getNodeEndPos(l[len(l)-1])
}

func (b *Builder) getNodeEndPos(n node.Node) endPos {
	el := -1
	ep := -1

	if n == nil {
		return endPos{-1, -1}
	}

	p := (*b.Positions)[n]
	if p != nil {
		el = p.EndLine
		ep = p.EndPos
	}

	return endPos{el, ep}
}

// NewNodeListPosition returns new Position
func (b *Builder) NewNodeListPosition(list []node.Node) *position.Position {
	return &position.Position{
		b.getListStartPos(list).startLine,
		b.getListEndPos(list).endLine,
		b.getListStartPos(list).startPos,
		b.getListEndPos(list).endPos,
	}
}

// NewNodePosition returns new Position
func (b *Builder) NewNodePosition(n node.Node) *position.Position {
	return &position.Position{
		b.getNodeStartPos(n).startLine,
		b.getNodeEndPos(n).endLine,
		b.getNodeStartPos(n).startPos,
		b.getNodeEndPos(n).endPos,
	}
}

// NewTokenPosition returns new Position
func (b *Builder) NewTokenPosition(t scanner.Token) *position.Position {
	return &position.Position{
		t.StartLine,
		t.EndLine,
		t.StartPos,
		t.EndPos,
	}
}

// NewTokensPosition returns new Position
func (b *Builder) NewTokensPosition(startToken scanner.Token, endToken scanner.Token) *position.Position {
	return &position.Position{
		startToken.StartLine,
		endToken.EndLine,
		startToken.StartPos,
		endToken.EndPos,
	}
}

// NewTokenNodePosition returns new Position
func (b *Builder) NewTokenNodePosition(t scanner.Token, n node.Node) *position.Position {
	return &position.Position{
		t.StartLine,
		b.getNodeEndPos(n).endLine,
		t.StartPos,
		b.getNodeEndPos(n).endPos,
	}
}

// NewNodeTokenPosition returns new Position
func (b *Builder) NewNodeTokenPosition(n node.Node, t scanner.Token) *position.Position {
	return &position.Position{
		b.getNodeStartPos(n).startLine,
		t.EndLine,
		b.getNodeStartPos(n).startPos,
		t.EndPos,
	}
}

// NewNodesPosition returns new Position
func (b *Builder) NewNodesPosition(startNode node.Node, endNode node.Node) *position.Position {
	return &position.Position{
		b.getNodeStartPos(startNode).startLine,
		b.getNodeEndPos(endNode).endLine,
		b.getNodeStartPos(startNode).startPos,
		b.getNodeEndPos(endNode).endPos,
	}
}

// NewNodeListTokenPosition returns new Position
func (b *Builder) NewNodeListTokenPosition(list []node.Node, t scanner.Token) *position.Position {
	return &position.Position{
		b.getListStartPos(list).startLine,
		t.EndLine,
		b.getListStartPos(list).startPos,
		t.EndPos,
	}
}

// NewTokenNodeListPosition returns new Position
func (b *Builder) NewTokenNodeListPosition(t scanner.Token, list []node.Node) *position.Position {
	return &position.Position{
		t.StartLine,
		b.getListEndPos(list).endLine,
		t.StartPos,
		b.getListEndPos(list).endPos,
	}
}

// NewNodeNodeListPosition returns new Position
func (b *Builder) NewNodeNodeListPosition(n node.Node, list []node.Node) *position.Position {
	return &position.Position{
		b.getNodeStartPos(n).startLine,
		b.getListEndPos(list).endLine,
		b.getNodeStartPos(n).startPos,
		b.getListEndPos(list).endPos,
	}
}

// NewOptionalListTokensPosition returns new Position
func (b *Builder) NewOptionalListTokensPosition(list []node.Node, t scanner.Token, endToken scanner.Token) *position.Position {
	if list == nil {
		return &position.Position{
			t.StartLine,
			endToken.EndLine,
			t.StartPos,
			endToken.EndPos,
		}
	}

	return &position.Position{
		b.getListStartPos(list).startLine,
		endToken.EndLine,
		b.getListStartPos(list).startPos,
		endToken.EndPos,
	}
}
