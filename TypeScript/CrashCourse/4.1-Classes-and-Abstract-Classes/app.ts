class Person {
    protected name: string;

    constructor(name: string) {
        this.name = name
    }

    greet() {
        console.log(`Hello, my nam eis ${this.name}`);
    }

    getName() {
        if (this.name.length < 2) {
            return ""
        }
        return this.name
    }

    setName(name: string) {
        if (name.length < 5) {
            return
        } else {
            this.name = name
        }
    }
}

// Protected methods


class Employee extends Person{
    callMe() {
        console.log(this.name)
    }
}

const p1 = new Person("Ayden");
p1.getName()
