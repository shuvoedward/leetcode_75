function minWindow(s, t) {
    const map = Array(128).fill(0); // Create a map of size 128 for ASCII characters
    for (const char of t) {
        map[char.charCodeAt(0)]++; // Increment the count for characters in `t`
    }

    let counter = t.length, // Count of characters we need to match
        begin = 0, // Start of the window
        end = 0, // End of the window
        d = Infinity, // Minimum length of the substring
        head = 0; // Start of the valid substring

    while (end < s.length) {
        if (map[s.charCodeAt(end++)]-- > 0) {
            counter--; // Decrease counter if `s[end]` is part of `t`
        }
        /** 
        // Check if the character contributes to substring validity
        if (map[s.charCodeAt(end)] > 0) {
            counter--; // Reduce the count of required characters
        }

        // Decrement the frequency of the character in the map
        map[s.charCodeAt(end)]--;

        // Move the `end` pointer forward
        end++;

         */

        while (counter === 0) {
            // If we have a valid substring
            if (end - begin < d) {
                d = end - begin; // Update the minimum length
                head = begin; // Update the starting index of the substring
            }

            if (map[s.charCodeAt(begin++)]++ === 0) {
                counter++; // Make the substring invalid again
            }
            // s.charCodeAt(begin)++;
            // begin++;
        }
    }

    return d === Infinity ? "" : s.substring(head, head + d);
}

console.log(minWindow("OUZODYXAZV", "XYZ"));
