var Utility;
(function (Utility) {
    var multipliers = /** @class */ (function () {
        function multipliers() {
        }
        multipliers.prototype.timesTwo = function (n) {
            return n * 2;
        };
        return multipliers;
    }());
    Utility.multipliers = multipliers;
})(Utility || (Utility = {}));
