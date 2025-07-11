// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Apig.Outputs
{

    [OutputType]
    public sealed class GetRoutesRouteMatchRuleQueryStringResult
    {
        /// <summary>
        /// The key of the query string.
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// The path of the api gateway route.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRoutesRouteMatchRuleQueryStringValueResult> Values;

        [OutputConstructor]
        private GetRoutesRouteMatchRuleQueryStringResult(
            string key,

            ImmutableArray<Outputs.GetRoutesRouteMatchRuleQueryStringValueResult> values)
        {
            Key = key;
            Values = values;
        }
    }
}
