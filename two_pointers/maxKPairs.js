function maxOperations(nums, k) {
    const m = new Map();
    let maxOperations = 0,
        i = 0;

    while (i < nums.length) {
        // if value is greater than k, then skip it.
        if (nums[i] >= k) {
            i++;
            continue;
        }
        const remainder = k - nums[i];

        if (m.has(remainder)) {
            maxOperations++;
            nums.splice(i, 1);
            const lastIndex = m.get(remainder).pop();
            nums.splice(lastIndex, 1);
            if (m.get(remainder).length < 1) {
                m.delete(remainder);
            }
            i--;
        } else {
            if (m.has(nums[i])) {
                m.get(nums[i]).push(i);
                i++;
            } else {
                m.set(nums[i], [i]);
                i++;
            }
        }
    }
    // console.log(m);
    // console.log(nums);

    return maxOperations;
}
// console.log(maxOperations([1, 2, 3, 4], 5));
// console.log(maxOperations([3, 1, 5, 1, 1, 1, 1, 1, 2, 2, 3, 2, 2], 1));
// console.log(
//     maxOperations(
//         [2, 5, 4, 4, 1, 3, 4, 4, 1, 4, 4, 1, 2, 1, 2, 2, 3, 2, 4, 2],
//         3
//     )
// );

console.log(
    maxOperations(
        [
            13, 16, 49, 4, 56, 64, 83, 35, 20, 73, 53, 67, 73, 73, 17, 28, 20,
            16, 55, 16, 20, 66, 13, 46, 9, 14, 52, 70, 46, 66, 40, 21, 5, 88,
            48, 21, 21, 44, 27, 56, 75, 58, 57, 15, 27, 4, 51, 77, 17, 21, 65,
            17, 62, 84, 71, 78, 10, 67, 49, 8, 47, 55, 41, 86, 43, 48, 69, 58,
            62, 27, 38, 24, 12, 82, 38, 62, 82, 32, 29, 27, 38, 37, 78, 9, 74,
            90, 64, 16, 37, 22, 37, 46, 20, 47, 31, 16, 81, 28, 82, 39, 86, 59,
            11, 78, 12, 13, 71, 49, 69, 1, 37, 24, 79, 32, 25, 67, 42, 30, 16,
            23, 51, 66, 72, 20, 11, 90, 34, 81, 10, 86, 51, 68, 10, 62, 59, 33,
            49, 30, 81, 69, 80, 79, 54, 78, 87, 44, 40, 47, 78, 44, 30, 23, 41,
            89, 35, 6, 88, 79, 14, 10, 27, 54, 83, 36, 78, 82, 51, 1, 48, 28,
            72, 34, 41, 32, 47, 32, 42, 3, 25, 78, 28, 37, 27, 77, 32, 83, 29,
            56, 86, 80, 50, 59, 44, 51, 25, 41, 18, 83, 62, 4, 16, 80, 72, 7,
            34, 21, 81, 15, 20, 35, 15, 46, 55, 81, 2, 36, 70, 87, 52, 19, 76,
            18, 27, 81, 19, 78, 36, 6, 84, 32, 27, 7, 70, 67, 87, 90, 37, 75,
            80, 72, 60, 68, 6, 72, 12, 90, 83, 20, 42, 36, 62, 45, 49, 45, 56,
            51, 66, 48, 30, 49, 58, 9, 4, 56, 53, 30, 22, 7, 43, 23, 89, 46, 81,
            61, 37, 78, 30, 9, 55, 43, 76, 65, 68, 64, 31, 1, 80, 39, 72, 45,
            37, 88, 54, 23, 89, 13, 68, 26, 75, 86, 82, 69, 15, 25, 57, 38, 89,
            70, 47, 4, 7, 11, 57, 64, 10, 73, 15, 16, 58, 5, 39, 60, 86, 50, 1,
            85, 7, 40, 37, 58, 57, 52, 2, 13, 14, 73, 83, 29, 35, 76, 24, 55,
            52, 75, 74, 38, 59, 73, 36, 90, 66, 61, 74, 74,
        ],
        77
    )
);
// console.log(maxOperations([3, 1, 3, 4, 3], 6));

function maxOperationsTwo(nums, k) {
    let l = 0,
        r = nums.length - 1;
    let maxOperations = 0;
    while (l < r) {
        if (k === nums[l] + nums[r]) {
            maxOperations++;
            nums.splice(r, 1);
            nums.splice(l, 1);
            r -= 2;
        } else if (nums[l] + nums[r] > k) {
            if (l > r) {
                l++;
            } else {
                r--;
            }
        } else {
            if (l < r) {
                l++;
            } else {
                r--;
            }
        }
        console.log(nums);
    }
    return maxOperations;
}
// console.log(
//     maxOperationsTwo(
//         [2, 5, 4, 4, 1, 3, 4, 4, 1, 4, 4, 1, 2, 1, 2, 2, 3, 2, 4, 2],
//         3
//     )
// );
