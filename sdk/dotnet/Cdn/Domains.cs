// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Cdn
{
    [Obsolete(@"volcengine.cdn.Domains has been deprecated in favor of volcengine.cdn.getDomains")]
    public static class Domains
    {
        /// <summary>
        /// Use this data source to query detailed information of cdn domains
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
        ///     var fooDomains = Volcengine.Cdn.GetDomains.Invoke(new()
        ///     {
        ///         Domain = fooCdnDomain.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<DomainsResult> InvokeAsync(DomainsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<DomainsResult>("volcengine:cdn/domains:Domains", args ?? new DomainsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of cdn domains
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
        ///     var fooDomains = Volcengine.Cdn.GetDomains.Invoke(new()
        ///     {
        ///         Domain = fooCdnDomain.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<DomainsResult> Invoke(DomainsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<DomainsResult>("volcengine:cdn/domains:Domains", args ?? new DomainsInvokeArgs(), options.WithDefaults());
    }


    public sealed class DomainsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Search by specifying domain name keywords, with fuzzy matching.
        /// </summary>
        [Input("domain")]
        public string? Domain { get; set; }

        /// <summary>
        /// Specify HTTPS configuration to filter accelerated domains. The optional values for this parameter are as follows: `true`: Indicates that the accelerated domain has enabled HTTPS function.`false`: Indicates that the accelerated domain has not enabled HTTPS function.
        /// </summary>
        [Input("https")]
        public bool? Https { get; set; }

        /// <summary>
        /// Specify IPv6 configuration to filter accelerated domain names. The optional values for this parameter are as follows: `true`: Indicates that the accelerated domain name supports requests using IPv6 addresses.`false`: Indicates that the accelerated domain name does not support requests using IPv6 addresses.
        /// </summary>
        [Input("ipv6")]
        public bool? Ipv6 { get; set; }

        /// <summary>
        /// Configure the origin protocol for the accelerated domain.
        /// </summary>
        [Input("originProtocol")]
        public string? OriginProtocol { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// Specify a primary origin server for filtering accelerated domains.
        /// </summary>
        [Input("primaryOrigin")]
        public string? PrimaryOrigin { get; set; }

        /// <summary>
        /// The project name of the domain.
        /// </summary>
        [Input("project")]
        public string? Project { get; set; }

        /// <summary>
        /// The business type of the domain name is indicated by this parameter. The possible values are: `download`: for file downloads. `web`: for web pages. `video`: for audio and video on demand.
        /// </summary>
        [Input("serviceType")]
        public string? ServiceType { get; set; }

        /// <summary>
        /// The status of the domain.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        [Input("tags")]
        private List<string>? _tags;

        /// <summary>
        /// Filter by specified domain name tags, up to 10 tags can be specified. Each tag is entered as a string in the format of key:value.
        /// </summary>
        public List<string> Tags
        {
            get => _tags ?? (_tags = new List<string>());
            set => _tags = value;
        }

        public DomainsArgs()
        {
        }
        public static new DomainsArgs Empty => new DomainsArgs();
    }

    public sealed class DomainsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Search by specifying domain name keywords, with fuzzy matching.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// Specify HTTPS configuration to filter accelerated domains. The optional values for this parameter are as follows: `true`: Indicates that the accelerated domain has enabled HTTPS function.`false`: Indicates that the accelerated domain has not enabled HTTPS function.
        /// </summary>
        [Input("https")]
        public Input<bool>? Https { get; set; }

        /// <summary>
        /// Specify IPv6 configuration to filter accelerated domain names. The optional values for this parameter are as follows: `true`: Indicates that the accelerated domain name supports requests using IPv6 addresses.`false`: Indicates that the accelerated domain name does not support requests using IPv6 addresses.
        /// </summary>
        [Input("ipv6")]
        public Input<bool>? Ipv6 { get; set; }

        /// <summary>
        /// Configure the origin protocol for the accelerated domain.
        /// </summary>
        [Input("originProtocol")]
        public Input<string>? OriginProtocol { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// Specify a primary origin server for filtering accelerated domains.
        /// </summary>
        [Input("primaryOrigin")]
        public Input<string>? PrimaryOrigin { get; set; }

        /// <summary>
        /// The project name of the domain.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The business type of the domain name is indicated by this parameter. The possible values are: `download`: for file downloads. `web`: for web pages. `video`: for audio and video on demand.
        /// </summary>
        [Input("serviceType")]
        public Input<string>? ServiceType { get; set; }

        /// <summary>
        /// The status of the domain.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// Filter by specified domain name tags, up to 10 tags can be specified. Each tag is entered as a string in the format of key:value.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        public DomainsInvokeArgs()
        {
        }
        public static new DomainsInvokeArgs Empty => new DomainsInvokeArgs();
    }


    [OutputType]
    public sealed class DomainsResult
    {
        /// <summary>
        /// Search by specifying domain name keywords, with fuzzy matching.
        /// </summary>
        public readonly string? Domain;
        /// <summary>
        /// The collection of query.
        /// </summary>
        public readonly ImmutableArray<Outputs.DomainsDomainResult> Domains;
        /// <summary>
        /// Specify HTTPS configuration to filter accelerated domains. The optional values for this parameter are as follows: `true`: Indicates that the accelerated domain has enabled HTTPS function.`false`: Indicates that the accelerated domain has not enabled HTTPS function.
        /// </summary>
        public readonly bool? Https;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Specify IPv6 configuration to filter accelerated domain names. The optional values for this parameter are as follows: `true`: Indicates that the accelerated domain name supports requests using IPv6 addresses.`false`: Indicates that the accelerated domain name does not support requests using IPv6 addresses.
        /// </summary>
        public readonly bool? Ipv6;
        /// <summary>
        /// Configure the origin protocol for the accelerated domain.
        /// </summary>
        public readonly string? OriginProtocol;
        public readonly string? OutputFile;
        /// <summary>
        /// List of primary source servers to accelerate the domain name.
        /// </summary>
        public readonly string? PrimaryOrigin;
        /// <summary>
        /// The project name of the domain.
        /// </summary>
        public readonly string? Project;
        /// <summary>
        /// The business type of the domain name is indicated by this parameter. The possible values are: `download`: for file downloads. `web`: for web pages. `video`: for audio and video on demand.
        /// </summary>
        public readonly string? ServiceType;
        /// <summary>
        /// The status of the domain.
        /// </summary>
        public readonly string? Status;
        /// <summary>
        /// Indicate the tags you have set for this domain name. You can set up to 10 tags.
        /// </summary>
        public readonly ImmutableArray<string> Tags;
        /// <summary>
        /// The total count of query.
        /// </summary>
        public readonly int TotalCount;

        [OutputConstructor]
        private DomainsResult(
            string? domain,

            ImmutableArray<Outputs.DomainsDomainResult> domains,

            bool? https,

            string id,

            bool? ipv6,

            string? originProtocol,

            string? outputFile,

            string? primaryOrigin,

            string? project,

            string? serviceType,

            string? status,

            ImmutableArray<string> tags,

            int totalCount)
        {
            Domain = domain;
            Domains = domains;
            Https = https;
            Id = id;
            Ipv6 = ipv6;
            OriginProtocol = originProtocol;
            OutputFile = outputFile;
            PrimaryOrigin = primaryOrigin;
            Project = project;
            ServiceType = serviceType;
            Status = status;
            Tags = tags;
            TotalCount = totalCount;
        }
    }
}
