let x: any = 1


// Unknown
let y: unknown = 1;

if (typeof y == "number"){
    const result = y + 1
} else if (typeof y == "string"){
    const result = y.length
}

// Type Cast
let z: unknown = 1
const result2 = (z as number) + 1


function processFeedback(input: any): void{
    console.log(`Processing:  ${input}`);
}
processFeedback("Great Service!");
processFeedback(5);
processFeedback(new Blob());