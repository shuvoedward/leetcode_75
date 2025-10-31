function lengthOfLongestSubstring(s) {
    // Optimal Solution.
    let mp = new Map();

    let l = 0,
        res = 0;

    for (let r = 0; r < s.length; r++) {
        if (mp.has(s[r])) {
            l = Math.max(mp.get(s[r]) + 1, l);
        }

        mp.set(s[r], r);
        res = Math.max(res, r - l + 1);
    }
    return res;
}

function usingSet(s) {
    const charSet = new Set();
    let l = 0;
    let res = 0;

    for (let r = 0; r < s.length; r++) {
        while (charSet.has(s[r])) {
            charSet.delete(s[l]);
            l++;
        }
        charSet.add(s[r]);
        res = Math.max(res, r - l + 1);
    }
    return res;
}

function bruteForce(s) {
    let res = 0;
    for (let i = 0; i < s.length; i++) {
        let charSet = new Set();
        for (let j = i; j < s.length; j++) {
            if (charSet.has(s[j])) {
                break;
            }
            charSet.add(s[j]);
        }
        res = Math.max(res, charSet.size);
    }
    return res;
}
