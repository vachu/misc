package avlbst

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type node struct {
	Data        interface{} `xml:"data,attr"`
	Left, Right *node
	parent      *node
}

type root struct {
	*node
}

// Comparer ...
type Comparer func(data1, data2 interface{}) int

// BinarySearchTree ...
type BinarySearchTree struct {
	root             *root
	IsHeightBalanced bool
	NodeCount        uint32
	cmp              Comparer
}

// NewBinarySearchTree ...
func NewBinarySearchTree(isAvlTree bool, cmp func(d1, d2 interface{}) int) *BinarySearchTree {
	if cmp != nil {
		return &BinarySearchTree{nil, isAvlTree, 0, cmp}
	}
	return nil
}

// IsAvlTree ...
func (bt *BinarySearchTree) IsAvlTree() bool {
	return bt.IsHeightBalanced
}

// IsEmpty ...
func (bt *BinarySearchTree) IsEmpty() bool {
	return bt.root == nil
}

func getParent(n *node) *node {
	if n != nil {
		return n.parent
	}
	return nil
}

// Add ...
func (bt *BinarySearchTree) Add(data interface{}) error {
	if bt.cmp == nil {
		return fmt.Errorf("Comparer function is 'nil'")
	}
	if data == nil {
		return fmt.Errorf("'nil' data provided")
	}

	newNode := &node{Data: data}
	if bt.IsEmpty() {
		bt.root = &root{newNode}
	} else {
	FOR:
		for n := bt.root.node; n != nil; {
			switch bt.cmp(newNode.Data, n.Data) {
			case -1:
				if n.Left == nil {
					newNode.parent = n
					n.Left = newNode
					bt.NodeCount++
					break FOR
				}
				n = n.Left
			case 1:
				if n.Right == nil {
					newNode.parent = n
					n.Right = newNode
					bt.NodeCount++
					break FOR
				}
				n = n.Right
			} // switch
		} // for
	} // else
	return nil
}

func (bt *BinarySearchTree) rebalance(n *node) {
	if !bt.IsAvlTree() || n == nil {
		return
	}

}

// ToXML ...
func (bt *BinarySearchTree) ToXML(w io.Writer) error {
	fmt.Fprintln(w, xml.Header)
	fmt.Fprintf(w, "<BinarySearchTree isAVLTree=\"%v\" nodeCount=\"%d\">\n", bt.IsAvlTree(), bt.NodeCount)

	enc := xml.NewEncoder(w)
	enc.Indent("\t", "\t")
	if err := enc.Encode(bt.root); err != nil {
		return err
	}

	fmt.Fprintln(w, "\n</BinarySearchTree>")
	return nil
}

// Run ...
func Run() {
	bst := NewBinarySearchTree(false, func(d1, d2 interface{}) int {
		v1, v2 := d1.(int), d2.(int)
		if v1 < v2 {
			return -1
		} else if v1 == v2 {
			return 0
		} else {
			return 1
		}
	})

	for _, n := range []int{1, 2, 3, 4, 5} {
		bst.Add(n)
	}
	bst.ToXML(os.Stdout)
}
