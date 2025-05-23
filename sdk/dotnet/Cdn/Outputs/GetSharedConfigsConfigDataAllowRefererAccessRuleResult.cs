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
    public sealed class GetSharedConfigsConfigDataAllowRefererAccessRuleResult
    {
        /// <summary>
        /// Indicates whether an empty Referer header, or a request without a Referer header, is not allowed. Default is false.
        /// </summary>
        public readonly bool AllowEmpty;
        /// <summary>
        /// The content indicating the Referer blacklist.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSharedConfigsConfigDataAllowRefererAccessRuleCommonTypeResult> CommonTypes;

        [OutputConstructor]
        private GetSharedConfigsConfigDataAllowRefererAccessRuleResult(
            bool allowEmpty,

            ImmutableArray<Outputs.GetSharedConfigsConfigDataAllowRefererAccessRuleCommonTypeResult> commonTypes)
        {
            AllowEmpty = allowEmpty;
            CommonTypes = commonTypes;
        }
    }
}
