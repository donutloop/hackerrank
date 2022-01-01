

// Complete the findMergeNode function below.

/*
 * For your reference:
 *
 * SinglyLinkedListNode {
 *     int data;
 *     SinglyLinkedListNode* next;
 * };
 *
 */

#include <set>
int findMergeNode(SinglyLinkedListNode* head1, SinglyLinkedListNode* head2) {

  SinglyLinkedListNode *current;
  std::set<SinglyLinkedListNode *> pointers;
  current = head1;
  while (current != nullptr) {
    pointers.insert(current);
    current = current->next;
  }

  current = head2;
  while (current != nullptr) {
    std::set<SinglyLinkedListNode *>::iterator pointer = pointers.find(current);
    if (pointer != pointers.end())
        return current->data;
    else 
        current = current->next;
  }
  return -1;
}

