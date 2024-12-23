// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Rds
{
    /// <summary>
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
    ///     var foo = new Volcengine.Rds.Instance("foo", new()
    ///     {
    ///         ChargeType = "PostPaid",
    ///         DbEngine = "MySQL",
    ///         DbEngineVersion = "MySQL_Community_5_7",
    ///         InstanceName = "tf-test",
    ///         InstanceSpecName = "rds.mysql.1c2g",
    ///         InstanceType = "HA",
    ///         Region = "cn-north-4",
    ///         StorageSpaceGb = 100,
    ///         StorageType = "LocalSSD",
    ///         SubnetId = "subnet-1g0d4fkh1nabk8ibuxx1jtvss",
    ///         VpcId = "vpc-3cj17x7u9bzeo6c6rrtzfpaeb",
    ///         Zone = "cn-langfang-b",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// RDS Instance can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import volcengine:rds/instance:Instance default mysql-42b38c769c4b
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:rds/instance:Instance")]
    public partial class Instance : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Whether to automatically renew. Default: false. Value:
        /// true: yes.
        /// false: no. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        /// </summary>
        [Output("autoRenew")]
        public Output<bool?> AutoRenew { get; private set; } = null!;

        /// <summary>
        /// Billing type. Value:
        /// PostPaid: Postpaid (pay-as-you-go).
        /// Prepaid: Prepaid (yearly and monthly).
        /// </summary>
        [Output("chargeType")]
        public Output<string> ChargeType { get; private set; } = null!;

        /// <summary>
        /// The connection info ot the RDS instance.
        /// </summary>
        [Output("connectionInfo")]
        public Output<Outputs.InstanceConnectionInfo> ConnectionInfo { get; private set; } = null!;

        /// <summary>
        /// Database type. Value:
        /// MySQL (default).
        /// </summary>
        [Output("dbEngine")]
        public Output<string?> DbEngine { get; private set; } = null!;

        /// <summary>
        /// Instance type. Value:
        /// MySQL_Community_5_7
        /// MySQL_8_0.
        /// </summary>
        [Output("dbEngineVersion")]
        public Output<string> DbEngineVersion { get; private set; } = null!;

        /// <summary>
        /// Set the name of the instance. The naming rules are as follows:
        /// 
        /// Cannot start with a number, a dash (-).
        /// It can only contain Chinese characters, letters, numbers, underscores (_) and underscores (-).
        /// The length needs to be within 1~128 characters.
        /// </summary>
        [Output("instanceName")]
        public Output<string?> InstanceName { get; private set; } = null!;

        /// <summary>
        /// Instance specification name, you can specify the specification name of the instance to be created. Value:
        /// rds.mysql.1c2g
        /// rds.mysql.2c4g
        /// rds.mysql.4c8g
        /// rds.mysql.4c16g
        /// rds.mysql.8c32g
        /// rds.mysql.16c64g
        /// rds.mysql.16c128g
        /// rds.mysql.32c128g
        /// rds.mysql.32c256g.
        /// </summary>
        [Output("instanceSpecName")]
        public Output<string> InstanceSpecName { get; private set; } = null!;

        /// <summary>
        /// Instance type. Value:
        /// HA: High availability version.
        /// </summary>
        [Output("instanceType")]
        public Output<string> InstanceType { get; private set; } = null!;

        /// <summary>
        /// The purchase cycle in the prepaid scenario. Value:
        /// Month: monthly subscription.
        /// Year: yearly subscription. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        /// </summary>
        [Output("prepaidPeriod")]
        public Output<string?> PrepaidPeriod { get; private set; } = null!;

        /// <summary>
        /// Select the project to which the instance belongs. If this parameter is left blank, the new instance will not be added to any project. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        /// </summary>
        [Output("projectName")]
        public Output<string?> ProjectName { get; private set; } = null!;

        /// <summary>
        /// The region of the RDS instance.
        /// </summary>
        [Output("region")]
        public Output<string?> Region { get; private set; } = null!;

        /// <summary>
        /// The storage space(GB) of the RDS instance.
        /// </summary>
        [Output("storageSpaceGb")]
        public Output<int> StorageSpaceGb { get; private set; } = null!;

        /// <summary>
        /// Instance storage type. Value:
        /// LocalSSD: Local SSD disk.
        /// </summary>
        [Output("storageType")]
        public Output<string> StorageType { get; private set; } = null!;

        /// <summary>
        /// Subnet ID. The subnet must belong to the selected Availability Zone.
        /// </summary>
        [Output("subnetId")]
        public Output<string> SubnetId { get; private set; } = null!;

        /// <summary>
        /// Fill in the high-privileged user account name. The naming rules are as follows:
        /// Unique name.
        /// Start with a letter and end with a letter or number.
        /// Consists of lowercase letters, numbers, or underscores (_).
        /// The length is 2~32 characters.
        /// [Keywords](https://www.volcengine.com/docs/6313/66162) are not allowed for account names.
        /// </summary>
        [Output("superAccountName")]
        public Output<string?> SuperAccountName { get; private set; } = null!;

        /// <summary>
        /// Set a high-privilege account password. The rules are as follows:
        /// Only uppercase and lowercase letters, numbers and the following special characters _#!@$%^*()+=-.
        /// The length needs to be within 8~32 characters.
        /// Contains at least 3 of uppercase letters, lowercase letters, numbers or special characters.
        /// </summary>
        [Output("superAccountPassword")]
        public Output<string?> SuperAccountPassword { get; private set; } = null!;

        /// <summary>
        /// supper_account_password is deprecated, use super_account_password instead Set a high-privilege account password. The rules are as follows:
        /// Only uppercase and lowercase letters, numbers and the following special characters _#!@$%^*()+=-.
        /// The length needs to be within 8~32 characters.
        /// Contains at least 3 of uppercase letters, lowercase letters, numbers or special characters.
        /// </summary>
        [Output("supperAccountPassword")]
        public Output<string?> SupperAccountPassword { get; private set; } = null!;

        /// <summary>
        /// The purchase time of RDS instance. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        /// </summary>
        [Output("usedTime")]
        public Output<int?> UsedTime { get; private set; } = null!;

        /// <summary>
        /// The vpc ID of the RDS instance.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;

        /// <summary>
        /// The available zone of the RDS instance.
        /// </summary>
        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a Instance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Instance(string name, InstanceArgs args, CustomResourceOptions? options = null)
            : base("volcengine:rds/instance:Instance", name, args ?? new InstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Instance(string name, Input<string> id, InstanceState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:rds/instance:Instance", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Instance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Instance Get(string name, Input<string> id, InstanceState? state = null, CustomResourceOptions? options = null)
        {
            return new Instance(name, id, state, options);
        }
    }

    public sealed class InstanceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to automatically renew. Default: false. Value:
        /// true: yes.
        /// false: no. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        /// </summary>
        [Input("autoRenew")]
        public Input<bool>? AutoRenew { get; set; }

        /// <summary>
        /// Billing type. Value:
        /// PostPaid: Postpaid (pay-as-you-go).
        /// Prepaid: Prepaid (yearly and monthly).
        /// </summary>
        [Input("chargeType", required: true)]
        public Input<string> ChargeType { get; set; } = null!;

        /// <summary>
        /// Database type. Value:
        /// MySQL (default).
        /// </summary>
        [Input("dbEngine")]
        public Input<string>? DbEngine { get; set; }

        /// <summary>
        /// Instance type. Value:
        /// MySQL_Community_5_7
        /// MySQL_8_0.
        /// </summary>
        [Input("dbEngineVersion", required: true)]
        public Input<string> DbEngineVersion { get; set; } = null!;

        /// <summary>
        /// Set the name of the instance. The naming rules are as follows:
        /// 
        /// Cannot start with a number, a dash (-).
        /// It can only contain Chinese characters, letters, numbers, underscores (_) and underscores (-).
        /// The length needs to be within 1~128 characters.
        /// </summary>
        [Input("instanceName")]
        public Input<string>? InstanceName { get; set; }

        /// <summary>
        /// Instance specification name, you can specify the specification name of the instance to be created. Value:
        /// rds.mysql.1c2g
        /// rds.mysql.2c4g
        /// rds.mysql.4c8g
        /// rds.mysql.4c16g
        /// rds.mysql.8c32g
        /// rds.mysql.16c64g
        /// rds.mysql.16c128g
        /// rds.mysql.32c128g
        /// rds.mysql.32c256g.
        /// </summary>
        [Input("instanceSpecName", required: true)]
        public Input<string> InstanceSpecName { get; set; } = null!;

        /// <summary>
        /// Instance type. Value:
        /// HA: High availability version.
        /// </summary>
        [Input("instanceType", required: true)]
        public Input<string> InstanceType { get; set; } = null!;

        /// <summary>
        /// The purchase cycle in the prepaid scenario. Value:
        /// Month: monthly subscription.
        /// Year: yearly subscription. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        /// </summary>
        [Input("prepaidPeriod")]
        public Input<string>? PrepaidPeriod { get; set; }

        /// <summary>
        /// Select the project to which the instance belongs. If this parameter is left blank, the new instance will not be added to any project. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        /// </summary>
        [Input("projectName")]
        public Input<string>? ProjectName { get; set; }

        /// <summary>
        /// The region of the RDS instance.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The storage space(GB) of the RDS instance.
        /// </summary>
        [Input("storageSpaceGb", required: true)]
        public Input<int> StorageSpaceGb { get; set; } = null!;

        /// <summary>
        /// Instance storage type. Value:
        /// LocalSSD: Local SSD disk.
        /// </summary>
        [Input("storageType", required: true)]
        public Input<string> StorageType { get; set; } = null!;

        /// <summary>
        /// Subnet ID. The subnet must belong to the selected Availability Zone.
        /// </summary>
        [Input("subnetId", required: true)]
        public Input<string> SubnetId { get; set; } = null!;

        /// <summary>
        /// Fill in the high-privileged user account name. The naming rules are as follows:
        /// Unique name.
        /// Start with a letter and end with a letter or number.
        /// Consists of lowercase letters, numbers, or underscores (_).
        /// The length is 2~32 characters.
        /// [Keywords](https://www.volcengine.com/docs/6313/66162) are not allowed for account names.
        /// </summary>
        [Input("superAccountName")]
        public Input<string>? SuperAccountName { get; set; }

        /// <summary>
        /// Set a high-privilege account password. The rules are as follows:
        /// Only uppercase and lowercase letters, numbers and the following special characters _#!@$%^*()+=-.
        /// The length needs to be within 8~32 characters.
        /// Contains at least 3 of uppercase letters, lowercase letters, numbers or special characters.
        /// </summary>
        [Input("superAccountPassword")]
        public Input<string>? SuperAccountPassword { get; set; }

        /// <summary>
        /// supper_account_password is deprecated, use super_account_password instead Set a high-privilege account password. The rules are as follows:
        /// Only uppercase and lowercase letters, numbers and the following special characters _#!@$%^*()+=-.
        /// The length needs to be within 8~32 characters.
        /// Contains at least 3 of uppercase letters, lowercase letters, numbers or special characters.
        /// </summary>
        [Input("supperAccountPassword")]
        public Input<string>? SupperAccountPassword { get; set; }

        /// <summary>
        /// The purchase time of RDS instance. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        /// </summary>
        [Input("usedTime")]
        public Input<int>? UsedTime { get; set; }

        /// <summary>
        /// The vpc ID of the RDS instance.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        /// <summary>
        /// The available zone of the RDS instance.
        /// </summary>
        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public InstanceArgs()
        {
        }
        public static new InstanceArgs Empty => new InstanceArgs();
    }

    public sealed class InstanceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to automatically renew. Default: false. Value:
        /// true: yes.
        /// false: no. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        /// </summary>
        [Input("autoRenew")]
        public Input<bool>? AutoRenew { get; set; }

        /// <summary>
        /// Billing type. Value:
        /// PostPaid: Postpaid (pay-as-you-go).
        /// Prepaid: Prepaid (yearly and monthly).
        /// </summary>
        [Input("chargeType")]
        public Input<string>? ChargeType { get; set; }

        /// <summary>
        /// The connection info ot the RDS instance.
        /// </summary>
        [Input("connectionInfo")]
        public Input<Inputs.InstanceConnectionInfoGetArgs>? ConnectionInfo { get; set; }

        /// <summary>
        /// Database type. Value:
        /// MySQL (default).
        /// </summary>
        [Input("dbEngine")]
        public Input<string>? DbEngine { get; set; }

        /// <summary>
        /// Instance type. Value:
        /// MySQL_Community_5_7
        /// MySQL_8_0.
        /// </summary>
        [Input("dbEngineVersion")]
        public Input<string>? DbEngineVersion { get; set; }

        /// <summary>
        /// Set the name of the instance. The naming rules are as follows:
        /// 
        /// Cannot start with a number, a dash (-).
        /// It can only contain Chinese characters, letters, numbers, underscores (_) and underscores (-).
        /// The length needs to be within 1~128 characters.
        /// </summary>
        [Input("instanceName")]
        public Input<string>? InstanceName { get; set; }

        /// <summary>
        /// Instance specification name, you can specify the specification name of the instance to be created. Value:
        /// rds.mysql.1c2g
        /// rds.mysql.2c4g
        /// rds.mysql.4c8g
        /// rds.mysql.4c16g
        /// rds.mysql.8c32g
        /// rds.mysql.16c64g
        /// rds.mysql.16c128g
        /// rds.mysql.32c128g
        /// rds.mysql.32c256g.
        /// </summary>
        [Input("instanceSpecName")]
        public Input<string>? InstanceSpecName { get; set; }

        /// <summary>
        /// Instance type. Value:
        /// HA: High availability version.
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        /// <summary>
        /// The purchase cycle in the prepaid scenario. Value:
        /// Month: monthly subscription.
        /// Year: yearly subscription. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        /// </summary>
        [Input("prepaidPeriod")]
        public Input<string>? PrepaidPeriod { get; set; }

        /// <summary>
        /// Select the project to which the instance belongs. If this parameter is left blank, the new instance will not be added to any project. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        /// </summary>
        [Input("projectName")]
        public Input<string>? ProjectName { get; set; }

        /// <summary>
        /// The region of the RDS instance.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The storage space(GB) of the RDS instance.
        /// </summary>
        [Input("storageSpaceGb")]
        public Input<int>? StorageSpaceGb { get; set; }

        /// <summary>
        /// Instance storage type. Value:
        /// LocalSSD: Local SSD disk.
        /// </summary>
        [Input("storageType")]
        public Input<string>? StorageType { get; set; }

        /// <summary>
        /// Subnet ID. The subnet must belong to the selected Availability Zone.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        /// <summary>
        /// Fill in the high-privileged user account name. The naming rules are as follows:
        /// Unique name.
        /// Start with a letter and end with a letter or number.
        /// Consists of lowercase letters, numbers, or underscores (_).
        /// The length is 2~32 characters.
        /// [Keywords](https://www.volcengine.com/docs/6313/66162) are not allowed for account names.
        /// </summary>
        [Input("superAccountName")]
        public Input<string>? SuperAccountName { get; set; }

        /// <summary>
        /// Set a high-privilege account password. The rules are as follows:
        /// Only uppercase and lowercase letters, numbers and the following special characters _#!@$%^*()+=-.
        /// The length needs to be within 8~32 characters.
        /// Contains at least 3 of uppercase letters, lowercase letters, numbers or special characters.
        /// </summary>
        [Input("superAccountPassword")]
        public Input<string>? SuperAccountPassword { get; set; }

        /// <summary>
        /// supper_account_password is deprecated, use super_account_password instead Set a high-privilege account password. The rules are as follows:
        /// Only uppercase and lowercase letters, numbers and the following special characters _#!@$%^*()+=-.
        /// The length needs to be within 8~32 characters.
        /// Contains at least 3 of uppercase letters, lowercase letters, numbers or special characters.
        /// </summary>
        [Input("supperAccountPassword")]
        public Input<string>? SupperAccountPassword { get; set; }

        /// <summary>
        /// The purchase time of RDS instance. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        /// </summary>
        [Input("usedTime")]
        public Input<int>? UsedTime { get; set; }

        /// <summary>
        /// The vpc ID of the RDS instance.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        /// <summary>
        /// The available zone of the RDS instance.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public InstanceState()
        {
        }
        public static new InstanceState Empty => new InstanceState();
    }
}
