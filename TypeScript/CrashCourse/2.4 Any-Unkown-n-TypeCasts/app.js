"use strict";
let x = 1;
// Unknown
let y = 1;
if (typeof y == "number") {
    const result = y + 1;
}
else if (typeof y == "string") {
    const result = y.length;
}
// Type Cast
let z = 1;
const result2 = z + 1;
function processFeedback(input) {
    console.log(`Processing:  ${input}`);
}
processFeedback("Great Service!");
processFeedback(5);
processFeedback(new Blob());
