function maxArea(heights) {
    // Todo: THIS IS NOT THE SOLUTION
    let totalArea = 0;
    let left = 0,
        right = 1;
    let leftBorder = 0;
    while (left < heights.length) {
        if (heights[left] > heights[right]) {
            leftBorder = heights[left];

            let heightInBetween = 0;
            console.log(`Left value ${left}`);
            console.log(`left border ${leftBorder}`);

            while (leftBorder > heights[right] && right < heights.length - 1) {
                heightInBetween += heights[right];
                right++;
                console.log(`Left ${left}, right ${right}`);
            }

            let rightBorder = heights[right];
            if (leftBorder <= rightBorder) {
                let height = Math.min(leftBorder, rightBorder);
                let width = right - left - 1;
                let area = height * width - heightInBetween;
                totalArea += area;
                console.log(`total area ${totalArea}`);
            }
        }
        left = right;
        right = left + 1;
    }
    return totalArea;
}
const heights = [0, 2, 0, 3, 1, 0, 1, 3, 2, 1];
// console.log(maxArea(heights));
// console.log(maxArea([4, 2, 0, 3, 2, 5]));
console.log(maxArea([0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1]));

function solution(height) {
    if (!height || height.length === 0) {
        return 0;
    }

    let l = 0;
    let r = height.length - 1;
    let leftMax = height[l];
    let rightMax = height[r];
    let res = 0;

    while (l < r) {
        if (leftMax < rightMax) {
            l++;
            leftMax = Math.max(leftMax, height[l]);
            res += leftMax - height[l];
        } else {
            r--;
            rightMax = Math.max(rightMax, height[r]);
            res += rightMax - height[r];
        }
    }
    return res;
}

/**
 * for height[i],
 * find leftmax and rightmax
 * take min of left and right max
 * if height[i] is greater than min, no water can be stored
 * if height[i] is less than min, water can be stored (min - height[i])
 *
 */
