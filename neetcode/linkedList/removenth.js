function removeNthFromEnd(head, n) {
    const dummy = {
        value: 0,
        next: head,
    };
    left = dummy;
    let right = head;

    while (n > 0) {
        right = right.next;
        n--;
    }

    while (right !== null) {
        left = left.next;
        right = right.next;
    }

    left.next = left.next.next;
    return dummy.next;
}
