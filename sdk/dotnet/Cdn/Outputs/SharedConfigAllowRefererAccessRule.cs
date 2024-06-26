// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Cdn.Outputs
{

    [OutputType]
    public sealed class SharedConfigAllowRefererAccessRule
    {
        /// <summary>
        /// Indicates whether an empty Referer header, or a request without a Referer header, is not allowed. Default is false.
        /// </summary>
        public readonly bool? AllowEmpty;
        /// <summary>
        /// The content indicating the Referer whitelist.
        /// </summary>
        public readonly Outputs.SharedConfigAllowRefererAccessRuleCommonType CommonType;

        [OutputConstructor]
        private SharedConfigAllowRefererAccessRule(
            bool? allowEmpty,

            Outputs.SharedConfigAllowRefererAccessRuleCommonType commonType)
        {
            AllowEmpty = allowEmpty;
            CommonType = commonType;
        }
    }
}
