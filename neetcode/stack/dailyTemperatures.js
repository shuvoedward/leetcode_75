function dailyTemperatures(temp) {
    const result = new Array(temp.length).fill(0);

    for (let i = 0; i < temp.length; i++) {
        for (let j = i + 1; j < temp.length; j++) {
            if (temp[i] < temp[j]) {
                result[i] = j - i;
                break;
            }
        }
    }
    return result;
}
// console.log(dailyTemperatures([30, 38, 30, 36, 35, 40, 28]));
// console.log(dailyTemperatures([22, 21, 20]));

const solution = (temp) => {
    const result = new Array(temp.length).fill(0);
    const stack = [];

    for (let i = 0; i < temp.length; i++) {
        const t = temp[i];
        while (stack.length > 0 && stack[stack.length - 1][0] < t) {
            const [stackT, stackIndex] = stack.pop();
            result[stackIndex] = i - stackIndex;
        }
        stack.push([t, i]);
    }
    return result;
    // Time complexity - O(n)
    // Space complexity - O(n)
};

console.log(solution([30, 38, 30, 36, 35, 40, 28]));
