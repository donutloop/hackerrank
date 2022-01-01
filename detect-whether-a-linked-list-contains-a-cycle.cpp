

// Complete the has_cycle function below.

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
bool has_cycle(SinglyLinkedListNode* head) {
    SinglyLinkedListNode* current;
    std::set<SinglyLinkedListNode*> pointers;
    current = head;
    while (current != nullptr){
      std::set<SinglyLinkedListNode *>::iterator pointer = pointers.find(current);
      if (pointer != pointers.end())
            return true;
      else
        pointers.insert(current);

      current = current->next;
    }
    return false;
}

