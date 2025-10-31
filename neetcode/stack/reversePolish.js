function reversePolishNotation(tokens) {
    const stack = [];
    const operators = new Map([
        ["+", (a, b) => a + b],
        ["-", (a, b) => a - b],

        ["*", (a, b) => a * b],
        ["/", (a, b) => parseInt(a / b)],
    ]);

    for (let i of tokens) {
        if (operators.has(i)) {
            let b = stack.pop();
            let a = stack.pop();

            const result = operators.get(i)(a, b);
            stack.push(result);

            console.log(`Operator  ${stack}`);
        } else {
            stack.push(parseInt(i, 10));
            console.log(stack);
        }
    }

    return stack.pop();
}

// console.log(reversePolishNotation(["2", "1", "+", "3", "*"])); // 9
// console.log(reversePolishNotation(["4", "13", "5", "/", "+"])); // 6
// console.log(
//     reversePolishNotation([
//         "10",
//         "6",
//         "9",
//         "3",
//         "+",
//         "-11",
//         "*",
//         "/",
//         "*",
//         "17",
//         "+",
//         "5",
//         "+",
//     ])
// ); // 6

console.log(reversePolishNotation(["4", "-2", "/", "2", "-3", "-", "-"]));

function neetcode(tokens) {
    const stack = [];
    for (const c of tokens) {
        if (c === "+") {
            stack.push(stack.pop() + stack.pop());
        } else if (c === "-") {
            const a = stack.pop();
            const b = stack.pop();
            stack.push(b - a);
        } else if (c === "*") {
            stack.push(stack.pop() * stack.pop());
        } else if (c === "/") {
            const a = stack.pop();
            const b = stack.pop();
            stack.push(Math.trunc(b / a));
        } else {
            stack.push(parseInt(c));
        }
    }
    return stack.pop();
}
