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
    public sealed class GetSharedConfigsConfigDataDenyRefererAccessRuleCommonTypeResult
    {
        /// <summary>
        /// This list is case-sensitive when matching requests. Default is true.
        /// </summary>
        public readonly bool IgnoreCase;
        /// <summary>
        /// The entries in this list are an array of IP addresses and CIDR network segments. The total number of entries cannot exceed 3,000. The IP addresses and segments can be in IPv4 and IPv6 format. Duplicate entries in the list will be removed and will not count towards the limit.
        /// </summary>
        public readonly ImmutableArray<string> Rules;

        [OutputConstructor]
        private GetSharedConfigsConfigDataDenyRefererAccessRuleCommonTypeResult(
            bool ignoreCase,

            ImmutableArray<string> rules)
        {
            IgnoreCase = ignoreCase;
            Rules = rules;
        }
    }
}
