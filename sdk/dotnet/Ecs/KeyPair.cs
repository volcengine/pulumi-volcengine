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
    /// Provides a resource to manage ecs key pair
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
    ///     var foo = new Volcengine.Ecs.KeyPair("foo", new()
    ///     {
    ///         Description = "acc-test",
    ///         KeyPairName = "acc-test-key-name",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ECS key pair can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import volcengine:ecs/keyPair:KeyPair default kp-mizl7m1kqccg5smt1bdpijuj
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:ecs/keyPair:KeyPair")]
    public partial class KeyPair : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The description of key pair.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The finger print info.
        /// </summary>
        [Output("fingerPrint")]
        public Output<string> FingerPrint { get; private set; } = null!;

        /// <summary>
        /// Target file to save private key. It is recommended that the value not be empty. You only have one chance to download the private key, the volcengine will not save your private key, please keep it safe. In the TF import scenario, this field will not write the private key locally.
        /// </summary>
        [Output("keyFile")]
        public Output<string?> KeyFile { get; private set; } = null!;

        /// <summary>
        /// The id of key pair.
        /// </summary>
        [Output("keyPairId")]
        public Output<string> KeyPairId { get; private set; } = null!;

        /// <summary>
        /// The name of key pair.
        /// </summary>
        [Output("keyPairName")]
        public Output<string> KeyPairName { get; private set; } = null!;

        /// <summary>
        /// Public key string.
        /// </summary>
        [Output("publicKey")]
        public Output<string?> PublicKey { get; private set; } = null!;


        /// <summary>
        /// Create a KeyPair resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public KeyPair(string name, KeyPairArgs args, CustomResourceOptions? options = null)
            : base("volcengine:ecs/keyPair:KeyPair", name, args ?? new KeyPairArgs(), MakeResourceOptions(options, ""))
        {
        }

        private KeyPair(string name, Input<string> id, KeyPairState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:ecs/keyPair:KeyPair", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing KeyPair resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static KeyPair Get(string name, Input<string> id, KeyPairState? state = null, CustomResourceOptions? options = null)
        {
            return new KeyPair(name, id, state, options);
        }
    }

    public sealed class KeyPairArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of key pair.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Target file to save private key. It is recommended that the value not be empty. You only have one chance to download the private key, the volcengine will not save your private key, please keep it safe. In the TF import scenario, this field will not write the private key locally.
        /// </summary>
        [Input("keyFile")]
        public Input<string>? KeyFile { get; set; }

        /// <summary>
        /// The name of key pair.
        /// </summary>
        [Input("keyPairName", required: true)]
        public Input<string> KeyPairName { get; set; } = null!;

        /// <summary>
        /// Public key string.
        /// </summary>
        [Input("publicKey")]
        public Input<string>? PublicKey { get; set; }

        public KeyPairArgs()
        {
        }
        public static new KeyPairArgs Empty => new KeyPairArgs();
    }

    public sealed class KeyPairState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of key pair.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The finger print info.
        /// </summary>
        [Input("fingerPrint")]
        public Input<string>? FingerPrint { get; set; }

        /// <summary>
        /// Target file to save private key. It is recommended that the value not be empty. You only have one chance to download the private key, the volcengine will not save your private key, please keep it safe. In the TF import scenario, this field will not write the private key locally.
        /// </summary>
        [Input("keyFile")]
        public Input<string>? KeyFile { get; set; }

        /// <summary>
        /// The id of key pair.
        /// </summary>
        [Input("keyPairId")]
        public Input<string>? KeyPairId { get; set; }

        /// <summary>
        /// The name of key pair.
        /// </summary>
        [Input("keyPairName")]
        public Input<string>? KeyPairName { get; set; }

        /// <summary>
        /// Public key string.
        /// </summary>
        [Input("publicKey")]
        public Input<string>? PublicKey { get; set; }

        public KeyPairState()
        {
        }
        public static new KeyPairState Empty => new KeyPairState();
    }
}