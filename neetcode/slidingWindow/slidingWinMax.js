// From leetcode
const maxSlidingWindow = function (nums, k) {
    const q = []; // stores *indices*
    const res = [];
    for (let i = 0; i < nums.length; i++) {
        while (q && nums[q[q.length - 1]] <= nums[i]) {
            q.pop();
        }
        q.push(i);
        // remove first element if it's outside the window
        if (q[0] === i - k) {
            q.shift();
        }
        // if window has k elements add to results
        // (first k-1 windows have < k elements because we start
        // from empty window and add 1 element each iteration)
        if (i >= k - 1) {
            res.push(nums[q[0]]);
        }
    }
    return res;
};

function maxSlidingWindowNeetcode(nums, k) {
    const n = nums.length;
    const output = new Array(n - k + 1);
    const q = new Deque();
    let l = 0,
        r = 0;

    while (r < n) {
        while (q.size() && nums[q.back()] < nums[r]) {
            q.popBack();
        }
        q.pushBack(r);

        if (l > q.front()) {
            // checks of the maximum value which is the first value of the q
            // is outside of the window range
            q.popFront();
        }

        if (r + 1 >= k) {
            output[l] = nums[q.front()];
            l++;
        }
        r++;
    }

    return output;
}
