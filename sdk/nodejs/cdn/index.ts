// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { CdnCertificateArgs, CdnCertificateState } from "./cdnCertificate";
export type CdnCertificate = import("./cdnCertificate").CdnCertificate;
export const CdnCertificate: typeof import("./cdnCertificate").CdnCertificate = null as any;
utilities.lazyLoad(exports, ["CdnCertificate"], () => require("./cdnCertificate"));

export { CdnDomainArgs, CdnDomainState } from "./cdnDomain";
export type CdnDomain = import("./cdnDomain").CdnDomain;
export const CdnDomain: typeof import("./cdnDomain").CdnDomain = null as any;
utilities.lazyLoad(exports, ["CdnDomain"], () => require("./cdnDomain"));

export { CertificatesArgs, CertificatesResult, CertificatesOutputArgs } from "./certificates";
export const certificates: typeof import("./certificates").certificates = null as any;
export const certificatesOutput: typeof import("./certificates").certificatesOutput = null as any;
utilities.lazyLoad(exports, ["certificates","certificatesOutput"], () => require("./certificates"));

export { ConfigsArgs, ConfigsResult, ConfigsOutputArgs } from "./configs";
export const configs: typeof import("./configs").configs = null as any;
export const configsOutput: typeof import("./configs").configsOutput = null as any;
utilities.lazyLoad(exports, ["configs","configsOutput"], () => require("./configs"));

export { DomainsArgs, DomainsResult, DomainsOutputArgs } from "./domains";
export const domains: typeof import("./domains").domains = null as any;
export const domainsOutput: typeof import("./domains").domainsOutput = null as any;
utilities.lazyLoad(exports, ["domains","domainsOutput"], () => require("./domains"));

export { GetCertificatesArgs, GetCertificatesResult, GetCertificatesOutputArgs } from "./getCertificates";
export const getCertificates: typeof import("./getCertificates").getCertificates = null as any;
export const getCertificatesOutput: typeof import("./getCertificates").getCertificatesOutput = null as any;
utilities.lazyLoad(exports, ["getCertificates","getCertificatesOutput"], () => require("./getCertificates"));

export { GetConfigsArgs, GetConfigsResult, GetConfigsOutputArgs } from "./getConfigs";
export const getConfigs: typeof import("./getConfigs").getConfigs = null as any;
export const getConfigsOutput: typeof import("./getConfigs").getConfigsOutput = null as any;
utilities.lazyLoad(exports, ["getConfigs","getConfigsOutput"], () => require("./getConfigs"));

export { GetDomainsArgs, GetDomainsResult, GetDomainsOutputArgs } from "./getDomains";
export const getDomains: typeof import("./getDomains").getDomains = null as any;
export const getDomainsOutput: typeof import("./getDomains").getDomainsOutput = null as any;
utilities.lazyLoad(exports, ["getDomains","getDomainsOutput"], () => require("./getDomains"));

export { GetSharedConfigsArgs, GetSharedConfigsResult, GetSharedConfigsOutputArgs } from "./getSharedConfigs";
export const getSharedConfigs: typeof import("./getSharedConfigs").getSharedConfigs = null as any;
export const getSharedConfigsOutput: typeof import("./getSharedConfigs").getSharedConfigsOutput = null as any;
utilities.lazyLoad(exports, ["getSharedConfigs","getSharedConfigsOutput"], () => require("./getSharedConfigs"));

export { SharedConfigArgs, SharedConfigState } from "./sharedConfig";
export type SharedConfig = import("./sharedConfig").SharedConfig;
export const SharedConfig: typeof import("./sharedConfig").SharedConfig = null as any;
utilities.lazyLoad(exports, ["SharedConfig"], () => require("./sharedConfig"));

export { SharedConfigsArgs, SharedConfigsResult, SharedConfigsOutputArgs } from "./sharedConfigs";
export const sharedConfigs: typeof import("./sharedConfigs").sharedConfigs = null as any;
export const sharedConfigsOutput: typeof import("./sharedConfigs").sharedConfigsOutput = null as any;
utilities.lazyLoad(exports, ["sharedConfigs","sharedConfigsOutput"], () => require("./sharedConfigs"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "volcengine:cdn/cdnCertificate:CdnCertificate":
                return new CdnCertificate(name, <any>undefined, { urn })
            case "volcengine:cdn/cdnDomain:CdnDomain":
                return new CdnDomain(name, <any>undefined, { urn })
            case "volcengine:cdn/sharedConfig:SharedConfig":
                return new SharedConfig(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("volcengine", "cdn/cdnCertificate", _module)
pulumi.runtime.registerResourceModule("volcengine", "cdn/cdnDomain", _module)
pulumi.runtime.registerResourceModule("volcengine", "cdn/sharedConfig", _module)
