// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Apig.Inputs
{

    public sealed class ApigRouteMatchRuleQueryStringGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The key of the query string.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        /// <summary>
        /// The path of the api gateway route.
        /// </summary>
        [Input("value", required: true)]
        public Input<Inputs.ApigRouteMatchRuleQueryStringValueGetArgs> Value { get; set; } = null!;

        public ApigRouteMatchRuleQueryStringGetArgs()
        {
        }
        public static new ApigRouteMatchRuleQueryStringGetArgs Empty => new ApigRouteMatchRuleQueryStringGetArgs();
    }
}
