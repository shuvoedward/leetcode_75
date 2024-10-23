const mergeAlternately = function (word1, word2) {
    let result = "";
    let l = 0;
    let length = Math.max(word1.length, word2.length);

    while (l < length) {
        let x = word1[l] || "";
        let y = word2[l] || "";
        result += x + y;
        l++;
    }

    return result;
};

function twoPointer(word1, word2) {
    const word1Length = word1.length;
    const word2Length = word2.length;

    let result = "";
    let i = 0,
        j = 0;

    while (i < word1Length || j < word2Length) {
        if (i < word1Length) {
            result += word1[i++];
        }
        if (j < word2Length) {
            result += word2[j++];
        }
    }
    return result;
    /**
     * Complexity Analysis
    Here, m is the length of word1 and n is the length of word2.

    Time complexity: O(m+n)

    We iterate over word1 and word2 once and push their letters 
    into result. It would take O(m+n) time.
    Space complexity: O(1)

    Without considering the space consumed by the input strings
    (word1 and word2) and the output string (result), we do not use more than constant space.
     */
}

// console.log(twoPointer("abc", "pqrstuv"));

function onePointer(word1, word2) {
    const word1Length = word1.length;
    const word2Length = word2.length;

    let result = "";

    for (let i = 0; i < Math.max(word1Length, word2Length); i++) {
        if (i < word1Length) result += word1[i];
        if (i < word2Length) result += word2[i];
    }
    return result;

    /**
     * Complexity Analysis
    Here, m is the length of word1 and n is the length of word2.

    Time complexity: O(m+n)

    We iterate over word1 and word2 once pushing their letters
     into result. It would take O(m+n) time.
    Space complexity: O(1)

    Without considering the space consumed by the input strings 
    (word1 and word2) and the output string (result), 
    we do not use more than constant space.
     */
}
console.log(onePointer("abc", "pqrstuv"));
