function findMin(nums) {
    if (nums.length < 2) {
        return nums[nums.length - 1];
    }
    let l = 0,
        r = nums.length - 1;

    let minimum = nums[0];
    while (l <= r) {
        let mid = Math.ceil((l + r) / 2);
        if (nums[l] < nums[mid]) {
            minimum = nums[l] < minimum ? nums[l] : minimum;

            l = mid + 1;
        } else {
            minimum = nums[mid] < minimum ? nums[mid] : minimum;

            r = mid - 1;
        }
    }
    return minimum;
}

console.log(findMin([3, 4, 5, 6, 1, 2])); // 1
console.log(findMin([4, 5, 0, 1, 2, 3])); // 0
console.log(findMin([4, 5, 6, 7])); // 4
console.log(findMin([9, 10, 0, 1, 2, 3, 4, 5, 6, 7, 8])); // 0
console.log(findMin([2, 3, 4, 5, 6, 7, 8, 9, 10, 0, 1])); // 0
console.log(findMin([1, 2])); // 1
console.log(findMin([2, 1])); // 1
