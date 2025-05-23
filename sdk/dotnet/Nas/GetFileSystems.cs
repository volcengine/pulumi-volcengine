// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Nas
{
    public static class GetFileSystems
    {
        /// <summary>
        /// Use this data source to query detailed information of nas file systems
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
        ///     var fooZones = Volcengine.Nas.GetZones.Invoke();
        /// 
        ///     var fooFileSystem = new List&lt;Volcengine.Nas.FileSystem&gt;();
        ///     for (var rangeIndex = 0; rangeIndex &lt; 3; rangeIndex++)
        ///     {
        ///         var range = new { Value = rangeIndex };
        ///         fooFileSystem.Add(new Volcengine.Nas.FileSystem($"fooFileSystem-{range.Value}", new()
        ///         {
        ///             FileSystemName = $"acc-test-fs-{range.Value}",
        ///             Description = "acc-test",
        ///             ZoneId = fooZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
        ///             Capacity = 103,
        ///             ProjectName = "default",
        ///             Tags = new[]
        ///             {
        ///                 new Volcengine.Nas.Inputs.FileSystemTagArgs
        ///                 {
        ///                     Key = "k1",
        ///                     Value = "v1",
        ///                 },
        ///             },
        ///         }));
        ///     }
        ///     var fooFileSystems = Volcengine.Nas.GetFileSystems.Invoke(new()
        ///     {
        ///         Ids = fooFileSystem.Select(__item =&gt; __item.Id).ToList(),
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetFileSystemsResult> InvokeAsync(GetFileSystemsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFileSystemsResult>("volcengine:nas/getFileSystems:getFileSystems", args ?? new GetFileSystemsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of nas file systems
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
        ///     var fooZones = Volcengine.Nas.GetZones.Invoke();
        /// 
        ///     var fooFileSystem = new List&lt;Volcengine.Nas.FileSystem&gt;();
        ///     for (var rangeIndex = 0; rangeIndex &lt; 3; rangeIndex++)
        ///     {
        ///         var range = new { Value = rangeIndex };
        ///         fooFileSystem.Add(new Volcengine.Nas.FileSystem($"fooFileSystem-{range.Value}", new()
        ///         {
        ///             FileSystemName = $"acc-test-fs-{range.Value}",
        ///             Description = "acc-test",
        ///             ZoneId = fooZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
        ///             Capacity = 103,
        ///             ProjectName = "default",
        ///             Tags = new[]
        ///             {
        ///                 new Volcengine.Nas.Inputs.FileSystemTagArgs
        ///                 {
        ///                     Key = "k1",
        ///                     Value = "v1",
        ///                 },
        ///             },
        ///         }));
        ///     }
        ///     var fooFileSystems = Volcengine.Nas.GetFileSystems.Invoke(new()
        ///     {
        ///         Ids = fooFileSystem.Select(__item =&gt; __item.Id).ToList(),
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetFileSystemsResult> Invoke(GetFileSystemsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFileSystemsResult>("volcengine:nas/getFileSystems:getFileSystems", args ?? new GetFileSystemsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFileSystemsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The charge type of nas file system.
        /// </summary>
        [Input("chargeType")]
        public string? ChargeType { get; set; }

        /// <summary>
        /// The name of nas file system. This field supports fuzzy queries.
        /// </summary>
        [Input("fileSystemName")]
        public string? FileSystemName { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of nas file system ids.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// The mount point id of nas file system.
        /// </summary>
        [Input("mountPointId")]
        public string? MountPointId { get; set; }

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
        /// The permission group id of nas file system.
        /// </summary>
        [Input("permissionGroupId")]
        public string? PermissionGroupId { get; set; }

        /// <summary>
        /// The project name of nas file system.
        /// </summary>
        [Input("projectName")]
        public string? ProjectName { get; set; }

        /// <summary>
        /// The protocol type of nas file system.
        /// </summary>
        [Input("protocolType")]
        public string? ProtocolType { get; set; }

        [Input("statuses")]
        private List<string>? _statuses;

        /// <summary>
        /// The status of nas file system.
        /// </summary>
        public List<string> Statuses
        {
            get => _statuses ?? (_statuses = new List<string>());
            set => _statuses = value;
        }

        /// <summary>
        /// The storage type of nas file system.
        /// </summary>
        [Input("storageType")]
        public string? StorageType { get; set; }

        [Input("tags")]
        private List<Inputs.GetFileSystemsTagArgs>? _tags;

        /// <summary>
        /// Tags.
        /// </summary>
        public List<Inputs.GetFileSystemsTagArgs> Tags
        {
            get => _tags ?? (_tags = new List<Inputs.GetFileSystemsTagArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// The zone id of nas file system.
        /// </summary>
        [Input("zoneId")]
        public string? ZoneId { get; set; }

        public GetFileSystemsArgs()
        {
        }
        public static new GetFileSystemsArgs Empty => new GetFileSystemsArgs();
    }

    public sealed class GetFileSystemsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The charge type of nas file system.
        /// </summary>
        [Input("chargeType")]
        public Input<string>? ChargeType { get; set; }

        /// <summary>
        /// The name of nas file system. This field supports fuzzy queries.
        /// </summary>
        [Input("fileSystemName")]
        public Input<string>? FileSystemName { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of nas file system ids.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// The mount point id of nas file system.
        /// </summary>
        [Input("mountPointId")]
        public Input<string>? MountPointId { get; set; }

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
        /// The permission group id of nas file system.
        /// </summary>
        [Input("permissionGroupId")]
        public Input<string>? PermissionGroupId { get; set; }

        /// <summary>
        /// The project name of nas file system.
        /// </summary>
        [Input("projectName")]
        public Input<string>? ProjectName { get; set; }

        /// <summary>
        /// The protocol type of nas file system.
        /// </summary>
        [Input("protocolType")]
        public Input<string>? ProtocolType { get; set; }

        [Input("statuses")]
        private InputList<string>? _statuses;

        /// <summary>
        /// The status of nas file system.
        /// </summary>
        public InputList<string> Statuses
        {
            get => _statuses ?? (_statuses = new InputList<string>());
            set => _statuses = value;
        }

        /// <summary>
        /// The storage type of nas file system.
        /// </summary>
        [Input("storageType")]
        public Input<string>? StorageType { get; set; }

        [Input("tags")]
        private InputList<Inputs.GetFileSystemsTagInputArgs>? _tags;

        /// <summary>
        /// Tags.
        /// </summary>
        public InputList<Inputs.GetFileSystemsTagInputArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.GetFileSystemsTagInputArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// The zone id of nas file system.
        /// </summary>
        [Input("zoneId")]
        public Input<string>? ZoneId { get; set; }

        public GetFileSystemsInvokeArgs()
        {
        }
        public static new GetFileSystemsInvokeArgs Empty => new GetFileSystemsInvokeArgs();
    }


    [OutputType]
    public sealed class GetFileSystemsResult
    {
        /// <summary>
        /// The charge type of the nas file system.
        /// </summary>
        public readonly string? ChargeType;
        /// <summary>
        /// The name of the nas file system.
        /// </summary>
        public readonly string? FileSystemName;
        /// <summary>
        /// The collection of query.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFileSystemsFileSystemResult> FileSystems;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? MountPointId;
        public readonly string? NameRegex;
        public readonly string? OutputFile;
        public readonly string? PermissionGroupId;
        /// <summary>
        /// The project name of the nas file system.
        /// </summary>
        public readonly string? ProjectName;
        /// <summary>
        /// The protocol type of the nas file system.
        /// </summary>
        public readonly string? ProtocolType;
        /// <summary>
        /// The status of the nas file system.
        /// </summary>
        public readonly ImmutableArray<string> Statuses;
        /// <summary>
        /// The storage type of the nas file system.
        /// </summary>
        public readonly string? StorageType;
        /// <summary>
        /// Tags of the nas file system.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFileSystemsTagResult> Tags;
        /// <summary>
        /// The total count of query.
        /// </summary>
        public readonly int TotalCount;
        /// <summary>
        /// The zone id of the nas file system.
        /// </summary>
        public readonly string? ZoneId;

        [OutputConstructor]
        private GetFileSystemsResult(
            string? chargeType,

            string? fileSystemName,

            ImmutableArray<Outputs.GetFileSystemsFileSystemResult> fileSystems,

            string id,

            ImmutableArray<string> ids,

            string? mountPointId,

            string? nameRegex,

            string? outputFile,

            string? permissionGroupId,

            string? projectName,

            string? protocolType,

            ImmutableArray<string> statuses,

            string? storageType,

            ImmutableArray<Outputs.GetFileSystemsTagResult> tags,

            int totalCount,

            string? zoneId)
        {
            ChargeType = chargeType;
            FileSystemName = fileSystemName;
            FileSystems = fileSystems;
            Id = id;
            Ids = ids;
            MountPointId = mountPointId;
            NameRegex = nameRegex;
            OutputFile = outputFile;
            PermissionGroupId = permissionGroupId;
            ProjectName = projectName;
            ProtocolType = protocolType;
            Statuses = statuses;
            StorageType = storageType;
            Tags = tags;
            TotalCount = totalCount;
            ZoneId = zoneId;
        }
    }
}
