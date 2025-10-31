// https://neetcode.io/problems/is-anagram
/**
 * Is Anagram
Given two strings s and t, return true if the two strings are anagrams of each other, otherwise return false.

An anagram is a string that contains the exact same characters as another string, but the order of the characters can be different.

Example 1:

Input: s = "racecar", t = "carrace"

Output: true
Example 2:

Input: s = "jar", t = "jam"

Output: false
Constraints:

s and t consist of lowercase English letters.
 */

// console.log(isAnagram("true", "ruta"));
// ttrue

function isAnagramHash(s, t) {
    if (s.length !== t.length) return false;

    const countS = new Map();
    const countT = new Map();

    for (let i = 0; i < s.length; i++) {
        countS.set(s[i], 1 + (countS.get(s[i]) || 0));
        countT.set(t[i], 1 + (countT.get(t[i]) || 0));
    }

    for (let [key, value] of countS) {
        if (countT.get(key) !== value) {
            return false;
        }
    }

    return true;
}
const result = isAnagramHash("ttruue", "ttruue");
console.log(result);

function isAnagramObj(s, t) {
    if (s.length !== t.length) {
        return false;
    }

    const countS = {};
    const countT = {};

    for (let i = 0; i < s.length; i++) {
        countS[s[i]] = 1 + (countS[s[i]] || 0);
        countT[t[i]] = 1 + (countT[t[i]] || 0);
    }

    for (const key in countS) {
        if (countS[key] !== countT[key]) {
            return false;
        }
    }

    return true;
}

// Time complexity using hashmap is O(s+t), space complexity, O(S+T)
//TODO: For space complexity O(1), sorting is the solution, sort first then compare

function isAnagramSorted(s, t) {
    s = s.split("").sort().join("");
    t = t.split("").sort().join("");

    return s === t;
}

console.log(isAnagramSorted("array", "yarraa"));
