function search(matrix, target) {
    const ROWS = matrix.length;
    const COLS = matrix[0].length;

    let top = 0;
    let bot = ROWS - 1;
    while (top <= bot) {
        const row = Math.floor((top + bot) / 2);
        if (target > matrix[row][COLS - 1]) {
            top = row + 1;
        } else if (target < matrix[row][0]) {
            bot = row - 1;
        } else {
            break;
        }
    }

    if (!(top <= bot)) {
        return false;
    }
    const row = Math.floor((top + bot) / 2);
    let l = 0;
    let r = COLS - 1;
    while (l <= r) {
        const m = Math.floor((l + r) / 2);
        if (target > matrix[row][m]) {
            l = m + 1;
        } else if (target < matrix[row][m]) {
            r = m - 1;
        } else {
            return true;
        }
    }
    return false;
}

const matrix = [
    [1, 2, 4, 8],
    [10, 11, 12, 13],
    [14, 20, 30, 40],
];
const target = 10;
console.log(search(matrix, target));
