class Person {
	name: string;
	constructor(name: string) {
		this.name = name;
	}
	dance(){
		console.log(this.name +  " is dancing");
	}
}

var bran = new Person("Bran");
bran.dance()

class AwesomePerson extends Person {
	dance() {
		console.log("SO AWESOME");
	}
}

var robb = new AwesomePerson("Robb")
robb.dance()
