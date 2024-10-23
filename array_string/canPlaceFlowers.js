function canPlaceFlowers(flowerbed, t) {
    const n = flowerbed.length;
    let flowerAlreadyPlanted = 0;
    let totalFlower = 0;

    if (n % 2 === 0) {
        totalFlower = n / 2;
    } else {
        for (let i = 0; i < n; i++) {
            if (flowerbed[i] === 1) {
                totalFlower =
                    i % 2 === 0 ? Math.ceil(n / 2) : Math.floor(n / 2);
            }
        }
    }

    totalFlower = totalFlower === 0 ? Math.ceil(n / 2) : totalFlower;

    for (let i of flowerbed) {
        if (i === 1) {
            flowerAlreadyPlanted++;
        }
    }
    console.log(flowerAlreadyPlanted, totalFlower, t);

    return totalFlower - flowerAlreadyPlanted >= t;
}
const flowerBed = [
    0, 0, 1, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 1, 0, 0, 1, 0, 1, 0, 1, 0, 0, 0, 1,
    0, 1, 0, 1, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0,
    1, 0, 0, 1, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 1, 0,
    0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0,
];

// console.log(canPlaceFlowers([1, 0, 0, 0, 1], 1)); // true
// console.log(canPlaceFlowers([1, 0, 0, 0, 1], 2)); // false
// console.log(canPlaceFlowers([1, 0, 1, 0, 1], 1)); // false
// console.log(canPlaceFlowers([0, 0, 1, 0, 0], 1)); // true
// console.log(canPlaceFlowers([0, 0, 0, 0, 0], 2)); // true
// console.log(canPlaceFlowers([1, 0, 1, 0, 1, 0, 1], 0)); // true
// console.log(canPlaceFlowers([0], 1)); // true
console.log(canPlaceFlowers(flowerBed, 17)); // false fails
// console.log(flowerBed.length);
