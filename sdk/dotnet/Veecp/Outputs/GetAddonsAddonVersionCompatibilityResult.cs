// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Veecp.Outputs
{

    [OutputType]
    public sealed class GetAddonsAddonVersionCompatibilityResult
    {
        /// <summary>
        /// The Kubernetes Version of addon.
        /// </summary>
        public readonly string KubernetesVersion;

        [OutputConstructor]
        private GetAddonsAddonVersionCompatibilityResult(string kubernetesVersion)
        {
            KubernetesVersion = kubernetesVersion;
        }
    }
}
