var n = 1;
n = 123;
var isWinter = false;
var count = 7;
var Name = "Nin";
var names = ["Va", "Lau", "Gre"];
var ferals;
(function (ferals) {
    ferals[ferals["Nin"] = 0] = "Nin";
    ferals[ferals["Va"] = 1] = "Va";
    ferals[ferals["Gre"] = 2] = "Gre";
    ferals[ferals["Lau"] = 3] = "Lau";
})(ferals || (ferals = {}));
;
var g = ferals.Gre;
function getName() {
    console.log("My name is, what?");
}
