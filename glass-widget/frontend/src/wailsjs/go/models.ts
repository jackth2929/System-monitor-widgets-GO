export namespace main {
	
	export class SystemStats {
	    hostname: string;
	    platform: string;
	    username: string;
	    userType: string;
	    cpuUsage: number;
	    cpuCores: number;
	    ramUsage: number;
	    ramUsed: number;
	    ramTotal: number;
	    diskCUsage: number;
	    diskDUsage: number;
	    netSentSpeed: number;
	    netRecvSpeed: number;
	
	    static createFrom(source: any = {}) {
	        return new SystemStats(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.hostname = source["hostname"];
	        this.platform = source["platform"];
	        this.username = source["username"];
	        this.userType = source["userType"];
	        this.cpuUsage = source["cpuUsage"];
	        this.cpuCores = source["cpuCores"];
	        this.ramUsage = source["ramUsage"];
	        this.ramUsed = source["ramUsed"];
	        this.ramTotal = source["ramTotal"];
	        this.diskCUsage = source["diskCUsage"];
	        this.diskDUsage = source["diskDUsage"];
	        this.netSentSpeed = source["netSentSpeed"];
	        this.netRecvSpeed = source["netRecvSpeed"];
	    }
	}
	export class TodoItem {
	    id: number;
	    text: string;
	    completed: boolean;
	
	    static createFrom(source: any = {}) {
	        return new TodoItem(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.text = source["text"];
	        this.completed = source["completed"];
	    }
	}

}

