// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Vke.Outputs
{

    [OutputType]
    public sealed class SupportAddonsAddonVersionCompatibilityResult
    {
        /// <summary>
        /// The Kubernetes Version of addon.
        /// </summary>
        public readonly string KubernetesVersion;

        [OutputConstructor]
        private SupportAddonsAddonVersionCompatibilityResult(string kubernetesVersion)
        {
            KubernetesVersion = kubernetesVersion;
        }
    }
}