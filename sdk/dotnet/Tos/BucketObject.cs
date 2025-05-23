// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Tos
{
    /// <summary>
    /// Provides a resource to manage tos object
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
    ///     var @default = new Volcengine.Tos.BucketObject("default", new()
    ///     {
    ///         AccountAcls = new[]
    ///         {
    ///             new Volcengine.Tos.Inputs.BucketObjectAccountAclArgs
    ///             {
    ///                 AccountId = "1",
    ///                 Permission = "READ",
    ///             },
    ///             new Volcengine.Tos.Inputs.BucketObjectAccountAclArgs
    ///             {
    ///                 AccountId = "2001",
    ///                 Permission = "WRITE_ACP",
    ///             },
    ///         },
    ///         BucketName = "tf-acc-test-bucket",
    ///         Encryption = "AES256",
    ///         FilePath = "/Users/bytedance/Work/Go/build/test.txt",
    ///         ObjectName = "tf-acc-test-object",
    ///         PublicAcl = "private",
    ///         Tags = new[]
    ///         {
    ///             new Volcengine.Tos.Inputs.BucketObjectTagArgs
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
    /// TOS Object can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import volcengine:tos/bucketObject:BucketObject default bucketName:objectName
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:tos/bucketObject:BucketObject")]
    public partial class BucketObject : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The user set of grant full control.
        /// </summary>
        [Output("accountAcls")]
        public Output<ImmutableArray<Outputs.BucketObjectAccountAcl>> AccountAcls { get; private set; } = null!;

        /// <summary>
        /// The name of the bucket.
        /// </summary>
        [Output("bucketName")]
        public Output<string> BucketName { get; private set; } = null!;

        /// <summary>
        /// The content of the TOS Object when content type is json or text and xml. Only one of `file_path,content` can be specified.
        /// </summary>
        [Output("content")]
        public Output<string> Content { get; private set; } = null!;

        /// <summary>
        /// The file md5 sum (32-bit hexadecimal string) for upload.
        /// </summary>
        [Output("contentMd5")]
        public Output<string?> ContentMd5 { get; private set; } = null!;

        /// <summary>
        /// The content type of the object.
        /// </summary>
        [Output("contentType")]
        public Output<string> ContentType { get; private set; } = null!;

        /// <summary>
        /// The flag of enable tos version.
        /// </summary>
        [Output("enableVersion")]
        public Output<bool> EnableVersion { get; private set; } = null!;

        /// <summary>
        /// The encryption of the object.Valid value is AES256.
        /// </summary>
        [Output("encryption")]
        public Output<string?> Encryption { get; private set; } = null!;

        /// <summary>
        /// The file path for upload. Only one of `file_path,content` can be specified.
        /// </summary>
        [Output("filePath")]
        public Output<string?> FilePath { get; private set; } = null!;

        /// <summary>
        /// Whether to enable the default inheritance bucket ACL function for the object.
        /// </summary>
        [Output("isDefault")]
        public Output<bool> IsDefault { get; private set; } = null!;

        /// <summary>
        /// The name of the object.
        /// </summary>
        [Output("objectName")]
        public Output<string> ObjectName { get; private set; } = null!;

        /// <summary>
        /// The public acl control of object. Valid value is private|public-read|public-read-write|authenticated-read|bucket-owner-read|default. `default` means to enable the default inheritance bucket ACL function for the object.
        /// </summary>
        [Output("publicAcl")]
        public Output<string?> PublicAcl { get; private set; } = null!;

        /// <summary>
        /// The storage type of the object.Valid value is STANDARD|IA.
        /// </summary>
        [Output("storageClass")]
        public Output<string?> StorageClass { get; private set; } = null!;

        /// <summary>
        /// Tos Bucket Tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.BucketObjectTag>> Tags { get; private set; } = null!;

        /// <summary>
        /// The version ids of the object if exist.
        /// </summary>
        [Output("versionIds")]
        public Output<ImmutableArray<string>> VersionIds { get; private set; } = null!;


        /// <summary>
        /// Create a BucketObject resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BucketObject(string name, BucketObjectArgs args, CustomResourceOptions? options = null)
            : base("volcengine:tos/bucketObject:BucketObject", name, args ?? new BucketObjectArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BucketObject(string name, Input<string> id, BucketObjectState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:tos/bucketObject:BucketObject", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BucketObject resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BucketObject Get(string name, Input<string> id, BucketObjectState? state = null, CustomResourceOptions? options = null)
        {
            return new BucketObject(name, id, state, options);
        }
    }

    public sealed class BucketObjectArgs : global::Pulumi.ResourceArgs
    {
        [Input("accountAcls")]
        private InputList<Inputs.BucketObjectAccountAclArgs>? _accountAcls;

        /// <summary>
        /// The user set of grant full control.
        /// </summary>
        public InputList<Inputs.BucketObjectAccountAclArgs> AccountAcls
        {
            get => _accountAcls ?? (_accountAcls = new InputList<Inputs.BucketObjectAccountAclArgs>());
            set => _accountAcls = value;
        }

        /// <summary>
        /// The name of the bucket.
        /// </summary>
        [Input("bucketName", required: true)]
        public Input<string> BucketName { get; set; } = null!;

        /// <summary>
        /// The content of the TOS Object when content type is json or text and xml. Only one of `file_path,content` can be specified.
        /// </summary>
        [Input("content")]
        public Input<string>? Content { get; set; }

        /// <summary>
        /// The file md5 sum (32-bit hexadecimal string) for upload.
        /// </summary>
        [Input("contentMd5")]
        public Input<string>? ContentMd5 { get; set; }

        /// <summary>
        /// The content type of the object.
        /// </summary>
        [Input("contentType")]
        public Input<string>? ContentType { get; set; }

        /// <summary>
        /// The encryption of the object.Valid value is AES256.
        /// </summary>
        [Input("encryption")]
        public Input<string>? Encryption { get; set; }

        /// <summary>
        /// The file path for upload. Only one of `file_path,content` can be specified.
        /// </summary>
        [Input("filePath")]
        public Input<string>? FilePath { get; set; }

        /// <summary>
        /// The name of the object.
        /// </summary>
        [Input("objectName", required: true)]
        public Input<string> ObjectName { get; set; } = null!;

        /// <summary>
        /// The public acl control of object. Valid value is private|public-read|public-read-write|authenticated-read|bucket-owner-read|default. `default` means to enable the default inheritance bucket ACL function for the object.
        /// </summary>
        [Input("publicAcl")]
        public Input<string>? PublicAcl { get; set; }

        /// <summary>
        /// The storage type of the object.Valid value is STANDARD|IA.
        /// </summary>
        [Input("storageClass")]
        public Input<string>? StorageClass { get; set; }

        [Input("tags")]
        private InputList<Inputs.BucketObjectTagArgs>? _tags;

        /// <summary>
        /// Tos Bucket Tags.
        /// </summary>
        public InputList<Inputs.BucketObjectTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.BucketObjectTagArgs>());
            set => _tags = value;
        }

        public BucketObjectArgs()
        {
        }
        public static new BucketObjectArgs Empty => new BucketObjectArgs();
    }

    public sealed class BucketObjectState : global::Pulumi.ResourceArgs
    {
        [Input("accountAcls")]
        private InputList<Inputs.BucketObjectAccountAclGetArgs>? _accountAcls;

        /// <summary>
        /// The user set of grant full control.
        /// </summary>
        public InputList<Inputs.BucketObjectAccountAclGetArgs> AccountAcls
        {
            get => _accountAcls ?? (_accountAcls = new InputList<Inputs.BucketObjectAccountAclGetArgs>());
            set => _accountAcls = value;
        }

        /// <summary>
        /// The name of the bucket.
        /// </summary>
        [Input("bucketName")]
        public Input<string>? BucketName { get; set; }

        /// <summary>
        /// The content of the TOS Object when content type is json or text and xml. Only one of `file_path,content` can be specified.
        /// </summary>
        [Input("content")]
        public Input<string>? Content { get; set; }

        /// <summary>
        /// The file md5 sum (32-bit hexadecimal string) for upload.
        /// </summary>
        [Input("contentMd5")]
        public Input<string>? ContentMd5 { get; set; }

        /// <summary>
        /// The content type of the object.
        /// </summary>
        [Input("contentType")]
        public Input<string>? ContentType { get; set; }

        /// <summary>
        /// The flag of enable tos version.
        /// </summary>
        [Input("enableVersion")]
        public Input<bool>? EnableVersion { get; set; }

        /// <summary>
        /// The encryption of the object.Valid value is AES256.
        /// </summary>
        [Input("encryption")]
        public Input<string>? Encryption { get; set; }

        /// <summary>
        /// The file path for upload. Only one of `file_path,content` can be specified.
        /// </summary>
        [Input("filePath")]
        public Input<string>? FilePath { get; set; }

        /// <summary>
        /// Whether to enable the default inheritance bucket ACL function for the object.
        /// </summary>
        [Input("isDefault")]
        public Input<bool>? IsDefault { get; set; }

        /// <summary>
        /// The name of the object.
        /// </summary>
        [Input("objectName")]
        public Input<string>? ObjectName { get; set; }

        /// <summary>
        /// The public acl control of object. Valid value is private|public-read|public-read-write|authenticated-read|bucket-owner-read|default. `default` means to enable the default inheritance bucket ACL function for the object.
        /// </summary>
        [Input("publicAcl")]
        public Input<string>? PublicAcl { get; set; }

        /// <summary>
        /// The storage type of the object.Valid value is STANDARD|IA.
        /// </summary>
        [Input("storageClass")]
        public Input<string>? StorageClass { get; set; }

        [Input("tags")]
        private InputList<Inputs.BucketObjectTagGetArgs>? _tags;

        /// <summary>
        /// Tos Bucket Tags.
        /// </summary>
        public InputList<Inputs.BucketObjectTagGetArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.BucketObjectTagGetArgs>());
            set => _tags = value;
        }

        [Input("versionIds")]
        private InputList<string>? _versionIds;

        /// <summary>
        /// The version ids of the object if exist.
        /// </summary>
        public InputList<string> VersionIds
        {
            get => _versionIds ?? (_versionIds = new InputList<string>());
            set => _versionIds = value;
        }

        public BucketObjectState()
        {
        }
        public static new BucketObjectState Empty => new BucketObjectState();
    }
}
