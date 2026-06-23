function add(x: number ,y:number){
    if(x == 0){
        return  "Invalid"
    }
    return x + y
}
const result = add(19,5)
console.log(result)

function makeName(firstName: string, lastName: string, middleName?: string){
    if(middleName){
        return firstName + " " + " " + middleName + " " + lastName + " "
    }else{
        return firstName + " " + lastName
    }
}


const fullName = makeName("Ayden", "Lee")
console.log(fullName)


function callFunc(func: (f: string, l: string, m: string) => string,
                  param1: string,
                  param2: string){

}
callFunc(makeName, "Tim", "Ruscica")


// Deeper function inbeded calls

function mul(x: number, y: number): number {
    return  x * y
}

function div(x: number, y: number): number {
    return x / y
}

function applyFunc(funcs: ((a: number, b: number) => number)[] , values: [number, number][] ): number[] {
    const results: number[] = [];
    for (let i=0; i < funcs.length; i++){
        const args = values[i]
        const result = funcs[i](args[0], args[1])
        results.push(result)
    }
    return results
}

applyFunc([mul, div], [[1,2],[4,5]])