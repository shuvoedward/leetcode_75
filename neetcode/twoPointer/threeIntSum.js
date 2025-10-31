function threeSum(nums) {
    nums.sort((a, b) => a - b);
    const result = [];

    let i = 0;
    while (i < nums.length - 3) {
        if (nums[i] === 0) {
            break;
        }
        if (nums[i] === nums[i - 1]) {
            i++;
        }
        let j = i + 1;
        let k = nums.length - 1;

        console.log("here");
        while (j < k) {
            const sum = nums[i] + nums[j] + nums[k];
            if (sum === 0) {
                result.push([nums[i] + nums[j] + nums[k]]);
                j++;

                k--;
                while (j < k && nums[j] === nums[j - 1]) {
                    j++;
                }
            } else if (sum > 0) {
                j++;
            } else {
                k--;
            }
        }
        i++;
    }
    return result;
}

const nums = [-1, 0, 1, 2, -1, -4];
console.log(threeSum(nums));
