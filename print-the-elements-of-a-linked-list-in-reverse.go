

// Complete the reversePrint function below.

/*
 * For your reference:
 *
 * SinglyLinkedListNode {
 *     data int32
 *     next *SinglyLinkedListNode
 * }
 *
 */
func reversePrint(llist *SinglyLinkedListNode) {
    data := make([]int32, 0)
    current := llist
    for current != nil {
        data = append(data, current.data)
        current = current.next
    }
    for i := len(data)-1; i >= 0; i-- {
        fmt.Println(data[i])
    } 
}

