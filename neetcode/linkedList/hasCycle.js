function hasCycle(head) {
    // slow and fast pointer
    let slow = head;
    let fast = head;

    while ((fast !== null) & (fast.next !== null)) {
        fast = fast.next.next;
        slow = slow.next;

        if (fast === slow) {
            return true;
        }
    }

    return false;
}

function hasCyclehash(head) {
    let seen = new Set();
    let cur = head;

    while (cur) {
        if (seen.has(cur)) {
            return true;
        }
        seen.add(cur);
        cur = cur.next;
    }

    return false;
}
