

// Complete the reverse function below.

/*
 * For your reference:
 *
 * SinglyLinkedListNode {
 *     data int32
 *     next *SinglyLinkedListNode
 * }
 *
 */
func reverse(head *SinglyLinkedListNode) *SinglyLinkedListNode {
    current := head
    data := make([]int32, 0)
    for current != nil {
        data = append(data, current.data)
        current = current.next 
    }
    count := len(data)-1
    current = head
    for current != nil {
        v := data[count]
        current.data = v
        current = current.next
        count-- 
    }
    return head
}

