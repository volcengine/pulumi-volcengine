// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { ConsumedPartitionsArgs, ConsumedPartitionsResult, ConsumedPartitionsOutputArgs } from "./consumedPartitions";
export const consumedPartitions: typeof import("./consumedPartitions").consumedPartitions = null as any;
export const consumedPartitionsOutput: typeof import("./consumedPartitions").consumedPartitionsOutput = null as any;
utilities.lazyLoad(exports, ["consumedPartitions","consumedPartitionsOutput"], () => require("./consumedPartitions"));

export { ConsumedTopicsArgs, ConsumedTopicsResult, ConsumedTopicsOutputArgs } from "./consumedTopics";
export const consumedTopics: typeof import("./consumedTopics").consumedTopics = null as any;
export const consumedTopicsOutput: typeof import("./consumedTopics").consumedTopicsOutput = null as any;
utilities.lazyLoad(exports, ["consumedTopics","consumedTopicsOutput"], () => require("./consumedTopics"));

export { GroupArgs, GroupState } from "./group";
export type Group = import("./group").Group;
export const Group: typeof import("./group").Group = null as any;
utilities.lazyLoad(exports, ["Group"], () => require("./group"));

export { GroupsArgs, GroupsResult, GroupsOutputArgs } from "./groups";
export const groups: typeof import("./groups").groups = null as any;
export const groupsOutput: typeof import("./groups").groupsOutput = null as any;
utilities.lazyLoad(exports, ["groups","groupsOutput"], () => require("./groups"));

export { InstanceArgs, InstanceState } from "./instance";
export type Instance = import("./instance").Instance;
export const Instance: typeof import("./instance").Instance = null as any;
utilities.lazyLoad(exports, ["Instance"], () => require("./instance"));

export { InstancesArgs, InstancesResult, InstancesOutputArgs } from "./instances";
export const instances: typeof import("./instances").instances = null as any;
export const instancesOutput: typeof import("./instances").instancesOutput = null as any;
utilities.lazyLoad(exports, ["instances","instancesOutput"], () => require("./instances"));

export { PublicAddressArgs, PublicAddressState } from "./publicAddress";
export type PublicAddress = import("./publicAddress").PublicAddress;
export const PublicAddress: typeof import("./publicAddress").PublicAddress = null as any;
utilities.lazyLoad(exports, ["PublicAddress"], () => require("./publicAddress"));

export { RegionsArgs, RegionsResult, RegionsOutputArgs } from "./regions";
export const regions: typeof import("./regions").regions = null as any;
export const regionsOutput: typeof import("./regions").regionsOutput = null as any;
utilities.lazyLoad(exports, ["regions","regionsOutput"], () => require("./regions"));

export { SaslUserArgs, SaslUserState } from "./saslUser";
export type SaslUser = import("./saslUser").SaslUser;
export const SaslUser: typeof import("./saslUser").SaslUser = null as any;
utilities.lazyLoad(exports, ["SaslUser"], () => require("./saslUser"));

export { SaslUsersArgs, SaslUsersResult, SaslUsersOutputArgs } from "./saslUsers";
export const saslUsers: typeof import("./saslUsers").saslUsers = null as any;
export const saslUsersOutput: typeof import("./saslUsers").saslUsersOutput = null as any;
utilities.lazyLoad(exports, ["saslUsers","saslUsersOutput"], () => require("./saslUsers"));

export { TopicArgs, TopicState } from "./topic";
export type Topic = import("./topic").Topic;
export const Topic: typeof import("./topic").Topic = null as any;
utilities.lazyLoad(exports, ["Topic"], () => require("./topic"));

export { TopicPartitionsArgs, TopicPartitionsResult, TopicPartitionsOutputArgs } from "./topicPartitions";
export const topicPartitions: typeof import("./topicPartitions").topicPartitions = null as any;
export const topicPartitionsOutput: typeof import("./topicPartitions").topicPartitionsOutput = null as any;
utilities.lazyLoad(exports, ["topicPartitions","topicPartitionsOutput"], () => require("./topicPartitions"));

export { TopicsArgs, TopicsResult, TopicsOutputArgs } from "./topics";
export const topics: typeof import("./topics").topics = null as any;
export const topicsOutput: typeof import("./topics").topicsOutput = null as any;
utilities.lazyLoad(exports, ["topics","topicsOutput"], () => require("./topics"));

export { ZonesArgs, ZonesResult, ZonesOutputArgs } from "./zones";
export const zones: typeof import("./zones").zones = null as any;
export const zonesOutput: typeof import("./zones").zonesOutput = null as any;
utilities.lazyLoad(exports, ["zones","zonesOutput"], () => require("./zones"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "volcengine:kafka/group:Group":
                return new Group(name, <any>undefined, { urn })
            case "volcengine:kafka/instance:Instance":
                return new Instance(name, <any>undefined, { urn })
            case "volcengine:kafka/publicAddress:PublicAddress":
                return new PublicAddress(name, <any>undefined, { urn })
            case "volcengine:kafka/saslUser:SaslUser":
                return new SaslUser(name, <any>undefined, { urn })
            case "volcengine:kafka/topic:Topic":
                return new Topic(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("volcengine", "kafka/group", _module)
pulumi.runtime.registerResourceModule("volcengine", "kafka/instance", _module)
pulumi.runtime.registerResourceModule("volcengine", "kafka/publicAddress", _module)
pulumi.runtime.registerResourceModule("volcengine", "kafka/saslUser", _module)
pulumi.runtime.registerResourceModule("volcengine", "kafka/topic", _module)