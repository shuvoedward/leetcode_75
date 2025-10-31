const board = [
    ["1", "2", ".", ".", "3", ".", ".", ".", "."],
    ["4", ".", ".", "5", ".", ".", ".", ".", "."],
    [".", "9", "8", ".", ".", ".", ".", ".", "3"],
    ["5", ".", ".", ".", "6", ".", ".", ".", "4"],
    [".", ".", ".", "8", ".", "3", ".", ".", "5"],
    ["7", ".", ".", ".", "2", ".", ".", ".", "6"],
    [".", ".", ".", ".", ".", ".", "2", ".", "."],
    [".", ".", ".", "4", "1", "9", ".", ".", "8"],
    [".", ".", ".", ".", "8", ".", ".", "7", "9"],
];

function isValid(board) {
    const cols = new Map();
    const rows = new Map();
    const squares = new Map(); // key = (r / 3) * 3 + c / 3

    for (let r = 0; r < 9; r++) {
        for (let c = 0; c < 9; c++) {
            const cell = board[r][c];
            if (cell === ".") {
                continue;
            }
            if (
                rows.get(r)?.has(cell) ||
                cols.get(c)?.has(cell) ||
                squares
                    .get(Math.floor(r / 3) * 3 + Math.floor(c / 3))
                    ?.has(cell)
            ) {
                return false;
            }
            cols.set(c, new Set(cols.get(c)).add(cell));
            rows.set(r, new Set(rows.get(r)).add(cell));
            squares.set(
                Math.floor(r / 3) * 3 + Math.floor(c / 3),
                new Set(
                    squares.get(Math.floor(r / 3) * 3 + Math.floor(c / 3))
                ).add(cell)
            );
        }
    }
    console.log(squares);
    console.log("------------------------");
    console.log(rows);
    return true;
}

const result = isValid(board);
console.log(result);

const board2 = [
    ["1", "2", ".", ".", "3", ".", ".", ".", "."],
    ["4", ".", ".", "5", ".", ".", ".", ".", "."],
    [".", "9", "8", ".", ".", ".", ".", ".", "3"],
    ["5", ".", ".", ".", "6", ".", ".", ".", "4"],
    [".", ".", ".", "8", ".", "3", ".", ".", "5"],
    ["7", ".", ".", ".", "2", ".", ".", ".", "6"],
    [".", ".", ".", ".", ".", ".", "2", ".", "."],
    [".", ".", ".", "4", "1", "9", ".", ".", "8"],
    [".", ".", ".", ".", "8", ".", ".", "7", "9"],
];

function practice(board) {
    // At the core, it is find the duplicates - use Hashmap
    // if a value exists in the hash - duplicate found - return false
    // if not push in the hashmap

    // how to navigate through the sudoku board
    // map = (key value) pair
    // rows = 0 row = ([values in the first row])
    // cols = 0 col = ([values in the first col])
    // squares = 0 square = ([values in the first square])

    const rows = new Map();
    const cols = new Map();
    const squares = new Map();

    for (let r = 0; r < 9; r++) {
        for (let c = 0; c < 9; c++) {
            // what to do when coming accross a dot, empty space - do nothing
            const cell = board[r][c];
            if (cell === ".") {
                continue;
            }

            // check if row, columns, and squares duplicate exists
            if (
                rows.get(r)?.has(cell) ||
                cols.get(c)?.has(cell) ||
                squares
                    .get(Math.floor(r / 3) * 3 + Math.floor(c / 3))
                    ?.has(cell)
            ) {
                return false;
            }
            cols.set(c, new Set(cols.get(c)).add(cell));
            rows.set(r, new Set(rows.get(r)).add(cell));
            squares.set(
                Math.floor(r / 3) * 3 + Math.floor(c / 3),
                new Set(
                    squares
                        .get(Math.floor(r / 3) * 3 + Math.floor(c / 3))
                        .add(cell)
                )
            );
        }
    }
    return true;
}
