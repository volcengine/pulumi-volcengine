// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.PulumiPackage.Volcengine.Cen
{
    /// <summary>
    /// Provides a resource to manage cen
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
    ///         var foo = new Volcengine.Cen.Cen("foo", new Volcengine.Cen.CenArgs
    ///         {
    ///             CenName = "tf-test-3",
    ///             Description = "tf-test",
    ///             ProjectName = "default",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Cen can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import volcengine:cen/cen:Cen default cen-7qthudw0ll6jmc****
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:cen/cen:Cen")]
    public partial class Cen : Pulumi.CustomResource
    {
        /// <summary>
        /// The account ID of the cen.
        /// </summary>
        [Output("accountId")]
        public Output<string> AccountId { get; private set; } = null!;

        /// <summary>
        /// A list of bandwidth package IDs of the cen.
        /// </summary>
        [Output("cenBandwidthPackageIds")]
        public Output<ImmutableArray<string>> CenBandwidthPackageIds { get; private set; } = null!;

        /// <summary>
        /// The ID of the cen.
        /// </summary>
        [Output("cenId")]
        public Output<string> CenId { get; private set; } = null!;

        /// <summary>
        /// The name of the cen.
        /// </summary>
        [Output("cenName")]
        public Output<string> CenName { get; private set; } = null!;

        /// <summary>
        /// The create time of the cen.
        /// </summary>
        [Output("creationTime")]
        public Output<string> CreationTime { get; private set; } = null!;

        /// <summary>
        /// The description of the cen.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// The ProjectName of the cen instance.
        /// </summary>
        [Output("projectName")]
        public Output<string?> ProjectName { get; private set; } = null!;

        /// <summary>
        /// The status of the cen.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.CenTag>> Tags { get; private set; } = null!;

        /// <summary>
        /// The update time of the cen.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a Cen resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Cen(string name, CenArgs? args = null, CustomResourceOptions? options = null)
            : base("volcengine:cen/cen:Cen", name, args ?? new CenArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Cen(string name, Input<string> id, CenState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:cen/cen:Cen", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Cen resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Cen Get(string name, Input<string> id, CenState? state = null, CustomResourceOptions? options = null)
        {
            return new Cen(name, id, state, options);
        }
    }

    public sealed class CenArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the cen.
        /// </summary>
        [Input("cenName")]
        public Input<string>? CenName { get; set; }

        /// <summary>
        /// The description of the cen.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The ProjectName of the cen instance.
        /// </summary>
        [Input("projectName")]
        public Input<string>? ProjectName { get; set; }

        [Input("tags")]
        private InputList<Inputs.CenTagArgs>? _tags;

        /// <summary>
        /// Tags.
        /// </summary>
        public InputList<Inputs.CenTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.CenTagArgs>());
            set => _tags = value;
        }

        public CenArgs()
        {
        }
    }

    public sealed class CenState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The account ID of the cen.
        /// </summary>
        [Input("accountId")]
        public Input<string>? AccountId { get; set; }

        [Input("cenBandwidthPackageIds")]
        private InputList<string>? _cenBandwidthPackageIds;

        /// <summary>
        /// A list of bandwidth package IDs of the cen.
        /// </summary>
        public InputList<string> CenBandwidthPackageIds
        {
            get => _cenBandwidthPackageIds ?? (_cenBandwidthPackageIds = new InputList<string>());
            set => _cenBandwidthPackageIds = value;
        }

        /// <summary>
        /// The ID of the cen.
        /// </summary>
        [Input("cenId")]
        public Input<string>? CenId { get; set; }

        /// <summary>
        /// The name of the cen.
        /// </summary>
        [Input("cenName")]
        public Input<string>? CenName { get; set; }

        /// <summary>
        /// The create time of the cen.
        /// </summary>
        [Input("creationTime")]
        public Input<string>? CreationTime { get; set; }

        /// <summary>
        /// The description of the cen.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The ProjectName of the cen instance.
        /// </summary>
        [Input("projectName")]
        public Input<string>? ProjectName { get; set; }

        /// <summary>
        /// The status of the cen.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputList<Inputs.CenTagGetArgs>? _tags;

        /// <summary>
        /// Tags.
        /// </summary>
        public InputList<Inputs.CenTagGetArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.CenTagGetArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// The update time of the cen.
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        public CenState()
        {
        }
    }
}