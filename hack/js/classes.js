var FilterOptions = /** @class */ (function () {
    function FilterOptions() {
    }
    return FilterOptions;
}());
var Filter = /** @class */ (function () {
    function Filter(options) {
        this.name = options.name;
        this.func = options.func;
    }
    Filter.prototype.prepare = function () {
        return {
            name: this.name,
            func: this.func,
        };
    };
    return Filter;
}());
var Config = /** @class */ (function () {
    function Config() {
        this.filters = [];
    }
    Config.prototype.toJSONString = function () {
        var res = {
            filters: this.filters.map(function (f) { return f.prepare(); }),
        };
        return JSON.stringify(res);
    };
    return Config;
}());
