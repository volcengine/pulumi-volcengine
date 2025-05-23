// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Cr.Outputs
{

    [OutputType]
    public sealed class GetRegistriesRegistryStatusResult
    {
        /// <summary>
        /// The condition of registry.
        /// </summary>
        public readonly ImmutableArray<string> Conditions;
        /// <summary>
        /// The phase status of registry.
        /// </summary>
        public readonly string Phase;

        [OutputConstructor]
        private GetRegistriesRegistryStatusResult(
            ImmutableArray<string> conditions,

            string phase)
        {
            Conditions = conditions;
            Phase = phase;
        }
    }
}
