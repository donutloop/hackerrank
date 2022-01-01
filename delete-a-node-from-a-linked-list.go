

// Complete the deleteNode function below.

/*
 * For your reference:
 *
 * SinglyLinkedListNode {
 *     data int32
 *     next *SinglyLinkedListNode
 * }
 *
 */
func deleteNode(head *SinglyLinkedListNode, position int32) *SinglyLinkedListNode {
    if position == 0 { 
        return head.next       
    }
    var count int32
    current := head.next
    pre := head        
    for current != nil {
        count++
        if count == position {
            if current.next != nil {
                pre.next = current.next
            } else {
                pre.next = nil
            }
            break
        }
        pre = current
        current = current.next  
    }
    return head
}

