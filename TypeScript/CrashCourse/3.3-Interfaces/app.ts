interface Person{
    name: string;
    age: number;
    height?: number;
    hello: () => void;
}

const person: Person = {
    name: "Ayden",
    age: 23,
    hello: function(){
        console.log(this.name + " says hi")
    }
}

// More advance interfaces

interface Employee extends Person {
    employeeId: number;

}

const worker: Employee = {
    name: "Tim",
    age: 23,
    height: 175,
    hello: function() {
        console.log(name + "says hi")
    },
    employeeId: 101
}

interface Manager extends Employee, Person{
    employees: Person[]

}


// Different instances

function getPerson(p: Person): Person{
    return{
        name: "Ayden",
        age: 25,
        hello: function(){
            console.log(name + " says hi")
        }
    }
}