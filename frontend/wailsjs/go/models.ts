export namespace main {
	
	export class CPUUsage {
	    avg: number;
	
	    static createFrom(source: any = {}) {
	        return new CPUUsage(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.avg = source["avg"];
	    }
	}

}

export namespace sys {
	
	export class MemoryInfo {
	    total: string;
	    available: string;
	    used: string;
	    used_percent: number;
	
	    static createFrom(source: any = {}) {
	        return new MemoryInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.total = source["total"];
	        this.available = source["available"];
	        this.used = source["used"];
	        this.used_percent = source["used_percent"];
	    }
	}

}

