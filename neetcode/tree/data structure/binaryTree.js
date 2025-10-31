class TreeNode {
    constructor(val) {
        this.val = val;
        this.left = null;
        this.right = null;
    }
}

// Binary search trees are a variation of binary tree with the addition of sorted property.
class BinarySearchTree {
    constructor(root) {
        this.root = root;
    }

    add(newNode) {
        if (node === null) {
            return newNode;
        }

        if (this.root.val > newNode.val) {
            this.root.left = add(newNode);
        } else if (this.root.val < newNode.val) {
            this.root.right = add(newNode);
        } else {
            return; // duplicate
        }
    }
}

function search(root, target) {
    if (root === null) {
        return false;
    }

    if (target > root.val) {
        return search(root.right, target);
    } else if (target < root.val) {
        return search(root.left, target);
    } else {
        return true;
    }
}

function findMinimum(root) {
    let cur = root;
    while (cur && cur.left) {
        cur = cur.left;
    }
    return cur; // return cur.val;
}

function insert(root, val) {
    if (root === null) {
        return new TreeNode(val);
    }

    if (val > root.val) {
        root.right = insert(root.right, val);
    } else if (val < root.val) {
        root.left = insert(root.left, val);
    }

    return root;
}

function remove(root, target) {
    // case1: 0 or 1 child
    // case2: 2 child or more

    // find target
    if (root === null) {
        return null;
    }

    if (target > root.val) {
        root.right = remove(root.right, target);
    } else if (target < root.val) {
        root.left = remove(root.left, target);
    } else {
        if (root.left == null) {
            return root.right;
        } else if (root.right == null) {
            return root.left;
        } else {
            let minNode = findMinimum(root.right);
            root.val = minNode.val;
            root.right = remove(root.right, minNode);
        }
    }
}
