function copyRandomList(head) {
    // two pass
    class Node {
        constructor(val, next = null, random = null) {
            this.val = val;
            this.next = next;
            this.random = next;
        }
    }

    const oldToCopy = new Map();

    let cur = head;
    while (cur) {
        const copy = new Node(cur.val);
        oldToCopy.set(cur, copy);
        cur = cur.next;
    }

    cur = head;
    while (cur) {
        const copy = oldToCopy.get(cur);
        copy.next = oldToCopy.get(cur.next);
        copy.random = oldToCopy.get(cur.random);
        cur = cur.next;
    }

    return oldToCopy.get(head);
}
