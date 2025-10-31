// String Encode and Decode
// Design an algorithm to encode a list of strings to a single string. The encoded string is then decoded back to the original list of strings.

// Please implement encode and decode

// Example 1:

// Input: ["neet","code","love","you"]

// Output:["neet","code","love","you"]
// Example 2:

// Input: ["we","say",":","yes"]

// Output: ["we","say",":","yes"]

function encode(arr) {
    let createStr = "";
    for (let str of arr) {
        createStr += `${str.length}#${str}`;
    }
    return createStr;
}

const arr = ["neet", "#ode", ";", "you"];
const str = encode(arr);
console.log(str);

function decode(str) {
    // 4#neet4#code1#;3#you
    let arrStr = str.split("");
    let result = [];

    console.log(arrStr);
    /**
     * [
  '4', '#', 'n', 'e', 'e',
  't', '4', '#', 'c', 'o',
  'd', 'e', '1', '#', ';',
  '3', '#', 'y', 'o', 'u'
]
 
     */
    let i = 0;

    while (i < arrStr.length) {
        if (arrStr[i] === "#") {
            let wordLength = parseInt(arrStr[i - 1]);

            if (!isNaN(wordLength)) {
                let word = arrStr.slice(i + 1, wordLength + i + 1).join("");
                result.push(word);
                i += wordLength;
            }
        }
        i++;
    }
    // time complexity is = O(n)
    // TODO: I SOLVED THIS ON MY OWN. WELL, Got help from chatgpt but idea was mine.
    console.log(result);
}

function decodeStr(str) {
    // Time Complexity O(n)
    // Space Complexity O(1)

    // 4#neet4#code1#;3#you
    let result = [];
    let i = 0;

    while (i < str.length) {
        let j = i;
        while (str[j] !== "#") {
            j++;
        }
        let length = parseInt(str.substring(i, j), 10);
        i = j + 1;
        j = i + length;
        result.push(str.substring(i, j));
        i = j;
    }
    return result;
}

console.log(decodeStr(str));
