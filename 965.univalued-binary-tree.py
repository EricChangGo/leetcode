#
# @lc app=leetcode id=965 lang=python3
#
# [965] Univalued Binary Tree
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

#include <iostream>
using namespace std;

class Solution {
public:
    bool isUnivalTree(TreeNode* root) {
      bool is_univalued = true;
      int the_value = root -> val;
      TreeNode *check = NULL;
      queue<TreeNode*> unvisit;
      // start with root
      unvisit.push(root);
      while(!unvisit.empty()) {
        check = unvisit.front();
        // check value
        if (check -> val != the_value) {
          is_univalued = false;
          break;
        }
        // insert children
        if (check -> left) unvisit.push(check -> left);
        if (check -> right) unvisit.push(check -> right);
        unvisit.pop();
      }
            
      cout << unvisit.size();
      return is_univalued;
    }
};
# @lc code=end

