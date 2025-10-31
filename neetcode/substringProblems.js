//  the template.

// For most substring problem, we are given a string 
// and need to find a substring of it which
// satisfy some restrictions. A general way is to use 
// a hashmap assisted with two pointers. The template is given below.


function findSubstring(s) {
    const map = Array(128).fill(0); // Create a map of size 128 for ASCII characters
    let counter; // Check whether the substring is valid
    let begin = 0, end = 0; // Two pointers: one for the head and one for the tail
    let d; // The length of the substring (result)

    // Initialize the hash map here
    for (/* initialization logic */) {
        // Populate the map as needed
    }

    while (end < s.length) {
        // Modify `counter` based on map and the current character
        if (map[s.charCodeAt(end++)]-- /* condition */) {
            /* modify counter here */
        }

        // Check the condition of the counter to determine substring validity
        while (/* counter condition */) {
            // Update `d` if finding the minimum
            /* update d here */

            // Increase `begin` to shrink the window and make it invalid/valid again
            if (map[s.charCodeAt(begin++)]++ /* condition */) {
                /* modify counter here */
            }
        }

        // Update `d` here if finding the maximum
    }

    return d;
}

// This is the implementation of above template. It is to find minimum window with character problem.
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
        
        while (counter === 0) { // If we have a valid substring
            if (end - begin < d) {
                d = end - begin; // Update the minimum length
                head = begin; // Update the starting index of the substring
            }
            
            if (map[s.charCodeAt(begin++)]++ === 0) {
                counter++; // Make the substring invalid again
            }
        }
    }
    
    return d === Infinity ? "" : s.substring(head, head + d);
}


// code for finding the length of the longest substring without repeating characters:


function lengthOfLongestSubstring(s) {
    const map = Array(128).fill(0); // Create a map for ASCII characters
    let counter = 0, // Tracks the number of repeating characters
        begin = 0, // Start of the window
        end = 0, // End of the window
        d = 0; // Maximum length of substring

    while (end < s.length) {
        if (map[s.charCodeAt(end++)]++ > 0) counter++; // Increment counter if character is a repeat

        while (counter > 0) { // If there's a repeating character, shrink the window
            if (map[s.charCodeAt(begin++)]-- > 1) counter--; // Remove the character from the window
        }

        d = Math.max(d, end - begin); // Update the maximum length while the substring is valid
    }

    return d;
}