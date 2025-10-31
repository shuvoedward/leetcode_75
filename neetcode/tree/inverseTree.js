function breadthFirst(root) {
    let queue = [];
    if (root !== null) {
        queue.push(root);
    }

    while (queue.length > 0) {
        let curNode = queue.shift(); // O(n)
        [curNode.left, curNode.right] = [curNode.right, curNode.left];
        if (curNode.left !== null) queue.push(curNode.left);
        if (curNode.right !== null) queue.push(curNode.right);
    }

    return root;
}

class TreeNode {
    constructor(val, left = null, right = null) {
        this.val = val;
        this.left = left;
        this.right = right;
    }
}

function depthFirst(root) {
    if (root === null) return null;

    const node = new TreeNode(root.val);

    node.right = depthFirst(root.left);
    node.left = depthFirst(root.right);

    return node;
}

function depthFirstStack(root) {
    if (!root) return null;
    const stack = [root];
    while (stack.length) {
        const node = stack.pop();
        [node.left, node.right] = [node.right, node.left];
        if (node.left) stack.push(node.left);
        if (node.right) stack.push(node.right);
    }
    return root;
}
