// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.Pulumi.Volcengine.Direct_connect.Outputs
{

    [OutputType]
    public sealed class GatewayTag
    {
        /// <summary>
        /// The tag key.
        /// </summary>
        public readonly string? Key;
        /// <summary>
        /// The tag value.
        /// </summary>
        public readonly string? Value;

        [OutputConstructor]
        private GatewayTag(
            string? key,

            string? value)
        {
            Key = key;
            Value = value;
        }
    }
}