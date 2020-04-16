/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
#include<iostream>
using namespace std;

class Solution {
public:
    ListNode* mergeTwoLists(ListNode* l1, ListNode* l2) {
      
      if (!l1) return l2;
      if (!l2) return l1;
      // travel through l1
      ListNode *head  = NULL;
      ListNode *target_node = NULL;
      // insert l2 into l1
      ListNode *insert_node = NULL;
      ListNode *next_insert = NULL;

      head = l1;
      target_node = head;
      insert_node = l2;
      
      /* stategy:
        merge l2 into l1
        1. insert l2 object in front, back, or keep on finding insert point though l1.
        2. end when all l2 have been insert.
      */

      while(true){
        next_insert = insert_node -> next;
        if(insert_node -> val < target_node -> val) {
          // head point move forward
          head = insert_node;
          insert_before(target_node, insert_node);
          // move target to the insert point
          target_node = insert_node;
          // execute next insert object
          insert_node = next_insert;
        }
        else {
          // head point remain
          if (!target_node -> next || target_node -> next -> val >= insert_node -> val) {
            insert_after(target_node, insert_node);
            // move target to the insert point
            target_node = insert_node;
            // execute next insert object
            insert_node = next_insert;
          }
          else {
            // no insert happens, keep finding insert point
            target_node = target_node -> next;
            continue;
          }
        }
        
        if (next_insert == NULL) {
          break;
        }
      }

      return head;
    };
    
    void print_list(ListNode *head) {
      ListNode *default_start_point = head;
      
      while(true){
        cout << head -> val;
        if (head -> next == NULL) {
          head = default_start_point;
          cout << "\n";
          break;
        }
        cout << "->";
        head = head -> next;
      }
    }

    void insert_after(ListNode *target_node, ListNode *insert_node) {
      ListNode *temp_node = NULL;
      temp_node = target_node -> next;
      target_node -> next = insert_node;
      insert_node -> next = temp_node;
    };

    void insert_before(ListNode *target_node, ListNode *insert_node) {
      insert_node -> next = target_node;
    };
};