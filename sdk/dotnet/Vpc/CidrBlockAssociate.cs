// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Vpc
{
    /// <summary>
    /// Provides a resource to manage vpc cidr block associate
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
    ///     var fooVpc = new Volcengine.Vpc.Vpc("fooVpc", new()
    ///     {
    ///         VpcName = "acc-test-vpc",
    ///         CidrBlock = "192.168.0.0/20",
    ///         ProjectName = "default",
    ///     });
    /// 
    ///     var fooCidrBlockAssociate = new Volcengine.Vpc.CidrBlockAssociate("fooCidrBlockAssociate", new()
    ///     {
    ///         VpcId = fooVpc.Id,
    ///         SecondaryCidrBlock = "192.168.16.0/20",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// The VpcCidrBlockAssociate is not support import.
    /// </summary>
    [VolcengineResourceType("volcengine:vpc/cidrBlockAssociate:CidrBlockAssociate")]
    public partial class CidrBlockAssociate : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The secondary cidr block of the VPC.
        /// </summary>
        [Output("secondaryCidrBlock")]
        public Output<string> SecondaryCidrBlock { get; private set; } = null!;

        /// <summary>
        /// The id of the VPC.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a CidrBlockAssociate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CidrBlockAssociate(string name, CidrBlockAssociateArgs args, CustomResourceOptions? options = null)
            : base("volcengine:vpc/cidrBlockAssociate:CidrBlockAssociate", name, args ?? new CidrBlockAssociateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CidrBlockAssociate(string name, Input<string> id, CidrBlockAssociateState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:vpc/cidrBlockAssociate:CidrBlockAssociate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CidrBlockAssociate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CidrBlockAssociate Get(string name, Input<string> id, CidrBlockAssociateState? state = null, CustomResourceOptions? options = null)
        {
            return new CidrBlockAssociate(name, id, state, options);
        }
    }

    public sealed class CidrBlockAssociateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The secondary cidr block of the VPC.
        /// </summary>
        [Input("secondaryCidrBlock", required: true)]
        public Input<string> SecondaryCidrBlock { get; set; } = null!;

        /// <summary>
        /// The id of the VPC.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public CidrBlockAssociateArgs()
        {
        }
        public static new CidrBlockAssociateArgs Empty => new CidrBlockAssociateArgs();
    }

    public sealed class CidrBlockAssociateState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The secondary cidr block of the VPC.
        /// </summary>
        [Input("secondaryCidrBlock")]
        public Input<string>? SecondaryCidrBlock { get; set; }

        /// <summary>
        /// The id of the VPC.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public CidrBlockAssociateState()
        {
        }
        public static new CidrBlockAssociateState Empty => new CidrBlockAssociateState();
    }
}
