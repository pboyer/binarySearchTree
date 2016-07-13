package bst

import ("fmt")

type Node struct {
    val int
    left, right *Node
}

func (n *Node) String() string {
    l := "nil"
    r := "nil"

    if n.left != nil {
        l = fmt.Sprintf("%v", n.left.val)
    }

    if n.right != nil {
        r = fmt.Sprintf("%v", n.right.val)
    }

    return fmt.Sprintf("%v - (%s, %s)", n.val, l, r)
}

func (n *Node) Insert(val int) error {
    if n.val == val {
        return fmt.Errorf("value already present")
    }

    if val < n.val {
        if n.left != nil {
            return n.left.Insert(val)
        }

        n.left = &Node{ val : val }
        return nil
    } 

    if n.right != nil {
        return n.right.Insert(val)
    }

    n.right = &Node{ val : val }
    return nil
}

func (n *Node) Remove(val int) error {
    return n.removeInner(&Node{}, false, val)
}

func (n *Node) removeInner(parent *Node, isLeft bool, val int) error {
    if n.val == val {
        numChild := 0

        if n.left != nil {
            numChild++
        }

        if n.right != nil {
            numChild++
        }

        switch numChild {
        case 0:
            if isLeft {
                parent.left = nil
            } else {
                parent.right = nil
            }
        case 1:
            if isLeft {
                if n.left != nil {
                    parent.left = n.left
                } else {
                    parent.left = n.right
                }
            } else {
                if n.left != nil {
                    parent.right = n.left
                } else {
                    parent.right = n.right
                }
            }
        case 2:
            n.val = n.left.val
            return n.left.removeInner(n, true, n.val)
        }

        return nil
    }

    if n.right == nil && n.left == nil {
        return fmt.Errorf("value not found")
    }

    if val < n.val {
        if n.left == nil {
            return fmt.Errorf("value not found")
        }

        return n.left.removeInner(n, true, val)
    } else {
        if n.right == nil {
            return fmt.Errorf("value not found")
        }

        return n.right.removeInner(n, false, val)
    }
}