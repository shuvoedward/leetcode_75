function binarySearch(nums, target) {
    const left = 0;
    const right = nums.length - 1;

    while (left <= right) {
        // const mid = (left + right) / 2; // TODO: Update this code
        const mid = left + Math.floor((r - l) / 2);
        if (target < nums[mid]) {
            right = mid - 1;
        } else if (target > nums[mid]) {
            left = mid + 1;
        } else {
            return mid;
        }
    }
    return -1;
}
