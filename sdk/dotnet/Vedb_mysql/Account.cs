// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Vedb_mysql
{
    /// <summary>
    /// Provides a resource to manage vedb mysql account
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
    ///         ZoneId = fooZones.Apply(zonesResult =&gt; zonesResult.Zones[2]?.Id),
    ///         VpcId = fooVpc.Id,
    ///     });
    /// 
    ///     var fooInstance = new Volcengine.Vedb_mysql.Instance("fooInstance", new()
    ///     {
    ///         ChargeType = "PostPaid",
    ///         StorageChargeType = "PostPaid",
    ///         DbEngineVersion = "MySQL_8_0",
    ///         DbMinorVersion = "3.0",
    ///         NodeNumber = 2,
    ///         NodeSpec = "vedb.mysql.x4.large",
    ///         SubnetId = fooSubnet.Id,
    ///         InstanceName = "tf-test",
    ///         ProjectName = "testA",
    ///         Tags = new[]
    ///         {
    ///             new Volcengine.Vedb_mysql.Inputs.InstanceTagArgs
    ///             {
    ///                 Key = "tftest",
    ///                 Value = "tftest",
    ///             },
    ///             new Volcengine.Vedb_mysql.Inputs.InstanceTagArgs
    ///             {
    ///                 Key = "tftest2",
    ///                 Value = "tftest2",
    ///             },
    ///         },
    ///     });
    /// 
    ///     var fooDatabase = new Volcengine.Vedb_mysql.Database("fooDatabase", new()
    ///     {
    ///         DbName = "tf-table",
    ///         InstanceId = fooInstance.Id,
    ///     });
    /// 
    ///     var fooAccount = new Volcengine.Vedb_mysql.Account("fooAccount", new()
    ///     {
    ///         AccountName = "tftest",
    ///         AccountPassword = "93f0cb0614Aab12",
    ///         AccountType = "Normal",
    ///         InstanceId = fooInstance.Id,
    ///         AccountPrivileges = new[]
    ///         {
    ///             new Volcengine.Vedb_mysql.Inputs.AccountAccountPrivilegeArgs
    ///             {
    ///                 DbName = fooDatabase.DbName,
    ///                 AccountPrivilege = "Custom",
    ///                 AccountPrivilegeDetail = "SELECT,INSERT,DELETE",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// VedbMysqlAccount can be imported using the instance id and account name, e.g.
    /// 
    /// ```sh
    /// $ pulumi import volcengine:vedb_mysql/account:Account default vedbm-r3xq0zdl****:testuser
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:vedb_mysql/account:Account")]
    public partial class Account : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Database account name. The account name must meet the following requirements:
        /// The name is unique and within 2 to 32 characters in length.
        /// Consists of lowercase letters, numbers, or underscores (_).
        /// Starts with a lowercase letter and ends with a letter or number.
        /// The name cannot contain certain prohibited words. For detailed information, please refer to prohibited keywords. And certain reserved words such as root, admin, etc. cannot be used.
        /// </summary>
        [Output("accountName")]
        public Output<string> AccountName { get; private set; } = null!;

        /// <summary>
        /// Password of database account. The account password must meet the following requirements:
        /// It can only contain upper and lower case letters, numbers and the following special characters _#!@$%^&amp;*()+=-.
        /// It must be within 8 to 32 characters in length.
        /// It must contain at least three of upper case letters, lower case letters, numbers or special characters.
        /// </summary>
        [Output("accountPassword")]
        public Output<string> AccountPassword { get; private set; } = null!;

        /// <summary>
        /// Database permission information. When the value of AccountType is Super, this parameter does not need to be passed. High-privilege accounts by default have all permissions for all databases under this instance. When the value of AccountType is Normal, it is recommended to pass this parameter to grant specified permissions for specified databases to ordinary accounts. If not set, this account does not have any permissions for any database.
        /// </summary>
        [Output("accountPrivileges")]
        public Output<ImmutableArray<Outputs.AccountAccountPrivilege>> AccountPrivileges { get; private set; } = null!;

        /// <summary>
        /// Database account type. Values: 
        /// Super: High-privilege account. Only one high-privilege account can be created for an instance. It has all permissions for all databases under this instance and can manage all ordinary accounts and databases.
        /// Normal: Multiple ordinary accounts can be created for an instance. Specific database permissions need to be manually granted to ordinary accounts.
        /// </summary>
        [Output("accountType")]
        public Output<string> AccountType { get; private set; } = null!;

        /// <summary>
        /// The id of the instance.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;


        /// <summary>
        /// Create a Account resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Account(string name, AccountArgs args, CustomResourceOptions? options = null)
            : base("volcengine:vedb_mysql/account:Account", name, args ?? new AccountArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Account(string name, Input<string> id, AccountState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:vedb_mysql/account:Account", name, state, MakeResourceOptions(options, id))
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
                    "accountPassword",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Account resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Account Get(string name, Input<string> id, AccountState? state = null, CustomResourceOptions? options = null)
        {
            return new Account(name, id, state, options);
        }
    }

    public sealed class AccountArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Database account name. The account name must meet the following requirements:
        /// The name is unique and within 2 to 32 characters in length.
        /// Consists of lowercase letters, numbers, or underscores (_).
        /// Starts with a lowercase letter and ends with a letter or number.
        /// The name cannot contain certain prohibited words. For detailed information, please refer to prohibited keywords. And certain reserved words such as root, admin, etc. cannot be used.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        [Input("accountPassword", required: true)]
        private Input<string>? _accountPassword;

        /// <summary>
        /// Password of database account. The account password must meet the following requirements:
        /// It can only contain upper and lower case letters, numbers and the following special characters _#!@$%^&amp;*()+=-.
        /// It must be within 8 to 32 characters in length.
        /// It must contain at least three of upper case letters, lower case letters, numbers or special characters.
        /// </summary>
        public Input<string>? AccountPassword
        {
            get => _accountPassword;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _accountPassword = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("accountPrivileges")]
        private InputList<Inputs.AccountAccountPrivilegeArgs>? _accountPrivileges;

        /// <summary>
        /// Database permission information. When the value of AccountType is Super, this parameter does not need to be passed. High-privilege accounts by default have all permissions for all databases under this instance. When the value of AccountType is Normal, it is recommended to pass this parameter to grant specified permissions for specified databases to ordinary accounts. If not set, this account does not have any permissions for any database.
        /// </summary>
        public InputList<Inputs.AccountAccountPrivilegeArgs> AccountPrivileges
        {
            get => _accountPrivileges ?? (_accountPrivileges = new InputList<Inputs.AccountAccountPrivilegeArgs>());
            set => _accountPrivileges = value;
        }

        /// <summary>
        /// Database account type. Values: 
        /// Super: High-privilege account. Only one high-privilege account can be created for an instance. It has all permissions for all databases under this instance and can manage all ordinary accounts and databases.
        /// Normal: Multiple ordinary accounts can be created for an instance. Specific database permissions need to be manually granted to ordinary accounts.
        /// </summary>
        [Input("accountType", required: true)]
        public Input<string> AccountType { get; set; } = null!;

        /// <summary>
        /// The id of the instance.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        public AccountArgs()
        {
        }
        public static new AccountArgs Empty => new AccountArgs();
    }

    public sealed class AccountState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Database account name. The account name must meet the following requirements:
        /// The name is unique and within 2 to 32 characters in length.
        /// Consists of lowercase letters, numbers, or underscores (_).
        /// Starts with a lowercase letter and ends with a letter or number.
        /// The name cannot contain certain prohibited words. For detailed information, please refer to prohibited keywords. And certain reserved words such as root, admin, etc. cannot be used.
        /// </summary>
        [Input("accountName")]
        public Input<string>? AccountName { get; set; }

        [Input("accountPassword")]
        private Input<string>? _accountPassword;

        /// <summary>
        /// Password of database account. The account password must meet the following requirements:
        /// It can only contain upper and lower case letters, numbers and the following special characters _#!@$%^&amp;*()+=-.
        /// It must be within 8 to 32 characters in length.
        /// It must contain at least three of upper case letters, lower case letters, numbers or special characters.
        /// </summary>
        public Input<string>? AccountPassword
        {
            get => _accountPassword;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _accountPassword = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("accountPrivileges")]
        private InputList<Inputs.AccountAccountPrivilegeGetArgs>? _accountPrivileges;

        /// <summary>
        /// Database permission information. When the value of AccountType is Super, this parameter does not need to be passed. High-privilege accounts by default have all permissions for all databases under this instance. When the value of AccountType is Normal, it is recommended to pass this parameter to grant specified permissions for specified databases to ordinary accounts. If not set, this account does not have any permissions for any database.
        /// </summary>
        public InputList<Inputs.AccountAccountPrivilegeGetArgs> AccountPrivileges
        {
            get => _accountPrivileges ?? (_accountPrivileges = new InputList<Inputs.AccountAccountPrivilegeGetArgs>());
            set => _accountPrivileges = value;
        }

        /// <summary>
        /// Database account type. Values: 
        /// Super: High-privilege account. Only one high-privilege account can be created for an instance. It has all permissions for all databases under this instance and can manage all ordinary accounts and databases.
        /// Normal: Multiple ordinary accounts can be created for an instance. Specific database permissions need to be manually granted to ordinary accounts.
        /// </summary>
        [Input("accountType")]
        public Input<string>? AccountType { get; set; }

        /// <summary>
        /// The id of the instance.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        public AccountState()
        {
        }
        public static new AccountState Empty => new AccountState();
    }
}