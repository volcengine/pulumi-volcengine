// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

declare var exports: any;
const __config = new pulumi.Config("volcengine");

/**
 * The Access Key for Volcengine Provider
 */
export declare const accessKey: string | undefined;
Object.defineProperty(exports, "accessKey", {
    get() {
        return __config.get("accessKey") ?? utilities.getEnv("VOLCENGINE_ACCESS_KEY");
    },
    enumerable: true,
});

export declare const assumeRole: outputs.config.AssumeRole | undefined;
Object.defineProperty(exports, "assumeRole", {
    get() {
        return __config.getObject<outputs.config.AssumeRole>("assumeRole");
    },
    enumerable: true,
});

/**
 * CUSTOMER ENDPOINT SUFFIX for Volcengine Provider
 */
export declare const customerEndpointSuffix: string | undefined;
Object.defineProperty(exports, "customerEndpointSuffix", {
    get() {
        return __config.get("customerEndpointSuffix");
    },
    enumerable: true,
});

/**
 * CUSTOMER ENDPOINTS for Volcengine Provider
 */
export declare const customerEndpoints: string | undefined;
Object.defineProperty(exports, "customerEndpoints", {
    get() {
        return __config.get("customerEndpoints");
    },
    enumerable: true,
});

/**
 * CUSTOMER HEADERS for Volcengine Provider
 */
export declare const customerHeaders: string | undefined;
Object.defineProperty(exports, "customerHeaders", {
    get() {
        return __config.get("customerHeaders");
    },
    enumerable: true,
});

/**
 * Disable SSL for Volcengine Provider
 */
export declare const disableSsl: boolean | undefined;
Object.defineProperty(exports, "disableSsl", {
    get() {
        return __config.getObject<boolean>("disableSsl");
    },
    enumerable: true,
});

/**
 * The Customer Endpoint for Volcengine Provider
 */
export declare const endpoint: string | undefined;
Object.defineProperty(exports, "endpoint", {
    get() {
        return __config.get("endpoint") ?? utilities.getEnv("VOLCENGINE_ENDPOINT");
    },
    enumerable: true,
});

/**
 * PROXY URL for Volcengine Provider
 */
export declare const proxyUrl: string | undefined;
Object.defineProperty(exports, "proxyUrl", {
    get() {
        return __config.get("proxyUrl");
    },
    enumerable: true,
});

/**
 * The Region for Volcengine Provider
 */
export declare const region: string | undefined;
Object.defineProperty(exports, "region", {
    get() {
        return __config.get("region") ?? utilities.getEnv("VOLCENGINE_REGION");
    },
    enumerable: true,
});

/**
 * The Secret Key for Volcengine Provider
 */
export declare const secretKey: string | undefined;
Object.defineProperty(exports, "secretKey", {
    get() {
        return __config.get("secretKey") ?? utilities.getEnv("VOLCENGINE_SECRET_KEY");
    },
    enumerable: true,
});

/**
 * The Session Token for Volcengine Provider
 */
export declare const sessionToken: string | undefined;
Object.defineProperty(exports, "sessionToken", {
    get() {
        return __config.get("sessionToken");
    },
    enumerable: true,
});

