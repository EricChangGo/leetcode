/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */

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

