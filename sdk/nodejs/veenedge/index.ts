// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { AvailableResourcesArgs, AvailableResourcesResult, AvailableResourcesOutputArgs } from "./availableResources";
export const availableResources: typeof import("./availableResources").availableResources = null as any;
export const availableResourcesOutput: typeof import("./availableResources").availableResourcesOutput = null as any;
utilities.lazyLoad(exports, ["availableResources","availableResourcesOutput"], () => require("./availableResources"));

export { CloudServerArgs, CloudServerState } from "./cloudServer";
export type CloudServer = import("./cloudServer").CloudServer;
export const CloudServer: typeof import("./cloudServer").CloudServer = null as any;
utilities.lazyLoad(exports, ["CloudServer"], () => require("./cloudServer"));

export { CloudServersArgs, CloudServersResult, CloudServersOutputArgs } from "./cloudServers";
export const cloudServers: typeof import("./cloudServers").cloudServers = null as any;
export const cloudServersOutput: typeof import("./cloudServers").cloudServersOutput = null as any;
utilities.lazyLoad(exports, ["cloudServers","cloudServersOutput"], () => require("./cloudServers"));

export { GetAvailableResourcesArgs, GetAvailableResourcesResult, GetAvailableResourcesOutputArgs } from "./getAvailableResources";
export const getAvailableResources: typeof import("./getAvailableResources").getAvailableResources = null as any;
export const getAvailableResourcesOutput: typeof import("./getAvailableResources").getAvailableResourcesOutput = null as any;
utilities.lazyLoad(exports, ["getAvailableResources","getAvailableResourcesOutput"], () => require("./getAvailableResources"));

export { GetCloudServersArgs, GetCloudServersResult, GetCloudServersOutputArgs } from "./getCloudServers";
export const getCloudServers: typeof import("./getCloudServers").getCloudServers = null as any;
export const getCloudServersOutput: typeof import("./getCloudServers").getCloudServersOutput = null as any;
utilities.lazyLoad(exports, ["getCloudServers","getCloudServersOutput"], () => require("./getCloudServers"));

export { GetInstanceTypesArgs, GetInstanceTypesResult, GetInstanceTypesOutputArgs } from "./getInstanceTypes";
export const getInstanceTypes: typeof import("./getInstanceTypes").getInstanceTypes = null as any;
export const getInstanceTypesOutput: typeof import("./getInstanceTypes").getInstanceTypesOutput = null as any;
utilities.lazyLoad(exports, ["getInstanceTypes","getInstanceTypesOutput"], () => require("./getInstanceTypes"));

export { GetInstancesArgs, GetInstancesResult, GetInstancesOutputArgs } from "./getInstances";
export const getInstances: typeof import("./getInstances").getInstances = null as any;
export const getInstancesOutput: typeof import("./getInstances").getInstancesOutput = null as any;
utilities.lazyLoad(exports, ["getInstances","getInstancesOutput"], () => require("./getInstances"));

export { GetVpcsArgs, GetVpcsResult, GetVpcsOutputArgs } from "./getVpcs";
export const getVpcs: typeof import("./getVpcs").getVpcs = null as any;
export const getVpcsOutput: typeof import("./getVpcs").getVpcsOutput = null as any;
utilities.lazyLoad(exports, ["getVpcs","getVpcsOutput"], () => require("./getVpcs"));

export { InstanceArgs, InstanceState } from "./instance";
export type Instance = import("./instance").Instance;
export const Instance: typeof import("./instance").Instance = null as any;
utilities.lazyLoad(exports, ["Instance"], () => require("./instance"));

export { InstanceTypesArgs, InstanceTypesResult, InstanceTypesOutputArgs } from "./instanceTypes";
export const instanceTypes: typeof import("./instanceTypes").instanceTypes = null as any;
export const instanceTypesOutput: typeof import("./instanceTypes").instanceTypesOutput = null as any;
utilities.lazyLoad(exports, ["instanceTypes","instanceTypesOutput"], () => require("./instanceTypes"));

export { InstancesArgs, InstancesResult, InstancesOutputArgs } from "./instances";
export const instances: typeof import("./instances").instances = null as any;
export const instancesOutput: typeof import("./instances").instancesOutput = null as any;
utilities.lazyLoad(exports, ["instances","instancesOutput"], () => require("./instances"));

export { VpcArgs, VpcState } from "./vpc";
export type Vpc = import("./vpc").Vpc;
export const Vpc: typeof import("./vpc").Vpc = null as any;
utilities.lazyLoad(exports, ["Vpc"], () => require("./vpc"));

export { VpcsArgs, VpcsResult, VpcsOutputArgs } from "./vpcs";
export const vpcs: typeof import("./vpcs").vpcs = null as any;
export const vpcsOutput: typeof import("./vpcs").vpcsOutput = null as any;
utilities.lazyLoad(exports, ["vpcs","vpcsOutput"], () => require("./vpcs"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "volcengine:veenedge/cloudServer:CloudServer":
                return new CloudServer(name, <any>undefined, { urn })
            case "volcengine:veenedge/instance:Instance":
                return new Instance(name, <any>undefined, { urn })
            case "volcengine:veenedge/vpc:Vpc":
                return new Vpc(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("volcengine", "veenedge/cloudServer", _module)
pulumi.runtime.registerResourceModule("volcengine", "veenedge/instance", _module)
pulumi.runtime.registerResourceModule("volcengine", "veenedge/vpc", _module)
