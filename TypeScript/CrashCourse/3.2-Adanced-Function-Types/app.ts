function sum(...number: number[]){

}

sum(1,2,3,4,)
sum(1,2)
sum(1,2,3,4,6,8,2)


function getItemLength(name: string): number;

function getItemLength(names: string[]): number;

function getItemLength(nameOrNames: unknown): number{
    if (typeof nameOrNames === "string"){
        return nameOrNames.length
    } else if(Array.isArray(nameOrNames)){
        return nameOrNames.length
    }
    return 0
}


const x = getItemLength(["ayden","matt", "joe"]);
console.log("This is how many names I have: " + x);