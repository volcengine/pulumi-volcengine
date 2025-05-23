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
    /// Provides a resource to manage ecs key pair associate
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
    ///     var fooKeyPair = new Volcengine.Ecs.KeyPair("fooKeyPair", new()
    ///     {
    ///         KeyPairName = "acc-test-key-name",
    ///         Description = "acc-test",
    ///     });
    /// 
    ///     var fooZones = Volcengine.Ecs.GetZones.Invoke();
    /// 
    ///     var fooImages = Volcengine.Ecs.GetImages.Invoke(new()
    ///     {
    ///         OsType = "Linux",
    ///         Visibility = "public",
    ///         InstanceTypeId = "ecs.g1.large",
    ///     });
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
    ///     var fooSecurityGroup = new Volcengine.Vpc.SecurityGroup("fooSecurityGroup", new()
    ///     {
    ///         VpcId = fooVpc.Id,
    ///         SecurityGroupName = "acc-test-security-group",
    ///     });
    /// 
    ///     var fooInstance = new Volcengine.Ecs.Instance("fooInstance", new()
    ///     {
    ///         ImageId = fooImages.Apply(getImagesResult =&gt; getImagesResult.Images[0]?.ImageId),
    ///         InstanceType = "ecs.g1.large",
    ///         InstanceName = "acc-test-ecs-name",
    ///         Password = "your password",
    ///         InstanceChargeType = "PostPaid",
    ///         SystemVolumeType = "ESSD_PL0",
    ///         SystemVolumeSize = 40,
    ///         SubnetId = fooSubnet.Id,
    ///         SecurityGroupIds = new[]
    ///         {
    ///             fooSecurityGroup.Id,
    ///         },
    ///     });
    /// 
    ///     var fooKeyPairAssociate = new Volcengine.Ecs.KeyPairAssociate("fooKeyPairAssociate", new()
    ///     {
    ///         InstanceId = fooInstance.Id,
    ///         KeyPairId = fooKeyPair.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ECS key pair associate can be imported using the id, e.g.
    /// 
    /// After binding the key pair, the instance needs to be restarted for the key pair to take effect.
    /// 
    /// After the key pair is bound, the password login method will automatically become invalid. If your instance has been set for password login, after the key pair is bound, you will no longer be able to use the password login method.
    /// 
    /// ```sh
    /// $ pulumi import volcengine:ecs/keyPairAssociate:KeyPairAssociate default kp-ybti5tkpkv2udbfolrft:i-mizl7m1kqccg5smt1bdpijuj
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:ecs/keyPairAssociate:KeyPairAssociate")]
    public partial class KeyPairAssociate : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of ECS Instance.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// The ID of ECS KeyPair Associate.
        /// </summary>
        [Output("keyPairId")]
        public Output<string> KeyPairId { get; private set; } = null!;


        /// <summary>
        /// Create a KeyPairAssociate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public KeyPairAssociate(string name, KeyPairAssociateArgs args, CustomResourceOptions? options = null)
            : base("volcengine:ecs/keyPairAssociate:KeyPairAssociate", name, args ?? new KeyPairAssociateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private KeyPairAssociate(string name, Input<string> id, KeyPairAssociateState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:ecs/keyPairAssociate:KeyPairAssociate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing KeyPairAssociate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static KeyPairAssociate Get(string name, Input<string> id, KeyPairAssociateState? state = null, CustomResourceOptions? options = null)
        {
            return new KeyPairAssociate(name, id, state, options);
        }
    }

    public sealed class KeyPairAssociateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of ECS Instance.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// The ID of ECS KeyPair Associate.
        /// </summary>
        [Input("keyPairId", required: true)]
        public Input<string> KeyPairId { get; set; } = null!;

        public KeyPairAssociateArgs()
        {
        }
        public static new KeyPairAssociateArgs Empty => new KeyPairAssociateArgs();
    }

    public sealed class KeyPairAssociateState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of ECS Instance.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// The ID of ECS KeyPair Associate.
        /// </summary>
        [Input("keyPairId")]
        public Input<string>? KeyPairId { get; set; }

        public KeyPairAssociateState()
        {
        }
        public static new KeyPairAssociateState Empty => new KeyPairAssociateState();
    }
}
