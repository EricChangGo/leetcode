/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */

/* Key: input is DPS (preorder traversal)*/

class Solution {
public:
    int check_index = 0;
    TreeNode* bstFromPreorder(vector<int>& preorder) {
      int checkt_index = 0;
      return build_tree(preorder, INT_MIN, INT_MAX); 
    }
    
    // construct the rest of the tree;
    TreeNode* build_tree(vector<int>& preorder, int min, int max) {
      // end of the leaf
      if(check_index == preorder.size() || preorder[check_index] >= max || preorder[check_index] <= min) return NULL;

      int check_value = preorder[check_index];
      TreeNode* root = NULL;
      root = new TreeNode(check_value);
      check_index++;
      // use next value, build subtree, their will be only one suitible spot 0987654`12for insert, rest will be null
      // left tree recursion
      root -> left = build_tree(preorder, min, check_value); 
      // right tree recursion
      root -> right = build_tree(preorder, check_value, max);
      return root;
    }
    
    /* 
      47% faster O(n^2), only 10% space.
      try to not using left right vector.
      
      TreeNode* bstFromPreorder(vector<int>& preorder) {
      if(preorder.size() == 0) return NULL;
      
      TreeNode* root = new TreeNode(preorder[0]);
      vector<int> left; // used as queue;
      vector<int> right; // used as queue;
      
      for(int i = 1; i < preorder.size(); i++) {
        if(root -> val > preorder[i])
          left.push_back(preorder[i]);
        else
          right.push_back(preorder[i]);
      }
      
      root -> left = bstFromPreorder(left);
      root -> right = bstFromPreorder(right);
      return root;
    }*/
};
# @lc code=end

