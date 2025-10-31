function negativeMarking(nums) {
    for (let num of nums) {
        let idx = Math.abs(num) - 1;
        if (nums[idx] < 0) {
            return Math.abs(num);
        }
        nums[idx] *= -1;
    }
    return -1;
}

function hashSet(nums) {
    const seen = new Set();

    for (let num of nums) {
        if (seen.has(num)) {
            return num;
        }
        seen.add(num);
    }
    return -1;
}

function usingArray(nums) {
    let seen = new Array(nums.length).fill(0);

    for (let num of nums) {
        if (seen[num - 1]) {
            return nums;
        }
        seen[num - 1] = 1;
    }
    return -1;
}

function fastAndSlowPointer(nums) {
    let slow = 0;
    let fast = 0;
    while (true) {
        slow = nums[slow];
        fast = nums[nums[fast]];
        if (slow === fast) {
            break;
        }
    }

    let slow2 = 0;
    while (true) {
        slow = nums[slow];
        slow2 = nums[slow2];
        if (slow === slow2) {
            return slow;
        }
    }
}
