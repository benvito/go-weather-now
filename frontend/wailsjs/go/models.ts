export namespace parser {
	
	export class WeatherInfo {
	    // Go type: time
	    time: any;
	    temperature: number;
	    rainly_chance: number;
	    wet: number;
	    wind: number;
	    // Go type: Image
	    image: any;
	
	    static createFrom(source: any = {}) {
	        return new WeatherInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.time = this.convertValues(source["time"], null);
	        this.temperature = source["temperature"];
	        this.rainly_chance = source["rainly_chance"];
	        this.wet = source["wet"];
	        this.wind = source["wind"];
	        this.image = this.convertValues(source["image"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

