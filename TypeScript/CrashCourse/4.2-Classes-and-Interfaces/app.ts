interface Animal{
    speak(): void
}

class Dog implements Animal{
    private name: string;
    private color: string;

    constructor(name: string, color: string){
        this.name = name;
        this.color = color;
    }

    speak(){
        console.log(`I am ${this.name} the dog and I am ${this.color}`)
    }

    test(){
        return 1
    }
}

class Cat implements Animal{
    private name: string;
    private color;

    constructor(name: string, color: string) {
        this.name = name;
        this.color = color;
    }
    speak(){
        console.log(`I am ${this.name} the cat and I am ${this.color}`)
    }
}

const dog: Animal = new Dog("Jim", "brown");
dog.speak()

console.log("\n")
const cat: Animal = new Cat("Tom", "orange");
cat.speak()




