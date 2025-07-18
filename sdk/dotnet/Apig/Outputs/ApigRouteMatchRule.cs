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
    public sealed class ApigRouteMatchRule
    {
        /// <summary>
        /// The header of the api gateway route.
        /// </summary>
        public readonly ImmutableArray<Outputs.ApigRouteMatchRuleHeader> Headers;
        /// <summary>
        /// The method of the api gateway route. Valid values: `GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONS`, `CONNECT`.
        /// </summary>
        public readonly ImmutableArray<string> Methods;
        /// <summary>
        /// The path of the api gateway route.
        /// </summary>
        public readonly Outputs.ApigRouteMatchRulePath Path;
        /// <summary>
        /// The query string of the api gateway route.
        /// </summary>
        public readonly ImmutableArray<Outputs.ApigRouteMatchRuleQueryString> QueryStrings;

        [OutputConstructor]
        private ApigRouteMatchRule(
            ImmutableArray<Outputs.ApigRouteMatchRuleHeader> headers,

            ImmutableArray<string> methods,

            Outputs.ApigRouteMatchRulePath path,

            ImmutableArray<Outputs.ApigRouteMatchRuleQueryString> queryStrings)
        {
            Headers = headers;
            Methods = methods;
            Path = path;
            QueryStrings = queryStrings;
        }
    }
}
