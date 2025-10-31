function carFleet(target, position, speed) {
    const n = position.length;
    const pairs = position.map((p, i) => [p, speed[i]]);

    pairs.sort((a, b) => b[0] - a[0]); // This is nlogn. Sorting takes up all the time

    let fleetCount = 0;
    const timeToReach = new Array(n); // Stack

    for (let i = 0; i < n; i++) {
        timeToReach[i] = (target - pairs[i][0]) / pairs[i][1];

        if (i >= 1 && timeToReach[i] <= timeToReach[i - 1]) {
            timeToReach[i] = timeToReach[i - 1];
        } else {
            fleetCount++;
        }
    }
    return fleetCount;
}
console.log(carFleet(10, [1, 4], [3, 2]));
console.log(carFleet(10, [4, 1, 0, 7], [2, 2, 1, 1]));
console.log(carFleet(10, [4, 1, 0, 3, 7], [2, 2, 1, 3, 1]));
