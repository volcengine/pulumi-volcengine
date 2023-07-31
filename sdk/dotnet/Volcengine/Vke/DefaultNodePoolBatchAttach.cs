// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.PulumiPackage.Volcengine.Vke
{
    /// <summary>
    /// Provides a resource to manage vke default node pool batch attach
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Volcengine = Volcengine.PulumiPackage.Volcengine;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var @default = new Volcengine.Vke.DefaultNodePoolBatchAttach("default", new Volcengine.Vke.DefaultNodePoolBatchAttachArgs
    ///         {
    ///             ClusterId = "ccc2umdnqtoflv91lqtq0",
    ///             DefaultNodePoolId = "11111",
    ///             Instances = 
    ///             {
    ///                 new Volcengine.Vke.Inputs.DefaultNodePoolBatchAttachInstanceArgs
    ///                 {
    ///                     AdditionalContainerStorageEnabled = false,
    ///                     InstanceId = "i-ybvza90ohwexzk8emaa3",
    ///                     KeepInstanceName = false,
    ///                 },
    ///                 new Volcengine.Vke.Inputs.DefaultNodePoolBatchAttachInstanceArgs
    ///                 {
    ///                     AdditionalContainerStorageEnabled = true,
    ///                     ContainerStoragePath = "/",
    ///                     InstanceId = "i-ybvza90ohxexzkm4zihf",
    ///                     KeepInstanceName = false,
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:vke/defaultNodePoolBatchAttach:DefaultNodePoolBatchAttach")]
    public partial class DefaultNodePoolBatchAttach : Pulumi.CustomResource
    {
        /// <summary>
        /// The ClusterId of NodePool.
        /// </summary>
        [Output("clusterId")]
        public Output<string> ClusterId { get; private set; } = null!;

        /// <summary>
        /// The default NodePool ID.
        /// </summary>
        [Output("defaultNodePoolId")]
        public Output<string> DefaultNodePoolId { get; private set; } = null!;

        /// <summary>
        /// The ECS InstanceIds add to NodePool.
        /// </summary>
        [Output("instances")]
        public Output<ImmutableArray<Outputs.DefaultNodePoolBatchAttachInstance>> Instances { get; private set; } = null!;

        /// <summary>
        /// Is import of the DefaultNodePool. It only works when imported, set to true.
        /// </summary>
        [Output("isImport")]
        public Output<bool> IsImport { get; private set; } = null!;

        /// <summary>
        /// The KubernetesConfig of NodeConfig.
        /// </summary>
        [Output("kubernetesConfigs")]
        public Output<ImmutableArray<Outputs.DefaultNodePoolBatchAttachKubernetesConfig>> KubernetesConfigs { get; private set; } = null!;

        /// <summary>
        /// The Config of NodePool.
        /// </summary>
        [Output("nodeConfigs")]
        public Output<ImmutableArray<Outputs.DefaultNodePoolBatchAttachNodeConfig>> NodeConfigs { get; private set; } = null!;

        /// <summary>
        /// Tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.DefaultNodePoolBatchAttachTag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a DefaultNodePoolBatchAttach resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DefaultNodePoolBatchAttach(string name, DefaultNodePoolBatchAttachArgs args, CustomResourceOptions? options = null)
            : base("volcengine:vke/defaultNodePoolBatchAttach:DefaultNodePoolBatchAttach", name, args ?? new DefaultNodePoolBatchAttachArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DefaultNodePoolBatchAttach(string name, Input<string> id, DefaultNodePoolBatchAttachState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:vke/defaultNodePoolBatchAttach:DefaultNodePoolBatchAttach", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DefaultNodePoolBatchAttach resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DefaultNodePoolBatchAttach Get(string name, Input<string> id, DefaultNodePoolBatchAttachState? state = null, CustomResourceOptions? options = null)
        {
            return new DefaultNodePoolBatchAttach(name, id, state, options);
        }
    }

    public sealed class DefaultNodePoolBatchAttachArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ClusterId of NodePool.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// The default NodePool ID.
        /// </summary>
        [Input("defaultNodePoolId", required: true)]
        public Input<string> DefaultNodePoolId { get; set; } = null!;

        [Input("instances")]
        private InputList<Inputs.DefaultNodePoolBatchAttachInstanceArgs>? _instances;

        /// <summary>
        /// The ECS InstanceIds add to NodePool.
        /// </summary>
        public InputList<Inputs.DefaultNodePoolBatchAttachInstanceArgs> Instances
        {
            get => _instances ?? (_instances = new InputList<Inputs.DefaultNodePoolBatchAttachInstanceArgs>());
            set => _instances = value;
        }

        public DefaultNodePoolBatchAttachArgs()
        {
        }
    }

    public sealed class DefaultNodePoolBatchAttachState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ClusterId of NodePool.
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        /// <summary>
        /// The default NodePool ID.
        /// </summary>
        [Input("defaultNodePoolId")]
        public Input<string>? DefaultNodePoolId { get; set; }

        [Input("instances")]
        private InputList<Inputs.DefaultNodePoolBatchAttachInstanceGetArgs>? _instances;

        /// <summary>
        /// The ECS InstanceIds add to NodePool.
        /// </summary>
        public InputList<Inputs.DefaultNodePoolBatchAttachInstanceGetArgs> Instances
        {
            get => _instances ?? (_instances = new InputList<Inputs.DefaultNodePoolBatchAttachInstanceGetArgs>());
            set => _instances = value;
        }

        /// <summary>
        /// Is import of the DefaultNodePool. It only works when imported, set to true.
        /// </summary>
        [Input("isImport")]
        public Input<bool>? IsImport { get; set; }

        [Input("kubernetesConfigs")]
        private InputList<Inputs.DefaultNodePoolBatchAttachKubernetesConfigGetArgs>? _kubernetesConfigs;

        /// <summary>
        /// The KubernetesConfig of NodeConfig.
        /// </summary>
        public InputList<Inputs.DefaultNodePoolBatchAttachKubernetesConfigGetArgs> KubernetesConfigs
        {
            get => _kubernetesConfigs ?? (_kubernetesConfigs = new InputList<Inputs.DefaultNodePoolBatchAttachKubernetesConfigGetArgs>());
            set => _kubernetesConfigs = value;
        }

        [Input("nodeConfigs")]
        private InputList<Inputs.DefaultNodePoolBatchAttachNodeConfigGetArgs>? _nodeConfigs;

        /// <summary>
        /// The Config of NodePool.
        /// </summary>
        public InputList<Inputs.DefaultNodePoolBatchAttachNodeConfigGetArgs> NodeConfigs
        {
            get => _nodeConfigs ?? (_nodeConfigs = new InputList<Inputs.DefaultNodePoolBatchAttachNodeConfigGetArgs>());
            set => _nodeConfigs = value;
        }

        [Input("tags")]
        private InputList<Inputs.DefaultNodePoolBatchAttachTagGetArgs>? _tags;

        /// <summary>
        /// Tags.
        /// </summary>
        public InputList<Inputs.DefaultNodePoolBatchAttachTagGetArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.DefaultNodePoolBatchAttachTagGetArgs>());
            set => _tags = value;
        }

        public DefaultNodePoolBatchAttachState()
        {
        }
    }
}