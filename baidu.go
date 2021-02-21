package main

import "container/list"

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

// 树的中序遍历
func inorder_print(root TreeNode) {
	println(root.val)
	x := list.List{}
	if root.left != nil {
		x.PushBack(root.left)

	}
	if root.right != nil {
		x.PushBack(root.right)
	}
	for {
		if x.Len() == 0 {
			break
		}
		y := x.Front()
		x.Remove(y)
		print(y.Value)
		if y.Value.left != nil {
			x.PushBack(y.Value.left)

		}
		if y.Value.right != nil {
			x.PushBack(y.Value.right)
		}
	}
}

//找两个节点的最近公共祖先
// 如果这两个节点其中一个是根节点，直接返回根节点
//如果一个结点在左子树，另一个在右子树，直接返回根节点
//如果两个都在左子树或者都在右子树上，直接使用递归返回最近公共祖先
func search(r TreeNode, l TreeNode) bool {
	if r == l {
		return true
	}
	if search(*r.left, l) {
		return true
	}
	if search(*r.right, l) {
		return true
	}
	return false
}

func lowest(root TreeNode, p TreeNode, q TreeNode) TreeNode {
	if p == root || q == root {
		return root
	}
	pl := search(*root.left, p)
	pr := search(*root.right, p)
	qr := search(*root.right, q)
	ql := search(*root.left, q)
	if qr && pl || ql && pr {
		return root
	}

	if qr && pr {
		return lowest(*root.left, p, q)
	}

	return lowest(*root.right, p, q)

}
