// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Cdn
{
    /// <summary>
    /// Provides a resource to manage cdn domain
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using System.Text.Json;
    /// using Pulumi;
    /// using Volcengine = Pulumi.Volcengine;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var fooCdnCertificate = new Volcengine.Cdn.CdnCertificate("fooCdnCertificate", new()
    ///     {
    ///         Certificate = "",
    ///         PrivateKey = "",
    ///         Desc = "tftest",
    ///         Source = "cdn_cert_hosting",
    ///     });
    /// 
    ///     var fooCdnDomain = new Volcengine.Cdn.CdnDomain("fooCdnDomain", new()
    ///     {
    ///         Domain = "tftest.byte-test.com",
    ///         ServiceType = "web",
    ///         Tags = new[]
    ///         {
    ///             new Volcengine.Cdn.Inputs.CdnDomainTagArgs
    ///             {
    ///                 Key = "tfkey1",
    ///                 Value = "tfvalue1",
    ///             },
    ///             new Volcengine.Cdn.Inputs.CdnDomainTagArgs
    ///             {
    ///                 Key = "tfkey2",
    ///                 Value = "tfvalue2",
    ///             },
    ///         },
    ///         DomainConfig = Output.JsonSerialize(Output.Create(new Dictionary&lt;string, object?&gt;
    ///         {
    ///             ["OriginProtocol"] = "https",
    ///             ["Origin"] = new[]
    ///             {
    ///                 new Dictionary&lt;string, object?&gt;
    ///                 {
    ///                     ["OriginAction"] = new Dictionary&lt;string, object?&gt;
    ///                     {
    ///                         ["OriginLines"] = new[]
    ///                         {
    ///                             new Dictionary&lt;string, object?&gt;
    ///                             {
    ///                                 ["Address"] = "1.1.1.1",
    ///                                 ["HttpPort"] = "80",
    ///                                 ["HttpsPort"] = "443",
    ///                                 ["InstanceType"] = "ip",
    ///                                 ["OriginType"] = "primary",
    ///                                 ["PrivateBucketAccess"] = false,
    ///                                 ["Weight"] = "2",
    ///                             },
    ///                         },
    ///                     },
    ///                 },
    ///             },
    ///             ["HTTPS"] = new Dictionary&lt;string, object?&gt;
    ///             {
    ///                 ["CertInfo"] = new Dictionary&lt;string, object?&gt;
    ///                 {
    ///                     ["CertId"] = fooCdnCertificate.Id,
    ///                 },
    ///                 ["DisableHttp"] = false,
    ///                 ["HTTP2"] = true,
    ///                 ["Switch"] = true,
    ///                 ["Ocsp"] = false,
    ///                 ["TlsVersion"] = new[]
    ///                 {
    ///                     "tlsv1.1",
    ///                     "tlsv1.2",
    ///                 },
    ///             },
    ///         })),
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// CdnDomain can be imported using the domain, e.g.
    /// 
    /// ```sh
    /// $ pulumi import volcengine:cdn/cdnDomain:CdnDomain default www.volcengine.com
    /// ```
    /// Please note that when you execute destroy, we will first take the domain name offline and then delete it.
    /// </summary>
    [VolcengineResourceType("volcengine:cdn/cdnDomain:CdnDomain")]
    public partial class CdnDomain : global::Pulumi.CustomResource
    {
        /// <summary>
        /// You need to add a domain. The main account can add up to 200 accelerated domains.
        /// </summary>
        [Output("domain")]
        public Output<string> Domain { get; private set; } = null!;

        /// <summary>
        /// Accelerate domain configuration. Please convert the configuration module structure into json and pass it into a string. You must specify the Origin module. The OriginProtocol parameter, OriginHost parameter, and other domain configuration modules are optional.
        /// </summary>
        [Output("domainConfig")]
        public Output<string> DomainConfig { get; private set; } = null!;

        /// <summary>
        /// The project to which this domain name belongs. Default is `default`.
        /// </summary>
        [Output("project")]
        public Output<string?> Project { get; private set; } = null!;

        /// <summary>
        /// Indicates the acceleration area. The parameter can take the following values: `chinese_mainland`: Indicates mainland China. `global`: Indicates global. `outside_chinese_mainland`: Indicates global (excluding mainland China).
        /// </summary>
        [Output("serviceRegion")]
        public Output<string> ServiceRegion { get; private set; } = null!;

        /// <summary>
        /// The business type of the domain name is indicated by this parameter. The possible values are: `download`: for file downloads. `web`: for web pages. `video`: for audio and video on demand.
        /// </summary>
        [Output("serviceType")]
        public Output<string> ServiceType { get; private set; } = null!;

        /// <summary>
        /// Configuration for sharing CNAME.
        /// </summary>
        [Output("sharedCname")]
        public Output<Outputs.CdnDomainSharedCname?> SharedCname { get; private set; } = null!;

        /// <summary>
        /// The status of the domain.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Indicate the tags you have set for this domain name. You can set up to 10 tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.CdnDomainTag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a CdnDomain resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CdnDomain(string name, CdnDomainArgs args, CustomResourceOptions? options = null)
            : base("volcengine:cdn/cdnDomain:CdnDomain", name, args ?? new CdnDomainArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CdnDomain(string name, Input<string> id, CdnDomainState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:cdn/cdnDomain:CdnDomain", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CdnDomain resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CdnDomain Get(string name, Input<string> id, CdnDomainState? state = null, CustomResourceOptions? options = null)
        {
            return new CdnDomain(name, id, state, options);
        }
    }

    public sealed class CdnDomainArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// You need to add a domain. The main account can add up to 200 accelerated domains.
        /// </summary>
        [Input("domain", required: true)]
        public Input<string> Domain { get; set; } = null!;

        /// <summary>
        /// Accelerate domain configuration. Please convert the configuration module structure into json and pass it into a string. You must specify the Origin module. The OriginProtocol parameter, OriginHost parameter, and other domain configuration modules are optional.
        /// </summary>
        [Input("domainConfig", required: true)]
        public Input<string> DomainConfig { get; set; } = null!;

        /// <summary>
        /// The project to which this domain name belongs. Default is `default`.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Indicates the acceleration area. The parameter can take the following values: `chinese_mainland`: Indicates mainland China. `global`: Indicates global. `outside_chinese_mainland`: Indicates global (excluding mainland China).
        /// </summary>
        [Input("serviceRegion")]
        public Input<string>? ServiceRegion { get; set; }

        /// <summary>
        /// The business type of the domain name is indicated by this parameter. The possible values are: `download`: for file downloads. `web`: for web pages. `video`: for audio and video on demand.
        /// </summary>
        [Input("serviceType", required: true)]
        public Input<string> ServiceType { get; set; } = null!;

        /// <summary>
        /// Configuration for sharing CNAME.
        /// </summary>
        [Input("sharedCname")]
        public Input<Inputs.CdnDomainSharedCnameArgs>? SharedCname { get; set; }

        [Input("tags")]
        private InputList<Inputs.CdnDomainTagArgs>? _tags;

        /// <summary>
        /// Indicate the tags you have set for this domain name. You can set up to 10 tags.
        /// </summary>
        public InputList<Inputs.CdnDomainTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.CdnDomainTagArgs>());
            set => _tags = value;
        }

        public CdnDomainArgs()
        {
        }
        public static new CdnDomainArgs Empty => new CdnDomainArgs();
    }

    public sealed class CdnDomainState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// You need to add a domain. The main account can add up to 200 accelerated domains.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// Accelerate domain configuration. Please convert the configuration module structure into json and pass it into a string. You must specify the Origin module. The OriginProtocol parameter, OriginHost parameter, and other domain configuration modules are optional.
        /// </summary>
        [Input("domainConfig")]
        public Input<string>? DomainConfig { get; set; }

        /// <summary>
        /// The project to which this domain name belongs. Default is `default`.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Indicates the acceleration area. The parameter can take the following values: `chinese_mainland`: Indicates mainland China. `global`: Indicates global. `outside_chinese_mainland`: Indicates global (excluding mainland China).
        /// </summary>
        [Input("serviceRegion")]
        public Input<string>? ServiceRegion { get; set; }

        /// <summary>
        /// The business type of the domain name is indicated by this parameter. The possible values are: `download`: for file downloads. `web`: for web pages. `video`: for audio and video on demand.
        /// </summary>
        [Input("serviceType")]
        public Input<string>? ServiceType { get; set; }

        /// <summary>
        /// Configuration for sharing CNAME.
        /// </summary>
        [Input("sharedCname")]
        public Input<Inputs.CdnDomainSharedCnameGetArgs>? SharedCname { get; set; }

        /// <summary>
        /// The status of the domain.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputList<Inputs.CdnDomainTagGetArgs>? _tags;

        /// <summary>
        /// Indicate the tags you have set for this domain name. You can set up to 10 tags.
        /// </summary>
        public InputList<Inputs.CdnDomainTagGetArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.CdnDomainTagGetArgs>());
            set => _tags = value;
        }

        public CdnDomainState()
        {
        }
        public static new CdnDomainState Empty => new CdnDomainState();
    }
}
