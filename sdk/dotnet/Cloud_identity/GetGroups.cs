// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Cloud_identity
{
    public static class GetGroups
    {
        /// <summary>
        /// Use this data source to query detailed information of cloud identity groups
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
        ///     var fooGroup = new List&lt;Volcengine.Cloud_identity.Group&gt;();
        ///     for (var rangeIndex = 0; rangeIndex &lt; 2; rangeIndex++)
        ///     {
        ///         var range = new { Value = rangeIndex };
        ///         fooGroup.Add(new Volcengine.Cloud_identity.Group($"fooGroup-{range.Value}", new()
        ///         {
        ///             Description = "tf",
        ///             DisplayName = $"tf-test-group-{range.Value}",
        ///             GroupName = $"acc-test-group-{range.Value}",
        ///             JoinType = "Manual",
        ///         }));
        ///     }
        ///     var fooGroups = Volcengine.Cloud_identity.GetGroups.Invoke(new()
        ///     {
        ///         GroupName = "acc-test-group",
        ///         JoinType = "Manual",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetGroupsResult> InvokeAsync(GetGroupsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetGroupsResult>("volcengine:cloud_identity/getGroups:getGroups", args ?? new GetGroupsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of cloud identity groups
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
        ///     var fooGroup = new List&lt;Volcengine.Cloud_identity.Group&gt;();
        ///     for (var rangeIndex = 0; rangeIndex &lt; 2; rangeIndex++)
        ///     {
        ///         var range = new { Value = rangeIndex };
        ///         fooGroup.Add(new Volcengine.Cloud_identity.Group($"fooGroup-{range.Value}", new()
        ///         {
        ///             Description = "tf",
        ///             DisplayName = $"tf-test-group-{range.Value}",
        ///             GroupName = $"acc-test-group-{range.Value}",
        ///             JoinType = "Manual",
        ///         }));
        ///     }
        ///     var fooGroups = Volcengine.Cloud_identity.GetGroups.Invoke(new()
        ///     {
        ///         GroupName = "acc-test-group",
        ///         JoinType = "Manual",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetGroupsResult> Invoke(GetGroupsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetGroupsResult>("volcengine:cloud_identity/getGroups:getGroups", args ?? new GetGroupsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGroupsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The display name of cloud identity group.
        /// </summary>
        [Input("displayName")]
        public string? DisplayName { get; set; }

        /// <summary>
        /// The name of cloud identity group.
        /// </summary>
        [Input("groupName")]
        public string? GroupName { get; set; }

        /// <summary>
        /// The join type of cloud identity group. Valid values: `Auto`, `Manual`.
        /// </summary>
        [Input("joinType")]
        public string? JoinType { get; set; }

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

        public GetGroupsArgs()
        {
        }
        public static new GetGroupsArgs Empty => new GetGroupsArgs();
    }

    public sealed class GetGroupsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The display name of cloud identity group.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The name of cloud identity group.
        /// </summary>
        [Input("groupName")]
        public Input<string>? GroupName { get; set; }

        /// <summary>
        /// The join type of cloud identity group. Valid values: `Auto`, `Manual`.
        /// </summary>
        [Input("joinType")]
        public Input<string>? JoinType { get; set; }

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

        public GetGroupsInvokeArgs()
        {
        }
        public static new GetGroupsInvokeArgs Empty => new GetGroupsInvokeArgs();
    }


    [OutputType]
    public sealed class GetGroupsResult
    {
        /// <summary>
        /// The display name of the cloud identity group.
        /// </summary>
        public readonly string? DisplayName;
        /// <summary>
        /// The name of the cloud identity group.
        /// </summary>
        public readonly string? GroupName;
        /// <summary>
        /// The collection of query.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetGroupsGroupResult> Groups;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The email of the cloud identity group.
        /// </summary>
        public readonly string? JoinType;
        public readonly string? NameRegex;
        public readonly string? OutputFile;
        /// <summary>
        /// The total count of query.
        /// </summary>
        public readonly int TotalCount;

        [OutputConstructor]
        private GetGroupsResult(
            string? displayName,

            string? groupName,

            ImmutableArray<Outputs.GetGroupsGroupResult> groups,

            string id,

            string? joinType,

            string? nameRegex,

            string? outputFile,

            int totalCount)
        {
            DisplayName = displayName;
            GroupName = groupName;
            Groups = groups;
            Id = id;
            JoinType = joinType;
            NameRegex = nameRegex;
            OutputFile = outputFile;
            TotalCount = totalCount;
        }
    }
}
