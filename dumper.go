package main

import (
	"fmt"
	"reflect"

	"github.com/z7zmey/php-parser/node"
)

type dumper struct {
	indent string
}

func (d dumper) EnterNode(n node.Node) bool {

	fmt.Printf("%v%v", d.indent, reflect.TypeOf(n))
	if p := n.Position(); p != nil {
		fmt.Printf(" %v", *p)
	}
	if a := n.Attributes(); len(a) > 0 {
		fmt.Printf(" %v", a)
	}
	fmt.Println()

	return true
}

func (d dumper) GetChildrenVisitor(key string) node.Visitor {
	fmt.Printf("%v%q:\n", d.indent+"  ", key)
	return dumper{d.indent + "    "}
}

func (d dumper) LeaveNode(n node.Node) {
	// do nothing
}