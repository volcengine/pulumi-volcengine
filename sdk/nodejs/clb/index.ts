// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { AclArgs, AclState } from "./acl";
export type Acl = import("./acl").Acl;
export const Acl: typeof import("./acl").Acl = null as any;
utilities.lazyLoad(exports, ["Acl"], () => require("./acl"));

export { AclEntryArgs, AclEntryState } from "./aclEntry";
export type AclEntry = import("./aclEntry").AclEntry;
export const AclEntry: typeof import("./aclEntry").AclEntry = null as any;
utilities.lazyLoad(exports, ["AclEntry"], () => require("./aclEntry"));

export { AclsArgs, AclsResult, AclsOutputArgs } from "./acls";
export const acls: typeof import("./acls").acls = null as any;
export const aclsOutput: typeof import("./acls").aclsOutput = null as any;
utilities.lazyLoad(exports, ["acls","aclsOutput"], () => require("./acls"));

export { CertificateArgs, CertificateState } from "./certificate";
export type Certificate = import("./certificate").Certificate;
export const Certificate: typeof import("./certificate").Certificate = null as any;
utilities.lazyLoad(exports, ["Certificate"], () => require("./certificate"));

export { CertificatesArgs, CertificatesResult, CertificatesOutputArgs } from "./certificates";
export const certificates: typeof import("./certificates").certificates = null as any;
export const certificatesOutput: typeof import("./certificates").certificatesOutput = null as any;
utilities.lazyLoad(exports, ["certificates","certificatesOutput"], () => require("./certificates"));

export { ClbArgs, ClbState } from "./clb";
export type Clb = import("./clb").Clb;
export const Clb: typeof import("./clb").Clb = null as any;
utilities.lazyLoad(exports, ["Clb"], () => require("./clb"));

export { ClbsArgs, ClbsResult, ClbsOutputArgs } from "./clbs";
export const clbs: typeof import("./clbs").clbs = null as any;
export const clbsOutput: typeof import("./clbs").clbsOutput = null as any;
utilities.lazyLoad(exports, ["clbs","clbsOutput"], () => require("./clbs"));

export { GetAclsArgs, GetAclsResult, GetAclsOutputArgs } from "./getAcls";
export const getAcls: typeof import("./getAcls").getAcls = null as any;
export const getAclsOutput: typeof import("./getAcls").getAclsOutput = null as any;
utilities.lazyLoad(exports, ["getAcls","getAclsOutput"], () => require("./getAcls"));

export { GetCertificatesArgs, GetCertificatesResult, GetCertificatesOutputArgs } from "./getCertificates";
export const getCertificates: typeof import("./getCertificates").getCertificates = null as any;
export const getCertificatesOutput: typeof import("./getCertificates").getCertificatesOutput = null as any;
utilities.lazyLoad(exports, ["getCertificates","getCertificatesOutput"], () => require("./getCertificates"));

export { GetClbsArgs, GetClbsResult, GetClbsOutputArgs } from "./getClbs";
export const getClbs: typeof import("./getClbs").getClbs = null as any;
export const getClbsOutput: typeof import("./getClbs").getClbsOutput = null as any;
utilities.lazyLoad(exports, ["getClbs","getClbsOutput"], () => require("./getClbs"));

export { GetListenersArgs, GetListenersResult, GetListenersOutputArgs } from "./getListeners";
export const getListeners: typeof import("./getListeners").getListeners = null as any;
export const getListenersOutput: typeof import("./getListeners").getListenersOutput = null as any;
utilities.lazyLoad(exports, ["getListeners","getListenersOutput"], () => require("./getListeners"));

export { GetRulesArgs, GetRulesResult, GetRulesOutputArgs } from "./getRules";
export const getRules: typeof import("./getRules").getRules = null as any;
export const getRulesOutput: typeof import("./getRules").getRulesOutput = null as any;
utilities.lazyLoad(exports, ["getRules","getRulesOutput"], () => require("./getRules"));

export { GetServerGroupServersArgs, GetServerGroupServersResult, GetServerGroupServersOutputArgs } from "./getServerGroupServers";
export const getServerGroupServers: typeof import("./getServerGroupServers").getServerGroupServers = null as any;
export const getServerGroupServersOutput: typeof import("./getServerGroupServers").getServerGroupServersOutput = null as any;
utilities.lazyLoad(exports, ["getServerGroupServers","getServerGroupServersOutput"], () => require("./getServerGroupServers"));

export { GetServerGroupsArgs, GetServerGroupsResult, GetServerGroupsOutputArgs } from "./getServerGroups";
export const getServerGroups: typeof import("./getServerGroups").getServerGroups = null as any;
export const getServerGroupsOutput: typeof import("./getServerGroups").getServerGroupsOutput = null as any;
utilities.lazyLoad(exports, ["getServerGroups","getServerGroupsOutput"], () => require("./getServerGroups"));

export { GetZonesArgs, GetZonesResult, GetZonesOutputArgs } from "./getZones";
export const getZones: typeof import("./getZones").getZones = null as any;
export const getZonesOutput: typeof import("./getZones").getZonesOutput = null as any;
utilities.lazyLoad(exports, ["getZones","getZonesOutput"], () => require("./getZones"));

export { ListenerArgs, ListenerState } from "./listener";
export type Listener = import("./listener").Listener;
export const Listener: typeof import("./listener").Listener = null as any;
utilities.lazyLoad(exports, ["Listener"], () => require("./listener"));

export { ListenersArgs, ListenersResult, ListenersOutputArgs } from "./listeners";
export const listeners: typeof import("./listeners").listeners = null as any;
export const listenersOutput: typeof import("./listeners").listenersOutput = null as any;
utilities.lazyLoad(exports, ["listeners","listenersOutput"], () => require("./listeners"));

export { RuleArgs, RuleState } from "./rule";
export type Rule = import("./rule").Rule;
export const Rule: typeof import("./rule").Rule = null as any;
utilities.lazyLoad(exports, ["Rule"], () => require("./rule"));

export { RulesArgs, RulesResult, RulesOutputArgs } from "./rules";
export const rules: typeof import("./rules").rules = null as any;
export const rulesOutput: typeof import("./rules").rulesOutput = null as any;
utilities.lazyLoad(exports, ["rules","rulesOutput"], () => require("./rules"));

export { ServerGroupArgs, ServerGroupState } from "./serverGroup";
export type ServerGroup = import("./serverGroup").ServerGroup;
export const ServerGroup: typeof import("./serverGroup").ServerGroup = null as any;
utilities.lazyLoad(exports, ["ServerGroup"], () => require("./serverGroup"));

export { ServerGroupServerArgs, ServerGroupServerState } from "./serverGroupServer";
export type ServerGroupServer = import("./serverGroupServer").ServerGroupServer;
export const ServerGroupServer: typeof import("./serverGroupServer").ServerGroupServer = null as any;
utilities.lazyLoad(exports, ["ServerGroupServer"], () => require("./serverGroupServer"));

export { ServerGroupServersArgs, ServerGroupServersResult, ServerGroupServersOutputArgs } from "./serverGroupServers";
export const serverGroupServers: typeof import("./serverGroupServers").serverGroupServers = null as any;
export const serverGroupServersOutput: typeof import("./serverGroupServers").serverGroupServersOutput = null as any;
utilities.lazyLoad(exports, ["serverGroupServers","serverGroupServersOutput"], () => require("./serverGroupServers"));

export { ServerGroupsArgs, ServerGroupsResult, ServerGroupsOutputArgs } from "./serverGroups";
export const serverGroups: typeof import("./serverGroups").serverGroups = null as any;
export const serverGroupsOutput: typeof import("./serverGroups").serverGroupsOutput = null as any;
utilities.lazyLoad(exports, ["serverGroups","serverGroupsOutput"], () => require("./serverGroups"));

export { ZonesArgs, ZonesResult, ZonesOutputArgs } from "./zones";
export const zones: typeof import("./zones").zones = null as any;
export const zonesOutput: typeof import("./zones").zonesOutput = null as any;
utilities.lazyLoad(exports, ["zones","zonesOutput"], () => require("./zones"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "volcengine:clb/acl:Acl":
                return new Acl(name, <any>undefined, { urn })
            case "volcengine:clb/aclEntry:AclEntry":
                return new AclEntry(name, <any>undefined, { urn })
            case "volcengine:clb/certificate:Certificate":
                return new Certificate(name, <any>undefined, { urn })
            case "volcengine:clb/clb:Clb":
                return new Clb(name, <any>undefined, { urn })
            case "volcengine:clb/listener:Listener":
                return new Listener(name, <any>undefined, { urn })
            case "volcengine:clb/rule:Rule":
                return new Rule(name, <any>undefined, { urn })
            case "volcengine:clb/serverGroup:ServerGroup":
                return new ServerGroup(name, <any>undefined, { urn })
            case "volcengine:clb/serverGroupServer:ServerGroupServer":
                return new ServerGroupServer(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("volcengine", "clb/acl", _module)
pulumi.runtime.registerResourceModule("volcengine", "clb/aclEntry", _module)
pulumi.runtime.registerResourceModule("volcengine", "clb/certificate", _module)
pulumi.runtime.registerResourceModule("volcengine", "clb/clb", _module)
pulumi.runtime.registerResourceModule("volcengine", "clb/listener", _module)
pulumi.runtime.registerResourceModule("volcengine", "clb/rule", _module)
pulumi.runtime.registerResourceModule("volcengine", "clb/serverGroup", _module)
pulumi.runtime.registerResourceModule("volcengine", "clb/serverGroupServer", _module)
