// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./accounts";
export * from "./endpoint";
export * from "./endpoints";
export * from "./instance";
export * from "./instanceParameter";
export * from "./instanceParameterLogs";
export * from "./instanceParameters";
export * from "./instances";
export * from "./mongoAllowList";
export * from "./mongoAllowListAssociate";
export * from "./mongoAllowLists";
export * from "./regions";
export * from "./specs";
export * from "./sslState";
export * from "./sslStates";
export * from "./zones";

// Import resources to register:
import { Endpoint } from "./endpoint";
import { Instance } from "./instance";
import { InstanceParameter } from "./instanceParameter";
import { MongoAllowList } from "./mongoAllowList";
import { MongoAllowListAssociate } from "./mongoAllowListAssociate";
import { SslState } from "./sslState";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "volcengine:mongodb/endpoint:Endpoint":
                return new Endpoint(name, <any>undefined, { urn })
            case "volcengine:mongodb/instance:Instance":
                return new Instance(name, <any>undefined, { urn })
            case "volcengine:mongodb/instanceParameter:InstanceParameter":
                return new InstanceParameter(name, <any>undefined, { urn })
            case "volcengine:mongodb/mongoAllowList:MongoAllowList":
                return new MongoAllowList(name, <any>undefined, { urn })
            case "volcengine:mongodb/mongoAllowListAssociate:MongoAllowListAssociate":
                return new MongoAllowListAssociate(name, <any>undefined, { urn })
            case "volcengine:mongodb/sslState:SslState":
                return new SslState(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("volcengine", "mongodb/endpoint", _module)
pulumi.runtime.registerResourceModule("volcengine", "mongodb/instance", _module)
pulumi.runtime.registerResourceModule("volcengine", "mongodb/instanceParameter", _module)
pulumi.runtime.registerResourceModule("volcengine", "mongodb/mongoAllowList", _module)
pulumi.runtime.registerResourceModule("volcengine", "mongodb/mongoAllowListAssociate", _module)
pulumi.runtime.registerResourceModule("volcengine", "mongodb/sslState", _module)