// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Escloud.Outputs
{

    [OutputType]
    public sealed class GetInstancesInstanceInstanceConfigurationResult
    {
        /// <summary>
        /// The user name of instance.
        /// </summary>
        public readonly string AdminUserName;
        /// <summary>
        /// The charge type of instance.
        /// </summary>
        public readonly string ChargeType;
        /// <summary>
        /// whether enable https.
        /// </summary>
        public readonly bool EnableHttps;
        /// <summary>
        /// Whether enable pure master.
        /// </summary>
        public readonly bool EnablePureMaster;
        /// <summary>
        /// The node number of host.
        /// </summary>
        public readonly int HotNodeNumber;
        /// <summary>
        /// The node resource spec of host.
        /// </summary>
        public readonly Outputs.GetInstancesInstanceInstanceConfigurationHotNodeResourceSpecResult HotNodeResourceSpec;
        /// <summary>
        /// The node storage spec of host.
        /// </summary>
        public readonly Outputs.GetInstancesInstanceInstanceConfigurationHotNodeStorageSpecResult HotNodeStorageSpec;
        /// <summary>
        /// The name of instance.
        /// </summary>
        public readonly string InstanceName;
        /// <summary>
        /// The node number of kibana.
        /// </summary>
        public readonly int KibanaNodeNumber;
        /// <summary>
        /// The node resource spec of kibana.
        /// </summary>
        public readonly Outputs.GetInstancesInstanceInstanceConfigurationKibanaNodeResourceSpecResult KibanaNodeResourceSpec;
        /// <summary>
        /// The node storage spec of kibana.
        /// </summary>
        public readonly Outputs.GetInstancesInstanceInstanceConfigurationKibanaNodeStorageSpecResult KibanaNodeStorageSpec;
        /// <summary>
        /// The node number of master.
        /// </summary>
        public readonly int MasterNodeNumber;
        /// <summary>
        /// The node resource spec of master.
        /// </summary>
        public readonly Outputs.GetInstancesInstanceInstanceConfigurationMasterNodeResourceSpecResult MasterNodeResourceSpec;
        /// <summary>
        /// The node storage spec of master.
        /// </summary>
        public readonly Outputs.GetInstancesInstanceInstanceConfigurationMasterNodeStorageSpecResult MasterNodeStorageSpec;
        /// <summary>
        /// The period of project.
        /// </summary>
        public readonly int Period;
        /// <summary>
        /// The name of project.
        /// </summary>
        public readonly string ProjectName;
        /// <summary>
        /// The region info of instance.
        /// </summary>
        public readonly string RegionId;
        /// <summary>
        /// The subnet info.
        /// </summary>
        public readonly Outputs.GetInstancesInstanceInstanceConfigurationSubnetResult Subnet;
        /// <summary>
        /// The version of plugin.
        /// </summary>
        public readonly string Version;
        /// <summary>
        /// The vpc info.
        /// </summary>
        public readonly Outputs.GetInstancesInstanceInstanceConfigurationVpcResult Vpc;
        /// <summary>
        /// The zoneId of instance.
        /// </summary>
        public readonly string ZoneId;
        /// <summary>
        /// The zone number of instance.
        /// </summary>
        public readonly int ZoneNumber;

        [OutputConstructor]
        private GetInstancesInstanceInstanceConfigurationResult(
            string adminUserName,

            string chargeType,

            bool enableHttps,

            bool enablePureMaster,

            int hotNodeNumber,

            Outputs.GetInstancesInstanceInstanceConfigurationHotNodeResourceSpecResult hotNodeResourceSpec,

            Outputs.GetInstancesInstanceInstanceConfigurationHotNodeStorageSpecResult hotNodeStorageSpec,

            string instanceName,

            int kibanaNodeNumber,

            Outputs.GetInstancesInstanceInstanceConfigurationKibanaNodeResourceSpecResult kibanaNodeResourceSpec,

            Outputs.GetInstancesInstanceInstanceConfigurationKibanaNodeStorageSpecResult kibanaNodeStorageSpec,

            int masterNodeNumber,

            Outputs.GetInstancesInstanceInstanceConfigurationMasterNodeResourceSpecResult masterNodeResourceSpec,

            Outputs.GetInstancesInstanceInstanceConfigurationMasterNodeStorageSpecResult masterNodeStorageSpec,

            int period,

            string projectName,

            string regionId,

            Outputs.GetInstancesInstanceInstanceConfigurationSubnetResult subnet,

            string version,

            Outputs.GetInstancesInstanceInstanceConfigurationVpcResult vpc,

            string zoneId,

            int zoneNumber)
        {
            AdminUserName = adminUserName;
            ChargeType = chargeType;
            EnableHttps = enableHttps;
            EnablePureMaster = enablePureMaster;
            HotNodeNumber = hotNodeNumber;
            HotNodeResourceSpec = hotNodeResourceSpec;
            HotNodeStorageSpec = hotNodeStorageSpec;
            InstanceName = instanceName;
            KibanaNodeNumber = kibanaNodeNumber;
            KibanaNodeResourceSpec = kibanaNodeResourceSpec;
            KibanaNodeStorageSpec = kibanaNodeStorageSpec;
            MasterNodeNumber = masterNodeNumber;
            MasterNodeResourceSpec = masterNodeResourceSpec;
            MasterNodeStorageSpec = masterNodeStorageSpec;
            Period = period;
            ProjectName = projectName;
            RegionId = regionId;
            Subnet = subnet;
            Version = version;
            Vpc = vpc;
            ZoneId = zoneId;
            ZoneNumber = zoneNumber;
        }
    }
}
