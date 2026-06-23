//literals
let directions: "north" | "south" | "west" | "east";
directions = "north"
// Another example of literals
let responseCode: 200| 400 | 201;
// Another
let result = true

// ENUM examples
enum Size {
    Small,
    Medium,
    Large
}

var size: Size = Size.Small

enum Directions {
    Up = "UP",
    Down = "DOWN",
    Left = "LEFT",
    Right = "RIGHT"
}

enum Description {
    SmallText = "this is some sub text to"
}

console.log(Description.SmallText)