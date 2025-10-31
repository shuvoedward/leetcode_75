function isPalindrome(sentence) {
    let L = 0,
        R = sentence.length - 1;

    while (L < R) {
        while (L < R && !isAlphaNum(sentence[L])) {
            L++;
        }
        while (L < R && !isAlphaNum(sentence[R])) {
            R--;
        }

        if (sentence[L].toLowerCase() !== sentence[R].toLowerCase()) {
            return false;
        }
        L++;
        R--;
    }
    return true;
}

function isAlphaNum(c) {
    const charCode = c.charCodeAt(0);
    // charCode would either be (A-Z) or (a-z) or (0-9)
    // if either three return true or else false

    return (
        (65 <= charCode && charCode <= 90) || // charCode is A-Z
        (97 <= charCode && charCode <= 122) || // a-z
        (48 <= charCode && charCode <= 57) // 0-9
    );
}

const result = isPalindrome("was it a car or a cat I saw?");
console.log(result);
