// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { PrivateZoneArgs, PrivateZoneState } from "./privateZone";
export type PrivateZone = import("./privateZone").PrivateZone;
export const PrivateZone: typeof import("./privateZone").PrivateZone = null as any;
utilities.lazyLoad(exports, ["PrivateZone"], () => require("./privateZone"));

export { PrivateZonesArgs, PrivateZonesResult, PrivateZonesOutputArgs } from "./privateZones";
export const privateZones: typeof import("./privateZones").privateZones = null as any;
export const privateZonesOutput: typeof import("./privateZones").privateZonesOutput = null as any;
utilities.lazyLoad(exports, ["privateZones","privateZonesOutput"], () => require("./privateZones"));

export { RecordArgs, RecordState } from "./record";
export type Record = import("./record").Record;
export const Record: typeof import("./record").Record = null as any;
utilities.lazyLoad(exports, ["Record"], () => require("./record"));

export { RecordSetsArgs, RecordSetsResult, RecordSetsOutputArgs } from "./recordSets";
export const recordSets: typeof import("./recordSets").recordSets = null as any;
export const recordSetsOutput: typeof import("./recordSets").recordSetsOutput = null as any;
utilities.lazyLoad(exports, ["recordSets","recordSetsOutput"], () => require("./recordSets"));

export { RecordWeightEnablerArgs, RecordWeightEnablerState } from "./recordWeightEnabler";
export type RecordWeightEnabler = import("./recordWeightEnabler").RecordWeightEnabler;
export const RecordWeightEnabler: typeof import("./recordWeightEnabler").RecordWeightEnabler = null as any;
utilities.lazyLoad(exports, ["RecordWeightEnabler"], () => require("./recordWeightEnabler"));

export { RecordsArgs, RecordsResult, RecordsOutputArgs } from "./records";
export const records: typeof import("./records").records = null as any;
export const recordsOutput: typeof import("./records").recordsOutput = null as any;
utilities.lazyLoad(exports, ["records","recordsOutput"], () => require("./records"));

export { ResolverEndpointArgs, ResolverEndpointState } from "./resolverEndpoint";
export type ResolverEndpoint = import("./resolverEndpoint").ResolverEndpoint;
export const ResolverEndpoint: typeof import("./resolverEndpoint").ResolverEndpoint = null as any;
utilities.lazyLoad(exports, ["ResolverEndpoint"], () => require("./resolverEndpoint"));

export { ResolverEndpointsArgs, ResolverEndpointsResult, ResolverEndpointsOutputArgs } from "./resolverEndpoints";
export const resolverEndpoints: typeof import("./resolverEndpoints").resolverEndpoints = null as any;
export const resolverEndpointsOutput: typeof import("./resolverEndpoints").resolverEndpointsOutput = null as any;
utilities.lazyLoad(exports, ["resolverEndpoints","resolverEndpointsOutput"], () => require("./resolverEndpoints"));

export { ResolverRuleArgs, ResolverRuleState } from "./resolverRule";
export type ResolverRule = import("./resolverRule").ResolverRule;
export const ResolverRule: typeof import("./resolverRule").ResolverRule = null as any;
utilities.lazyLoad(exports, ["ResolverRule"], () => require("./resolverRule"));

export { ResolverRulesArgs, ResolverRulesResult, ResolverRulesOutputArgs } from "./resolverRules";
export const resolverRules: typeof import("./resolverRules").resolverRules = null as any;
export const resolverRulesOutput: typeof import("./resolverRules").resolverRulesOutput = null as any;
utilities.lazyLoad(exports, ["resolverRules","resolverRulesOutput"], () => require("./resolverRules"));

export { UserVpcAuthorizationArgs, UserVpcAuthorizationState } from "./userVpcAuthorization";
export type UserVpcAuthorization = import("./userVpcAuthorization").UserVpcAuthorization;
export const UserVpcAuthorization: typeof import("./userVpcAuthorization").UserVpcAuthorization = null as any;
utilities.lazyLoad(exports, ["UserVpcAuthorization"], () => require("./userVpcAuthorization"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "volcengine:private_zone/privateZone:PrivateZone":
                return new PrivateZone(name, <any>undefined, { urn })
            case "volcengine:private_zone/record:Record":
                return new Record(name, <any>undefined, { urn })
            case "volcengine:private_zone/recordWeightEnabler:RecordWeightEnabler":
                return new RecordWeightEnabler(name, <any>undefined, { urn })
            case "volcengine:private_zone/resolverEndpoint:ResolverEndpoint":
                return new ResolverEndpoint(name, <any>undefined, { urn })
            case "volcengine:private_zone/resolverRule:ResolverRule":
                return new ResolverRule(name, <any>undefined, { urn })
            case "volcengine:private_zone/userVpcAuthorization:UserVpcAuthorization":
                return new UserVpcAuthorization(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("volcengine", "private_zone/privateZone", _module)
pulumi.runtime.registerResourceModule("volcengine", "private_zone/record", _module)
pulumi.runtime.registerResourceModule("volcengine", "private_zone/recordWeightEnabler", _module)
pulumi.runtime.registerResourceModule("volcengine", "private_zone/resolverEndpoint", _module)
pulumi.runtime.registerResourceModule("volcengine", "private_zone/resolverRule", _module)
pulumi.runtime.registerResourceModule("volcengine", "private_zone/userVpcAuthorization", _module)