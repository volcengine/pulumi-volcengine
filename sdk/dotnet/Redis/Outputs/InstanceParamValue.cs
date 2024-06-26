// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Redis.Outputs
{

    [OutputType]
    public sealed class InstanceParamValue
    {
        /// <summary>
        /// The name of configuration parameter.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The value of configuration parameter.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private InstanceParamValue(
            string name,

            string value)
        {
            Name = name;
            Value = value;
        }
    }
}
