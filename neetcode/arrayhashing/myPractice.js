// TODO: How to check if an array has duplicates? - HashMap js set

const nums = [1, 2, 3, 1];
function hasDuplicates(nums) {
    const count = new Set();
    for (let n of nums) {
        if (count.has(n)) return true;
        count.add(n);
    }
    return false;
}

// TODO: How to count frequency? Use map or obj
function isAnagram(s, t) {
    if (s.length !== t.length) {
        return false;
    }

    const countS = new Map();
    const countT = new Map();

    for (let i = 0; i < s.length; i++) {
        countS.set(s[i], (countS.get(s[i]) || 0) + 1);
        countT.set(t[i], (countT.get(t[i]) || 0) + 1);
    }

    for (let c of s) {
        if (countS.get(c) !== countT.get(c)) {
            // when two values are not same return false
            return false;
        }
    }
    return true;
}
// TODO: How to store index? use map or obj
function twoSum(nums, target) {
    const count = new Map();

    for (let i = 0; i < nums.length; i++) {
        const remainder = target - count[i];
        if (count.has(remainder)) {
            return [count.get(remainder), count.get(count[i])];
        }
    }
}

const strs = ["act", "pots", "tops", "cat", "stop", "hat"];

// ans = { 0#0# : [cat, tac, ...]}
// ans = {0#1# : []}
// TODO: time complexity, loop through the whole array, n elements,
// TODO: loop through each string, m length. O(n * m) m = average length of strings
function groupAnagram(strs) {
    const ans = {};

    for (let str of strs) {
        const count = new Array(26).fill(0);

        for (let c of str) {
            count[c.charCodeAt(0) - "a".charCodeAt(0)]++;
        }
        const key = count.join("#");

        if (!ans[key]) {
            ans[key] = [];
        }
        ans[key].push(str);
    }
    return Object.values(ans);
}

function kFrequent(nums, k) {
    const ans = new Array.from({ length: nums.length + 1 }, () => []);
    const result = [];

    const count = new Map();

    for (let n of nums) {
        count.set(n, (count.get(n) || 0) + 1);
    }

    for (let [key, values] of count) {
        ans[values].push(key);
    }

    for (let i = ans.length - 1; i > 0; i--) {
        for (n of ans[i]) {
            result.push(n);
            if (result.length === k) return result;
        }
    }
}

function productExceptSelf(nums) {
    const result = [];

    let prefix = 1,
        postFix = 1;

    // Prefix
    for (let i = 0; i < nums.length; i++) {
        result.push(prefix);
        prefix *= nums[i];
    }

    // Postfix
    for (let i = nums.length - 1; i >= 0; i--) {
        result[i] = result[i] * postFix;
        postFix *= nums[i];
    }
    return result;
}
console.log(productExceptSelf([2, 3, 4, 5]));
console.log(productExceptSelf([1, 2, 4, 6]));
console.log(productExceptSelf([-1, 0, 1, 2, 3]));
