// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Transit_router
{
    /// <summary>
    /// Provides a resource to manage transit router bandwidth package
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Volcengine = Pulumi.Volcengine;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var foo = new Volcengine.Transit_router.BandwidthPackage("foo", new()
    ///     {
    ///         Bandwidth = 2,
    ///         Description = "acc-test",
    ///         Period = 1,
    ///         ProjectName = "default",
    ///         RenewType = "Manual",
    ///         Tags = new[]
    ///         {
    ///             new Volcengine.Transit_router.Inputs.BandwidthPackageTagArgs
    ///             {
    ///                 Key = "k1",
    ///                 Value = "v1",
    ///             },
    ///         },
    ///         TransitRouterBandwidthPackageName = "acc-tf-test",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// TransitRouterBandwidthPackage can be imported using the Id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import volcengine:transit_router/bandwidthPackage:BandwidthPackage default tbp-cd-2felfww0i6pkw59gp68bq****
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:transit_router/bandwidthPackage:BandwidthPackage")]
    public partial class BandwidthPackage : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The detailed information on cross regional connections associated with bandwidth packets.
        /// </summary>
        [Output("allocations")]
        public Output<ImmutableArray<Outputs.BandwidthPackageAllocation>> Allocations { get; private set; } = null!;

        /// <summary>
        /// The bandwidth peak of the transit router bandwidth package. Unit: Mbps. Valid values: 2-10000. Default is 2 Mbps.
        /// </summary>
        [Output("bandwidth")]
        public Output<int?> Bandwidth { get; private set; } = null!;

        /// <summary>
        /// The business status of the transit router bandwidth package.
        /// </summary>
        [Output("businessStatus")]
        public Output<string> BusinessStatus { get; private set; } = null!;

        /// <summary>
        /// The create time of the transit router bandwidth package.
        /// </summary>
        [Output("creationTime")]
        public Output<string> CreationTime { get; private set; } = null!;

        /// <summary>
        /// The delete time of the transit router bandwidth package.
        /// </summary>
        [Output("deleteTime")]
        public Output<string> DeleteTime { get; private set; } = null!;

        /// <summary>
        /// The description of the transit router bandwidth package.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// The expired time of the transit router bandwidth package.
        /// </summary>
        [Output("expiredTime")]
        public Output<string> ExpiredTime { get; private set; } = null!;

        /// <summary>
        /// The period of the transit router bandwidth package, the valid value range in 1~9 or 12 or 36. Default value is 12. The period unit defaults to `Month`.The modification of this field only takes effect when the value of the `renew_type` is `Manual`.
        /// </summary>
        [Output("period")]
        public Output<int?> Period { get; private set; } = null!;

        /// <summary>
        /// The ProjectName of the transit router bandwidth package.
        /// </summary>
        [Output("projectName")]
        public Output<string> ProjectName { get; private set; } = null!;

        /// <summary>
        /// The remaining renewal times of of the transit router bandwidth package. Valid values: -1 or 1~100. Default value is -1, means unlimited renewal.This field is only effective when the value of the `renew_type` is `Auto`.
        /// </summary>
        [Output("remainRenewTimes")]
        public Output<int?> RemainRenewTimes { get; private set; } = null!;

        /// <summary>
        /// The remaining bandwidth of the transit router bandwidth package. Unit: Mbps.
        /// </summary>
        [Output("remainingBandwidth")]
        public Output<int> RemainingBandwidth { get; private set; } = null!;

        /// <summary>
        /// The auto renewal period of the transit router bandwidth package. Valid values: 1,2,3,6,12. Default value is 1. Unit: Month.This field is only effective when the value of the `renew_type` is `Auto`. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        /// </summary>
        [Output("renewPeriod")]
        public Output<int?> RenewPeriod { get; private set; } = null!;

        /// <summary>
        /// The renewal type of the transit router bandwidth package. Valid values: `Manual`, `Auto`, `NoRenew`. Default is `Manual`.This field is only effective when modifying the bandwidth package.
        /// </summary>
        [Output("renewType")]
        public Output<string?> RenewType { get; private set; } = null!;

        /// <summary>
        /// The status of the transit router bandwidth package.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.BandwidthPackageTag>> Tags { get; private set; } = null!;

        /// <summary>
        /// The name of the transit router bandwidth package.
        /// </summary>
        [Output("transitRouterBandwidthPackageName")]
        public Output<string> TransitRouterBandwidthPackageName { get; private set; } = null!;

        /// <summary>
        /// The update time of the transit router bandwidth package.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a BandwidthPackage resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BandwidthPackage(string name, BandwidthPackageArgs? args = null, CustomResourceOptions? options = null)
            : base("volcengine:transit_router/bandwidthPackage:BandwidthPackage", name, args ?? new BandwidthPackageArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BandwidthPackage(string name, Input<string> id, BandwidthPackageState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:transit_router/bandwidthPackage:BandwidthPackage", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/volcengine",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing BandwidthPackage resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BandwidthPackage Get(string name, Input<string> id, BandwidthPackageState? state = null, CustomResourceOptions? options = null)
        {
            return new BandwidthPackage(name, id, state, options);
        }
    }

    public sealed class BandwidthPackageArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The bandwidth peak of the transit router bandwidth package. Unit: Mbps. Valid values: 2-10000. Default is 2 Mbps.
        /// </summary>
        [Input("bandwidth")]
        public Input<int>? Bandwidth { get; set; }

        /// <summary>
        /// The description of the transit router bandwidth package.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The period of the transit router bandwidth package, the valid value range in 1~9 or 12 or 36. Default value is 12. The period unit defaults to `Month`.The modification of this field only takes effect when the value of the `renew_type` is `Manual`.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        /// <summary>
        /// The ProjectName of the transit router bandwidth package.
        /// </summary>
        [Input("projectName")]
        public Input<string>? ProjectName { get; set; }

        /// <summary>
        /// The remaining renewal times of of the transit router bandwidth package. Valid values: -1 or 1~100. Default value is -1, means unlimited renewal.This field is only effective when the value of the `renew_type` is `Auto`.
        /// </summary>
        [Input("remainRenewTimes")]
        public Input<int>? RemainRenewTimes { get; set; }

        /// <summary>
        /// The auto renewal period of the transit router bandwidth package. Valid values: 1,2,3,6,12. Default value is 1. Unit: Month.This field is only effective when the value of the `renew_type` is `Auto`. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        /// </summary>
        [Input("renewPeriod")]
        public Input<int>? RenewPeriod { get; set; }

        /// <summary>
        /// The renewal type of the transit router bandwidth package. Valid values: `Manual`, `Auto`, `NoRenew`. Default is `Manual`.This field is only effective when modifying the bandwidth package.
        /// </summary>
        [Input("renewType")]
        public Input<string>? RenewType { get; set; }

        [Input("tags")]
        private InputList<Inputs.BandwidthPackageTagArgs>? _tags;

        /// <summary>
        /// Tags.
        /// </summary>
        public InputList<Inputs.BandwidthPackageTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.BandwidthPackageTagArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// The name of the transit router bandwidth package.
        /// </summary>
        [Input("transitRouterBandwidthPackageName")]
        public Input<string>? TransitRouterBandwidthPackageName { get; set; }

        public BandwidthPackageArgs()
        {
        }
        public static new BandwidthPackageArgs Empty => new BandwidthPackageArgs();
    }

    public sealed class BandwidthPackageState : global::Pulumi.ResourceArgs
    {
        [Input("allocations")]
        private InputList<Inputs.BandwidthPackageAllocationGetArgs>? _allocations;

        /// <summary>
        /// The detailed information on cross regional connections associated with bandwidth packets.
        /// </summary>
        public InputList<Inputs.BandwidthPackageAllocationGetArgs> Allocations
        {
            get => _allocations ?? (_allocations = new InputList<Inputs.BandwidthPackageAllocationGetArgs>());
            set => _allocations = value;
        }

        /// <summary>
        /// The bandwidth peak of the transit router bandwidth package. Unit: Mbps. Valid values: 2-10000. Default is 2 Mbps.
        /// </summary>
        [Input("bandwidth")]
        public Input<int>? Bandwidth { get; set; }

        /// <summary>
        /// The business status of the transit router bandwidth package.
        /// </summary>
        [Input("businessStatus")]
        public Input<string>? BusinessStatus { get; set; }

        /// <summary>
        /// The create time of the transit router bandwidth package.
        /// </summary>
        [Input("creationTime")]
        public Input<string>? CreationTime { get; set; }

        /// <summary>
        /// The delete time of the transit router bandwidth package.
        /// </summary>
        [Input("deleteTime")]
        public Input<string>? DeleteTime { get; set; }

        /// <summary>
        /// The description of the transit router bandwidth package.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The expired time of the transit router bandwidth package.
        /// </summary>
        [Input("expiredTime")]
        public Input<string>? ExpiredTime { get; set; }

        /// <summary>
        /// The period of the transit router bandwidth package, the valid value range in 1~9 or 12 or 36. Default value is 12. The period unit defaults to `Month`.The modification of this field only takes effect when the value of the `renew_type` is `Manual`.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        /// <summary>
        /// The ProjectName of the transit router bandwidth package.
        /// </summary>
        [Input("projectName")]
        public Input<string>? ProjectName { get; set; }

        /// <summary>
        /// The remaining renewal times of of the transit router bandwidth package. Valid values: -1 or 1~100. Default value is -1, means unlimited renewal.This field is only effective when the value of the `renew_type` is `Auto`.
        /// </summary>
        [Input("remainRenewTimes")]
        public Input<int>? RemainRenewTimes { get; set; }

        /// <summary>
        /// The remaining bandwidth of the transit router bandwidth package. Unit: Mbps.
        /// </summary>
        [Input("remainingBandwidth")]
        public Input<int>? RemainingBandwidth { get; set; }

        /// <summary>
        /// The auto renewal period of the transit router bandwidth package. Valid values: 1,2,3,6,12. Default value is 1. Unit: Month.This field is only effective when the value of the `renew_type` is `Auto`. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        /// </summary>
        [Input("renewPeriod")]
        public Input<int>? RenewPeriod { get; set; }

        /// <summary>
        /// The renewal type of the transit router bandwidth package. Valid values: `Manual`, `Auto`, `NoRenew`. Default is `Manual`.This field is only effective when modifying the bandwidth package.
        /// </summary>
        [Input("renewType")]
        public Input<string>? RenewType { get; set; }

        /// <summary>
        /// The status of the transit router bandwidth package.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputList<Inputs.BandwidthPackageTagGetArgs>? _tags;

        /// <summary>
        /// Tags.
        /// </summary>
        public InputList<Inputs.BandwidthPackageTagGetArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.BandwidthPackageTagGetArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// The name of the transit router bandwidth package.
        /// </summary>
        [Input("transitRouterBandwidthPackageName")]
        public Input<string>? TransitRouterBandwidthPackageName { get; set; }

        /// <summary>
        /// The update time of the transit router bandwidth package.
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        public BandwidthPackageState()
        {
        }
        public static new BandwidthPackageState Empty => new BandwidthPackageState();
    }
}