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
    /// Provides a resource to manage tos bucket
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
    ///     var @default = new Volcengine.Tos.Bucket("default", new()
    ///     {
    ///         AccountAcls = new[]
    ///         {
    ///             new Volcengine.Tos.Inputs.BucketAccountAclArgs
    ///             {
    ///                 AccountId = "1",
    ///                 Permission = "READ",
    ///             },
    ///             new Volcengine.Tos.Inputs.BucketAccountAclArgs
    ///             {
    ///                 AccountId = "2001",
    ///                 Permission = "WRITE_ACP",
    ///             },
    ///         },
    ///         BucketName = "tf-acc-test-bucket",
    ///         EnableVersion = true,
    ///         ProjectName = "default",
    ///         PublicAcl = "private",
    ///         Tags = new[]
    ///         {
    ///             new Volcengine.Tos.Inputs.BucketTagArgs
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
    /// Tos Bucket can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import volcengine:tos/bucket:Bucket default bucketName
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:tos/bucket:Bucket")]
    public partial class Bucket : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The user set of grant full control.
        /// </summary>
        [Output("accountAcls")]
        public Output<ImmutableArray<Outputs.BucketAccountAcl>> AccountAcls { get; private set; } = null!;

        /// <summary>
        /// The name of the bucket.
        /// </summary>
        [Output("bucketName")]
        public Output<string> BucketName { get; private set; } = null!;

        /// <summary>
        /// The create date of the TOS bucket.
        /// </summary>
        [Output("creationDate")]
        public Output<string> CreationDate { get; private set; } = null!;

        /// <summary>
        /// The flag of enable tos version.
        /// </summary>
        [Output("enableVersion")]
        public Output<bool?> EnableVersion { get; private set; } = null!;

        /// <summary>
        /// The extranet endpoint of the TOS bucket.
        /// </summary>
        [Output("extranetEndpoint")]
        public Output<string> ExtranetEndpoint { get; private set; } = null!;

        /// <summary>
        /// The intranet endpoint the TOS bucket.
        /// </summary>
        [Output("intranetEndpoint")]
        public Output<string> IntranetEndpoint { get; private set; } = null!;

        /// <summary>
        /// The location of the TOS bucket.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The ProjectName of the Tos Bucket. Default is `default`.
        /// </summary>
        [Output("projectName")]
        public Output<string?> ProjectName { get; private set; } = null!;

        /// <summary>
        /// The public acl control of object.Valid value is private|public-read|public-read-write|authenticated-read|bucket-owner-read.
        /// </summary>
        [Output("publicAcl")]
        public Output<string?> PublicAcl { get; private set; } = null!;

        /// <summary>
        /// The storage type of the object.Valid value is STANDARD|IA|ARCHIVE_FR.Default is STANDARD.
        /// </summary>
        [Output("storageClass")]
        public Output<string?> StorageClass { get; private set; } = null!;

        /// <summary>
        /// Tos Bucket Tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.BucketTag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Bucket resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Bucket(string name, BucketArgs args, CustomResourceOptions? options = null)
            : base("volcengine:tos/bucket:Bucket", name, args ?? new BucketArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Bucket(string name, Input<string> id, BucketState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:tos/bucket:Bucket", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Bucket resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Bucket Get(string name, Input<string> id, BucketState? state = null, CustomResourceOptions? options = null)
        {
            return new Bucket(name, id, state, options);
        }
    }

    public sealed class BucketArgs : global::Pulumi.ResourceArgs
    {
        [Input("accountAcls")]
        private InputList<Inputs.BucketAccountAclArgs>? _accountAcls;

        /// <summary>
        /// The user set of grant full control.
        /// </summary>
        public InputList<Inputs.BucketAccountAclArgs> AccountAcls
        {
            get => _accountAcls ?? (_accountAcls = new InputList<Inputs.BucketAccountAclArgs>());
            set => _accountAcls = value;
        }

        /// <summary>
        /// The name of the bucket.
        /// </summary>
        [Input("bucketName", required: true)]
        public Input<string> BucketName { get; set; } = null!;

        /// <summary>
        /// The flag of enable tos version.
        /// </summary>
        [Input("enableVersion")]
        public Input<bool>? EnableVersion { get; set; }

        /// <summary>
        /// The ProjectName of the Tos Bucket. Default is `default`.
        /// </summary>
        [Input("projectName")]
        public Input<string>? ProjectName { get; set; }

        /// <summary>
        /// The public acl control of object.Valid value is private|public-read|public-read-write|authenticated-read|bucket-owner-read.
        /// </summary>
        [Input("publicAcl")]
        public Input<string>? PublicAcl { get; set; }

        /// <summary>
        /// The storage type of the object.Valid value is STANDARD|IA|ARCHIVE_FR.Default is STANDARD.
        /// </summary>
        [Input("storageClass")]
        public Input<string>? StorageClass { get; set; }

        [Input("tags")]
        private InputList<Inputs.BucketTagArgs>? _tags;

        /// <summary>
        /// Tos Bucket Tags.
        /// </summary>
        public InputList<Inputs.BucketTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.BucketTagArgs>());
            set => _tags = value;
        }

        public BucketArgs()
        {
        }
        public static new BucketArgs Empty => new BucketArgs();
    }

    public sealed class BucketState : global::Pulumi.ResourceArgs
    {
        [Input("accountAcls")]
        private InputList<Inputs.BucketAccountAclGetArgs>? _accountAcls;

        /// <summary>
        /// The user set of grant full control.
        /// </summary>
        public InputList<Inputs.BucketAccountAclGetArgs> AccountAcls
        {
            get => _accountAcls ?? (_accountAcls = new InputList<Inputs.BucketAccountAclGetArgs>());
            set => _accountAcls = value;
        }

        /// <summary>
        /// The name of the bucket.
        /// </summary>
        [Input("bucketName")]
        public Input<string>? BucketName { get; set; }

        /// <summary>
        /// The create date of the TOS bucket.
        /// </summary>
        [Input("creationDate")]
        public Input<string>? CreationDate { get; set; }

        /// <summary>
        /// The flag of enable tos version.
        /// </summary>
        [Input("enableVersion")]
        public Input<bool>? EnableVersion { get; set; }

        /// <summary>
        /// The extranet endpoint of the TOS bucket.
        /// </summary>
        [Input("extranetEndpoint")]
        public Input<string>? ExtranetEndpoint { get; set; }

        /// <summary>
        /// The intranet endpoint the TOS bucket.
        /// </summary>
        [Input("intranetEndpoint")]
        public Input<string>? IntranetEndpoint { get; set; }

        /// <summary>
        /// The location of the TOS bucket.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The ProjectName of the Tos Bucket. Default is `default`.
        /// </summary>
        [Input("projectName")]
        public Input<string>? ProjectName { get; set; }

        /// <summary>
        /// The public acl control of object.Valid value is private|public-read|public-read-write|authenticated-read|bucket-owner-read.
        /// </summary>
        [Input("publicAcl")]
        public Input<string>? PublicAcl { get; set; }

        /// <summary>
        /// The storage type of the object.Valid value is STANDARD|IA|ARCHIVE_FR.Default is STANDARD.
        /// </summary>
        [Input("storageClass")]
        public Input<string>? StorageClass { get; set; }

        [Input("tags")]
        private InputList<Inputs.BucketTagGetArgs>? _tags;

        /// <summary>
        /// Tos Bucket Tags.
        /// </summary>
        public InputList<Inputs.BucketTagGetArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.BucketTagGetArgs>());
            set => _tags = value;
        }

        public BucketState()
        {
        }
        public static new BucketState Empty => new BucketState();
    }
}