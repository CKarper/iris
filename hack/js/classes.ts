class FilterOptions {
    name: string;
    func: string;
}

class Filter {
    name: string;
    func: string;
    constructor(options: FilterOptions) {
        this.name = options.name;
        this.func = options.func;
    }
    prepare() : any {
        return {
            name: this.name,
            func: this.func,
        }
    }
}

class Config {
    filters: Filter[];
    constructor() {
        this.filters = [];
    }
    toJSONString() : string {
        const res = {
            filters: this.filters.map(f => f.prepare()),
        };
        return JSON.stringify(res)
    }
}