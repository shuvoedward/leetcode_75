/**
  * Anagram Groups
Given an array of strings strs, group all anagrams 
together into sublists. You may return the output in any order.

An anagram is a string that contains the exact same characters 
as another string, but the order of the characters can be different.

Example 1:

Input: strs = ["act","pots","tops","cat","stop","hat"]

Output: [["hat"],["act", "cat"],["stop", "pots", "tops"]]
Example 2:

Input: strs = ["x"]

Output: [["x"]]
Example 3:

Input: strs = [""]

Output: [[""]]

*/

function groupAnagram(strs) {
    const ans = {};

    for (const str of strs) {
        const count = Array(26).fill(0);

        for (const c of str) {
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
strs = ["act", "pots", "tops", "cat", "stop", "hat"];
const answer = groupAnagram(strs);

console.log(answer);
