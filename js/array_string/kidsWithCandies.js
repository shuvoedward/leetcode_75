function kidsWithCanides(candies, extraCandies) {
    const output = [];
    let maxCandy = 0;
    for (let i of candies) {
        if (i > maxCandy) {
            maxCandy = i;
        }
        // maxCandy = Math.max(maxCandy, i);
    }

    for (let i of candies) {
        // replacing if else using the below code
        // if(i+extraCandies < maxCandy) output.push(false) else output.push(true);
        output.push(i + extraCandies >= maxCandy);
    }
    return output;
}

console.log(kidsWithCanides([2, 3, 5, 1, 3], 3)); // [true, true, true, false, true]
console.log(kidsWithCanides([4, 2, 1, 1, 2], 1)); // [true, false, false, false, false]
console.log(kidsWithCanides([12, 1, 12], 10)); // [true, false, true]
