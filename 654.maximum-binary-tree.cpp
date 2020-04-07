/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */
class Solution {
public:
    TreeNode* constructMaximumBinaryTree(vector<int>& nums) {
      return build_tree(nums, 0, nums.size() - 1);
    }
    
    TreeNode* build_tree(vector<int>& nums, int start_index, int last_index) {      
      // root will be the maximum of the tree
      if(start_index == last_index) return new TreeNode(nums[start_index]);
      else if(start_index > last_index) return NULL; 
      
      int root_value_index = start_index;
      int root_value = nums[root_value_index];
      
      for(int i = start_index; i <= last_index; i++) {
        if(nums[i] > root_value) {
          root_value = nums[i];
          root_value_index = i;
        }
      }
      
      TreeNode* root = new TreeNode(root_value);
      // left will be the index smaller than root value
      root -> left = build_tree(nums, start_index, root_value_index - 1);
      // right will be the index smaller than root value
      root -> right = build_tree(nums, root_value_index+1, last_index);
      return root;
    }
    
    TreeNode* create_end_point(int value) {
      TreeNode* root = new TreeNode(value);
      // left will be the index smaller than root value
      root -> left = NULL;
      // right will be the index smaller than root value
      root -> right = NULL;
      return root;
    }
  
    /*void print_in_BST(TreeNode* root) {
      TreeNode* target_node = NULL;
      queue<TreeNode*> unvisit_nodes;
      unvisit_nodes.push(root);
      
      while(!unvisit_nodes.empty()) {
        target_node = unvisit_nodes.front();
        cout << target_node -> val;
        if(target_node -> left) unvisit_nodes.push(target_node -> left);
        if(target_node -> right) unvisit_nodes.push(target_node -> right);
        unvisit_nodes.pop();
      }
    }*/
};
// @lc code=end

