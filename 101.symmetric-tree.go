/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/**
            1
           / \
          2   2
         / \  / \
        3  4  4  3
      / \ / \ /\ /\  
      O X @ ! ! @ X O
    
    1. Start with level 1 (2,2) = (left, right) and next we compare (3,3) & (4,4)

    2. Target on (3,3), next will be check (O,O) (X,X)
        
		3. Then we give a recursive formula. 
		After we compare (a, b) = (left, right) nodes, we recursively check (a->left, b->right) && (a->right, b->left)

 */

 func compare(a *TreeNode, b *TreeNode) bool {
	if a == nil && b == nil {return true}
	if a == nil || b == nil {return false}
	if a.Val != b.Val {return false}
	
	return compare(a.Left, b.Right)&&compare(a.Right,b.Left)
}
func isSymmetric(root *TreeNode) bool {
	if root == nil {return true}
	return compare(root.Left, root.Right)
}

