// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Cloud_identity
{
    public static class Users
    {
        /// <summary>
        /// Use this data source to query detailed information of cloud identity users
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
        ///     var fooUser = new List&lt;Volcengine.Cloud_identity.User&gt;();
        ///     for (var rangeIndex = 0; rangeIndex &lt; 2; rangeIndex++)
        ///     {
        ///         var range = new { Value = rangeIndex };
        ///         fooUser.Add(new Volcengine.Cloud_identity.User($"fooUser-{range.Value}", new()
        ///         {
        ///             Description = "tf",
        ///             DisplayName = $"tf-test-user-{range.Value}",
        ///             Email = "88@qq.com",
        ///             Phone = "181",
        ///             UserName = $"acc-test-user-{range.Value}",
        ///         }));
        ///     }
        ///     var fooUsers = Volcengine.Cloud_identity.Users.Invoke(new()
        ///     {
        ///         Source = "Manual",
        ///         UserName = "acc-test-user",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<UsersResult> InvokeAsync(UsersArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<UsersResult>("volcengine:cloud_identity/users:Users", args ?? new UsersArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of cloud identity users
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
        ///     var fooUser = new List&lt;Volcengine.Cloud_identity.User&gt;();
        ///     for (var rangeIndex = 0; rangeIndex &lt; 2; rangeIndex++)
        ///     {
        ///         var range = new { Value = rangeIndex };
        ///         fooUser.Add(new Volcengine.Cloud_identity.User($"fooUser-{range.Value}", new()
        ///         {
        ///             Description = "tf",
        ///             DisplayName = $"tf-test-user-{range.Value}",
        ///             Email = "88@qq.com",
        ///             Phone = "181",
        ///             UserName = $"acc-test-user-{range.Value}",
        ///         }));
        ///     }
        ///     var fooUsers = Volcengine.Cloud_identity.Users.Invoke(new()
        ///     {
        ///         Source = "Manual",
        ///         UserName = "acc-test-user",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<UsersResult> Invoke(UsersInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<UsersResult>("volcengine:cloud_identity/users:Users", args ?? new UsersInvokeArgs(), options.WithDefaults());
    }


    public sealed class UsersArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The department id.
        /// </summary>
        [Input("departmentId")]
        public string? DepartmentId { get; set; }

        /// <summary>
        /// The display name of cloud identity user.
        /// </summary>
        [Input("displayName")]
        public string? DisplayName { get; set; }

        /// <summary>
        /// A Name Regex of Resource.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The source of cloud identity user. Valid values: `Sync`, `Manual`.
        /// </summary>
        [Input("source")]
        public string? Source { get; set; }

        /// <summary>
        /// The name of cloud identity user.
        /// </summary>
        [Input("userName")]
        public string? UserName { get; set; }

        public UsersArgs()
        {
        }
        public static new UsersArgs Empty => new UsersArgs();
    }

    public sealed class UsersInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The department id.
        /// </summary>
        [Input("departmentId")]
        public Input<string>? DepartmentId { get; set; }

        /// <summary>
        /// The display name of cloud identity user.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// A Name Regex of Resource.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// The source of cloud identity user. Valid values: `Sync`, `Manual`.
        /// </summary>
        [Input("source")]
        public Input<string>? Source { get; set; }

        /// <summary>
        /// The name of cloud identity user.
        /// </summary>
        [Input("userName")]
        public Input<string>? UserName { get; set; }

        public UsersInvokeArgs()
        {
        }
        public static new UsersInvokeArgs Empty => new UsersInvokeArgs();
    }


    [OutputType]
    public sealed class UsersResult
    {
        public readonly string? DepartmentId;
        /// <summary>
        /// The display name of the cloud identity user.
        /// </summary>
        public readonly string? DisplayName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? NameRegex;
        public readonly string? OutputFile;
        /// <summary>
        /// The source of the cloud identity user.
        /// </summary>
        public readonly string? Source;
        /// <summary>
        /// The total count of query.
        /// </summary>
        public readonly int TotalCount;
        /// <summary>
        /// The name of the cloud identity user.
        /// </summary>
        public readonly string? UserName;
        /// <summary>
        /// The collection of query.
        /// </summary>
        public readonly ImmutableArray<Outputs.UsersUserResult> Users;

        [OutputConstructor]
        private UsersResult(
            string? departmentId,

            string? displayName,

            string id,

            string? nameRegex,

            string? outputFile,

            string? source,

            int totalCount,

            string? userName,

            ImmutableArray<Outputs.UsersUserResult> users)
        {
            DepartmentId = departmentId;
            DisplayName = displayName;
            Id = id;
            NameRegex = nameRegex;
            OutputFile = outputFile;
            Source = source;
            TotalCount = totalCount;
            UserName = userName;
            Users = users;
        }
    }
}