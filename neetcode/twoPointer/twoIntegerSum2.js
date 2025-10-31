function twoIntSum(nums, target) {
    let l = 0,
        r = nums.length - 1;

    while (l < r) {
        let sum = nums[l] + nums[r];

        if (sum === target) {
            return [l, r];
        } else if (sum < target) {
            l++;
        } else {
            // there is always an answer
            r--;
        }
    }
    return [];
}

console.log(twoIntSum([1, 2, 3, 4], 5));
