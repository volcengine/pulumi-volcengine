// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Ecs
{
    /// <summary>
    /// Provides a resource to manage image import
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
    ///     var foo = new Volcengine.Ecs.ImageImport("foo", new()
    ///     {
    ///         BootMode = "UEFI",
    ///         Description = "acc-test",
    ///         ImageName = "acc-test-image",
    ///         Platform = "CentOS",
    ///         ProjectName = "default",
    ///         Tags = new[]
    ///         {
    ///             new Volcengine.Ecs.Inputs.ImageImportTagArgs
    ///             {
    ///                 Key = "k1",
    ///                 Value = "v1",
    ///             },
    ///         },
    ///         Url = "https://*****_system.qcow2",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ImageImport can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import volcengine:ecs/imageImport:ImageImport default resource_id
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:ecs/imageImport:ImageImport")]
    public partial class ImageImport : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The architecture of the custom image. Valid values: `amd64`, `arm64`.
        /// </summary>
        [Output("architecture")]
        public Output<string> Architecture { get; private set; } = null!;

        /// <summary>
        /// The boot mode of the custom image. Valid values: `BIOS`, `UEFI`.
        /// </summary>
        [Output("bootMode")]
        public Output<string> BootMode { get; private set; } = null!;

        /// <summary>
        /// The create time of Image.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// The description of the custom image.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The name of the custom image.
        /// </summary>
        [Output("imageName")]
        public Output<string> ImageName { get; private set; } = null!;

        /// <summary>
        /// Whether the Image support cloud-init.
        /// </summary>
        [Output("isSupportCloudInit")]
        public Output<bool> IsSupportCloudInit { get; private set; } = null!;

        /// <summary>
        /// The license type of the custom image. Valid values: `VolcanoEngine`.
        /// </summary>
        [Output("licenseType")]
        public Output<string> LicenseType { get; private set; } = null!;

        /// <summary>
        /// The name of Image operating system.
        /// </summary>
        [Output("osName")]
        public Output<string> OsName { get; private set; } = null!;

        /// <summary>
        /// The os type of the custom image. Valid values: `linux`, `Windows`.
        /// </summary>
        [Output("osType")]
        public Output<string> OsType { get; private set; } = null!;

        /// <summary>
        /// The platform of the custom image. Valid values: `CentOS`, `Debian`, `veLinux`, `Windows Server`, `Fedora`, `OpenSUSE`, `Ubuntu`, `Rocky Linux`, `AlmaLinux`.
        /// </summary>
        [Output("platform")]
        public Output<string> Platform { get; private set; } = null!;

        /// <summary>
        /// The platform version of the custom image.
        /// </summary>
        [Output("platformVersion")]
        public Output<string> PlatformVersion { get; private set; } = null!;

        /// <summary>
        /// The project name of the custom image.
        /// </summary>
        [Output("projectName")]
        public Output<string> ProjectName { get; private set; } = null!;

        /// <summary>
        /// The share mode of Image.
        /// </summary>
        [Output("shareStatus")]
        public Output<string> ShareStatus { get; private set; } = null!;

        /// <summary>
        /// The size(GiB) of Image.
        /// </summary>
        [Output("size")]
        public Output<int> Size { get; private set; } = null!;

        /// <summary>
        /// The status of Image.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.ImageImportTag>> Tags { get; private set; } = null!;

        /// <summary>
        /// The update time of Image.
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;

        /// <summary>
        /// The url of the custom image in tos bucket.When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;

        /// <summary>
        /// The visibility of Image.
        /// </summary>
        [Output("visibility")]
        public Output<string> Visibility { get; private set; } = null!;


        /// <summary>
        /// Create a ImageImport resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ImageImport(string name, ImageImportArgs args, CustomResourceOptions? options = null)
            : base("volcengine:ecs/imageImport:ImageImport", name, args ?? new ImageImportArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ImageImport(string name, Input<string> id, ImageImportState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:ecs/imageImport:ImageImport", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ImageImport resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ImageImport Get(string name, Input<string> id, ImageImportState? state = null, CustomResourceOptions? options = null)
        {
            return new ImageImport(name, id, state, options);
        }
    }

    public sealed class ImageImportArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The architecture of the custom image. Valid values: `amd64`, `arm64`.
        /// </summary>
        [Input("architecture")]
        public Input<string>? Architecture { get; set; }

        /// <summary>
        /// The boot mode of the custom image. Valid values: `BIOS`, `UEFI`.
        /// </summary>
        [Input("bootMode")]
        public Input<string>? BootMode { get; set; }

        /// <summary>
        /// The description of the custom image.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the custom image.
        /// </summary>
        [Input("imageName", required: true)]
        public Input<string> ImageName { get; set; } = null!;

        /// <summary>
        /// The license type of the custom image. Valid values: `VolcanoEngine`.
        /// </summary>
        [Input("licenseType")]
        public Input<string>? LicenseType { get; set; }

        /// <summary>
        /// The os type of the custom image. Valid values: `linux`, `Windows`.
        /// </summary>
        [Input("osType")]
        public Input<string>? OsType { get; set; }

        /// <summary>
        /// The platform of the custom image. Valid values: `CentOS`, `Debian`, `veLinux`, `Windows Server`, `Fedora`, `OpenSUSE`, `Ubuntu`, `Rocky Linux`, `AlmaLinux`.
        /// </summary>
        [Input("platform", required: true)]
        public Input<string> Platform { get; set; } = null!;

        /// <summary>
        /// The platform version of the custom image.
        /// </summary>
        [Input("platformVersion")]
        public Input<string>? PlatformVersion { get; set; }

        /// <summary>
        /// The project name of the custom image.
        /// </summary>
        [Input("projectName")]
        public Input<string>? ProjectName { get; set; }

        [Input("tags")]
        private InputList<Inputs.ImageImportTagArgs>? _tags;

        /// <summary>
        /// Tags.
        /// </summary>
        public InputList<Inputs.ImageImportTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.ImageImportTagArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// The url of the custom image in tos bucket.When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        /// </summary>
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        public ImageImportArgs()
        {
        }
        public static new ImageImportArgs Empty => new ImageImportArgs();
    }

    public sealed class ImageImportState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The architecture of the custom image. Valid values: `amd64`, `arm64`.
        /// </summary>
        [Input("architecture")]
        public Input<string>? Architecture { get; set; }

        /// <summary>
        /// The boot mode of the custom image. Valid values: `BIOS`, `UEFI`.
        /// </summary>
        [Input("bootMode")]
        public Input<string>? BootMode { get; set; }

        /// <summary>
        /// The create time of Image.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// The description of the custom image.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the custom image.
        /// </summary>
        [Input("imageName")]
        public Input<string>? ImageName { get; set; }

        /// <summary>
        /// Whether the Image support cloud-init.
        /// </summary>
        [Input("isSupportCloudInit")]
        public Input<bool>? IsSupportCloudInit { get; set; }

        /// <summary>
        /// The license type of the custom image. Valid values: `VolcanoEngine`.
        /// </summary>
        [Input("licenseType")]
        public Input<string>? LicenseType { get; set; }

        /// <summary>
        /// The name of Image operating system.
        /// </summary>
        [Input("osName")]
        public Input<string>? OsName { get; set; }

        /// <summary>
        /// The os type of the custom image. Valid values: `linux`, `Windows`.
        /// </summary>
        [Input("osType")]
        public Input<string>? OsType { get; set; }

        /// <summary>
        /// The platform of the custom image. Valid values: `CentOS`, `Debian`, `veLinux`, `Windows Server`, `Fedora`, `OpenSUSE`, `Ubuntu`, `Rocky Linux`, `AlmaLinux`.
        /// </summary>
        [Input("platform")]
        public Input<string>? Platform { get; set; }

        /// <summary>
        /// The platform version of the custom image.
        /// </summary>
        [Input("platformVersion")]
        public Input<string>? PlatformVersion { get; set; }

        /// <summary>
        /// The project name of the custom image.
        /// </summary>
        [Input("projectName")]
        public Input<string>? ProjectName { get; set; }

        /// <summary>
        /// The share mode of Image.
        /// </summary>
        [Input("shareStatus")]
        public Input<string>? ShareStatus { get; set; }

        /// <summary>
        /// The size(GiB) of Image.
        /// </summary>
        [Input("size")]
        public Input<int>? Size { get; set; }

        /// <summary>
        /// The status of Image.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputList<Inputs.ImageImportTagGetArgs>? _tags;

        /// <summary>
        /// Tags.
        /// </summary>
        public InputList<Inputs.ImageImportTagGetArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.ImageImportTagGetArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// The update time of Image.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        /// <summary>
        /// The url of the custom image in tos bucket.When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        /// <summary>
        /// The visibility of Image.
        /// </summary>
        [Input("visibility")]
        public Input<string>? Visibility { get; set; }

        public ImageImportState()
        {
        }
        public static new ImageImportState Empty => new ImageImportState();
    }
}