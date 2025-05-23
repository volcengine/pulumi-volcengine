// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Iam
{
    [Obsolete(@"volcengine.iam.UserGroupPolicyAttachments has been deprecated in favor of volcengine.iam.getUserGroupPolicyAttachments")]
    public static class UserGroupPolicyAttachments
    {
        /// <summary>
        /// Use this data source to query detailed information of iam user group policy attachments
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
        ///     var fooUserGroupPolicyAttachments = Volcengine.Iam.GetUserGroupPolicyAttachments.Invoke(new()
        ///     {
        ///         UserGroupName = fooUserGroupPolicyAttachment.UserGroupName,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<UserGroupPolicyAttachmentsResult> InvokeAsync(UserGroupPolicyAttachmentsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<UserGroupPolicyAttachmentsResult>("volcengine:iam/userGroupPolicyAttachments:UserGroupPolicyAttachments", args ?? new UserGroupPolicyAttachmentsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of iam user group policy attachments
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
        ///     var fooUserGroupPolicyAttachments = Volcengine.Iam.GetUserGroupPolicyAttachments.Invoke(new()
        ///     {
        ///         UserGroupName = fooUserGroupPolicyAttachment.UserGroupName,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<UserGroupPolicyAttachmentsResult> Invoke(UserGroupPolicyAttachmentsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<UserGroupPolicyAttachmentsResult>("volcengine:iam/userGroupPolicyAttachments:UserGroupPolicyAttachments", args ?? new UserGroupPolicyAttachmentsInvokeArgs(), options.WithDefaults());
    }


    public sealed class UserGroupPolicyAttachmentsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// A name of user group.
        /// </summary>
        [Input("userGroupName", required: true)]
        public string UserGroupName { get; set; } = null!;

        public UserGroupPolicyAttachmentsArgs()
        {
        }
        public static new UserGroupPolicyAttachmentsArgs Empty => new UserGroupPolicyAttachmentsArgs();
    }

    public sealed class UserGroupPolicyAttachmentsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// A name of user group.
        /// </summary>
        [Input("userGroupName", required: true)]
        public Input<string> UserGroupName { get; set; } = null!;

        public UserGroupPolicyAttachmentsInvokeArgs()
        {
        }
        public static new UserGroupPolicyAttachmentsInvokeArgs Empty => new UserGroupPolicyAttachmentsInvokeArgs();
    }


    [OutputType]
    public sealed class UserGroupPolicyAttachmentsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? OutputFile;
        /// <summary>
        /// The collection of query.
        /// </summary>
        public readonly ImmutableArray<Outputs.UserGroupPolicyAttachmentsPolicyResult> Policies;
        /// <summary>
        /// The total count of query.
        /// </summary>
        public readonly int TotalCount;
        public readonly string UserGroupName;

        [OutputConstructor]
        private UserGroupPolicyAttachmentsResult(
            string id,

            string? outputFile,

            ImmutableArray<Outputs.UserGroupPolicyAttachmentsPolicyResult> policies,

            int totalCount,

            string userGroupName)
        {
            Id = id;
            OutputFile = outputFile;
            Policies = policies;
            TotalCount = totalCount;
            UserGroupName = userGroupName;
        }
    }
}
