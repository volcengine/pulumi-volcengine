// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Cloud_monitor
{
    public static class GetContactGroups
    {
        /// <summary>
        /// Use this data source to query detailed information of cloud monitor contact groups
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
        ///     var foo = Volcengine.Cloud_monitor.GetContactGroups.Invoke(new()
        ///     {
        ///         Name = "tftest",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetContactGroupsResult> InvokeAsync(GetContactGroupsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetContactGroupsResult>("volcengine:cloud_monitor/getContactGroups:getContactGroups", args ?? new GetContactGroupsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of cloud monitor contact groups
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
        ///     var foo = Volcengine.Cloud_monitor.GetContactGroups.Invoke(new()
        ///     {
        ///         Name = "tftest",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetContactGroupsResult> Invoke(GetContactGroupsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetContactGroupsResult>("volcengine:cloud_monitor/getContactGroups:getContactGroups", args ?? new GetContactGroupsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetContactGroupsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Search for keywords in contact group names, supports fuzzy search.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public GetContactGroupsArgs()
        {
        }
        public static new GetContactGroupsArgs Empty => new GetContactGroupsArgs();
    }

    public sealed class GetContactGroupsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Search for keywords in contact group names, supports fuzzy search.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public GetContactGroupsInvokeArgs()
        {
        }
        public static new GetContactGroupsInvokeArgs Empty => new GetContactGroupsInvokeArgs();
    }


    [OutputType]
    public sealed class GetContactGroupsResult
    {
        /// <summary>
        /// The collection of query.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetContactGroupsGroupResult> Groups;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the contact group.
        /// </summary>
        public readonly string? Name;
        public readonly string? OutputFile;
        /// <summary>
        /// The total count of query.
        /// </summary>
        public readonly int TotalCount;

        [OutputConstructor]
        private GetContactGroupsResult(
            ImmutableArray<Outputs.GetContactGroupsGroupResult> groups,

            string id,

            string? name,

            string? outputFile,

            int totalCount)
        {
            Groups = groups;
            Id = id;
            Name = name;
            OutputFile = outputFile;
            TotalCount = totalCount;
        }
    }
}
