// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { GetInstancesArgs, GetInstancesResult, GetInstancesOutputArgs } from "./getInstances";
export const getInstances: typeof import("./getInstances").getInstances = null as any;
export const getInstancesOutput: typeof import("./getInstances").getInstancesOutput = null as any;
utilities.lazyLoad(exports, ["getInstances","getInstancesOutput"], () => require("./getInstances"));

export { GetRegionsArgs, GetRegionsResult, GetRegionsOutputArgs } from "./getRegions";
export const getRegions: typeof import("./getRegions").getRegions = null as any;
export const getRegionsOutput: typeof import("./getRegions").getRegionsOutput = null as any;
utilities.lazyLoad(exports, ["getRegions","getRegionsOutput"], () => require("./getRegions"));

export { GetZonesArgs, GetZonesResult, GetZonesOutputArgs } from "./getZones";
export const getZones: typeof import("./getZones").getZones = null as any;
export const getZonesOutput: typeof import("./getZones").getZonesOutput = null as any;
utilities.lazyLoad(exports, ["getZones","getZonesOutput"], () => require("./getZones"));

export { InstanceArgs, InstanceState } from "./instance";
export type Instance = import("./instance").Instance;
export const Instance: typeof import("./instance").Instance = null as any;
utilities.lazyLoad(exports, ["Instance"], () => require("./instance"));

export { InstancesArgs, InstancesResult, InstancesOutputArgs } from "./instances";
export const instances: typeof import("./instances").instances = null as any;
export const instancesOutput: typeof import("./instances").instancesOutput = null as any;
utilities.lazyLoad(exports, ["instances","instancesOutput"], () => require("./instances"));

export { RegionsArgs, RegionsResult, RegionsOutputArgs } from "./regions";
export const regions: typeof import("./regions").regions = null as any;
export const regionsOutput: typeof import("./regions").regionsOutput = null as any;
utilities.lazyLoad(exports, ["regions","regionsOutput"], () => require("./regions"));

export { ZonesArgs, ZonesResult, ZonesOutputArgs } from "./zones";
export const zones: typeof import("./zones").zones = null as any;
export const zonesOutput: typeof import("./zones").zonesOutput = null as any;
utilities.lazyLoad(exports, ["zones","zonesOutput"], () => require("./zones"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "volcengine:escloud/instance:Instance":
                return new Instance(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("volcengine", "escloud/instance", _module)
