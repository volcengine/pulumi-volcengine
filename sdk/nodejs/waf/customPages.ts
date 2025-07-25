// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of waf custom pages
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const foo = volcengine.waf.getCustomPages({
 *     host: "www.tf-test.com",
 * });
 * ```
 */
/** @deprecated volcengine.waf.CustomPages has been deprecated in favor of volcengine.waf.getCustomPages */
export function customPages(args: CustomPagesArgs, opts?: pulumi.InvokeOptions): Promise<CustomPagesResult> {
    pulumi.log.warn("customPages is deprecated: volcengine.waf.CustomPages has been deprecated in favor of volcengine.waf.getCustomPages")

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:waf/customPages:CustomPages", {
        "host": args.host,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "projectName": args.projectName,
        "ruleTag": args.ruleTag,
    }, opts);
}

/**
 * A collection of arguments for invoking CustomPages.
 */
export interface CustomPagesArgs {
    /**
     * The domain names that need to be viewed.
     */
    host: string;
    /**
     * A Name Regex of Resource.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
    /**
     * The name of the project to which your domain names belong.
     */
    projectName?: string;
    /**
     * Unique identification of the rules.
     */
    ruleTag?: string;
}

/**
 * A collection of values returned by CustomPages.
 */
export interface CustomPagesResult {
    /**
     * Details of the rules.
     */
    readonly datas: outputs.waf.CustomPagesData[];
    /**
     * Domain name to be protected.
     */
    readonly host: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly nameRegex?: string;
    readonly outputFile?: string;
    readonly projectName?: string;
    /**
     * Unique identification of the rules.
     */
    readonly ruleTag?: string;
    /**
     * The total count of query.
     */
    readonly totalCount: number;
}
/**
 * Use this data source to query detailed information of waf custom pages
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const foo = volcengine.waf.getCustomPages({
 *     host: "www.tf-test.com",
 * });
 * ```
 */
/** @deprecated volcengine.waf.CustomPages has been deprecated in favor of volcengine.waf.getCustomPages */
export function customPagesOutput(args: CustomPagesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<CustomPagesResult> {
    return pulumi.output(args).apply((a: any) => customPages(a, opts))
}

/**
 * A collection of arguments for invoking CustomPages.
 */
export interface CustomPagesOutputArgs {
    /**
     * The domain names that need to be viewed.
     */
    host: pulumi.Input<string>;
    /**
     * A Name Regex of Resource.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The name of the project to which your domain names belong.
     */
    projectName?: pulumi.Input<string>;
    /**
     * Unique identification of the rules.
     */
    ruleTag?: pulumi.Input<string>;
}
