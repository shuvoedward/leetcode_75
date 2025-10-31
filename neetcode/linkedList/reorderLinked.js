function reorderList(head) {
    // find the middle part of the list
    let slow = head;
    let fast = head.next;

    while (fast !== null && fast.next !== null) {
        slow = slow.next;
        fast = fast.next.next;
    }

    // reverse second part of the linked list
    let cur = slow.next;
    let prev = null; // let prev = slow.next = null;
    slow.next = null;
    while (cur !== null) {
        const tmp = cur.next;
        cur.next = prev;
        prev = cur;
        cur = tmp;
    }

    // now reorder the linked list
    let first = head;
    let second = prev;

    while (second !== null) {
        const tmp1 = first.next;
        const tmp2 = second.next;
        first.next = second;
        second.next = tmp1;
        first = tmp1;
        second = tmp2;
    }
}
