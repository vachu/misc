package avlbst

import (
	"container/list"
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

// Add ...
func (bt *BinarySearchTree) Add(data interface{}) error {
	if bt.cmp == nil {
		return fmt.Errorf("Comparer function is 'nil'")
	}
	if data == nil {
		return fmt.Errorf("'nil' data provided")
	}

	newNode := &node{Data: data}
	bt.NodeCount++
	if bt.IsEmpty() {
		bt.root = &root{newNode}
	} else {
	FOR:
		for n := bt.root.node; n != nil; {
			switch bt.cmp(newNode.Data, n.Data) {
			case 0:
				bt.NodeCount-- // as newNode is not added to tree
				return fmt.Errorf("data already available in BST")
			case -1:
				if n.Left == nil {
					newNode.parent = n
					n.Left = newNode
					break FOR
				}
				n = n.Left
			case 1:
				if n.Right == nil {
					newNode.parent = n
					n.Right = newNode
					break FOR
				}
				n = n.Right
			} // switch
		} // for
	} // else
	bt.rebalance(newNode)
	return nil // newNode added to BST
}

func getParent(n *node) *node {
	if n != nil {
		return n.parent
	}
	return nil
}

func getHeight(n *node) int {
	if n == nil {
		return 0
	}
	lHeight := getHeight(n.Left)
	rHeight := getHeight(n.Right)
	if lHeight > rHeight {
		return 1 + lHeight
	}
	return 1 + rHeight
}

func getBalanceFactor(n *node) int {
	if n == nil {
		return 0
	}
	return getHeight(n.Left) - getHeight(n.Right)
}

func (bt *BinarySearchTree) rebalance(newNode *node) {
	if !bt.IsAvlTree() || newNode == nil {
		return
	}

	for n := newNode; n != nil; n = getParent(n) {
		bf := getBalanceFactor(n)
		switch {
		case bf > 1:
			bfLeftChild := getBalanceFactor(n.Left)
			if bfLeftChild == -1 {
				bt.rotateLeft(n.Left)
			}
			bt.rotateRight(n)
			return
		case bf < -1:
			bfRightChild := getBalanceFactor(n.Right)
			if bfRightChild == 1 {
				bt.rotateRight(n.Right)
			}
			bt.rotateLeft(n)
			return
		}
	}
}

func (bt *BinarySearchTree) rotateLeft(n *node) *node {
	if n == nil || n.Right == nil {
		return n
	}
	nParent := n.parent
	rightChild := n.Right
	n.parent = rightChild
	n.Right = rightChild.Left
	rightChild.Left = n
	rightChild.parent = nParent // now rightChild is in old n's place

	if nParent == nil {
		bt.root.node = rightChild
	} else {
		if nParent.Right == n {
			nParent.Right = rightChild
		} else {
			nParent.Left = rightChild
		}
	}
	return rightChild
}

func (bt *BinarySearchTree) rotateRight(n *node) *node {
	if n == nil || n.Left == nil {
		return n
	}
	nParent := n.parent
	leftChild := n.Left
	n.parent = leftChild
	n.Left = leftChild.Right
	leftChild.Right = n
	leftChild.parent = nParent // now rightChild is in old n's place
	if nParent == nil {
		bt.root.node = leftChild
	} else {
		if nParent.Right == n {
			nParent.Right = leftChild
		} else {
			nParent.Left = leftChild
		}
	}
	return leftChild
}

// ToXML ...
func (bt *BinarySearchTree) ToXML(w io.Writer) error {
	const rootNodeName = "BinarySearchTree"

	fmt.Fprintln(w, xml.Header)
	fmt.Fprintf(w, "<%s", rootNodeName)
	fmt.Fprintf(w, " isAVLTree=\"%v\"", bt.IsAvlTree())
	fmt.Fprintf(w, " nodeCount=\"%d\"", bt.NodeCount)
	fmt.Fprintf(w, " height=\"%d\"", getHeight(bt.root.node))
	fmt.Fprintf(w, " balanceFactor=\"%d\"", getBalanceFactor(bt.root.node))
	fmt.Fprintln(w, ">")

	enc := xml.NewEncoder(w)
	enc.Indent("\t", "\t")
	if err := enc.Encode(bt.root); err != nil {
		return err
	}

	fmt.Fprintf(w, "\n</%s>", rootNodeName)
	return nil
}

// TraversalKind ...
type TraversalKind uint

// Depth-First traversal types
const (
	INORDER TraversalKind = iota
	PREORDER
	POSTORDER
	BREADTHFIRST
)

func traverseInorder(w io.Writer, n *node) {
	if n == nil {
		return
	}
	traverseInorder(w, n.Left)
	fmt.Fprintf(w, "%v ", n.Data)
	traverseInorder(w, n.Right)
}

func traversePreorder(w io.Writer, n *node) {
	if n == nil {
		return
	}
	fmt.Fprintf(w, "%v ", n.Data)
	traversePreorder(w, n.Left)
	traversePreorder(w, n.Right)
}

func traversePostorder(w io.Writer, n *node) {
	if n == nil {
		return
	}
	traversePostorder(w, n.Left)
	traversePostorder(w, n.Right)
	fmt.Fprintf(w, "%v ", n.Data)
}

func traverseBreadthFirst(w io.Writer, root *node) {
	queue := list.New()
	if root != nil {
		queue.PushBack(root)
	}
	for queue.Len() > 0 {
		n := queue.Remove(queue.Front()).(*node)
		fmt.Fprintf(w, "%v ", n.Data)

		if n.Left != nil {
			queue.PushBack(n.Left)
		}
		if n.Right != nil {
			queue.PushBack(n.Right)
		}
	}
}

// Traverse ...
func (bt *BinarySearchTree) Traverse(w io.Writer, k TraversalKind) {
	switch k {
	case INORDER:
		fmt.Fprint(w, "Inorder Traversal      : ")
		traverseInorder(w, bt.root.node)
		fmt.Fprintln(w)
	case PREORDER:
		fmt.Fprint(w, "Preorder Traversal     : ")
		traversePreorder(w, bt.root.node)
		fmt.Fprintln(w)
	case POSTORDER:
		fmt.Fprint(w, "Postorder Traversal    : ")
		traversePostorder(w, bt.root.node)
		fmt.Fprintln(w)
	case BREADTHFIRST:
		fmt.Fprint(w, "Breadth-first Traversal: ")
		traverseBreadthFirst(w, bt.root.node)
		fmt.Fprintln(w)
	}
}

func (bt *BinarySearchTree) has(n *node, value interface{}) *node {
	for n != nil {
		switch bt.cmp(n.Data, value) {
		case 0:
			return n
		case 1:
			n = n.Left
		case -1:
			n = n.Right
		}
	}
	return nil
}

// Has ...
func (bt *BinarySearchTree) Has(value interface{}) bool {
	return bt.has(bt.root.node, value) != nil
}

func getRightmostNode(n *node) (r *node) {
	for ; n != nil; n = n.Right {
		r = n
	}
	return
}

// RightmostValue ...
func (bt *BinarySearchTree) RightmostValue() (rightmost interface{}) {
	if r := getRightmostNode(bt.root.node); r != nil {
		rightmost = r.Data
	}
	return
}

func getLeftmostNode(n *node) (l *node) {
	for ; n != nil; n = n.Left {
		l = n
	}
	return
}

// LeftmostValue ...
func (bt *BinarySearchTree) LeftmostValue() (leftmost interface{}) {
	if r := getLeftmostNode(bt.root.node); r != nil {
		leftmost = r.Data
	}
	return
}

// Delete ....
func (bt *BinarySearchTree) Delete(value interface{}) error {
	n := bt.has(bt.root.node, value)
	if n == nil {
		return fmt.Errorf("value unavailable")
	}

	var replacingNode *node
	if getHeight(n.Left) > getHeight(n.Right) {
		replacingNode = getRightmostNode(n.Left)
	} else {
		replacingNode = getLeftmostNode(n.Right)
	}
	if replacingNode != nil {
		if replacingNode.parent.Left == replacingNode {
			replacingNode.parent.Left = nil
		} else {
			replacingNode.parent.Right = nil
		}
		bt.rebalance(replacingNode.parent)
		replacingNode.parent = n.parent
		replacingNode.Left = n.Left
		replacingNode.Right = n.Right
	}
	if n.parent != nil {
		if n == n.parent.Left {
			n.parent.Left = replacingNode
		} else {
			n.parent.Right = replacingNode
		}
		bt.rebalance(n.parent)
	}
	if n == bt.root.node {
		bt.root.node = replacingNode
	}
	n.Left, n.Right, n.parent = nil, nil, nil // orphaning the node to be deleted
	bt.NodeCount--
	return nil
}

// Run ...
func Run() {
	bst := NewBinarySearchTree(true, func(d1, d2 interface{}) int {
		v1, v2 := d1.(int), d2.(int)
		if v1 < v2 {
			return -1
		} else if v1 == v2 {
			return 0
		} else {
			return 1
		}
	})
	for i := 0; i < 15; i++ {
		bst.Add(i + 1)
	}

	bst.ToXML(os.Stdout)
	fmt.Println()
	bst.Traverse(os.Stdout, INORDER)
	fmt.Println("Deleting 8...")
	if err := bst.Delete(8); err != nil {
		fmt.Println("ERROR:", err.Error())
	} else {
		bst.ToXML(os.Stdout)
		bst.Traverse(os.Stdout, INORDER)
	}
	// bst.Traverse(os.Stdout, PREORDER)
	// bst.Traverse(os.Stdout, POSTORDER)
	// bst.Traverse(os.Stdout, BREADTHFIRST)
}
