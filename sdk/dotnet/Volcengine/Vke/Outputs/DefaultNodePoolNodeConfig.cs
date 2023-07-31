// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.PulumiPackage.Volcengine.Vke.Outputs
{

    [OutputType]
    public sealed class DefaultNodePoolNodeConfig
    {
        /// <summary>
        /// Tags for Ecs.
        /// </summary>
        public readonly ImmutableArray<Outputs.DefaultNodePoolNodeConfigEcsTag> EcsTags;
        /// <summary>
        /// The initializeScript of NodeConfig.
        /// </summary>
        public readonly string? InitializeScript;
        /// <summary>
        /// The NamePrefix of NodeConfig.
        /// </summary>
        public readonly string? NamePrefix;
        /// <summary>
        /// The Security of NodeConfig.
        /// </summary>
        public readonly Outputs.DefaultNodePoolNodeConfigSecurity Security;

        [OutputConstructor]
        private DefaultNodePoolNodeConfig(
            ImmutableArray<Outputs.DefaultNodePoolNodeConfigEcsTag> ecsTags,

            string? initializeScript,

            string? namePrefix,

            Outputs.DefaultNodePoolNodeConfigSecurity security)
        {
            EcsTags = ecsTags;
            InitializeScript = initializeScript;
            NamePrefix = namePrefix;
            Security = security;
        }
    }
}