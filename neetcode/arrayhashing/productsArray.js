//TODO: Brute force

function multiply(nums) {
    const output = [];

    for (let i = 0; i < nums.length; i++) {
        let product = 1;
        for (let j = 0; j < nums.length; j++) {
            if (!(i === j)) {
                product *= nums[j];
            }
        }
        output.push(product);
    }
    return output;
}

const nums = [1, 2, 4, 6];

function multiplyByDivision(nums) {
    //TODO: Time complexity O(n)
    //TODO: Space complexity O(n)

    let total = 1;
    const output = [];
    for (let n of nums) {
        total *= n;
    }

    for (let n of nums) {
        output.push(total / n);
    }

    console.log(output);
}

function withoutDivision(nums) {
    const output = new Array(nums.length).fill(1);
    let prefix = 1;
    let postfix = 1;

    for (let i = 0; i < nums.length; i++) {
        output[i] = prefix;
        prefix *= nums[i];
    }
    for (let i = nums.length - 1; i >= 0; i--) {
        output[i] *= postfix;
        postfix *= nums[i];
    }
    return output;
}

// multiplyByDivision(nums);
const result = withoutDivision(nums);
console.log(result);
