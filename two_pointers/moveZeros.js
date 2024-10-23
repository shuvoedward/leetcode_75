function moveZeros(nums) {
    if (nums.length <= 1) {
        return nums;
    }
    let l = 0,
        r = 1;

    while (l < nums.length && r < nums.length) {
        while (l < nums.length && nums[l] !== 0) {
            l++;
            r = l + 1;
        }
        while (r < nums.length && nums[r] === 0) {
            r++;
        }
        // console.log(r);
        if (l < nums.length && r < nums.length) {
            nums[l] = nums[r];
            nums[r] = 0;
            l++;
            r++;
        }
        // console.log(nums);
    }
    return nums;
}
// console.log(moveZeros([0, 1, 0, 3, 12]));
// console.log(moveZeros([1, 0, 0, 0, 12]));
// console.log(moveZeros([0, 0]));
// console.log(moveZeros([2, 1]));
console.log(moveZeros([4, 2, 4, 0, 0, 3, 0, 5, 1, 0]));
