// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Vpn
{
    [Obsolete(@"volcengine.vpn.SslVpnClientCerts has been deprecated in favor of volcengine.vpn.getSslVpnClientCerts")]
    public static class SslVpnClientCerts
    {
        /// <summary>
        /// Use this data source to query detailed information of ssl vpn client certs
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
        ///     var fooZones = Volcengine.Ecs.GetZones.Invoke();
        /// 
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
        ///         ZoneId = fooZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
        ///         VpcId = fooVpc.Id,
        ///     });
        /// 
        ///     var fooGateway = new Volcengine.Vpn.Gateway("fooGateway", new()
        ///     {
        ///         VpcId = fooVpc.Id,
        ///         SubnetId = fooSubnet.Id,
        ///         Bandwidth = 5,
        ///         VpnGatewayName = "acc-test1",
        ///         Description = "acc-test1",
        ///         Period = 7,
        ///         ProjectName = "default",
        ///         SslEnabled = true,
        ///         SslMaxConnections = 5,
        ///     });
        /// 
        ///     var fooSslVpnServer = new Volcengine.Vpn.SslVpnServer("fooSslVpnServer", new()
        ///     {
        ///         VpnGatewayId = fooGateway.Id,
        ///         LocalSubnets = new[]
        ///         {
        ///             fooSubnet.CidrBlock,
        ///         },
        ///         ClientIpPool = "172.16.2.0/24",
        ///         SslVpnServerName = "acc-test-ssl",
        ///         Description = "acc-test",
        ///         Protocol = "UDP",
        ///         Cipher = "AES-128-CBC",
        ///         Auth = "SHA1",
        ///         Compress = true,
        ///     });
        /// 
        ///     var fooSslVpnClientCert = new List&lt;Volcengine.Vpn.SslVpnClientCert&gt;();
        ///     for (var rangeIndex = 0; rangeIndex &lt; 5; rangeIndex++)
        ///     {
        ///         var range = new { Value = rangeIndex };
        ///         fooSslVpnClientCert.Add(new Volcengine.Vpn.SslVpnClientCert($"fooSslVpnClientCert-{range.Value}", new()
        ///         {
        ///             SslVpnServerId = fooSslVpnServer.Id,
        ///             SslVpnClientCertName = $"acc-test-client-cert-{range.Value}",
        ///             Description = "acc-test",
        ///         }));
        ///     }
        ///     var fooSslVpnClientCerts = Volcengine.Vpn.GetSslVpnClientCerts.Invoke(new()
        ///     {
        ///         Ids = fooSslVpnClientCert.Select(__item =&gt; __item.Id).ToList(),
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<SslVpnClientCertsResult> InvokeAsync(SslVpnClientCertsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<SslVpnClientCertsResult>("volcengine:vpn/sslVpnClientCerts:SslVpnClientCerts", args ?? new SslVpnClientCertsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of ssl vpn client certs
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
        ///     var fooZones = Volcengine.Ecs.GetZones.Invoke();
        /// 
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
        ///         ZoneId = fooZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
        ///         VpcId = fooVpc.Id,
        ///     });
        /// 
        ///     var fooGateway = new Volcengine.Vpn.Gateway("fooGateway", new()
        ///     {
        ///         VpcId = fooVpc.Id,
        ///         SubnetId = fooSubnet.Id,
        ///         Bandwidth = 5,
        ///         VpnGatewayName = "acc-test1",
        ///         Description = "acc-test1",
        ///         Period = 7,
        ///         ProjectName = "default",
        ///         SslEnabled = true,
        ///         SslMaxConnections = 5,
        ///     });
        /// 
        ///     var fooSslVpnServer = new Volcengine.Vpn.SslVpnServer("fooSslVpnServer", new()
        ///     {
        ///         VpnGatewayId = fooGateway.Id,
        ///         LocalSubnets = new[]
        ///         {
        ///             fooSubnet.CidrBlock,
        ///         },
        ///         ClientIpPool = "172.16.2.0/24",
        ///         SslVpnServerName = "acc-test-ssl",
        ///         Description = "acc-test",
        ///         Protocol = "UDP",
        ///         Cipher = "AES-128-CBC",
        ///         Auth = "SHA1",
        ///         Compress = true,
        ///     });
        /// 
        ///     var fooSslVpnClientCert = new List&lt;Volcengine.Vpn.SslVpnClientCert&gt;();
        ///     for (var rangeIndex = 0; rangeIndex &lt; 5; rangeIndex++)
        ///     {
        ///         var range = new { Value = rangeIndex };
        ///         fooSslVpnClientCert.Add(new Volcengine.Vpn.SslVpnClientCert($"fooSslVpnClientCert-{range.Value}", new()
        ///         {
        ///             SslVpnServerId = fooSslVpnServer.Id,
        ///             SslVpnClientCertName = $"acc-test-client-cert-{range.Value}",
        ///             Description = "acc-test",
        ///         }));
        ///     }
        ///     var fooSslVpnClientCerts = Volcengine.Vpn.GetSslVpnClientCerts.Invoke(new()
        ///     {
        ///         Ids = fooSslVpnClientCert.Select(__item =&gt; __item.Id).ToList(),
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<SslVpnClientCertsResult> Invoke(SslVpnClientCertsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<SslVpnClientCertsResult>("volcengine:vpn/sslVpnClientCerts:SslVpnClientCerts", args ?? new SslVpnClientCertsInvokeArgs(), options.WithDefaults());
    }


    public sealed class SslVpnClientCertsArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// The ids list of ssl vpn client cert.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A Name Regex of ssl vpn client cert.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The name of the ssl vpn client cert.
        /// </summary>
        [Input("sslVpnClientCertName")]
        public string? SslVpnClientCertName { get; set; }

        /// <summary>
        /// The id of the ssl vpn server.
        /// </summary>
        [Input("sslVpnServerId")]
        public string? SslVpnServerId { get; set; }

        public SslVpnClientCertsArgs()
        {
        }
        public static new SslVpnClientCertsArgs Empty => new SslVpnClientCertsArgs();
    }

    public sealed class SslVpnClientCertsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// The ids list of ssl vpn client cert.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A Name Regex of ssl vpn client cert.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// The name of the ssl vpn client cert.
        /// </summary>
        [Input("sslVpnClientCertName")]
        public Input<string>? SslVpnClientCertName { get; set; }

        /// <summary>
        /// The id of the ssl vpn server.
        /// </summary>
        [Input("sslVpnServerId")]
        public Input<string>? SslVpnServerId { get; set; }

        public SslVpnClientCertsInvokeArgs()
        {
        }
        public static new SslVpnClientCertsInvokeArgs Empty => new SslVpnClientCertsInvokeArgs();
    }


    [OutputType]
    public sealed class SslVpnClientCertsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        public readonly string? OutputFile;
        /// <summary>
        /// The name of the ssl vpn client cert.
        /// </summary>
        public readonly string? SslVpnClientCertName;
        /// <summary>
        /// The collection of of ssl vpn client certs.
        /// </summary>
        public readonly ImmutableArray<Outputs.SslVpnClientCertsSslVpnClientCertResult> SslVpnClientCerts;
        /// <summary>
        /// The id of the ssl vpn server.
        /// </summary>
        public readonly string? SslVpnServerId;
        /// <summary>
        /// The total count of ssl vpn client cert query.
        /// </summary>
        public readonly int TotalCount;

        [OutputConstructor]
        private SslVpnClientCertsResult(
            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            string? outputFile,

            string? sslVpnClientCertName,

            ImmutableArray<Outputs.SslVpnClientCertsSslVpnClientCertResult> sslVpnClientCerts,

            string? sslVpnServerId,

            int totalCount)
        {
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            OutputFile = outputFile;
            SslVpnClientCertName = sslVpnClientCertName;
            SslVpnClientCerts = sslVpnClientCerts;
            SslVpnServerId = sslVpnServerId;
            TotalCount = totalCount;
        }
    }
}
