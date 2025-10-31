// let question = "#javascript\nwhat is javascript?";
// // let question = "what is your name?\n#java";
// // let question = "what is your name? #hello";
// let start = 0,
//     end = 0;
// for (let i = 0; i < question.length; i++) {
//     if (question[i] === "#") {
//         start = i;
//         end = i + 1;
//         while (end < question.length) {
//             if (question[end] === "\n" || question[end] === " ") {
//                 break;
//             }
//             end++;
//         }
//         break;
//     }
// }

// console.log(
//     `word: ${question.substring(
//         start + 1,
//         end
//     )}, \n question: ${question.substring()}`
// );

// let cleanQuestion =
//     question.substring(0, start).trim() + question.substring(end).trim();
// console.log(cleanQuestion);

// let question = "what is your name?\n#java";
let question = "#javascript\nwhat is javascript?";
// let question = "what is your name? #hello";

// Extract the hashtagged word and clean the question
const match = question.match(/#(\S+)/); // Match hashtag and word after it
let hashtag = match ? match[1] : ""; // Extract hashtag word if exists
let cleanQuestion = question.replace(/#\S+/, "").trim(); // Remove hashtag and clean

console.log(match);
console.log(`Hashtag: ${hashtag}`);
console.log(`Question without hashtag: ${cleanQuestion}`);
