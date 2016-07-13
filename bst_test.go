package bst

import ("testing")

func confirmStructure(t *testing.T, n *Node) {
    if n.left != nil {
        if n.left.val >= n.val {
            t.Fatalf("invalid structure %v", n)
            return
        }

        confirmStructure(t, n.left)
    }

    if n.right != nil {
        if n.right.val <= n.val {
            t.Fatalf("invalid structure %v", n)
            return
        }

        confirmStructure(t, n.right)
    }
}

func TestBST(t *testing.T) {
    n := &Node{val : 5}

    n.Insert(3)
    n.Insert(8)
    n.Insert(9)
    n.Insert(1)
    n.Insert(-1)
    n.Insert(4)

    confirmStructure(t, n)

    n.Remove(5)

    confirmStructure(t, n)
}

