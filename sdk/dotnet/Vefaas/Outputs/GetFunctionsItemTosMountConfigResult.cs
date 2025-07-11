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
    public sealed class GetFunctionsItemTosMountConfigResult
    {
        /// <summary>
        /// After enabling TOS, you need to provide an AKSK with access rights to the TOS domain name.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFunctionsItemTosMountConfigCredentialResult> Credentials;
        /// <summary>
        /// Whether to enable TOS storage mounting.
        /// </summary>
        public readonly bool EnableTos;
        /// <summary>
        /// After enabling TOS, you need to provide a TOS storage configuration list, with a maximum of 5 items.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFunctionsItemTosMountConfigMountPointResult> MountPoints;

        [OutputConstructor]
        private GetFunctionsItemTosMountConfigResult(
            ImmutableArray<Outputs.GetFunctionsItemTosMountConfigCredentialResult> credentials,

            bool enableTos,

            ImmutableArray<Outputs.GetFunctionsItemTosMountConfigMountPointResult> mountPoints)
        {
            Credentials = credentials;
            EnableTos = enableTos;
            MountPoints = mountPoints;
        }
    }
}
