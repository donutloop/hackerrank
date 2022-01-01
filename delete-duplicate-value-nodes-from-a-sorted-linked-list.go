

// Complete the removeDuplicates function below.

/*
 * For your reference:
 *
 * SinglyLinkedListNode {
 *     data int32
 *     next *SinglyLinkedListNode
 * }
 *
 */
func removeDuplicates(head *SinglyLinkedListNode) *SinglyLinkedListNode {
    set := make(map[int32]int32)
    pre := head
    current := head.next
    set[pre.data] =+ 1
    for current != nil {
        set[current.data] = set[current.data] + 1 
        if set[current.data] > 1 {
            if current.next == nil {
                pre.next = nil
                break
            }
            set[current.data] = set[current.data] - 1  
            pre.next = current.next
            pre = pre
            current = pre.next
        } else {
            pre = current
            current = current.next
        } 
    }
    return head
}

