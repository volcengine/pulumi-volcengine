// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Vepfs
{
    /// <summary>
    /// Provides a resource to manage vepfs file system
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
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// VepfsFileSystem can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import volcengine:vepfs/fileSystem:FileSystem default resource_id
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:vepfs/fileSystem:FileSystem")]
    public partial class FileSystem : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The id of the account.
        /// </summary>
        [Output("accountId")]
        public Output<string> AccountId { get; private set; } = null!;

        /// <summary>
        /// The bandwidth info of the vepfs file system.
        /// </summary>
        [Output("bandwidth")]
        public Output<int> Bandwidth { get; private set; } = null!;

        /// <summary>
        /// The capacity of the vepfs file system.
        /// </summary>
        [Output("capacity")]
        public Output<int> Capacity { get; private set; } = null!;

        /// <summary>
        /// The charge status of the vepfs file system.
        /// </summary>
        [Output("chargeStatus")]
        public Output<string> ChargeStatus { get; private set; } = null!;

        /// <summary>
        /// The charge type of the vepfs file system.
        /// </summary>
        [Output("chargeType")]
        public Output<string> ChargeType { get; private set; } = null!;

        /// <summary>
        /// The create time of the vepfs file system.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// The description info of the vepfs file system.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Whether to enable data balance after capacity expansion. This filed is valid only when expanding capacity.
        /// </summary>
        [Output("enableRestripe")]
        public Output<bool?> EnableRestripe { get; private set; } = null!;

        /// <summary>
        /// The expire time of the vepfs file system.
        /// </summary>
        [Output("expireTime")]
        public Output<string> ExpireTime { get; private set; } = null!;

        /// <summary>
        /// The name of the vepfs file system.
        /// </summary>
        [Output("fileSystemName")]
        public Output<string> FileSystemName { get; private set; } = null!;

        /// <summary>
        /// The type of the vepfs file system.
        /// </summary>
        [Output("fileSystemType")]
        public Output<string> FileSystemType { get; private set; } = null!;

        /// <summary>
        /// The free time of the vepfs file system.
        /// </summary>
        [Output("freeTime")]
        public Output<string> FreeTime { get; private set; } = null!;

        /// <summary>
        /// The last modify time of the vepfs file system.
        /// </summary>
        [Output("lastModifyTime")]
        public Output<string> LastModifyTime { get; private set; } = null!;

        /// <summary>
        /// The project of the vepfs file system.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The protocol type of the vepfs file system.
        /// </summary>
        [Output("protocolType")]
        public Output<string> ProtocolType { get; private set; } = null!;

        /// <summary>
        /// The id of the region.
        /// </summary>
        [Output("regionId")]
        public Output<string> RegionId { get; private set; } = null!;

        /// <summary>
        /// The status of the vepfs file system.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The stop service time of the vepfs file system.
        /// </summary>
        [Output("stopServiceTime")]
        public Output<string> StopServiceTime { get; private set; } = null!;

        /// <summary>
        /// The store type of the vepfs file system. Valid values: `Advance_100`, `Performance`, `Intelligent_Computing`.
        /// </summary>
        [Output("storeType")]
        public Output<string> StoreType { get; private set; } = null!;

        /// <summary>
        /// The store type cn name of the vepfs file system.
        /// </summary>
        [Output("storeTypeCn")]
        public Output<string> StoreTypeCn { get; private set; } = null!;

        /// <summary>
        /// The subnet id of the vepfs file system.
        /// </summary>
        [Output("subnetId")]
        public Output<string> SubnetId { get; private set; } = null!;

        /// <summary>
        /// Tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.FileSystemTag>> Tags { get; private set; } = null!;

        /// <summary>
        /// The version info of the vepfs file system.
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;

        /// <summary>
        /// The id of the zone.
        /// </summary>
        [Output("zoneId")]
        public Output<string> ZoneId { get; private set; } = null!;

        /// <summary>
        /// The name of the zone.
        /// </summary>
        [Output("zoneName")]
        public Output<string> ZoneName { get; private set; } = null!;


        /// <summary>
        /// Create a FileSystem resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FileSystem(string name, FileSystemArgs args, CustomResourceOptions? options = null)
            : base("volcengine:vepfs/fileSystem:FileSystem", name, args ?? new FileSystemArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FileSystem(string name, Input<string> id, FileSystemState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:vepfs/fileSystem:FileSystem", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FileSystem resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FileSystem Get(string name, Input<string> id, FileSystemState? state = null, CustomResourceOptions? options = null)
        {
            return new FileSystem(name, id, state, options);
        }
    }

    public sealed class FileSystemArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The capacity of the vepfs file system.
        /// </summary>
        [Input("capacity", required: true)]
        public Input<int> Capacity { get; set; } = null!;

        /// <summary>
        /// The description info of the vepfs file system.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether to enable data balance after capacity expansion. This filed is valid only when expanding capacity.
        /// </summary>
        [Input("enableRestripe")]
        public Input<bool>? EnableRestripe { get; set; }

        /// <summary>
        /// The name of the vepfs file system.
        /// </summary>
        [Input("fileSystemName", required: true)]
        public Input<string> FileSystemName { get; set; } = null!;

        /// <summary>
        /// The project of the vepfs file system.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The store type of the vepfs file system. Valid values: `Advance_100`, `Performance`, `Intelligent_Computing`.
        /// </summary>
        [Input("storeType", required: true)]
        public Input<string> StoreType { get; set; } = null!;

        /// <summary>
        /// The subnet id of the vepfs file system.
        /// </summary>
        [Input("subnetId", required: true)]
        public Input<string> SubnetId { get; set; } = null!;

        [Input("tags")]
        private InputList<Inputs.FileSystemTagArgs>? _tags;

        /// <summary>
        /// Tags.
        /// </summary>
        public InputList<Inputs.FileSystemTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.FileSystemTagArgs>());
            set => _tags = value;
        }

        public FileSystemArgs()
        {
        }
        public static new FileSystemArgs Empty => new FileSystemArgs();
    }

    public sealed class FileSystemState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The id of the account.
        /// </summary>
        [Input("accountId")]
        public Input<string>? AccountId { get; set; }

        /// <summary>
        /// The bandwidth info of the vepfs file system.
        /// </summary>
        [Input("bandwidth")]
        public Input<int>? Bandwidth { get; set; }

        /// <summary>
        /// The capacity of the vepfs file system.
        /// </summary>
        [Input("capacity")]
        public Input<int>? Capacity { get; set; }

        /// <summary>
        /// The charge status of the vepfs file system.
        /// </summary>
        [Input("chargeStatus")]
        public Input<string>? ChargeStatus { get; set; }

        /// <summary>
        /// The charge type of the vepfs file system.
        /// </summary>
        [Input("chargeType")]
        public Input<string>? ChargeType { get; set; }

        /// <summary>
        /// The create time of the vepfs file system.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// The description info of the vepfs file system.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether to enable data balance after capacity expansion. This filed is valid only when expanding capacity.
        /// </summary>
        [Input("enableRestripe")]
        public Input<bool>? EnableRestripe { get; set; }

        /// <summary>
        /// The expire time of the vepfs file system.
        /// </summary>
        [Input("expireTime")]
        public Input<string>? ExpireTime { get; set; }

        /// <summary>
        /// The name of the vepfs file system.
        /// </summary>
        [Input("fileSystemName")]
        public Input<string>? FileSystemName { get; set; }

        /// <summary>
        /// The type of the vepfs file system.
        /// </summary>
        [Input("fileSystemType")]
        public Input<string>? FileSystemType { get; set; }

        /// <summary>
        /// The free time of the vepfs file system.
        /// </summary>
        [Input("freeTime")]
        public Input<string>? FreeTime { get; set; }

        /// <summary>
        /// The last modify time of the vepfs file system.
        /// </summary>
        [Input("lastModifyTime")]
        public Input<string>? LastModifyTime { get; set; }

        /// <summary>
        /// The project of the vepfs file system.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The protocol type of the vepfs file system.
        /// </summary>
        [Input("protocolType")]
        public Input<string>? ProtocolType { get; set; }

        /// <summary>
        /// The id of the region.
        /// </summary>
        [Input("regionId")]
        public Input<string>? RegionId { get; set; }

        /// <summary>
        /// The status of the vepfs file system.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The stop service time of the vepfs file system.
        /// </summary>
        [Input("stopServiceTime")]
        public Input<string>? StopServiceTime { get; set; }

        /// <summary>
        /// The store type of the vepfs file system. Valid values: `Advance_100`, `Performance`, `Intelligent_Computing`.
        /// </summary>
        [Input("storeType")]
        public Input<string>? StoreType { get; set; }

        /// <summary>
        /// The store type cn name of the vepfs file system.
        /// </summary>
        [Input("storeTypeCn")]
        public Input<string>? StoreTypeCn { get; set; }

        /// <summary>
        /// The subnet id of the vepfs file system.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        [Input("tags")]
        private InputList<Inputs.FileSystemTagGetArgs>? _tags;

        /// <summary>
        /// Tags.
        /// </summary>
        public InputList<Inputs.FileSystemTagGetArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.FileSystemTagGetArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// The version info of the vepfs file system.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        /// <summary>
        /// The id of the zone.
        /// </summary>
        [Input("zoneId")]
        public Input<string>? ZoneId { get; set; }

        /// <summary>
        /// The name of the zone.
        /// </summary>
        [Input("zoneName")]
        public Input<string>? ZoneName { get; set; }

        public FileSystemState()
        {
        }
        public static new FileSystemState Empty => new FileSystemState();
    }
}