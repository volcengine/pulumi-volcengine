// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Iam
{
    /// <summary>
    /// Provides a resource to manage iam user group policy attachment
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
    ///     var fooPolicy = new Volcengine.Iam.Policy("fooPolicy", new()
    ///     {
    ///         PolicyName = "acc-test-policy",
    ///         Description = "acc-test",
    ///         PolicyDocument = "{\"Statement\":[{\"Effect\":\"Allow\",\"Action\":[\"auto_scaling:DescribeScalingGroups\"],\"Resource\":[\"*\"]}]}",
    ///     });
    /// 
    ///     var fooUserGroup = new Volcengine.Iam.UserGroup("fooUserGroup", new()
    ///     {
    ///         UserGroupName = "acc-test-group",
    ///         Description = "acc-test",
    ///         DisplayName = "acc-test",
    ///     });
    /// 
    ///     var fooUserGroupPolicyAttachment = new Volcengine.Iam.UserGroupPolicyAttachment("fooUserGroupPolicyAttachment", new()
    ///     {
    ///         PolicyName = fooPolicy.PolicyName,
    ///         PolicyType = "Custom",
    ///         UserGroupName = fooUserGroup.UserGroupName,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// IamUserGroupPolicyAttachment can be imported using the user group name and policy name, e.g.
    /// 
    /// ```sh
    /// $ pulumi import volcengine:iam/userGroupPolicyAttachment:UserGroupPolicyAttachment default userGroupName:policyName
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:iam/userGroupPolicyAttachment:UserGroupPolicyAttachment")]
    public partial class UserGroupPolicyAttachment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The policy name.
        /// </summary>
        [Output("policyName")]
        public Output<string> PolicyName { get; private set; } = null!;

        /// <summary>
        /// Strategy types, System strategy, Custom strategy.
        /// </summary>
        [Output("policyType")]
        public Output<string> PolicyType { get; private set; } = null!;

        /// <summary>
        /// The user group name.
        /// </summary>
        [Output("userGroupName")]
        public Output<string> UserGroupName { get; private set; } = null!;


        /// <summary>
        /// Create a UserGroupPolicyAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public UserGroupPolicyAttachment(string name, UserGroupPolicyAttachmentArgs args, CustomResourceOptions? options = null)
            : base("volcengine:iam/userGroupPolicyAttachment:UserGroupPolicyAttachment", name, args ?? new UserGroupPolicyAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private UserGroupPolicyAttachment(string name, Input<string> id, UserGroupPolicyAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:iam/userGroupPolicyAttachment:UserGroupPolicyAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing UserGroupPolicyAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static UserGroupPolicyAttachment Get(string name, Input<string> id, UserGroupPolicyAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new UserGroupPolicyAttachment(name, id, state, options);
        }
    }

    public sealed class UserGroupPolicyAttachmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The policy name.
        /// </summary>
        [Input("policyName", required: true)]
        public Input<string> PolicyName { get; set; } = null!;

        /// <summary>
        /// Strategy types, System strategy, Custom strategy.
        /// </summary>
        [Input("policyType", required: true)]
        public Input<string> PolicyType { get; set; } = null!;

        /// <summary>
        /// The user group name.
        /// </summary>
        [Input("userGroupName", required: true)]
        public Input<string> UserGroupName { get; set; } = null!;

        public UserGroupPolicyAttachmentArgs()
        {
        }
        public static new UserGroupPolicyAttachmentArgs Empty => new UserGroupPolicyAttachmentArgs();
    }

    public sealed class UserGroupPolicyAttachmentState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The policy name.
        /// </summary>
        [Input("policyName")]
        public Input<string>? PolicyName { get; set; }

        /// <summary>
        /// Strategy types, System strategy, Custom strategy.
        /// </summary>
        [Input("policyType")]
        public Input<string>? PolicyType { get; set; }

        /// <summary>
        /// The user group name.
        /// </summary>
        [Input("userGroupName")]
        public Input<string>? UserGroupName { get; set; }

        public UserGroupPolicyAttachmentState()
        {
        }
        public static new UserGroupPolicyAttachmentState Empty => new UserGroupPolicyAttachmentState();
    }
}
