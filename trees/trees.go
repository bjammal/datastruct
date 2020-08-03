package trees

import "fmt"

//Tree is the data type of a Tree
type Tree struct {
	root *Node
}

//Btree is the data type of a binary tree
type Btree struct {
	root *Bnode
}

//Node is the data type of tree node
type Node struct {
	Val      int
	children []*Node
}

//Bnode is the data type of a binary tree node
type Bnode struct {
	Val   int
	Left  *Bnode
	Right *Bnode
}

//Insert a new node in a binary tree,
//it returns a tree so we can chain calls to insert new nodes
func (bst *Btree) Insert(v int) *Btree {
	if bst.root == nil {
		bst.root = &Bnode{Val: v}
	} else {
		bst.root.insert(v)
	}
	return bst
}

//insert a node in a binary search tree
func (n *Bnode) insert(v int) {
	if v <= n.Val {
		if n.Left == nil {
			n.Left = &Bnode{Val: v}
		} else {
			n.Left.insert(v)
		}
	} else {
		if n.Right == nil {
			n.Right = &Bnode{Val: v}
		} else {
			n.Right.insert(v)
		}
	}
}

//PrintNode prints the a node and all its descendents
func (n *Bnode) PrintNode() {
	if n == nil {
		return
	}
	fmt.Printf("Node %d : ", n.Val)
	if n.Left != nil {
		fmt.Printf("Left = %d ", n.Left.Val)
	}
	if n.Right != nil {
		fmt.Printf("Right = %d", n.Right.Val)
	}
	fmt.Printf("\n")
	n.Left.PrintNode()
	n.Right.PrintNode()
}

//FindNode finds a node in a binary search tree and returns it
func (bst *Btree) FindNode(Val int) *Bnode {
	if bst == nil || bst.root == nil {
		return nil
	}

	if bst.root != nil && bst.root.Val == Val {
		return bst.root
	}
	if Val <= bst.root.Val {
		return bst.root.Left.FindNode(Val)
	} else {
		return bst.root.Right.FindNode(Val)
	}
}

//FindNode checks if a node exists and returns a pointer to it
func (n *Bnode) FindNode(Val int) *Bnode {
	if n == nil {
		return nil
	}
	if n.Val == Val {
		return n
	}
	if Val <= n.Val {
		return n.Left.FindNode(Val)
	} else {
		return n.Right.FindNode(Val)
	}
}
