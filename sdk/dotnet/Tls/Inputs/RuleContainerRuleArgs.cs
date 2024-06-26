// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Tls.Inputs
{

    public sealed class RuleContainerRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the container to be collected.
        /// </summary>
        [Input("containerNameRegex")]
        public Input<string>? ContainerNameRegex { get; set; }

        [Input("envTag")]
        private InputMap<string>? _envTag;

        /// <summary>
        /// Whether to add environment variables as log tags to raw log data.
        /// </summary>
        public InputMap<string> EnvTag
        {
            get => _envTag ?? (_envTag = new InputMap<string>());
            set => _envTag = value;
        }

        [Input("excludeContainerEnvRegex")]
        private InputMap<string>? _excludeContainerEnvRegex;

        /// <summary>
        /// The container environment variable blacklist is used to specify the range of containers not to be collected.
        /// </summary>
        public InputMap<string> ExcludeContainerEnvRegex
        {
            get => _excludeContainerEnvRegex ?? (_excludeContainerEnvRegex = new InputMap<string>());
            set => _excludeContainerEnvRegex = value;
        }

        [Input("excludeContainerLabelRegex")]
        private InputMap<string>? _excludeContainerLabelRegex;

        /// <summary>
        /// The container Label blacklist is used to specify the range of containers not to be collected.
        /// </summary>
        public InputMap<string> ExcludeContainerLabelRegex
        {
            get => _excludeContainerLabelRegex ?? (_excludeContainerLabelRegex = new InputMap<string>());
            set => _excludeContainerLabelRegex = value;
        }

        [Input("includeContainerEnvRegex")]
        private InputMap<string>? _includeContainerEnvRegex;

        /// <summary>
        /// The container environment variable whitelist specifies the container to be collected through the container environment variable. If the whitelist is not enabled, it means that all containers are specified to be collected.
        /// </summary>
        public InputMap<string> IncludeContainerEnvRegex
        {
            get => _includeContainerEnvRegex ?? (_includeContainerEnvRegex = new InputMap<string>());
            set => _includeContainerEnvRegex = value;
        }

        [Input("includeContainerLabelRegex")]
        private InputMap<string>? _includeContainerLabelRegex;

        /// <summary>
        /// The container label whitelist specifies the containers to be collected through the container label. If the whitelist is not enabled, all containers are specified to be collected.
        /// </summary>
        public InputMap<string> IncludeContainerLabelRegex
        {
            get => _includeContainerLabelRegex ?? (_includeContainerLabelRegex = new InputMap<string>());
            set => _includeContainerLabelRegex = value;
        }

        /// <summary>
        /// Collection rules for Kubernetes containers.
        /// </summary>
        [Input("kubernetesRule")]
        public Input<Inputs.RuleContainerRuleKubernetesRuleArgs>? KubernetesRule { get; set; }

        /// <summary>
        /// The collection mode.
        /// </summary>
        [Input("stream", required: true)]
        public Input<string> Stream { get; set; } = null!;

        public RuleContainerRuleArgs()
        {
        }
        public static new RuleContainerRuleArgs Empty => new RuleContainerRuleArgs();
    }
}
