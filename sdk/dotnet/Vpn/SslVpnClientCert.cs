// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Vpn
{
    /// <summary>
    /// Provides a resource to manage ssl vpn client cert
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
    ///     var fooZones = Volcengine.Ecs.Zones.Invoke();
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
    ///         ZoneId = fooZones.Apply(zonesResult =&gt; zonesResult.Zones[0]?.Id),
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
    ///     var fooSslVpnClientCert = new Volcengine.Vpn.SslVpnClientCert("fooSslVpnClientCert", new()
    ///     {
    ///         SslVpnServerId = fooSslVpnServer.Id,
    ///         SslVpnClientCertName = "acc-test-client-cert",
    ///         Description = "acc-test",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// SSL VPN Client Cert can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import volcengine:vpn/sslVpnClientCert:SslVpnClientCert default vsc-2d6b7gjrzc2yo58ozfcx2****
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:vpn/sslVpnClientCert:SslVpnClientCert")]
    public partial class SslVpnClientCert : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The CA certificate.
        /// </summary>
        [Output("caCertificate")]
        public Output<string> CaCertificate { get; private set; } = null!;

        /// <summary>
        /// The status of the ssl vpn client cert.
        /// </summary>
        [Output("certificateStatus")]
        public Output<string> CertificateStatus { get; private set; } = null!;

        /// <summary>
        /// The client certificate.
        /// </summary>
        [Output("clientCertificate")]
        public Output<string> ClientCertificate { get; private set; } = null!;

        /// <summary>
        /// The key of the ssl vpn client.
        /// </summary>
        [Output("clientKey")]
        public Output<string> ClientKey { get; private set; } = null!;

        /// <summary>
        /// The creation time of the ssl vpn client cert.
        /// </summary>
        [Output("creationTime")]
        public Output<string> CreationTime { get; private set; } = null!;

        /// <summary>
        /// The description of the ssl vpn client cert.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// The expired time of the ssl vpn client cert.
        /// </summary>
        [Output("expiredTime")]
        public Output<string> ExpiredTime { get; private set; } = null!;

        /// <summary>
        /// The config of the open vpn client.
        /// </summary>
        [Output("openVpnClientConfig")]
        public Output<string> OpenVpnClientConfig { get; private set; } = null!;

        /// <summary>
        /// The name of the ssl vpn client cert.
        /// </summary>
        [Output("sslVpnClientCertName")]
        public Output<string> SslVpnClientCertName { get; private set; } = null!;

        /// <summary>
        /// The id of the ssl vpn server.
        /// </summary>
        [Output("sslVpnServerId")]
        public Output<string> SslVpnServerId { get; private set; } = null!;

        /// <summary>
        /// The status of the ssl vpn client.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The update time of the ssl vpn client cert.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a SslVpnClientCert resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SslVpnClientCert(string name, SslVpnClientCertArgs args, CustomResourceOptions? options = null)
            : base("volcengine:vpn/sslVpnClientCert:SslVpnClientCert", name, args ?? new SslVpnClientCertArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SslVpnClientCert(string name, Input<string> id, SslVpnClientCertState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:vpn/sslVpnClientCert:SslVpnClientCert", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SslVpnClientCert resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SslVpnClientCert Get(string name, Input<string> id, SslVpnClientCertState? state = null, CustomResourceOptions? options = null)
        {
            return new SslVpnClientCert(name, id, state, options);
        }
    }

    public sealed class SslVpnClientCertArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the ssl vpn client cert.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the ssl vpn client cert.
        /// </summary>
        [Input("sslVpnClientCertName")]
        public Input<string>? SslVpnClientCertName { get; set; }

        /// <summary>
        /// The id of the ssl vpn server.
        /// </summary>
        [Input("sslVpnServerId", required: true)]
        public Input<string> SslVpnServerId { get; set; } = null!;

        public SslVpnClientCertArgs()
        {
        }
        public static new SslVpnClientCertArgs Empty => new SslVpnClientCertArgs();
    }

    public sealed class SslVpnClientCertState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The CA certificate.
        /// </summary>
        [Input("caCertificate")]
        public Input<string>? CaCertificate { get; set; }

        /// <summary>
        /// The status of the ssl vpn client cert.
        /// </summary>
        [Input("certificateStatus")]
        public Input<string>? CertificateStatus { get; set; }

        /// <summary>
        /// The client certificate.
        /// </summary>
        [Input("clientCertificate")]
        public Input<string>? ClientCertificate { get; set; }

        /// <summary>
        /// The key of the ssl vpn client.
        /// </summary>
        [Input("clientKey")]
        public Input<string>? ClientKey { get; set; }

        /// <summary>
        /// The creation time of the ssl vpn client cert.
        /// </summary>
        [Input("creationTime")]
        public Input<string>? CreationTime { get; set; }

        /// <summary>
        /// The description of the ssl vpn client cert.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The expired time of the ssl vpn client cert.
        /// </summary>
        [Input("expiredTime")]
        public Input<string>? ExpiredTime { get; set; }

        /// <summary>
        /// The config of the open vpn client.
        /// </summary>
        [Input("openVpnClientConfig")]
        public Input<string>? OpenVpnClientConfig { get; set; }

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

        /// <summary>
        /// The status of the ssl vpn client.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The update time of the ssl vpn client cert.
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        public SslVpnClientCertState()
        {
        }
        public static new SslVpnClientCertState Empty => new SslVpnClientCertState();
    }
}