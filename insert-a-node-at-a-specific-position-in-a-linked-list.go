

// Complete the insertNodeAtPosition function below.

/*
 * For your reference:
 *
 * SinglyLinkedListNode {
 *     data int32
 *     next *SinglyLinkedListNode
 * }
 *
 */
func insertNodeAtPosition(llist *SinglyLinkedListNode, data int32, position int32) *SinglyLinkedListNode {
    if position == 0 {
        return &SinglyLinkedListNode{data: data, next: llist}
    }
    counter := int32(1)
    current := llist
    next := llist.next
    for current != nil {
        if position == counter {
            node := &SinglyLinkedListNode{data: data, next: next}
            current.next = node
            return llist
        }
        current = current.next
        next = current.next
        counter++
    }
    return llist
}

