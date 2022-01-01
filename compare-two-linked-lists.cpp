

// Complete the compare_lists function below.

/*
 * For your reference:
 *
 * SinglyLinkedListNode {
 *     int data;
 *     SinglyLinkedListNode* next;
 * };
 *
 */

bool check(SinglyLinkedListNode *currentOne, SinglyLinkedListNode *currentTwo, int count) {
  if (currentOne != nullptr) {
    count = count + 1;
  }
  if (currentTwo != nullptr) {
    count = count + 1;
  }
  if (count == 1) {
    return false;
  }
  return true;
}

bool compare_lists(SinglyLinkedListNode* head1, SinglyLinkedListNode* head2) {
    
  SinglyLinkedListNode * currentOne;
  SinglyLinkedListNode * currentTwo;
  int count;
  count = 0;

  currentOne = head1;
  currentTwo = head2;
  if (!check(currentOne, currentTwo, count)) {
      return false;
  }
  while (currentOne && currentTwo) {
      if (currentOne->data != currentTwo->data) {
          return false;
      }
      count = 0;
      currentOne = currentOne->next;
      currentTwo = currentTwo->next;
      if (!check(currentOne, currentTwo, count)) {
         return false;
      }
  }
  return true;
}

