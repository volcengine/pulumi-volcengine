// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Vepfs
{
    public static class FileSystems
    {
        /// <summary>
        /// Use this data source to query detailed information of vepfs file systems
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
        ///         CidrBlock = "172.16.0.0/16",
        ///     });
        /// 
        ///     var fooSubnet = new Volcengine.Vpc.Subnet("fooSubnet", new()
        ///     {
        ///         SubnetName = "acc-test-subnet",
        ///         CidrBlock = "172.16.0.0/24",
        ///         ZoneId = "cn-beijing-a",
        ///         VpcId = fooVpc.Id,
        ///     });
        /// 
        ///     var fooFileSystem = new Volcengine.Vepfs.FileSystem("fooFileSystem", new()
        ///     {
        ///         FileSystemName = "acc-test-file-system",
        ///         SubnetId = fooSubnet.Id,
        ///         StoreType = "Advance_100",
        ///         Description = "tf-test",
        ///         Capacity = 12,
        ///         Project = "default",
        ///         EnableRestripe = false,
        ///         Tags = new[]
        ///         {
        ///             new Volcengine.Vepfs.Inputs.FileSystemTagArgs
        ///             {
        ///                 Key = "k1",
        ///                 Value = "v1",
        ///             },
        ///         },
        ///     });
        /// 
        ///     var fooFileSystems = Volcengine.Vepfs.FileSystems.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             fooFileSystem.Id,
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<FileSystemsResult> InvokeAsync(FileSystemsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<FileSystemsResult>("volcengine:vepfs/fileSystems:FileSystems", args ?? new FileSystemsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of vepfs file systems
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
        ///         CidrBlock = "172.16.0.0/16",
        ///     });
        /// 
        ///     var fooSubnet = new Volcengine.Vpc.Subnet("fooSubnet", new()
        ///     {
        ///         SubnetName = "acc-test-subnet",
        ///         CidrBlock = "172.16.0.0/24",
        ///         ZoneId = "cn-beijing-a",
        ///         VpcId = fooVpc.Id,
        ///     });
        /// 
        ///     var fooFileSystem = new Volcengine.Vepfs.FileSystem("fooFileSystem", new()
        ///     {
        ///         FileSystemName = "acc-test-file-system",
        ///         SubnetId = fooSubnet.Id,
        ///         StoreType = "Advance_100",
        ///         Description = "tf-test",
        ///         Capacity = 12,
        ///         Project = "default",
        ///         EnableRestripe = false,
        ///         Tags = new[]
        ///         {
        ///             new Volcengine.Vepfs.Inputs.FileSystemTagArgs
        ///             {
        ///                 Key = "k1",
        ///                 Value = "v1",
        ///             },
        ///         },
        ///     });
        /// 
        ///     var fooFileSystems = Volcengine.Vepfs.FileSystems.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             fooFileSystem.Id,
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<FileSystemsResult> Invoke(FileSystemsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<FileSystemsResult>("volcengine:vepfs/fileSystems:FileSystems", args ?? new FileSystemsInvokeArgs(), options.WithDefaults());
    }


    public sealed class FileSystemsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Name of Vepfs File System. This field support fuzzy query.
        /// </summary>
        [Input("fileSystemName")]
        public string? FileSystemName { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Vepfs File System IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

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
        /// The project of Vepfs File System.
        /// </summary>
        [Input("project")]
        public string? Project { get; set; }

        [Input("statuses")]
        private List<string>? _statuses;

        /// <summary>
        /// The query status list of Vepfs File System.
        /// </summary>
        public List<string> Statuses
        {
            get => _statuses ?? (_statuses = new List<string>());
            set => _statuses = value;
        }

        /// <summary>
        /// The Store Type of Vepfs File System.
        /// </summary>
        [Input("storeType")]
        public string? StoreType { get; set; }

        /// <summary>
        /// The zone id of File System.
        /// </summary>
        [Input("zoneId")]
        public string? ZoneId { get; set; }

        public FileSystemsArgs()
        {
        }
        public static new FileSystemsArgs Empty => new FileSystemsArgs();
    }

    public sealed class FileSystemsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Name of Vepfs File System. This field support fuzzy query.
        /// </summary>
        [Input("fileSystemName")]
        public Input<string>? FileSystemName { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Vepfs File System IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

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
        /// The project of Vepfs File System.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("statuses")]
        private InputList<string>? _statuses;

        /// <summary>
        /// The query status list of Vepfs File System.
        /// </summary>
        public InputList<string> Statuses
        {
            get => _statuses ?? (_statuses = new InputList<string>());
            set => _statuses = value;
        }

        /// <summary>
        /// The Store Type of Vepfs File System.
        /// </summary>
        [Input("storeType")]
        public Input<string>? StoreType { get; set; }

        /// <summary>
        /// The zone id of File System.
        /// </summary>
        [Input("zoneId")]
        public Input<string>? ZoneId { get; set; }

        public FileSystemsInvokeArgs()
        {
        }
        public static new FileSystemsInvokeArgs Empty => new FileSystemsInvokeArgs();
    }


    [OutputType]
    public sealed class FileSystemsResult
    {
        /// <summary>
        /// The name of the vepfs file system.
        /// </summary>
        public readonly string? FileSystemName;
        /// <summary>
        /// The collection of query.
        /// </summary>
        public readonly ImmutableArray<Outputs.FileSystemsFileSystemResult> FileSystems;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        public readonly string? OutputFile;
        /// <summary>
        /// The project name of the vepfs file system.
        /// </summary>
        public readonly string? Project;
        /// <summary>
        /// The status of the vepfs file system.
        /// </summary>
        public readonly ImmutableArray<string> Statuses;
        /// <summary>
        /// The store type of the vepfs file system.
        /// </summary>
        public readonly string? StoreType;
        /// <summary>
        /// The total count of query.
        /// </summary>
        public readonly int TotalCount;
        /// <summary>
        /// The id of the zone.
        /// </summary>
        public readonly string? ZoneId;

        [OutputConstructor]
        private FileSystemsResult(
            string? fileSystemName,

            ImmutableArray<Outputs.FileSystemsFileSystemResult> fileSystems,

            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            string? outputFile,

            string? project,

            ImmutableArray<string> statuses,

            string? storeType,

            int totalCount,

            string? zoneId)
        {
            FileSystemName = fileSystemName;
            FileSystems = fileSystems;
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            OutputFile = outputFile;
            Project = project;
            Statuses = statuses;
            StoreType = storeType;
            TotalCount = totalCount;
            ZoneId = zoneId;
        }
    }
}