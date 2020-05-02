package avlbst

import (
	"encoding/xml"
	"fmt"
	"io"
)

type node struct {
	Data    interface{} `xml:"data,attr"`
	Count   uint8       `xml:"count,attr"`
	balance int8
	Left    *leftNode
	Right   *rightNode
	parent  *node
}

type rootNode struct {
	*node
}

type leftNode struct {
	*node
}

type rightNode struct {
	*node
}

// Comparer ...
type Comparer func(data1, data2 interface{}) int

// BinarySearchTree ...
type BinarySearchTree struct {
	root             *rootNode
	IsHeightBalanced bool   `xml:"isAVLTree,attr"`
	NodeCount        uint32 `xml:"noOfNodes,attr"`
	cmp              Comparer
}

// NewBinarySearchTree ...
func NewBinarySearchTree(isAvlTree bool, cmp func(d1, d2 interface{}) int) *BinarySearchTree {
	return &BinarySearchTree{nil, isAvlTree, 0, cmp}
}

// IsAvlTree ...
func (bt *BinarySearchTree) IsAvlTree() bool {
	return bt.IsHeightBalanced
}

// IsEmpty ...
func (bt *BinarySearchTree) IsEmpty() bool {
	return bt.root == nil
}

// Add ...
func (bt *BinarySearchTree) Add(data interface{}) error {
	if bt.cmp == nil {
		return fmt.Errorf("Comparer function is 'nil'")
	}
	if data == nil {
		return fmt.Errorf("'nil' data provided")
	}

	newNode := &node{Data: data, Count: 1}
	bt.NodeCount++
	if bt.IsEmpty() {
		bt.root = &rootNode{newNode}
	} else {
	FOR:
		for n := bt.root.node; n != nil; {
			switch bt.cmp(newNode.Data, n.Data) {
			case -1:
				if n.Left == nil {
					newNode.parent = n
					n.Left = &leftNode{newNode}
					break FOR
				}
				n = n.Left.node
			case 0:
				n.Count++
				bt.NodeCount-- // since no new node is created here
				break FOR
			case 1:
				if n.Right == nil {
					newNode.parent = n
					n.Right = &rightNode{newNode}
					break FOR
				}
				n = n.Right.node
			} // switch
		} // for
	} // else

	return nil
}

// ToXML ...
func (bt *BinarySearchTree) ToXML(w io.Writer) error {
	fmt.Fprintln(w, xml.Header)

	enc := xml.NewEncoder(w)
	enc.Indent("\t", "\t")
	if err := enc.Encode(bt.root); err != nil {
		return err
	}
	return nil
}
