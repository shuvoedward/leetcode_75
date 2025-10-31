function maxProfit(prices) {
    // Dynamic programming
    let maxProfit = 0;
    let minimum = prices[0];

    for (let i = 1; i < prices.length; i++) {
        minimum = Math.min(minimum, prices[i]);
        maxProfit = Math.max(maxProfit, prices[i] - minimum);
    }

    return maxProfit;
}

function maxProfit2(prices) {
    let minPrice = Infinity;
    let maxProfit = 0;

    for (let price of prices) {
        if (price < minPrice) {
            minPrice = price;
        } else if (price - minPrice > maxProfit) {
            maxProfit = price - minPrice;
        }
    }
    return maxProfit;
}

function twoPointers(prices) {
    // From neetcode
    let l = 0,
        r = 1;
    let maxP = 0;

    while (r < prices.length) {
        if (prices[l] < prices[r]) {
            let profit = prices[r] - prices[l];
            maxP = Math.max(maxP, profit);
        } else {
            l = r;
        }
        r++;
    }
    return maxP;
}

console.log(maxProfit([10, 1, 5, 6, 7, 1]));

console.log(maxProfit([10, 8, 7, 5, 2]));
console.log(maxProfit([7, 1, 5, 3, 6, 4]));
