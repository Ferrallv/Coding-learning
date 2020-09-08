"use strict";
function filterByTerm(input, searchTerm, lookupKey = "url") {
    if (!searchTerm)
        throw Error("searchTerm cannot be empty");
    if (!input.length)
        throw Error("inputArr cannot be empty");
    const regex = new RegExp(searchTerm, "i");
    return input
        .filter(function (arrayElement) {
        return arrayElement[lookupKey].match(regex);
    });
}
const obj1 = { url: "string1" };
const obj2 = { url: "string2" };
const obj3 = { url: "string3" };
const arrOfLinks = [obj1, obj2, obj3];
const term = "java";
filterByTerm(arrOfLinks, term);
const tom = {
    name: "Tom",
    city: "Munich",
    age: 33,
    printDetails: function () {
        console.log(`${this.name} - ${this.city}`);
    }
};
