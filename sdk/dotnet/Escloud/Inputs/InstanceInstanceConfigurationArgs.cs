// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Escloud.Inputs
{

    public sealed class InstanceInstanceConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("adminPassword", required: true)]
        private Input<string>? _adminPassword;

        /// <summary>
        /// The password of administrator account. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        /// </summary>
        public Input<string>? AdminPassword
        {
            get => _adminPassword;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _adminPassword = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The name of administrator account(should be admin).
        /// </summary>
        [Input("adminUserName", required: true)]
        public Input<string> AdminUserName { get; set; } = null!;

        /// <summary>
        /// The charge type of ESCloud instance, the value can be PostPaid or PrePaid.
        /// </summary>
        [Input("chargeType", required: true)]
        public Input<string> ChargeType { get; set; } = null!;

        /// <summary>
        /// Configuration code used for billing.
        /// </summary>
        [Input("configurationCode", required: true)]
        public Input<string> ConfigurationCode { get; set; } = null!;

        /// <summary>
        /// Whether Https access is enabled.
        /// </summary>
        [Input("enableHttps", required: true)]
        public Input<bool> EnableHttps { get; set; } = null!;

        /// <summary>
        /// Whether the Master node is independent.
        /// </summary>
        [Input("enablePureMaster", required: true)]
        public Input<bool> EnablePureMaster { get; set; } = null!;

        /// <summary>
        /// Whether to force restart when changes are made. If true, it means that the cluster will be forced to restart without paying attention to instance availability. Works only on modified the node_specs_assigns field.
        /// </summary>
        [Input("forceRestartAfterScale")]
        public Input<bool>? ForceRestartAfterScale { get; set; }

        /// <summary>
        /// The name of ESCloud instance.
        /// </summary>
        [Input("instanceName")]
        public Input<string>? InstanceName { get; set; }

        [Input("maintenanceDays")]
        private InputList<string>? _maintenanceDays;

        /// <summary>
        /// The maintainable date for the instance. Works only on modified scenes.
        /// </summary>
        public InputList<string> MaintenanceDays
        {
            get => _maintenanceDays ?? (_maintenanceDays = new InputList<string>());
            set => _maintenanceDays = value;
        }

        /// <summary>
        /// The maintainable time period for the instance. Works only on modified scenes.
        /// </summary>
        [Input("maintenanceTime")]
        public Input<string>? MaintenanceTime { get; set; }

        [Input("nodeSpecsAssigns", required: true)]
        private InputList<Inputs.InstanceInstanceConfigurationNodeSpecsAssignArgs>? _nodeSpecsAssigns;

        /// <summary>
        /// The number and configuration of various ESCloud instance node. Kibana NodeSpecsAssign should not be modified.
        /// </summary>
        public InputList<Inputs.InstanceInstanceConfigurationNodeSpecsAssignArgs> NodeSpecsAssigns
        {
            get => _nodeSpecsAssigns ?? (_nodeSpecsAssigns = new InputList<Inputs.InstanceInstanceConfigurationNodeSpecsAssignArgs>());
            set => _nodeSpecsAssigns = value;
        }

        /// <summary>
        /// The project name  to which the ESCloud instance belongs.
        /// </summary>
        [Input("projectName")]
        public Input<string>? ProjectName { get; set; }

        /// <summary>
        /// The region ID of ESCloud instance.
        /// </summary>
        [Input("regionId")]
        public Input<string>? RegionId { get; set; }

        /// <summary>
        /// The ID of subnet, the subnet must belong to the AZ selected.
        /// </summary>
        [Input("subnetId", required: true)]
        public Input<string> SubnetId { get; set; } = null!;

        /// <summary>
        /// The version of ESCloud instance, the value is V6_7 or V7_10.
        /// </summary>
        [Input("version", required: true)]
        public Input<string> Version { get; set; } = null!;

        /// <summary>
        /// The available zone ID of ESCloud instance.
        /// </summary>
        [Input("zoneId")]
        public Input<string>? ZoneId { get; set; }

        /// <summary>
        /// The zone count of the ESCloud instance used.
        /// </summary>
        [Input("zoneNumber", required: true)]
        public Input<int> ZoneNumber { get; set; } = null!;

        public InstanceInstanceConfigurationArgs()
        {
        }
        public static new InstanceInstanceConfigurationArgs Empty => new InstanceInstanceConfigurationArgs();
    }
}
