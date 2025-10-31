function maxArea(heights) {
    const areas = [];

    for (let i = 0; i < heights.length - 1; i++) {
        for (let k = i + 1; k < heights.length; k++) {
            const height = heights[i] < heights[k] ? heights[i] : heights[k];
            const width = k - i;

            areas.push(height * width);
        }
    }

    let maxArea = 0;
    for (let area of areas) {
        if (maxArea < area) {
            maxArea = area;
        }
    }
    console.log(areas);
    return maxArea;
}
const height = [1, 7, 2, 5, 4, 7, 3, 6];

function maxArea2(heights) {
    let maxArea = 0;

    for (let i = 0; i < heights.length - 1; i++) {
        for (let k = i + 1; k < heights.length; k++) {
            const height = heights[i] < heights[k] ? heights[i] : heights[k];
            const width = k - i;
            const area = height * width;
            if (maxArea < area) {
                maxArea = area;
            }
        }
    }

    console.log(maxArea);
}

function solution(height) {
    let maxArea = 0;

    let l = 0,
        r = height.length - 1;

    while (l < r) {
        // let h = height[l] < height[r] ? height[l] : height[r];
        let h = Math.max(height[l], height[r]);
        let area = h * (r - l);
        maxArea = maxArea < area ? area : maxArea;
        if (height[l] < height[r]) {
            l++;
        } else {
            r--;
        }
    }
    return maxArea;
}
/** 
 * I think a good explanation for why we move pointer with 
 * the lower height is because we already have the max area 
 * with that height - since it is the lower pointer that means 
 * that every other distance that is closer will always be a smaller 
 * distance with the same or less height which means smaller area. 
Therefore we do not need to look at every other combination with that pointer.
 */
console.log(solution(height));
console.log(solution([1, 8, 6, 2, 5, 4, 8, 3, 7]));
console.log(solution([3, 2, 4, 1, 3, 7]));
