// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Vefaas.Outputs
{

    [OutputType]
    public sealed class FunctionsItemNasStorageResult
    {
        /// <summary>
        /// Whether to enable NAS storage mounting.
        /// </summary>
        public readonly bool EnableNas;
        /// <summary>
        /// The configuration of NAS.
        /// </summary>
        public readonly ImmutableArray<Outputs.FunctionsItemNasStorageNasConfigResult> NasConfigs;

        [OutputConstructor]
        private FunctionsItemNasStorageResult(
            bool enableNas,

            ImmutableArray<Outputs.FunctionsItemNasStorageNasConfigResult> nasConfigs)
        {
            EnableNas = enableNas;
            NasConfigs = nasConfigs;
        }
    }
}
