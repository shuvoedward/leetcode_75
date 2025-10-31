function bfs(root) {
    let queue = [];

    if (root != null) {
        queue.push(root);
    }

    let level = 0;
    while (queue.length > 0) {
        console.log("level " + level + ": ");
        // using levelLength instead of queue.length directly in the for loop
        // that ensures the levelLenght does not update after each  queue.push(curr.left);
        // and for loop only runs for the first initial elements
        // if using queue.length directly in for loop it will evaluate each time its length
        let levelLength = queue.length;
        for (let i = 0; i < levelLength; i++) {
            // for loops allows to separate the tree by level
            let curr = queue.shift();
            console.log(curr.val + " ");
            if (curr.left != null) {
                queue.push(curr.left);
            }
            if (curr.right != null) {
                queue.push(curr.right);
            }
        }
        level++;
        console.log();
    }
}
