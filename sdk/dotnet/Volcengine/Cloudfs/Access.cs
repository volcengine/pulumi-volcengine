// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.Pulumi.Volcengine.Cloudfs
{
    /// <summary>
    /// Provides a resource to manage cloudfs access
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Volcengine = Volcengine.Pulumi.Volcengine;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var foo1 = new Volcengine.Cloudfs.Access("foo1", new()
    ///     {
    ///         FsName = "tftest2",
    ///         SecurityGroupId = "sg-rrv1klfg5s00v0x578mx14m",
    ///         SubnetId = "subnet-13fca1crr5d6o3n6nu46cyb5m",
    ///         VpcRouteEnabled = false,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// CloudFs Access can be imported using the FsName:AccessId, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import volcengine:cloudfs/access:Access default tfname:access-**rdgmedx3fow
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:cloudfs/access:Access")]
    public partial class Access : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The account id of access.
        /// </summary>
        [Output("accessAccountId")]
        public Output<int> AccessAccountId { get; private set; } = null!;

        /// <summary>
        /// The iam role of access. If the VPC of another account is attached, the other account needs to create a role with CFSCacheAccess permission, and enter the role name as a parameter.
        /// </summary>
        [Output("accessIamRole")]
        public Output<string?> AccessIamRole { get; private set; } = null!;

        /// <summary>
        /// The id of access.
        /// </summary>
        [Output("accessId")]
        public Output<string> AccessId { get; private set; } = null!;

        /// <summary>
        /// The service name of access.
        /// </summary>
        [Output("accessServiceName")]
        public Output<string> AccessServiceName { get; private set; } = null!;

        /// <summary>
        /// The creation time.
        /// </summary>
        [Output("createdTime")]
        public Output<string> CreatedTime { get; private set; } = null!;

        /// <summary>
        /// The name of file system.
        /// </summary>
        [Output("fsName")]
        public Output<string> FsName { get; private set; } = null!;

        /// <summary>
        /// Whether is default access.
        /// </summary>
        [Output("isDefault")]
        public Output<bool> IsDefault { get; private set; } = null!;

        /// <summary>
        /// The id of security group.
        /// </summary>
        [Output("securityGroupId")]
        public Output<string> SecurityGroupId { get; private set; } = null!;

        /// <summary>
        /// Status of access.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The id of subnet.
        /// </summary>
        [Output("subnetId")]
        public Output<string> SubnetId { get; private set; } = null!;

        /// <summary>
        /// The id of vpc.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;

        /// <summary>
        /// Whether enable all vpc route.
        /// </summary>
        [Output("vpcRouteEnabled")]
        public Output<bool?> VpcRouteEnabled { get; private set; } = null!;


        /// <summary>
        /// Create a Access resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Access(string name, AccessArgs args, CustomResourceOptions? options = null)
            : base("volcengine:cloudfs/access:Access", name, args ?? new AccessArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Access(string name, Input<string> id, AccessState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:cloudfs/access:Access", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Access resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Access Get(string name, Input<string> id, AccessState? state = null, CustomResourceOptions? options = null)
        {
            return new Access(name, id, state, options);
        }
    }

    public sealed class AccessArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The account id of access.
        /// </summary>
        [Input("accessAccountId")]
        public Input<int>? AccessAccountId { get; set; }

        /// <summary>
        /// The iam role of access. If the VPC of another account is attached, the other account needs to create a role with CFSCacheAccess permission, and enter the role name as a parameter.
        /// </summary>
        [Input("accessIamRole")]
        public Input<string>? AccessIamRole { get; set; }

        /// <summary>
        /// The name of file system.
        /// </summary>
        [Input("fsName", required: true)]
        public Input<string> FsName { get; set; } = null!;

        /// <summary>
        /// The id of security group.
        /// </summary>
        [Input("securityGroupId", required: true)]
        public Input<string> SecurityGroupId { get; set; } = null!;

        /// <summary>
        /// The id of subnet.
        /// </summary>
        [Input("subnetId", required: true)]
        public Input<string> SubnetId { get; set; } = null!;

        /// <summary>
        /// Whether enable all vpc route.
        /// </summary>
        [Input("vpcRouteEnabled")]
        public Input<bool>? VpcRouteEnabled { get; set; }

        public AccessArgs()
        {
        }
        public static new AccessArgs Empty => new AccessArgs();
    }

    public sealed class AccessState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The account id of access.
        /// </summary>
        [Input("accessAccountId")]
        public Input<int>? AccessAccountId { get; set; }

        /// <summary>
        /// The iam role of access. If the VPC of another account is attached, the other account needs to create a role with CFSCacheAccess permission, and enter the role name as a parameter.
        /// </summary>
        [Input("accessIamRole")]
        public Input<string>? AccessIamRole { get; set; }

        /// <summary>
        /// The id of access.
        /// </summary>
        [Input("accessId")]
        public Input<string>? AccessId { get; set; }

        /// <summary>
        /// The service name of access.
        /// </summary>
        [Input("accessServiceName")]
        public Input<string>? AccessServiceName { get; set; }

        /// <summary>
        /// The creation time.
        /// </summary>
        [Input("createdTime")]
        public Input<string>? CreatedTime { get; set; }

        /// <summary>
        /// The name of file system.
        /// </summary>
        [Input("fsName")]
        public Input<string>? FsName { get; set; }

        /// <summary>
        /// Whether is default access.
        /// </summary>
        [Input("isDefault")]
        public Input<bool>? IsDefault { get; set; }

        /// <summary>
        /// The id of security group.
        /// </summary>
        [Input("securityGroupId")]
        public Input<string>? SecurityGroupId { get; set; }

        /// <summary>
        /// Status of access.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The id of subnet.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        /// <summary>
        /// The id of vpc.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        /// <summary>
        /// Whether enable all vpc route.
        /// </summary>
        [Input("vpcRouteEnabled")]
        public Input<bool>? VpcRouteEnabled { get; set; }

        public AccessState()
        {
        }
        public static new AccessState Empty => new AccessState();
    }
}