// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Cloud_identity
{
    /// <summary>
    /// Provides a resource to manage cloud identity user
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
    ///     var foo = new Volcengine.Cloud_identity.User("foo", new()
    ///     {
    ///         Description = "tf",
    ///         DisplayName = "tf-test-user",
    ///         Email = "88@qq.com",
    ///         Phone = "1810000****",
    ///         UserName = "acc-test-user",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// CloudIdentityUser can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import volcengine:cloud_identity/user:User default resource_id
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:cloud_identity/user:User")]
    public partial class User : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The description of the cloud identity user.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The display name of the cloud identity user.
        /// </summary>
        [Output("displayName")]
        public Output<string?> DisplayName { get; private set; } = null!;

        /// <summary>
        /// The email of the cloud identity user.
        /// </summary>
        [Output("email")]
        public Output<string?> Email { get; private set; } = null!;

        /// <summary>
        /// The identity type of the cloud identity user.
        /// </summary>
        [Output("identityType")]
        public Output<string> IdentityType { get; private set; } = null!;

        /// <summary>
        /// The phone of the cloud identity user. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        /// </summary>
        [Output("phone")]
        public Output<string?> Phone { get; private set; } = null!;

        /// <summary>
        /// The source of the cloud identity user.
        /// </summary>
        [Output("source")]
        public Output<string> Source { get; private set; } = null!;

        /// <summary>
        /// The name of the cloud identity user.
        /// </summary>
        [Output("userName")]
        public Output<string> UserName { get; private set; } = null!;


        /// <summary>
        /// Create a User resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public User(string name, UserArgs args, CustomResourceOptions? options = null)
            : base("volcengine:cloud_identity/user:User", name, args ?? new UserArgs(), MakeResourceOptions(options, ""))
        {
        }

        private User(string name, Input<string> id, UserState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:cloud_identity/user:User", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/volcengine",
                AdditionalSecretOutputs =
                {
                    "phone",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing User resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static User Get(string name, Input<string> id, UserState? state = null, CustomResourceOptions? options = null)
        {
            return new User(name, id, state, options);
        }
    }

    public sealed class UserArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the cloud identity user.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The display name of the cloud identity user.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The email of the cloud identity user.
        /// </summary>
        [Input("email")]
        public Input<string>? Email { get; set; }

        [Input("phone")]
        private Input<string>? _phone;

        /// <summary>
        /// The phone of the cloud identity user. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        /// </summary>
        public Input<string>? Phone
        {
            get => _phone;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _phone = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The name of the cloud identity user.
        /// </summary>
        [Input("userName", required: true)]
        public Input<string> UserName { get; set; } = null!;

        public UserArgs()
        {
        }
        public static new UserArgs Empty => new UserArgs();
    }

    public sealed class UserState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the cloud identity user.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The display name of the cloud identity user.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The email of the cloud identity user.
        /// </summary>
        [Input("email")]
        public Input<string>? Email { get; set; }

        /// <summary>
        /// The identity type of the cloud identity user.
        /// </summary>
        [Input("identityType")]
        public Input<string>? IdentityType { get; set; }

        [Input("phone")]
        private Input<string>? _phone;

        /// <summary>
        /// The phone of the cloud identity user. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        /// </summary>
        public Input<string>? Phone
        {
            get => _phone;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _phone = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The source of the cloud identity user.
        /// </summary>
        [Input("source")]
        public Input<string>? Source { get; set; }

        /// <summary>
        /// The name of the cloud identity user.
        /// </summary>
        [Input("userName")]
        public Input<string>? UserName { get; set; }

        public UserState()
        {
        }
        public static new UserState Empty => new UserState();
    }
}
