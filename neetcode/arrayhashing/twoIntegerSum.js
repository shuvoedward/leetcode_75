// https://neetcode.io/problems/two-integer-sum

/**
 * Two Integer Sum
Given an array of integers nums and an integer target, return the indices i and j such that nums[i] + nums[j] == target and i != j.

You may assume that every input has exactly one pair of indices i and j that satisfy the condition.

Return the answer with the smaller index first.

Example 1:

Input: 
nums = [3,4,5,6], target = 7

Output: [0,1]
Explanation: nums[0] + nums[1] == 7, so we return [0, 1].

Example 2:

Input: nums = [4,5,6], target = 10

Output: [0,2]
Example 3:

Input: nums = [5,5], target = 10

Output: [0,1]
 */

function hasTwoSum(nums, target) {
    for (let i = 0; i < nums.length - 1; i++) {
        let result = target - nums[0];
        for (let j = i + 1; j < nums.length; j++) {
            if (nums[j] === result) {
                console.log(i, j);
                return;
            }
        }
    }
}

hasTwoSum([4, 5, 6], 10);

function hasTwoSumHash(nums, target) {
    const map = new Map();

    // take the compliment
    // check if the compliment in the hash.
    // if not then set the number and index

    for (let i = 0; i < nums.length; i++) {
        const compliment = target - nums[i];

        if (map.has(compliment)) {
            return [i, map.get(compliment)];
        }

        map.set(nums[i], i);
    }
    //TODO: Duplicate won't work

    return [-1, -1];
    // time complexity O(n)
    // space O(n)
}
