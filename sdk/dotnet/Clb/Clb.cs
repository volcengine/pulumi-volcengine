// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Clb
{
    /// <summary>
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Volcengine = Pulumi.Volcengine;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var foo = new Volcengine.Clb.Clb("foo", new Volcengine.Clb.ClbArgs
    ///         {
    ///             Description = "Demo",
    ///             LoadBalancerName = "terraform-auto-create",
    ///             LoadBalancerSpec = "small_1",
    ///             ProjectName = "yyy",
    ///             SubnetId = "subnet-mj92ij84m5fk5smt1arvwrtw",
    ///             Type = "public",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// CLB can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import volcengine:clb/clb:Clb default clb-273y2ok6ets007fap8txvf6us
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:clb/clb:Clb")]
    public partial class Clb : Pulumi.CustomResource
    {
        /// <summary>
        /// The description of the CLB.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The eni address of the CLB.
        /// </summary>
        [Output("eniAddress")]
        public Output<string> EniAddress { get; private set; } = null!;

        /// <summary>
        /// The billing type of the CLB, the value can be `PostPaid`.
        /// </summary>
        [Output("loadBalancerBillingType")]
        public Output<string> LoadBalancerBillingType { get; private set; } = null!;

        /// <summary>
        /// The name of the CLB.
        /// </summary>
        [Output("loadBalancerName")]
        public Output<string> LoadBalancerName { get; private set; } = null!;

        /// <summary>
        /// The specification of the CLB, the value can be `small_1`, `small_2`, `medium_1`, `medium_2`, `large_1`, `large_2`.
        /// </summary>
        [Output("loadBalancerSpec")]
        public Output<string> LoadBalancerSpec { get; private set; } = null!;

        /// <summary>
        /// The master zone ID of the CLB.
        /// </summary>
        [Output("masterZoneId")]
        public Output<string> MasterZoneId { get; private set; } = null!;

        /// <summary>
        /// The reason of the console modification protection.
        /// </summary>
        [Output("modificationProtectionReason")]
        public Output<string?> ModificationProtectionReason { get; private set; } = null!;

        /// <summary>
        /// The status of the console modification protection, the value can be `NonProtection` or `ConsoleProtection`.
        /// </summary>
        [Output("modificationProtectionStatus")]
        public Output<string?> ModificationProtectionStatus { get; private set; } = null!;

        /// <summary>
        /// The ProjectName of the CLB.
        /// </summary>
        [Output("projectName")]
        public Output<string?> ProjectName { get; private set; } = null!;

        /// <summary>
        /// The region of the request.
        /// </summary>
        [Output("regionId")]
        public Output<string> RegionId { get; private set; } = null!;

        /// <summary>
        /// The slave zone ID of the CLB.
        /// </summary>
        [Output("slaveZoneId")]
        public Output<string> SlaveZoneId { get; private set; } = null!;

        /// <summary>
        /// The id of the Subnet.
        /// </summary>
        [Output("subnetId")]
        public Output<string> SubnetId { get; private set; } = null!;

        /// <summary>
        /// Tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.ClbTag>> Tags { get; private set; } = null!;

        /// <summary>
        /// The type of the CLB. And optional choice contains `public` or `private`.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The id of the VPC.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a Clb resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Clb(string name, ClbArgs args, CustomResourceOptions? options = null)
            : base("volcengine:clb/clb:Clb", name, args ?? new ClbArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Clb(string name, Input<string> id, ClbState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:clb/clb:Clb", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Clb resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Clb Get(string name, Input<string> id, ClbState? state = null, CustomResourceOptions? options = null)
        {
            return new Clb(name, id, state, options);
        }
    }

    public sealed class ClbArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the CLB.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The eni address of the CLB.
        /// </summary>
        [Input("eniAddress")]
        public Input<string>? EniAddress { get; set; }

        /// <summary>
        /// The billing type of the CLB, the value can be `PostPaid`.
        /// </summary>
        [Input("loadBalancerBillingType")]
        public Input<string>? LoadBalancerBillingType { get; set; }

        /// <summary>
        /// The name of the CLB.
        /// </summary>
        [Input("loadBalancerName")]
        public Input<string>? LoadBalancerName { get; set; }

        /// <summary>
        /// The specification of the CLB, the value can be `small_1`, `small_2`, `medium_1`, `medium_2`, `large_1`, `large_2`.
        /// </summary>
        [Input("loadBalancerSpec", required: true)]
        public Input<string> LoadBalancerSpec { get; set; } = null!;

        /// <summary>
        /// The master zone ID of the CLB.
        /// </summary>
        [Input("masterZoneId")]
        public Input<string>? MasterZoneId { get; set; }

        /// <summary>
        /// The reason of the console modification protection.
        /// </summary>
        [Input("modificationProtectionReason")]
        public Input<string>? ModificationProtectionReason { get; set; }

        /// <summary>
        /// The status of the console modification protection, the value can be `NonProtection` or `ConsoleProtection`.
        /// </summary>
        [Input("modificationProtectionStatus")]
        public Input<string>? ModificationProtectionStatus { get; set; }

        /// <summary>
        /// The ProjectName of the CLB.
        /// </summary>
        [Input("projectName")]
        public Input<string>? ProjectName { get; set; }

        /// <summary>
        /// The region of the request.
        /// </summary>
        [Input("regionId")]
        public Input<string>? RegionId { get; set; }

        /// <summary>
        /// The slave zone ID of the CLB.
        /// </summary>
        [Input("slaveZoneId")]
        public Input<string>? SlaveZoneId { get; set; }

        /// <summary>
        /// The id of the Subnet.
        /// </summary>
        [Input("subnetId", required: true)]
        public Input<string> SubnetId { get; set; } = null!;

        [Input("tags")]
        private InputList<Inputs.ClbTagArgs>? _tags;

        /// <summary>
        /// Tags.
        /// </summary>
        public InputList<Inputs.ClbTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.ClbTagArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// The type of the CLB. And optional choice contains `public` or `private`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// The id of the VPC.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public ClbArgs()
        {
        }
    }

    public sealed class ClbState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the CLB.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The eni address of the CLB.
        /// </summary>
        [Input("eniAddress")]
        public Input<string>? EniAddress { get; set; }

        /// <summary>
        /// The billing type of the CLB, the value can be `PostPaid`.
        /// </summary>
        [Input("loadBalancerBillingType")]
        public Input<string>? LoadBalancerBillingType { get; set; }

        /// <summary>
        /// The name of the CLB.
        /// </summary>
        [Input("loadBalancerName")]
        public Input<string>? LoadBalancerName { get; set; }

        /// <summary>
        /// The specification of the CLB, the value can be `small_1`, `small_2`, `medium_1`, `medium_2`, `large_1`, `large_2`.
        /// </summary>
        [Input("loadBalancerSpec")]
        public Input<string>? LoadBalancerSpec { get; set; }

        /// <summary>
        /// The master zone ID of the CLB.
        /// </summary>
        [Input("masterZoneId")]
        public Input<string>? MasterZoneId { get; set; }

        /// <summary>
        /// The reason of the console modification protection.
        /// </summary>
        [Input("modificationProtectionReason")]
        public Input<string>? ModificationProtectionReason { get; set; }

        /// <summary>
        /// The status of the console modification protection, the value can be `NonProtection` or `ConsoleProtection`.
        /// </summary>
        [Input("modificationProtectionStatus")]
        public Input<string>? ModificationProtectionStatus { get; set; }

        /// <summary>
        /// The ProjectName of the CLB.
        /// </summary>
        [Input("projectName")]
        public Input<string>? ProjectName { get; set; }

        /// <summary>
        /// The region of the request.
        /// </summary>
        [Input("regionId")]
        public Input<string>? RegionId { get; set; }

        /// <summary>
        /// The slave zone ID of the CLB.
        /// </summary>
        [Input("slaveZoneId")]
        public Input<string>? SlaveZoneId { get; set; }

        /// <summary>
        /// The id of the Subnet.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        [Input("tags")]
        private InputList<Inputs.ClbTagGetArgs>? _tags;

        /// <summary>
        /// Tags.
        /// </summary>
        public InputList<Inputs.ClbTagGetArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.ClbTagGetArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// The type of the CLB. And optional choice contains `public` or `private`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// The id of the VPC.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public ClbState()
        {
        }
    }
}