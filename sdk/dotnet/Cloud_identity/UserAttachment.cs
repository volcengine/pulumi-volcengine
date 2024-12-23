// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Cloud_identity
{
    /// <summary>
    /// Provides a resource to manage cloud identity user attachment
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
    ///     var fooGroup = new Volcengine.Cloud_identity.Group("fooGroup", new()
    ///     {
    ///         GroupName = "acc-test-group",
    ///         DisplayName = "tf-test-group",
    ///         JoinType = "Manual",
    ///         Description = "tf",
    ///     });
    /// 
    ///     var fooUser = new Volcengine.Cloud_identity.User("fooUser", new()
    ///     {
    ///         UserName = "acc-test-user",
    ///         DisplayName = "tf-test-user",
    ///         Description = "tf",
    ///         Email = "88@qq.com",
    ///         Phone = "181",
    ///     });
    /// 
    ///     var fooUserAttachment = new Volcengine.Cloud_identity.UserAttachment("fooUserAttachment", new()
    ///     {
    ///         UserId = fooUser.Id,
    ///         GroupId = fooGroup.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// CloudIdentityUserAttachment can be imported using the group_id:user_id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import volcengine:cloud_identity/userAttachment:UserAttachment default resource_id
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:cloud_identity/userAttachment:UserAttachment")]
    public partial class UserAttachment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The id of the cloud identity group.
        /// </summary>
        [Output("groupId")]
        public Output<string> GroupId { get; private set; } = null!;

        /// <summary>
        /// The id of the cloud identity user.
        /// </summary>
        [Output("userId")]
        public Output<string> UserId { get; private set; } = null!;


        /// <summary>
        /// Create a UserAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public UserAttachment(string name, UserAttachmentArgs args, CustomResourceOptions? options = null)
            : base("volcengine:cloud_identity/userAttachment:UserAttachment", name, args ?? new UserAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private UserAttachment(string name, Input<string> id, UserAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:cloud_identity/userAttachment:UserAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing UserAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static UserAttachment Get(string name, Input<string> id, UserAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new UserAttachment(name, id, state, options);
        }
    }

    public sealed class UserAttachmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The id of the cloud identity group.
        /// </summary>
        [Input("groupId", required: true)]
        public Input<string> GroupId { get; set; } = null!;

        /// <summary>
        /// The id of the cloud identity user.
        /// </summary>
        [Input("userId", required: true)]
        public Input<string> UserId { get; set; } = null!;

        public UserAttachmentArgs()
        {
        }
        public static new UserAttachmentArgs Empty => new UserAttachmentArgs();
    }

    public sealed class UserAttachmentState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The id of the cloud identity group.
        /// </summary>
        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        /// <summary>
        /// The id of the cloud identity user.
        /// </summary>
        [Input("userId")]
        public Input<string>? UserId { get; set; }

        public UserAttachmentState()
        {
        }
        public static new UserAttachmentState Empty => new UserAttachmentState();
    }
}
