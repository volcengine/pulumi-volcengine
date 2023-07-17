// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./addon";
export * from "./addons";
export * from "./cluster";
export * from "./clusters";
export * from "./defaultNodePool";
export * from "./defaultNodePoolBatchAttach";
export * from "./kubeconfig";
export * from "./kubeconfigs";
export * from "./node";
export * from "./nodePool";
export * from "./nodePools";
export * from "./nodes";
export * from "./supportAddons";

// Import resources to register:
import { Addon } from "./addon";
import { Cluster } from "./cluster";
import { DefaultNodePool } from "./defaultNodePool";
import { DefaultNodePoolBatchAttach } from "./defaultNodePoolBatchAttach";
import { Kubeconfig } from "./kubeconfig";
import { Node } from "./node";
import { NodePool } from "./nodePool";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "volcengine:vke/addon:Addon":
                return new Addon(name, <any>undefined, { urn })
            case "volcengine:vke/cluster:Cluster":
                return new Cluster(name, <any>undefined, { urn })
            case "volcengine:vke/defaultNodePool:DefaultNodePool":
                return new DefaultNodePool(name, <any>undefined, { urn })
            case "volcengine:vke/defaultNodePoolBatchAttach:DefaultNodePoolBatchAttach":
                return new DefaultNodePoolBatchAttach(name, <any>undefined, { urn })
            case "volcengine:vke/kubeconfig:Kubeconfig":
                return new Kubeconfig(name, <any>undefined, { urn })
            case "volcengine:vke/node:Node":
                return new Node(name, <any>undefined, { urn })
            case "volcengine:vke/nodePool:NodePool":
                return new NodePool(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("volcengine", "vke/addon", _module)
pulumi.runtime.registerResourceModule("volcengine", "vke/cluster", _module)
pulumi.runtime.registerResourceModule("volcengine", "vke/defaultNodePool", _module)
pulumi.runtime.registerResourceModule("volcengine", "vke/defaultNodePoolBatchAttach", _module)
pulumi.runtime.registerResourceModule("volcengine", "vke/kubeconfig", _module)
pulumi.runtime.registerResourceModule("volcengine", "vke/node", _module)
pulumi.runtime.registerResourceModule("volcengine", "vke/nodePool", _module)