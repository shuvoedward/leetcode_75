/** 
 * Top K Elements in List
Given an integer array nums and an integer k, return the k most frequent elements within the array.

The test cases are generated such that the answer is always unique.

You may return the output in any order.

Example 1:

Input: nums = [1,2,2,3,3,3], k = 2

Output: [2,3]
Example 2:

Input: nums = [7,7], k = 1

Output: [7]
*/

function topKFrequent(nums, k) {
    // Bucket Sort
    const count = {};
    const freq = Array.from({ length: nums.length + 1 }, () => []);

    for (const n of nums) {
        count[n] = (count[n] || 0) + 1;
    }

    for (const n in count) {
        freq[count[n]].push(parseInt(n));
    }

    const res = [];
    for (let i = freq.length - 1; i > 0; i--) {
        for (const n of freq[i]) {
            res.push(n);
            if (res.length === k) {
                return res;
            }
        }
    }
}

const nums = [1, 1, 1, 2, 2, 3];
const k = 2;

const result = topKFrequent(nums, k);

const practice = (nums, k) => {
    // count how many times an element in the array appears
    const count = {};
    const freq = Array.from({ length: nums.length + 1 }, () => []);
    const result = [];
    // keys of an obj are unique
    // in obj keys are only string, so any other data types will be converted to string
    for (let n of nums) {
        count[n] = (count[n] || 0) + 1;
    }

    for (let n in count) {
        // count[n] is the value, n is the key
        freq[count[n]].push(parseInt(n));
    }
    for (let i = freq.length - 1; i > 0; i--) {
        for (c of freq[i]) {
            result.push(c);
            if (result.length === k) return result;
        }
    }
};

console.log(practice(nums, k));
