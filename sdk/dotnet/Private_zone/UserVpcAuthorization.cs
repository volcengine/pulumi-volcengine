// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Private_zone
{
    /// <summary>
    /// Provides a resource to manage private zone user vpc authorization
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
    ///     var foo = new Volcengine.Private_zone.UserVpcAuthorization("foo", new()
    ///     {
    ///         AccountId = "2100278462",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// PrivateZoneUserVpcAuthorization can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import volcengine:private_zone/userVpcAuthorization:UserVpcAuthorization default resource_id
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:private_zone/userVpcAuthorization:UserVpcAuthorization")]
    public partial class UserVpcAuthorization : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The account Id which authorizes the private zone resource.
        /// </summary>
        [Output("accountId")]
        public Output<string> AccountId { get; private set; } = null!;


        /// <summary>
        /// Create a UserVpcAuthorization resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public UserVpcAuthorization(string name, UserVpcAuthorizationArgs args, CustomResourceOptions? options = null)
            : base("volcengine:private_zone/userVpcAuthorization:UserVpcAuthorization", name, args ?? new UserVpcAuthorizationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private UserVpcAuthorization(string name, Input<string> id, UserVpcAuthorizationState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:private_zone/userVpcAuthorization:UserVpcAuthorization", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing UserVpcAuthorization resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static UserVpcAuthorization Get(string name, Input<string> id, UserVpcAuthorizationState? state = null, CustomResourceOptions? options = null)
        {
            return new UserVpcAuthorization(name, id, state, options);
        }
    }

    public sealed class UserVpcAuthorizationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The account Id which authorizes the private zone resource.
        /// </summary>
        [Input("accountId", required: true)]
        public Input<string> AccountId { get; set; } = null!;

        public UserVpcAuthorizationArgs()
        {
        }
        public static new UserVpcAuthorizationArgs Empty => new UserVpcAuthorizationArgs();
    }

    public sealed class UserVpcAuthorizationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The account Id which authorizes the private zone resource.
        /// </summary>
        [Input("accountId")]
        public Input<string>? AccountId { get; set; }

        public UserVpcAuthorizationState()
        {
        }
        public static new UserVpcAuthorizationState Empty => new UserVpcAuthorizationState();
    }
}