interface Stark {
	name: string;
}

function printName(stark : Stark) {
	console.log(stark.name);
}

printName({name:"Eddard"});
printName({name:"Joe"});
