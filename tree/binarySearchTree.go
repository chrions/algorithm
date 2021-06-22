package tree

//二叉查找树要求，在树中的任意一个节点，其左子树中的每个节点的值，都要小于这个节点的值，而右子树节点的值都大于这个节点的值

type BST struct {
	*BinaryTree
	//比对函数，0:v==nodeV,正数:v>nodeV,负数:v<nodeV
	compareFunc func(v, nodeV interface{}) int
}

func NewBST(rootV interface{}, compareFunc func(v, nodeV interface{}) int) *BST {
	if compareFunc == nil {
		return nil
	}

	return &BST{BinaryTree: NewBinaryTree(rootV), compareFunc: compareFunc}
}

func (this *BST) Find(v interface{}) *Node {
	p := this.root
	for p != nil {
		ret := this.compareFunc(v, p.data)
		if ret == 0 {
			return p
		}
		if ret > 0 {
			p = p.right
		}
		if ret < 0 {
			p = p.left
		}
	}
	return nil
}

func (this *BST) Insert(v interface{}) bool {
	p := this.root
	for p != nil {
		ret := this.compareFunc(v, p.data)
		if ret == 0 {
			return false
		}
		if ret > 0 {
			if p.right == nil {
				p.right = NewNode(v)
				break
			}
			p = p.right
		}
		if ret < 0 {
			if nil == p.left {
				p.left = NewNode(v)
				break
			}
			p = p.left

		}
	}
	return true
}

//二叉查找树的删除有三个要求：
//1.如果要删除的节点没有子节点，只需要将父节点中指向要删除节点的指针置为null
//2.如果要删除的节点只有一个节点（左右节点都行）只需要更新父节点中，指向要删除节点的指针，让他指向
//要删除节点的子节点
//3.如果要删除的节点有两个子节点，我们需要找到这个节点的右子树中的最小节点，把它替换到要删除的节点上，
//然后再删除掉这个最小节点

func (this *BST) Delete(v interface{}) bool {
	var pp *Node = nil
	p := this.root
	deleteLeft := false
	for nil != p {
		compareResult := this.compareFunc(v, p.data)
		if compareResult > 0 {
			pp = p
			p = p.right
			deleteLeft = false
		}
		if compareResult < 0 {
			pp = p
			p = p.left
			deleteLeft = true
		} else {
			break
		}
	}

	if nil == p { //需要删除的节点不存在
		return false
	} else if nil == p.left && nil == p.right { //删除的是一个叶子节点
		if nil != pp {
			if deleteLeft {
				pp.left = nil
			} else {
				pp.right = nil
			}
		} else { //根节点
			this.root = nil
		}
	} else if nil != p.right { //删除的是一个有右孩子，不一定有左孩子的节点
		//找到p节点右孩子的最小节点
		pq := p
		q := p.right //向右走一步
		fromRight := true
		for nil != q.left { //向左走到底
			pq = q
			q = q.left
			fromRight = false
		}
		if fromRight {
			pq.right = nil
		} else {
			pq.left = nil
		}
		q.left = p.left
		q.right = p.right
		if nil == pp { //根节点被删除
			this.root = q
		} else {
			if deleteLeft {
				pq.left = q
			} else {
				pq.right = q
			}
		}
	} else { //删除的是一个只有左孩子的节点
		if nil != pp {
			if deleteLeft {
				pp.left = p.left
			} else {
				pp.right = p.left
			}
		} else {
			if deleteLeft {
				this.root = p.left
			} else {
				this.root = p.left
			}
		}
	}

	return true

}

func (this *BST) Min() *Node {
	p := this.root
	for nil != p.left {
		p = p.left
	}
	return p
}

func (this *BST) Max() *Node {
	p := this.root
	for nil != p.right {
		p = p.right
	}
	return p
}
