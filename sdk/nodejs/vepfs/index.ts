// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { FileSystemArgs, FileSystemState } from "./fileSystem";
export type FileSystem = import("./fileSystem").FileSystem;
export const FileSystem: typeof import("./fileSystem").FileSystem = null as any;
utilities.lazyLoad(exports, ["FileSystem"], () => require("./fileSystem"));

export { FileSystemsArgs, FileSystemsResult, FileSystemsOutputArgs } from "./fileSystems";
export const fileSystems: typeof import("./fileSystems").fileSystems = null as any;
export const fileSystemsOutput: typeof import("./fileSystems").fileSystemsOutput = null as any;
utilities.lazyLoad(exports, ["fileSystems","fileSystemsOutput"], () => require("./fileSystems"));

export { FilesetArgs, FilesetState } from "./fileset";
export type Fileset = import("./fileset").Fileset;
export const Fileset: typeof import("./fileset").Fileset = null as any;
utilities.lazyLoad(exports, ["Fileset"], () => require("./fileset"));

export { FilesetsArgs, FilesetsResult, FilesetsOutputArgs } from "./filesets";
export const filesets: typeof import("./filesets").filesets = null as any;
export const filesetsOutput: typeof import("./filesets").filesetsOutput = null as any;
utilities.lazyLoad(exports, ["filesets","filesetsOutput"], () => require("./filesets"));

export { MountServiceArgs, MountServiceState } from "./mountService";
export type MountService = import("./mountService").MountService;
export const MountService: typeof import("./mountService").MountService = null as any;
utilities.lazyLoad(exports, ["MountService"], () => require("./mountService"));

export { MountServiceAttachmentArgs, MountServiceAttachmentState } from "./mountServiceAttachment";
export type MountServiceAttachment = import("./mountServiceAttachment").MountServiceAttachment;
export const MountServiceAttachment: typeof import("./mountServiceAttachment").MountServiceAttachment = null as any;
utilities.lazyLoad(exports, ["MountServiceAttachment"], () => require("./mountServiceAttachment"));

export { MountServicesArgs, MountServicesResult, MountServicesOutputArgs } from "./mountServices";
export const mountServices: typeof import("./mountServices").mountServices = null as any;
export const mountServicesOutput: typeof import("./mountServices").mountServicesOutput = null as any;
utilities.lazyLoad(exports, ["mountServices","mountServicesOutput"], () => require("./mountServices"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "volcengine:vepfs/fileSystem:FileSystem":
                return new FileSystem(name, <any>undefined, { urn })
            case "volcengine:vepfs/fileset:Fileset":
                return new Fileset(name, <any>undefined, { urn })
            case "volcengine:vepfs/mountService:MountService":
                return new MountService(name, <any>undefined, { urn })
            case "volcengine:vepfs/mountServiceAttachment:MountServiceAttachment":
                return new MountServiceAttachment(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("volcengine", "vepfs/fileSystem", _module)
pulumi.runtime.registerResourceModule("volcengine", "vepfs/fileset", _module)
pulumi.runtime.registerResourceModule("volcengine", "vepfs/mountService", _module)
pulumi.runtime.registerResourceModule("volcengine", "vepfs/mountServiceAttachment", _module)