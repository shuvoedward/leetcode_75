function maxDepthStack(root) {
    // iterative dfs

    // O(n) for time and space complexity
    const stack = [[root, 1]];
    let res = 0;

    while (stack.length > 0) {
        const current = stack.pop();
        const node = current[0];
        const depth = current[1];

        if (node !== null) {
            res = Math.max(res, depth);
            stack.push([node.left, depth + 1]);
            stack.push([node.right, depth + 1]);
        }
    }
    return res;
}

function maxDepthRecursion(root) {
    // O(n) for time and space complexity
    if (root === null) {
        return 0;
    }

    return (
        1 +
        Math.max(maxDepthRecursion(root.left), maxDepthRecursion(root.right))
    );
}

function maxDepthBFS(root) {
    // time complexity: O(n)
    // space complexity: O(n)
    const queue = [];

    if (root !== null) {
        queue.push(root);
    }

    let level = 0;
    while (queue.length > 0) {
        let levelLength = queue.length;
        for (let i = 0; i < levelLength; i++) {
            const current = queue.shift();
            if (current.left !== null) {
                queue.push(current.left);
            }
            if (current.right !== null) {
                queue.push(current.right);
            }
        }
        level++;
    }
    return level;
}
