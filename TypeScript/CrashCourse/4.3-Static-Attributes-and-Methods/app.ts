class Dog{
    static instanceCount: number = 0;
    // This is an instance attribute vvv
    name: string;

    constructor(name:string ) {
        Dog.instanceCount++;
        this.name = name;
    }
    static decreaseCount(){
        this.instanceCount--;
    }
}

const dog1 = new Dog("Tom")
console.log(Dog.instanceCount)

console.log("\n")

const dog2 = new Dog("Jerry")
console.log(Dog.instanceCount)

console.log("Below is the decrease count. \n")
Dog.decreaseCount()
console.log(Dog.instanceCount)
