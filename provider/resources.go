// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package volcengine

import (
	"fmt"
	"path/filepath"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	shimv1 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v1"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"github.com/volcengine/pulumi-volcengine/provider/pkg/version"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	volcenginePkg = "volcengine"
	// modules:
	//mainMod          = "index" // the volcengine module
	volcengineCdnMod               = "cdn"
	volcenginePrivatelinkMod       = "privatelink"
	volcengineRedisMod             = "redis"
	volcengineTransitRouterMod     = "transit_router"
	volcengineVkeMod               = "vke"
	volcengineBandwidthPackageMod  = "bandwidth_package"
	volcengineBioosMod             = "bioos"
	volcengineDnsMod               = "dns"
	volcengineEscloudV2Mod         = "escloud_v2"
	volcengineOrganizationMod      = "organization"
	volcengineTlsMod               = "tls"
	volcengineAlbMod               = "alb"
	volcengineCenMod               = "cen"
	volcengineCloudIdentityMod     = "cloud_identity"
	volcengineVedbMysqlMod         = "vedb_mysql"
	volcengineVmpMod               = "vmp"
	volcengineClbMod               = "clb"
	volcengineCrMod                = "cr"
	volcengineDirectConnectMod     = "direct_connect"
	volcengineEbsMod               = "ebs"
	volcengineRdsMod               = "rds"
	volcengineVeecpMod             = "veecp"
	volcengineWafMod               = "waf"
	volcengineEipMod               = "eip"
	volcenginePrivateZoneMod       = "private_zone"
	volcengineRdsV2Mod             = "rds_v2"
	volcengineVpnMod               = "vpn"
	volcengineCloudMonitorMod      = "cloud_monitor"
	volcengineFinancialRelationMod = "financial_relation"
	volcengineRdsMysqlMod          = "rds_mysql"
	volcengineRdsPostgresqlMod     = "rds_postgresql"
	volcengineTosMod               = "tos"
	volcengineVefaasMod            = "vefaas"
	volcengineVepfsMod             = "vepfs"
	volcengineAutoscalingMod       = "autoscaling"
	volcengineCloudFirewallMod     = "cloud_firewall"
	volcengineEcsMod               = "ecs"
	volcengineIamMod               = "iam"
	volcengineKafkaMod             = "kafka"
	volcengineMongodbMod           = "mongodb"
	volcengineNasMod               = "nas"
	volcengineVpcMod               = "vpc"
	volcengineApigMod              = "apig"
	volcengineCloudfsMod           = "cloudfs"
	volcengineEscloudMod           = "escloud"
	volcengineKmsMod               = "kms"
	volcengineNatMod               = "nat"
	volcengineRabbitmqMod          = "rabbitmq"
	volcengineRdsMssqlMod          = "rds_mssql"
	volcengineRocketmqMod          = "rocketmq"
	volcengineVeenedgeMod          = "veenedge"
)

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(vars resource.PropertyMap, c shim.ResourceConfig) error {
	return nil
}

func resourceType(mod string, res string) tokens.Type {
	return tfbridge.MakeResource(volcenginePkg, mod, res)
}

func dataSourceType(mod string, res string) tokens.ModuleMember {
	return tfbridge.MakeDataSource(volcenginePkg, mod, res)
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := shimv1.NewProvider(volcengine.Provider().(*schema.Provider))

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:    p,
		Name: "volcengine",
		// DisplayName is a way to be able to change the casing of the provider
		// name when being displayed on the Pulumi registry
		DisplayName: "Volcengine",
		// The default publisher for all packages is Pulumi.
		// Change this to your personal name (or a company name) that you
		// would like to be shown in the Pulumi Registry if this package is published
		// there.
		Publisher: "Volcengine",
		// LogoURL is optional but useful to help identify your package in the Pulumi Registry
		// if this package is published there.
		LogoURL: "https://avatars.githubusercontent.com/u/67365215",
		// PluginDownloadURL is an optional URL used to download the Provider
		// for use in Pulumi programs
		PluginDownloadURL: "github://api.github.com/volcengine",
		Description:       "A Pulumi package for creating and managing volcengine cloud resources.",
		// category/cloud tag helps with categorizing the package in the Pulumi Registry.
		// For all available categories, see `Keywords` in
		// https://www.pulumi.com/docs/guides/pulumi-packages/schema/#package.
		Keywords:   []string{"pulumi", "volcengine", "category/cloud"},
		License:    "Apache-2.0",
		Homepage:   "https://volcengine.com",
		Repository: "https://github.com/volcengine/pulumi-volcengine",
		GitHubOrg:  "volcengine",
		Config: map[string]*tfbridge.SchemaInfo{
			"region": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"VOLCENGINE_REGION"},
				},
			},
			"access_key": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"VOLCENGINE_ACCESS_KEY"},
				},
			},
			"secret_key": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"VOLCENGINE_SECRET_KEY"},
				},
			},
			"endpoint": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"VOLCENGINE_ENDPOINT"},
				},
			},
		},
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			"volcengine_traffic_mirror_filter_rule":                       {Tok: resourceType(volcengineVpcMod, "TrafficMirrorFilterRule")},
			"volcengine_network_interface_attach":                         {Tok: resourceType(volcengineVpcMod, "NetworkInterfaceAttach")},
			"volcengine_vpc_ipv6_gateway":                                 {Tok: resourceType(volcengineVpcMod, "Ipv6Gateway")},
			"volcengine_ha_vip_associate":                                 {Tok: resourceType(volcengineVpcMod, "HaVipAssociate")},
			"volcengine_vpc_cidr_block_associate":                         {Tok: resourceType(volcengineVpcMod, "CidrBlockAssociate")},
			"volcengine_flow_log":                                         {Tok: resourceType(volcengineVpcMod, "FlowLog")},
			"volcengine_route_entry":                                      {Tok: resourceType(volcengineVpcMod, "RouteEntry")},
			"volcengine_network_acl":                                      {Tok: resourceType(volcengineVpcMod, "NetworkAcl")},
			"volcengine_security_group_rule":                              {Tok: resourceType(volcengineVpcMod, "SecurityGroupRule")},
			"volcengine_subnet":                                           {Tok: resourceType(volcengineVpcMod, "Subnet")},
			"volcengine_traffic_mirror_filter":                            {Tok: resourceType(volcengineVpcMod, "TrafficMirrorFilter")},
			"volcengine_network_acl_associate":                            {Tok: resourceType(volcengineVpcMod, "NetworkAclAssociate")},
			"volcengine_vpc":                                              {Tok: resourceType(volcengineVpcMod, "Vpc")},
			"volcengine_security_group":                                   {Tok: resourceType(volcengineVpcMod, "SecurityGroup")},
			"volcengine_traffic_mirror_target":                            {Tok: resourceType(volcengineVpcMod, "TrafficMirrorTarget")},
			"volcengine_vpc_prefix_list":                                  {Tok: resourceType(volcengineVpcMod, "PrefixList")},
			"volcengine_vpc_ipv6_address_bandwidth":                       {Tok: resourceType(volcengineVpcMod, "Ipv6AddressBandwidth")},
			"volcengine_route_table":                                      {Tok: resourceType(volcengineVpcMod, "RouteTable")},
			"volcengine_traffic_mirror_session":                           {Tok: resourceType(volcengineVpcMod, "TrafficMirrorSession")},
			"volcengine_flow_log_active":                                  {Tok: resourceType(volcengineVpcMod, "FlowLogActive")},
			"volcengine_route_table_associate":                            {Tok: resourceType(volcengineVpcMod, "RouteTableAssociate")},
			"volcengine_network_interface":                                {Tok: resourceType(volcengineVpcMod, "NetworkInterface")},
			"volcengine_ha_vip":                                           {Tok: resourceType(volcengineVpcMod, "HaVip")},
			"volcengine_nas_mount_point":                                  {Tok: resourceType(volcengineNasMod, "MountPoint")},
			"volcengine_nas_permission_group":                             {Tok: resourceType(volcengineNasMod, "PermissionGroup")},
			"volcengine_nas_auto_snapshot_policy":                         {Tok: resourceType(volcengineNasMod, "AutoSnapshotPolicy")},
			"volcengine_nas_file_system":                                  {Tok: resourceType(volcengineNasMod, "FileSystem")},
			"volcengine_nas_auto_snapshot_policy_apply":                   {Tok: resourceType(volcengineNasMod, "AutoSnapshotPolicyApply")},
			"volcengine_nas_snapshot":                                     {Tok: resourceType(volcengineNasMod, "Snapshot")},
			"volcengine_rds_ip_list":                                      {Tok: resourceType(volcengineRdsMod, "IpList")},
			"volcengine_rds_database":                                     {Tok: resourceType(volcengineRdsMod, "Database")},
			"volcengine_rds_account":                                      {Tok: resourceType(volcengineRdsMod, "Account")},
			"volcengine_rds_parameter_template":                           {Tok: resourceType(volcengineRdsMod, "ParameterTemplate")},
			"volcengine_rds_instance":                                     {Tok: resourceType(volcengineRdsMod, "Instance")},
			"volcengine_rds_account_privilege":                            {Tok: resourceType(volcengineRdsMod, "AccountPrivilege")},
			"volcengine_listener":                                         {Tok: resourceType(volcengineClbMod, "Listener")},
			"volcengine_acl":                                              {Tok: resourceType(volcengineClbMod, "Acl")},
			"volcengine_acl_entry":                                        {Tok: resourceType(volcengineClbMod, "AclEntry")},
			"volcengine_clb_rule":                                         {Tok: resourceType(volcengineClbMod, "Rule")},
			"volcengine_clb":                                              {Tok: resourceType(volcengineClbMod, "Clb")},
			"volcengine_server_group":                                     {Tok: resourceType(volcengineClbMod, "ServerGroup")},
			"volcengine_server_group_server":                              {Tok: resourceType(volcengineClbMod, "ServerGroupServer")},
			"volcengine_certificate":                                      {Tok: resourceType(volcengineClbMod, "Certificate")},
			"volcengine_financial_relation":                               {Tok: resourceType(volcengineFinancialRelationMod, "FinancialRelation")},
			"volcengine_snat_entry":                                       {Tok: resourceType(volcengineNatMod, "SnatEntry")},
			"volcengine_nat_gateway":                                      {Tok: resourceType(volcengineNatMod, "Gateway")},
			"volcengine_dnat_entry":                                       {Tok: resourceType(volcengineNatMod, "DnatEntry")},
			"volcengine_direct_connect_gateway":                           {Tok: resourceType(volcengineDirectConnectMod, "Gateway")},
			"volcengine_direct_connect_gateway_route":                     {Tok: resourceType(volcengineDirectConnectMod, "GatewayRoute")},
			"volcengine_direct_connect_connection":                        {Tok: resourceType(volcengineDirectConnectMod, "Connection")},
			"volcengine_direct_connect_virtual_interface":                 {Tok: resourceType(volcengineDirectConnectMod, "VirtualInterface")},
			"volcengine_direct_connect_bgp_peer":                          {Tok: resourceType(volcengineDirectConnectMod, "BgpPeer")},
			"volcengine_vke_node_pool":                                    {Tok: resourceType(volcengineVkeMod, "NodePool")},
			"volcengine_vke_kubeconfig":                                   {Tok: resourceType(volcengineVkeMod, "Kubeconfig")},
			"volcengine_vke_default_node_pool_batch_attach":               {Tok: resourceType(volcengineVkeMod, "DefaultNodePoolBatchAttach")},
			"volcengine_vke_default_node_pool":                            {Tok: resourceType(volcengineVkeMod, "DefaultNodePool")},
			"volcengine_vke_node":                                         {Tok: resourceType(volcengineVkeMod, "Node")},
			"volcengine_vke_cluster":                                      {Tok: resourceType(volcengineVkeMod, "Cluster")},
			"volcengine_vke_addon":                                        {Tok: resourceType(volcengineVkeMod, "Addon")},
			"volcengine_vke_permission":                                   {Tok: resourceType(volcengineVkeMod, "Permission")},
			"volcengine_kms_key_rotation":                                 {Tok: resourceType(volcengineKmsMod, "KeyRotation")},
			"volcengine_kms_key_archive":                                  {Tok: resourceType(volcengineKmsMod, "KeyArchive")},
			"volcengine_kms_keyring":                                      {Tok: resourceType(volcengineKmsMod, "Keyring")},
			"volcengine_kms_key":                                          {Tok: resourceType(volcengineKmsMod, "Key")},
			"volcengine_kms_secret":                                       {Tok: resourceType(volcengineKmsMod, "Secret")},
			"volcengine_kms_key_enable":                                   {Tok: resourceType(volcengineKmsMod, "KeyEnable")},
			"volcengine_cloud_monitor_event_rule":                         {Tok: resourceType(volcengineCloudMonitorMod, "EventRule")},
			"volcengine_cloud_monitor_contact_group":                      {Tok: resourceType(volcengineCloudMonitorMod, "ContactGroup")},
			"volcengine_cloud_monitor_contact":                            {Tok: resourceType(volcengineCloudMonitorMod, "Contact")},
			"volcengine_cloud_monitor_rule":                               {Tok: resourceType(volcengineCloudMonitorMod, "Rule")},
			"volcengine_iam_user_group_attachment":                        {Tok: resourceType(volcengineIamMod, "UserGroupAttachment")},
			"volcengine_iam_user_policy_attachment":                       {Tok: resourceType(volcengineIamMod, "UserPolicyAttachment")},
			"volcengine_iam_login_profile":                                {Tok: resourceType(volcengineIamMod, "LoginProfile")},
			"volcengine_iam_user_group_policy_attachment":                 {Tok: resourceType(volcengineIamMod, "UserGroupPolicyAttachment")},
			"volcengine_iam_role":                                         {Tok: resourceType(volcengineIamMod, "Role")},
			"volcengine_iam_access_key":                                   {Tok: resourceType(volcengineIamMod, "AccessKey")},
			"volcengine_iam_saml_provider":                                {Tok: resourceType(volcengineIamMod, "SamlProvider")},
			"volcengine_iam_policy":                                       {Tok: resourceType(volcengineIamMod, "Policy")},
			"volcengine_iam_user_group":                                   {Tok: resourceType(volcengineIamMod, "UserGroup")},
			"volcengine_iam_role_policy_attachment":                       {Tok: resourceType(volcengineIamMod, "RolePolicyAttachment")},
			"volcengine_iam_user":                                         {Tok: resourceType(volcengineIamMod, "User")},
			"volcengine_privatelink_vpc_endpoint_service":                 {Tok: resourceType(volcenginePrivatelinkMod, "VpcEndpointService")},
			"volcengine_privatelink_vpc_endpoint_service_resource":        {Tok: resourceType(volcenginePrivatelinkMod, "VpcEndpointServiceResource")},
			"volcengine_privatelink_security_group":                       {Tok: resourceType(volcenginePrivatelinkMod, "SecurityGroup")},
			"volcengine_privatelink_vpc_endpoint_zone":                    {Tok: resourceType(volcenginePrivatelinkMod, "VpcEndpointZone")},
			"volcengine_privatelink_vpc_endpoint":                         {Tok: resourceType(volcenginePrivatelinkMod, "VpcEndpoint")},
			"volcengine_privatelink_vpc_endpoint_service_permission":      {Tok: resourceType(volcenginePrivatelinkMod, "VpcEndpointServicePermission")},
			"volcengine_privatelink_vpc_endpoint_connection":              {Tok: resourceType(volcenginePrivatelinkMod, "VpcEndpointConnection")},
			"volcengine_rocketmq_allow_list_associate":                    {Tok: resourceType(volcengineRocketmqMod, "RocketMQAllowListAssociate")},
			"volcengine_rocketmq_topic":                                   {Tok: resourceType(volcengineRocketmqMod, "RocketMQTopic")},
			"volcengine_rocketmq_instance":                                {Tok: resourceType(volcengineRocketmqMod, "RocketMQInstance")},
			"volcengine_rocketmq_group":                                   {Tok: resourceType(volcengineRocketmqMod, "RocketMQGroup")},
			"volcengine_rocketmq_public_address":                          {Tok: resourceType(volcengineRocketmqMod, "RocketMQPublicAddress")},
			"volcengine_rocketmq_access_key":                              {Tok: resourceType(volcengineRocketmqMod, "RocketMQAccessKey")},
			"volcengine_rocketmq_allow_list":                              {Tok: resourceType(volcengineRocketmqMod, "RocketMQAllowList")},
			"volcengine_transit_router_route_entry":                       {Tok: resourceType(volcengineTransitRouterMod, "RouteEntry")},
			"volcengine_transit_router_route_table":                       {Tok: resourceType(volcengineTransitRouterMod, "RouteTable")},
			"volcengine_transit_router":                                   {Tok: resourceType(volcengineTransitRouterMod, "TransitRouter")},
			"volcengine_transit_router_bandwidth_package":                 {Tok: resourceType(volcengineTransitRouterMod, "BandwidthPackage")},
			"volcengine_transit_router_route_table_association":           {Tok: resourceType(volcengineTransitRouterMod, "RouteTableAssociation")},
			"volcengine_transit_router_route_table_propagation":           {Tok: resourceType(volcengineTransitRouterMod, "RouteTablePropagation")},
			"volcengine_transit_router_grant_rule":                        {Tok: resourceType(volcengineTransitRouterMod, "GrantRule")},
			"volcengine_transit_router_vpc_attachment":                    {Tok: resourceType(volcengineTransitRouterMod, "VpcAttachment")},
			"volcengine_transit_router_direct_connect_gateway_attachment": {Tok: resourceType(volcengineTransitRouterMod, "DirectConnectGatewayAttachment")},
			"volcengine_transit_router_peer_attachment":                   {Tok: resourceType(volcengineTransitRouterMod, "PeerAttachment")},
			"volcengine_transit_router_shared_transit_router_state":       {Tok: resourceType(volcengineTransitRouterMod, "SharedTransitRouterState")},
			"volcengine_transit_router_vpn_attachment":                    {Tok: resourceType(volcengineTransitRouterMod, "VpnAttachment")},
			"volcengine_alb":                                              {Tok: resourceType(volcengineAlbMod, "Alb")},
			"volcengine_alb_listener_domain_extension":                    {Tok: resourceType(volcengineAlbMod, "ListenerDomainExtension")},
			"volcengine_alb_health_check_template":                        {Tok: resourceType(volcengineAlbMod, "HealthCheckTemplate")},
			"volcengine_alb_certificate":                                  {Tok: resourceType(volcengineAlbMod, "Certificate")},
			"volcengine_alb_ca_certificate":                               {Tok: resourceType(volcengineAlbMod, "CACertificate")},
			"volcengine_alb_rule":                                         {Tok: resourceType(volcengineAlbMod, "Rule")},
			"volcengine_alb_acl":                                          {Tok: resourceType(volcengineAlbMod, "Acl")},
			"volcengine_alb_listener":                                     {Tok: resourceType(volcengineAlbMod, "Listener")},
			"volcengine_alb_server_group_server":                          {Tok: resourceType(volcengineAlbMod, "ServerGroupServer")},
			"volcengine_alb_customized_cfg":                               {Tok: resourceType(volcengineAlbMod, "CustomizedCfg")},
			"volcengine_alb_server_group":                                 {Tok: resourceType(volcengineAlbMod, "ServerGroup")},
			"volcengine_dns_record":                                       {Tok: resourceType(volcengineDnsMod, "Record")},
			"volcengine_dns_backup_schedule":                              {Tok: resourceType(volcengineDnsMod, "BackupSchedule")},
			"volcengine_dns_zone":                                         {Tok: resourceType(volcengineDnsMod, "Zone")},
			"volcengine_dns_backup":                                       {Tok: resourceType(volcengineDnsMod, "Backup")},
			"volcengine_mongodb_instance":                                 {Tok: resourceType(volcengineMongodbMod, "Instance")},
			"volcengine_mongodb_ssl_state":                                {Tok: resourceType(volcengineMongodbMod, "SslState")},
			"volcengine_mongodb_endpoint":                                 {Tok: resourceType(volcengineMongodbMod, "Endpoint")},
			"volcengine_mongodb_instance_parameter":                       {Tok: resourceType(volcengineMongodbMod, "InstanceParameter")},
			"volcengine_mongodb_allow_list":                               {Tok: resourceType(volcengineMongodbMod, "MongoAllowList")},
			"volcengine_mongodb_allow_list_associate":                     {Tok: resourceType(volcengineMongodbMod, "MongoAllowListAssociate")},
			"volcengine_mongodb_account":                                  {Tok: resourceType(volcengineMongodbMod, "Account")},
			"volcengine_cdn_certificate":                                  {Tok: resourceType(volcengineCdnMod, "CdnCertificate")},
			"volcengine_cdn_shared_config":                                {Tok: resourceType(volcengineCdnMod, "SharedConfig")},
			"volcengine_cdn_domain":                                       {Tok: resourceType(volcengineCdnMod, "CdnDomain")},
			"volcengine_cfw_dns_control_policy":                           {Tok: resourceType(volcengineCloudFirewallMod, "CfwDnsControlPolicy")},
			"volcengine_cfw_vpc_firewall_acl_rule_priority":               {Tok: resourceType(volcengineCloudFirewallMod, "CfwVpcFirewallAclRulePriority")},
			"volcengine_cfw_address_book":                                 {Tok: resourceType(volcengineCloudFirewallMod, "CfwAddressBook")},
			"volcengine_cfw_control_policy_priority":                      {Tok: resourceType(volcengineCloudFirewallMod, "CfwControlPolicyPriority")},
			"volcengine_cfw_nat_firewall_control_policy_priority":         {Tok: resourceType(volcengineCloudFirewallMod, "CfwNatFirewallControlPolicyPriority")},
			"volcengine_cfw_vpc_firewall_acl_rule":                        {Tok: resourceType(volcengineCloudFirewallMod, "CfwVpcFirewallAclRule")},
			"volcengine_cfw_nat_firewall_control_policy":                  {Tok: resourceType(volcengineCloudFirewallMod, "CfwNatFirewallControlPolicy")},
			"volcengine_cfw_control_policy":                               {Tok: resourceType(volcengineCloudFirewallMod, "CfwControlPolicy")},
			"volcengine_ecs_invocation":                                   {Tok: resourceType(volcengineEcsMod, "Invocation")},
			"volcengine_image_share_permission":                           {Tok: resourceType(volcengineEcsMod, "ImageSharePermission")},
			"volcengine_ecs_key_pair_associate":                           {Tok: resourceType(volcengineEcsMod, "KeyPairAssociate")},
			"volcengine_image_import":                                     {Tok: resourceType(volcengineEcsMod, "ImageImport")},
			"volcengine_ecs_hpc_cluster":                                  {Tok: resourceType(volcengineEcsMod, "HpcCluster")},
			"volcengine_ecs_launch_template":                              {Tok: resourceType(volcengineEcsMod, "LaunchTemplate")},
			"volcengine_ecs_deployment_set_associate":                     {Tok: resourceType(volcengineEcsMod, "DeploymentSetAssociate")},
			"volcengine_ecs_command":                                      {Tok: resourceType(volcengineEcsMod, "Command")},
			"volcengine_image":                                            {Tok: resourceType(volcengineEcsMod, "Image")},
			"volcengine_ecs_key_pair":                                     {Tok: resourceType(volcengineEcsMod, "KeyPair")},
			"volcengine_ecs_instance":                                     {Tok: resourceType(volcengineEcsMod, "Instance")},
			"volcengine_ecs_deployment_set":                               {Tok: resourceType(volcengineEcsMod, "DeploymentSet")},
			"volcengine_ecs_instance_state":                               {Tok: resourceType(volcengineEcsMod, "State")},
			"volcengine_iam_role_attachment":                              {Tok: resourceType(volcengineEcsMod, "IamRoleAttachment")},
			"volcengine_rabbitmq_instance_plugin":                         {Tok: resourceType(volcengineRabbitmqMod, "InstancePlugin")},
			"volcengine_rabbitmq_public_address":                          {Tok: resourceType(volcengineRabbitmqMod, "PublicAddress")},
			"volcengine_rabbitmq_instance":                                {Tok: resourceType(volcengineRabbitmqMod, "Instance")},
			"volcengine_organization_unit":                                {Tok: resourceType(volcengineOrganizationMod, "Unit")},
			"volcengine_organization_service_control_policy_enabler":      {Tok: resourceType(volcengineOrganizationMod, "ServiceControlPolicyEnabler")},
			"volcengine_organization_account":                             {Tok: resourceType(volcengineOrganizationMod, "Account")},
			"volcengine_organization":                                     {Tok: resourceType(volcengineOrganizationMod, "Organization")},
			"volcengine_organization_service_control_policy":              {Tok: resourceType(volcengineOrganizationMod, "ServiceControlPolicy")},
			"volcengine_organization_service_control_policy_attachment":   {Tok: resourceType(volcengineOrganizationMod, "ServiceControlPolicyAttachment")},
			"volcengine_tos_bucket_inventory":                             {Tok: resourceType(volcengineTosMod, "BucketInventory")},
			"volcengine_tos_bucket":                                       {Tok: resourceType(volcengineTosMod, "Bucket")},
			"volcengine_tos_bucket_policy":                                {Tok: resourceType(volcengineTosMod, "BucketPolicy")},
			"volcengine_tos_bucket_realtime_log":                          {Tok: resourceType(volcengineTosMod, "BucketRealtimeLog")},
			"volcengine_tos_object":                                       {Tok: resourceType(volcengineTosMod, "BucketObject")},
			"volcengine_apig_upstream":                                    {Tok: resourceType(volcengineApigMod, "ApigUpstream")},
			"volcengine_apig_upstream_source":                             {Tok: resourceType(volcengineApigMod, "ApigUpstreamSource")},
			"volcengine_apig_upstream_version":                            {Tok: resourceType(volcengineApigMod, "ApigUpstreamVersion")},
			"volcengine_apig_gateway_service":                             {Tok: resourceType(volcengineApigMod, "ApigGatewayService")},
			"volcengine_apig_gateway":                                     {Tok: resourceType(volcengineApigMod, "ApigGateway")},
			"volcengine_apig_route":                                       {Tok: resourceType(volcengineApigMod, "ApigRoute")},
			"volcengine_apig_custom_domain":                               {Tok: resourceType(volcengineApigMod, "ApigCustomDomain")},
			"volcengine_rds_instance_v2":                                  {Tok: resourceType(volcengineRdsV2Mod, "RdsInstanceV2")},
			"volcengine_kafka_public_address":                             {Tok: resourceType(volcengineKafkaMod, "PublicAddress")},
			"volcengine_kafka_sasl_user":                                  {Tok: resourceType(volcengineKafkaMod, "SaslUser")},
			"volcengine_kafka_allow_list":                                 {Tok: resourceType(volcengineKafkaMod, "AllowList")},
			"volcengine_kafka_group":                                      {Tok: resourceType(volcengineKafkaMod, "Group")},
			"volcengine_kafka_topic":                                      {Tok: resourceType(volcengineKafkaMod, "Topic")},
			"volcengine_kafka_allow_list_associate":                       {Tok: resourceType(volcengineKafkaMod, "AllowListAssociate")},
			"volcengine_kafka_instance":                                   {Tok: resourceType(volcengineKafkaMod, "Instance")},
			"volcengine_private_zone_record":                              {Tok: resourceType(volcenginePrivateZoneMod, "Record")},
			"volcengine_private_zone_resolver_rule":                       {Tok: resourceType(volcenginePrivateZoneMod, "ResolverRule")},
			"volcengine_private_zone_user_vpc_authorization":              {Tok: resourceType(volcenginePrivateZoneMod, "UserVpcAuthorization")},
			"volcengine_private_zone_resolver_endpoint":                   {Tok: resourceType(volcenginePrivateZoneMod, "ResolverEndpoint")},
			"volcengine_private_zone":                                     {Tok: resourceType(volcenginePrivateZoneMod, "PrivateZone")},
			"volcengine_private_zone_record_weight_enabler":               {Tok: resourceType(volcenginePrivateZoneMod, "RecordWeightEnabler")},
			"volcengine_vefaas_release":                                   {Tok: resourceType(volcengineVefaasMod, "Release")},
			"volcengine_vefaas_kafka_trigger":                             {Tok: resourceType(volcengineVefaasMod, "KafkaTrigger")},
			"volcengine_vefaas_function":                                  {Tok: resourceType(volcengineVefaasMod, "Function")},
			"volcengine_vefaas_timer":                                     {Tok: resourceType(volcengineVefaasMod, "Timer")},
			"volcengine_waf_custom_bot":                                   {Tok: resourceType(volcengineWafMod, "CustomBot")},
			"volcengine_waf_acl_rule":                                     {Tok: resourceType(volcengineWafMod, "AclRule")},
			"volcengine_waf_instance_ctl":                                 {Tok: resourceType(volcengineWafMod, "InstanceCtl")},
			"volcengine_waf_system_bot":                                   {Tok: resourceType(volcengineWafMod, "SystemBot")},
			"volcengine_waf_custom_page":                                  {Tok: resourceType(volcengineWafMod, "CustomPage")},
			"volcengine_waf_vulnerability":                                {Tok: resourceType(volcengineWafMod, "Vulnerability")},
			"volcengine_waf_domain":                                       {Tok: resourceType(volcengineWafMod, "WafDomain")},
			"volcengine_waf_host_group":                                   {Tok: resourceType(volcengineWafMod, "HostGroup")},
			"volcengine_waf_bot_analyse_protect_rule":                     {Tok: resourceType(volcengineWafMod, "BotAnalyseProtectRule")},
			"volcengine_waf_cc_rule":                                      {Tok: resourceType(volcengineWafMod, "CcRule")},
			"volcengine_waf_ip_group":                                     {Tok: resourceType(volcengineWafMod, "IpGroup")},
			"volcengine_vepfs_fileset":                                    {Tok: resourceType(volcengineVepfsMod, "Fileset")},
			"volcengine_vepfs_mount_service_attachment":                   {Tok: resourceType(volcengineVepfsMod, "MountServiceAttachment")},
			"volcengine_vepfs_file_system":                                {Tok: resourceType(volcengineVepfsMod, "FileSystem")},
			"volcengine_vepfs_mount_service":                              {Tok: resourceType(volcengineVepfsMod, "MountService")},
			"volcengine_scaling_configuration":                            {Tok: resourceType(volcengineAutoscalingMod, "ScalingConfiguration")},
			"volcengine_scaling_group_enabler":                            {Tok: resourceType(volcengineAutoscalingMod, "ScalingGroupEnabler")},
			"volcengine_scaling_configuration_attachment":                 {Tok: resourceType(volcengineAutoscalingMod, "ScalingConfigurationAttachment")},
			"volcengine_scaling_instance_attachment":                      {Tok: resourceType(volcengineAutoscalingMod, "ScalingInstanceAttachment")},
			"volcengine_scaling_lifecycle_hook":                           {Tok: resourceType(volcengineAutoscalingMod, "ScalingLifecycleHook")},
			"volcengine_scaling_group":                                    {Tok: resourceType(volcengineAutoscalingMod, "ScalingGroup")},
			"volcengine_scaling_policy":                                   {Tok: resourceType(volcengineAutoscalingMod, "ScalingPolicy")},
			"volcengine_bandwidth_package_attachment":                     {Tok: resourceType(volcengineBandwidthPackageMod, "Attachment")},
			"volcengine_bandwidth_package":                                {Tok: resourceType(volcengineBandwidthPackageMod, "BandwidthPackage")},
			"volcengine_veecp_edge_node":                                  {Tok: resourceType(volcengineVeecpMod, "EdgeNode")},
			"volcengine_veecp_node_pool":                                  {Tok: resourceType(volcengineVeecpMod, "NodePool")},
			"volcengine_veecp_edge_node_pool":                             {Tok: resourceType(volcengineVeecpMod, "EdgeNodePool")},
			"volcengine_veecp_batch_edge_machine":                         {Tok: resourceType(volcengineVeecpMod, "BatchEdgeMachine")},
			"volcengine_veecp_addon":                                      {Tok: resourceType(volcengineVeecpMod, "Addon")},
			"volcengine_veecp_cluster":                                    {Tok: resourceType(volcengineVeecpMod, "Cluster")},
			"volcengine_veenedge_cloud_server":                            {Tok: resourceType(volcengineVeenedgeMod, "CloudServer")},
			"volcengine_veenedge_instance":                                {Tok: resourceType(volcengineVeenedgeMod, "Instance")},
			"volcengine_veenedge_vpc":                                     {Tok: resourceType(volcengineVeenedgeMod, "Vpc")},
			"volcengine_eip_address":                                      {Tok: resourceType(volcengineEipMod, "Address")},
			"volcengine_eip_associate":                                    {Tok: resourceType(volcengineEipMod, "Associate")},
			"volcengine_vedb_mysql_allowlist_associate":                   {Tok: resourceType(volcengineVedbMysqlMod, "AllowlistAssociate")},
			"volcengine_vedb_mysql_endpoint_public_address":               {Tok: resourceType(volcengineVedbMysqlMod, "EndpointPublicAddress")},
			"volcengine_vedb_mysql_endpoint":                              {Tok: resourceType(volcengineVedbMysqlMod, "Endpoint")},
			"volcengine_vedb_mysql_backup":                                {Tok: resourceType(volcengineVedbMysqlMod, "Backup")},
			"volcengine_vedb_mysql_allowlist":                             {Tok: resourceType(volcengineVedbMysqlMod, "Allowlist")},
			"volcengine_vedb_mysql_database":                              {Tok: resourceType(volcengineVedbMysqlMod, "Database")},
			"volcengine_vedb_mysql_instance":                              {Tok: resourceType(volcengineVedbMysqlMod, "Instance")},
			"volcengine_vedb_mysql_account":                               {Tok: resourceType(volcengineVedbMysqlMod, "Account")},
			"volcengine_cloud_identity_permission_set":                    {Tok: resourceType(volcengineCloudIdentityMod, "PermissionSet")},
			"volcengine_cloud_identity_permission_set_provisioning":       {Tok: resourceType(volcengineCloudIdentityMod, "PermissionSetProvisioning")},
			"volcengine_cloud_identity_user_provisioning":                 {Tok: resourceType(volcengineCloudIdentityMod, "UserProvisioning")},
			"volcengine_cloud_identity_user":                              {Tok: resourceType(volcengineCloudIdentityMod, "User")},
			"volcengine_cloud_identity_group":                             {Tok: resourceType(volcengineCloudIdentityMod, "Group")},
			"volcengine_cloud_identity_user_attachment":                   {Tok: resourceType(volcengineCloudIdentityMod, "UserAttachment")},
			"volcengine_cloud_identity_permission_set_assignment":         {Tok: resourceType(volcengineCloudIdentityMod, "PermissionSetAssignment")},
			"volcengine_cloudfs_namespace":                                {Tok: resourceType(volcengineCloudfsMod, "Namespace")},
			"volcengine_cloudfs_file_system":                              {Tok: resourceType(volcengineCloudfsMod, "FileSystem")},
			"volcengine_cloudfs_access":                                   {Tok: resourceType(volcengineCloudfsMod, "Access")},
			"volcengine_ebs_auto_snapshot_policy":                         {Tok: resourceType(volcengineEbsMod, "AutoSnapshotPolicy")},
			"volcengine_volume_attach":                                    {Tok: resourceType(volcengineEbsMod, "VolumeAttach")},
			"volcengine_ebs_snapshot_group":                               {Tok: resourceType(volcengineEbsMod, "SnapshotGroup")},
			"volcengine_ebs_auto_snapshot_policy_attachment":              {Tok: resourceType(volcengineEbsMod, "AutoSnapshotPolicyAttachment")},
			"volcengine_volume":                                           {Tok: resourceType(volcengineEbsMod, "Volume")},
			"volcengine_ebs_snapshot":                                     {Tok: resourceType(volcengineEbsMod, "Snapshot")},
			"volcengine_redis_account":                                    {Tok: resourceType(volcengineRedisMod, "Account")},
			"volcengine_redis_backup":                                     {Tok: resourceType(volcengineRedisMod, "Backup")},
			"volcengine_redis_endpoint":                                   {Tok: resourceType(volcengineRedisMod, "Endpoint")},
			"volcengine_redis_backup_restore":                             {Tok: resourceType(volcengineRedisMod, "BackupRestore")},
			"volcengine_redis_parameter_group":                            {Tok: resourceType(volcengineRedisMod, "ParameterGroup")},
			"volcengine_redis_continuous_backup":                          {Tok: resourceType(volcengineRedisMod, "ContinuousBackup")},
			"volcengine_redis_allow_list":                                 {Tok: resourceType(volcengineRedisMod, "AllowList")},
			"volcengine_redis_allow_list_associate":                       {Tok: resourceType(volcengineRedisMod, "AllowListAssociate")},
			"volcengine_redis_instance_state":                             {Tok: resourceType(volcengineRedisMod, "State")},
			"volcengine_redis_instance":                                   {Tok: resourceType(volcengineRedisMod, "Instance")},
			"volcengine_rds_postgresql_schema":                            {Tok: resourceType(volcengineRdsPostgresqlMod, "Schema")},
			"volcengine_rds_postgresql_instance_readonly_node":            {Tok: resourceType(volcengineRdsPostgresqlMod, "InstanceReadonlyNode")},
			"volcengine_rds_postgresql_database":                          {Tok: resourceType(volcengineRdsPostgresqlMod, "Database")},
			"volcengine_rds_postgresql_account":                           {Tok: resourceType(volcengineRdsPostgresqlMod, "Account")},
			"volcengine_rds_postgresql_allowlist_associate":               {Tok: resourceType(volcengineRdsPostgresqlMod, "AllowlistAssociate")},
			"volcengine_rds_postgresql_allowlist":                         {Tok: resourceType(volcengineRdsPostgresqlMod, "Allowlist")},
			"volcengine_rds_postgresql_instance":                          {Tok: resourceType(volcengineRdsPostgresqlMod, "Instance")},
			"volcengine_escloud_instance":                                 {Tok: resourceType(volcengineEscloudMod, "Instance")},
			"volcengine_cr_endpoint":                                      {Tok: resourceType(volcengineCrMod, "Endpoint")},
			"volcengine_cr_registry_state":                                {Tok: resourceType(volcengineCrMod, "State")},
			"volcengine_cr_tag":                                           {Tok: resourceType(volcengineCrMod, "Tag")},
			"volcengine_cr_registry":                                      {Tok: resourceType(volcengineCrMod, "Registry")},
			"volcengine_cr_repository":                                    {Tok: resourceType(volcengineCrMod, "Repository")},
			"volcengine_cr_endpoint_acl_policy":                           {Tok: resourceType(volcengineCrMod, "EndpointAclPolicy")},
			"volcengine_cr_namespace":                                     {Tok: resourceType(volcengineCrMod, "Namespace")},
			"volcengine_cr_vpc_endpoint":                                  {Tok: resourceType(volcengineCrMod, "VpcEndpoint")},
			"volcengine_vmp_notify_group_policy":                          {Tok: resourceType(volcengineVmpMod, "NotifyGroupPolicy")},
			"volcengine_vmp_contact_group":                                {Tok: resourceType(volcengineVmpMod, "ContactGroup")},
			"volcengine_vmp_rule_file":                                    {Tok: resourceType(volcengineVmpMod, "RuleFile")},
			"volcengine_vmp_notify_template":                              {Tok: resourceType(volcengineVmpMod, "NotifyTemplate")},
			"volcengine_vmp_alerting_rule":                                {Tok: resourceType(volcengineVmpMod, "AlertingRule")},
			"volcengine_vmp_workspace":                                    {Tok: resourceType(volcengineVmpMod, "Workspace")},
			"volcengine_vmp_contact":                                      {Tok: resourceType(volcengineVmpMod, "Contact")},
			"volcengine_vmp_notify_policy":                                {Tok: resourceType(volcengineVmpMod, "NotifyPolicy")},
			"volcengine_escloud_ip_white_list":                            {Tok: resourceType(volcengineEscloudV2Mod, "EscloudIpWhiteList")},
			"volcengine_escloud_instance_v2":                              {Tok: resourceType(volcengineEscloudV2Mod, "EscloudInstanceV2")},
			"volcengine_tls_alarm":                                        {Tok: resourceType(volcengineTlsMod, "Alarm")},
			"volcengine_tls_project":                                      {Tok: resourceType(volcengineTlsMod, "Project")},
			"volcengine_tls_topic":                                        {Tok: resourceType(volcengineTlsMod, "Topic")},
			"volcengine_tls_rule_applier":                                 {Tok: resourceType(volcengineTlsMod, "RuleApplier")},
			"volcengine_tls_etl_task":                                     {Tok: resourceType(volcengineTlsMod, "EtlTask")},
			"volcengine_tls_rule":                                         {Tok: resourceType(volcengineTlsMod, "Rule")},
			"volcengine_tls_host_group":                                   {Tok: resourceType(volcengineTlsMod, "HostGroup")},
			"volcengine_tls_shipper":                                      {Tok: resourceType(volcengineTlsMod, "Shipper")},
			"volcengine_tls_index":                                        {Tok: resourceType(volcengineTlsMod, "Index")},
			"volcengine_tls_consumer_group":                               {Tok: resourceType(volcengineTlsMod, "ConsumerGroup")},
			"volcengine_tls_host":                                         {Tok: resourceType(volcengineTlsMod, "Host")},
			"volcengine_tls_schedule_sql_task":                            {Tok: resourceType(volcengineTlsMod, "ScheduleSqlTask")},
			"volcengine_tls_alarm_notify_group":                           {Tok: resourceType(volcengineTlsMod, "AlarmNotifyGroup")},
			"volcengine_tls_kafka_consumer":                               {Tok: resourceType(volcengineTlsMod, "KafkaConsumer")},
			"volcengine_tls_import_task":                                  {Tok: resourceType(volcengineTlsMod, "ImportTask")},
			"volcengine_bioos_cluster":                                    {Tok: resourceType(volcengineBioosMod, "Cluster")},
			"volcengine_bioos_workspace":                                  {Tok: resourceType(volcengineBioosMod, "Workspace")},
			"volcengine_bioos_cluster_bind":                               {Tok: resourceType(volcengineBioosMod, "ClusterBind")},
			"volcengine_rds_mysql_instance_readonly_node":                 {Tok: resourceType(volcengineRdsMysqlMod, "InstanceReadonlyNode")},
			"volcengine_rds_mysql_endpoint_public_address":                {Tok: resourceType(volcengineRdsMysqlMod, "EndpointPublicAddress")},
			"volcengine_rds_mysql_allowlist_associate":                    {Tok: resourceType(volcengineRdsMysqlMod, "AllowlistAssociate")},
			"volcengine_rds_mysql_backup_policy":                          {Tok: resourceType(volcengineRdsMysqlMod, "BackupPolicy")},
			"volcengine_rds_mysql_account":                                {Tok: resourceType(volcengineRdsMysqlMod, "Account")},
			"volcengine_rds_mysql_backup":                                 {Tok: resourceType(volcengineRdsMysqlMod, "Backup")},
			"volcengine_rds_mysql_instance":                               {Tok: resourceType(volcengineRdsMysqlMod, "Instance")},
			"volcengine_rds_mysql_parameter_template":                     {Tok: resourceType(volcengineRdsMysqlMod, "ParameterTemplate")},
			"volcengine_rds_mysql_allowlist":                              {Tok: resourceType(volcengineRdsMysqlMod, "Allowlist")},
			"volcengine_rds_mysql_endpoint":                               {Tok: resourceType(volcengineRdsMysqlMod, "Endpoint")},
			"volcengine_rds_mysql_database":                               {Tok: resourceType(volcengineRdsMysqlMod, "Database")},
			"volcengine_rds_mssql_instance":                               {Tok: resourceType(volcengineRdsMssqlMod, "Instance")},
			"volcengine_rds_mssql_backup":                                 {Tok: resourceType(volcengineRdsMssqlMod, "Backup")},
			"volcengine_vpn_connection":                                   {Tok: resourceType(volcengineVpnMod, "Connection")},
			"volcengine_vpn_gateway_route":                                {Tok: resourceType(volcengineVpnMod, "GatewayRoute")},
			"volcengine_ssl_vpn_client_cert":                              {Tok: resourceType(volcengineVpnMod, "SslVpnClientCert")},
			"volcengine_vpn_gateway":                                      {Tok: resourceType(volcengineVpnMod, "Gateway")},
			"volcengine_ssl_vpn_server":                                   {Tok: resourceType(volcengineVpnMod, "SslVpnServer")},
			"volcengine_customer_gateway":                                 {Tok: resourceType(volcengineVpnMod, "CustomerGateway")},
			"volcengine_cen_route_entry":                                  {Tok: resourceType(volcengineCenMod, "RouteEntry")},
			"volcengine_cen_bandwidth_package":                            {Tok: resourceType(volcengineCenMod, "BandwidthPackage")},
			"volcengine_cen_attach_instance":                              {Tok: resourceType(volcengineCenMod, "AttachInstance")},
			"volcengine_cen_bandwidth_package_associate":                  {Tok: resourceType(volcengineCenMod, "BandwidthPackageAssociate")},
			"volcengine_cen_service_route_entry":                          {Tok: resourceType(volcengineCenMod, "ServiceRouteEntry")},
			"volcengine_cen_grant_instance":                               {Tok: resourceType(volcengineCenMod, "GrantInstance")},
			"volcengine_cen":                                              {Tok: resourceType(volcengineCenMod, "Cen")},
			"volcengine_cen_inter_region_bandwidth":                       {Tok: resourceType(volcengineCenMod, "InterRegionBandwidth")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"volcengine_traffic_mirror_filter_rules":                       {Tok: dataSourceType(volcengineVpcMod, "TrafficMirrorFilterRules")},
			"volcengine_vpc_ipv6_gateways":                                 {Tok: dataSourceType(volcengineVpcMod, "Ipv6Gateways")},
			"volcengine_vpc_ipv6_addresses":                                {Tok: dataSourceType(volcengineVpcMod, "Ipv6Addresses")},
			"volcengine_flow_logs":                                         {Tok: dataSourceType(volcengineVpcMod, "FlowLogs")},
			"volcengine_route_entries":                                     {Tok: dataSourceType(volcengineVpcMod, "RouteEntries")},
			"volcengine_network_acls":                                      {Tok: dataSourceType(volcengineVpcMod, "NetworkAcls")},
			"volcengine_security_group_rules":                              {Tok: dataSourceType(volcengineVpcMod, "SecurityGroupRules")},
			"volcengine_subnets":                                           {Tok: dataSourceType(volcengineVpcMod, "Subnets")},
			"volcengine_traffic_mirror_filters":                            {Tok: dataSourceType(volcengineVpcMod, "TrafficMirrorFilters")},
			"volcengine_vpcs":                                              {Tok: dataSourceType(volcengineVpcMod, "Vpcs")},
			"volcengine_security_groups":                                   {Tok: dataSourceType(volcengineVpcMod, "SecurityGroups")},
			"volcengine_traffic_mirror_targets":                            {Tok: dataSourceType(volcengineVpcMod, "TrafficMirrorTargets")},
			"volcengine_vpc_prefix_lists":                                  {Tok: dataSourceType(volcengineVpcMod, "PrefixLists")},
			"volcengine_vpc_ipv6_address_bandwidths":                       {Tok: dataSourceType(volcengineVpcMod, "Ipv6AddressBandwidths")},
			"volcengine_route_tables":                                      {Tok: dataSourceType(volcengineVpcMod, "RouteTables")},
			"volcengine_traffic_mirror_sessions":                           {Tok: dataSourceType(volcengineVpcMod, "TrafficMirrorSessions")},
			"volcengine_network_interfaces":                                {Tok: dataSourceType(volcengineVpcMod, "NetworkInterfaces")},
			"volcengine_ha_vips":                                           {Tok: dataSourceType(volcengineVpcMod, "HaVips")},
			"volcengine_nas_mount_points":                                  {Tok: dataSourceType(volcengineNasMod, "MountPoints")},
			"volcengine_nas_permission_groups":                             {Tok: dataSourceType(volcengineNasMod, "PermissionGroups")},
			"volcengine_nas_zones":                                         {Tok: dataSourceType(volcengineNasMod, "Zones")},
			"volcengine_nas_regions":                                       {Tok: dataSourceType(volcengineNasMod, "Regions")},
			"volcengine_nas_auto_snapshot_policies":                        {Tok: dataSourceType(volcengineNasMod, "AutoSnapshotPolicies")},
			"volcengine_nas_file_systems":                                  {Tok: dataSourceType(volcengineNasMod, "FileSystems")},
			"volcengine_nas_snapshots":                                     {Tok: dataSourceType(volcengineNasMod, "Snapshots")},
			"volcengine_rds_ip_lists":                                      {Tok: dataSourceType(volcengineRdsMod, "IpLists")},
			"volcengine_rds_databases":                                     {Tok: dataSourceType(volcengineRdsMod, "Databases")},
			"volcengine_rds_accounts":                                      {Tok: dataSourceType(volcengineRdsMod, "Accounts")},
			"volcengine_rds_parameter_templates":                           {Tok: dataSourceType(volcengineRdsMod, "ParameterTemplates")},
			"volcengine_rds_instances":                                     {Tok: dataSourceType(volcengineRdsMod, "Instances")},
			"volcengine_listeners":                                         {Tok: dataSourceType(volcengineClbMod, "Listeners")},
			"volcengine_acls":                                              {Tok: dataSourceType(volcengineClbMod, "Acls")},
			"volcengine_clb_rules":                                         {Tok: dataSourceType(volcengineClbMod, "Rules")},
			"volcengine_clbs":                                              {Tok: dataSourceType(volcengineClbMod, "Clbs")},
			"volcengine_server_groups":                                     {Tok: dataSourceType(volcengineClbMod, "ServerGroups")},
			"volcengine_server_group_servers":                              {Tok: dataSourceType(volcengineClbMod, "ServerGroupServers")},
			"volcengine_clb_zones":                                         {Tok: dataSourceType(volcengineClbMod, "Zones")},
			"volcengine_certificates":                                      {Tok: dataSourceType(volcengineClbMod, "Certificates")},
			"volcengine_financial_relations":                               {Tok: dataSourceType(volcengineFinancialRelationMod, "FinancialRelations")},
			"volcengine_snat_entries":                                      {Tok: dataSourceType(volcengineNatMod, "SnatEntries")},
			"volcengine_nat_gateways":                                      {Tok: dataSourceType(volcengineNatMod, "Gateways")},
			"volcengine_dnat_entries":                                      {Tok: dataSourceType(volcengineNatMod, "DnatEntries")},
			"volcengine_direct_connect_gateways":                           {Tok: dataSourceType(volcengineDirectConnectMod, "Gateways")},
			"volcengine_direct_connect_gateway_routes":                     {Tok: dataSourceType(volcengineDirectConnectMod, "GatewayRoutes")},
			"volcengine_direct_connect_connections":                        {Tok: dataSourceType(volcengineDirectConnectMod, "Connections")},
			"volcengine_direct_connect_virtual_interfaces":                 {Tok: dataSourceType(volcengineDirectConnectMod, "VirtualInterfaces")},
			"volcengine_direct_connect_bgp_peers":                          {Tok: dataSourceType(volcengineDirectConnectMod, "BgpPeers")},
			"volcengine_vke_node_pools":                                    {Tok: dataSourceType(volcengineVkeMod, "NodePools")},
			"volcengine_vke_kubeconfigs":                                   {Tok: dataSourceType(volcengineVkeMod, "Kubeconfigs")},
			"volcengine_vke_nodes":                                         {Tok: dataSourceType(volcengineVkeMod, "Nodes")},
			"volcengine_vke_clusters":                                      {Tok: dataSourceType(volcengineVkeMod, "Clusters")},
			"volcengine_vke_support_addons":                                {Tok: dataSourceType(volcengineVkeMod, "SupportAddons")},
			"volcengine_vke_addons":                                        {Tok: dataSourceType(volcengineVkeMod, "Addons")},
			"volcengine_vke_permissions":                                   {Tok: dataSourceType(volcengineVkeMod, "Permissions")},
			"volcengine_vke_support_resource_types":                        {Tok: dataSourceType(volcengineVkeMod, "SupportResourceTypes")},
			"volcengine_kms_keyrings":                                      {Tok: dataSourceType(volcengineKmsMod, "Keyrings")},
			"volcengine_kms_keys":                                          {Tok: dataSourceType(volcengineKmsMod, "Keys")},
			"volcengine_kms_secrets":                                       {Tok: dataSourceType(volcengineKmsMod, "Secrets")},
			"volcengine_cloud_monitor_event_rules":                         {Tok: dataSourceType(volcengineCloudMonitorMod, "EventRules")},
			"volcengine_cloud_monitor_contact_groups":                      {Tok: dataSourceType(volcengineCloudMonitorMod, "ContactGroups")},
			"volcengine_cloud_monitor_contacts":                            {Tok: dataSourceType(volcengineCloudMonitorMod, "Contacts")},
			"volcengine_cloud_monitor_rules":                               {Tok: dataSourceType(volcengineCloudMonitorMod, "Rules")},
			"volcengine_iam_user_group_policy_attachments":                 {Tok: dataSourceType(volcengineIamMod, "UserGroupPolicyAttachments")},
			"volcengine_iam_roles":                                         {Tok: dataSourceType(volcengineIamMod, "Roles")},
			"volcengine_iam_access_keys":                                   {Tok: dataSourceType(volcengineIamMod, "AccessKeys")},
			"volcengine_iam_saml_providers":                                {Tok: dataSourceType(volcengineIamMod, "SamlProviders")},
			"volcengine_iam_policies":                                      {Tok: dataSourceType(volcengineIamMod, "Policies")},
			"volcengine_iam_user_groups":                                   {Tok: dataSourceType(volcengineIamMod, "UserGroups")},
			"volcengine_iam_users":                                         {Tok: dataSourceType(volcengineIamMod, "Users")},
			"volcengine_privatelink_vpc_endpoint_services":                 {Tok: dataSourceType(volcenginePrivatelinkMod, "VpcEndpointServices")},
			"volcengine_privatelink_vpc_endpoint_zones":                    {Tok: dataSourceType(volcenginePrivatelinkMod, "VpcEndpointZones")},
			"volcengine_privatelink_vpc_endpoints":                         {Tok: dataSourceType(volcenginePrivatelinkMod, "VpcEndpoints")},
			"volcengine_privatelink_vpc_endpoint_service_permissions":      {Tok: dataSourceType(volcenginePrivatelinkMod, "VpcEndpointServicePermissions")},
			"volcengine_privatelink_vpc_endpoint_connections":              {Tok: dataSourceType(volcenginePrivatelinkMod, "VpcEndpointConnections")},
			"volcengine_rocketmq_topics":                                   {Tok: dataSourceType(volcengineRocketmqMod, "Topics")},
			"volcengine_rocketmq_instances":                                {Tok: dataSourceType(volcengineRocketmqMod, "Instances")},
			"volcengine_rocketmq_groups":                                   {Tok: dataSourceType(volcengineRocketmqMod, "Groups")},
			"volcengine_rocketmq_access_keys":                              {Tok: dataSourceType(volcengineRocketmqMod, "AccessKeys")},
			"volcengine_rocketmq_allow_lists":                              {Tok: dataSourceType(volcengineRocketmqMod, "AllowLists")},
			"volcengine_transit_router_route_entries":                      {Tok: dataSourceType(volcengineTransitRouterMod, "RouteEntries")},
			"volcengine_transit_router_route_tables":                       {Tok: dataSourceType(volcengineTransitRouterMod, "RouteTables")},
			"volcengine_transit_routers":                                   {Tok: dataSourceType(volcengineTransitRouterMod, "TransitRouters")},
			"volcengine_transit_router_bandwidth_packages":                 {Tok: dataSourceType(volcengineTransitRouterMod, "BandwidthPackages")},
			"volcengine_transit_router_route_table_associations":           {Tok: dataSourceType(volcengineTransitRouterMod, "RouteTableAssociations")},
			"volcengine_transit_router_route_table_propagations":           {Tok: dataSourceType(volcengineTransitRouterMod, "RouteTablePropagations")},
			"volcengine_transit_router_grant_rules":                        {Tok: dataSourceType(volcengineTransitRouterMod, "GrantRules")},
			"volcengine_transit_router_vpc_attachments":                    {Tok: dataSourceType(volcengineTransitRouterMod, "VpcAttachments")},
			"volcengine_transit_router_direct_connect_gateway_attachments": {Tok: dataSourceType(volcengineTransitRouterMod, "DirectConnectGatewayAttachments")},
			"volcengine_transit_router_peer_attachments":                   {Tok: dataSourceType(volcengineTransitRouterMod, "PeerAttachments")},
			"volcengine_transit_router_vpn_attachments":                    {Tok: dataSourceType(volcengineTransitRouterMod, "VpnAttachments")},
			"volcengine_albs":                                              {Tok: dataSourceType(volcengineAlbMod, "Albs")},
			"volcengine_alb_listener_domain_extensions":                    {Tok: dataSourceType(volcengineAlbMod, "ListenerDomainExtensions")},
			"volcengine_alb_health_check_templates":                        {Tok: dataSourceType(volcengineAlbMod, "HealthCheckTemplates")},
			"volcengine_alb_certificates":                                  {Tok: dataSourceType(volcengineAlbMod, "Certificates")},
			"volcengine_alb_ca_certificates":                               {Tok: dataSourceType(volcengineAlbMod, "CaCertificates")},
			"volcengine_alb_rules":                                         {Tok: dataSourceType(volcengineAlbMod, "Rules")},
			"volcengine_alb_acls":                                          {Tok: dataSourceType(volcengineAlbMod, "Acls")},
			"volcengine_alb_listeners":                                     {Tok: dataSourceType(volcengineAlbMod, "Listeners")},
			"volcengine_alb_server_group_servers":                          {Tok: dataSourceType(volcengineAlbMod, "ServerGroupServers")},
			"volcengine_alb_customized_cfgs":                               {Tok: dataSourceType(volcengineAlbMod, "CustomizedCfgs")},
			"volcengine_alb_zones":                                         {Tok: dataSourceType(volcengineAlbMod, "Zones")},
			"volcengine_alb_server_groups":                                 {Tok: dataSourceType(volcengineAlbMod, "ServerGroups")},
			"volcengine_dns_records":                                       {Tok: dataSourceType(volcengineDnsMod, "Records")},
			"volcengine_dns_zones":                                         {Tok: dataSourceType(volcengineDnsMod, "Zones")},
			"volcengine_dns_backups":                                       {Tok: dataSourceType(volcengineDnsMod, "Backups")},
			"volcengine_dns_record_sets":                                   {Tok: dataSourceType(volcengineDnsMod, "RecordSets")},
			"volcengine_mongodb_zones":                                     {Tok: dataSourceType(volcengineMongodbMod, "Zones")},
			"volcengine_mongodb_instances":                                 {Tok: dataSourceType(volcengineMongodbMod, "Instances")},
			"volcengine_mongodb_ssl_states":                                {Tok: dataSourceType(volcengineMongodbMod, "SslStates")},
			"volcengine_mongodb_endpoints":                                 {Tok: dataSourceType(volcengineMongodbMod, "Endpoints")},
			"volcengine_mongodb_instance_parameters":                       {Tok: dataSourceType(volcengineMongodbMod, "InstanceParameters")},
			"volcengine_mongodb_instance_parameter_logs":                   {Tok: dataSourceType(volcengineMongodbMod, "InstanceParameterLogs")},
			"volcengine_mongodb_specs":                                     {Tok: dataSourceType(volcengineMongodbMod, "Specs")},
			"volcengine_mongodb_regions":                                   {Tok: dataSourceType(volcengineMongodbMod, "Regions")},
			"volcengine_mongodb_allow_lists":                               {Tok: dataSourceType(volcengineMongodbMod, "MongoAllowLists")},
			"volcengine_mongodb_accounts":                                  {Tok: dataSourceType(volcengineMongodbMod, "Accounts")},
			"volcengine_cdn_certificates":                                  {Tok: dataSourceType(volcengineCdnMod, "Certificates")},
			"volcengine_cdn_shared_configs":                                {Tok: dataSourceType(volcengineCdnMod, "SharedConfigs")},
			"volcengine_cdn_domains":                                       {Tok: dataSourceType(volcengineCdnMod, "Domains")},
			"volcengine_cdn_configs":                                       {Tok: dataSourceType(volcengineCdnMod, "Configs")},
			"volcengine_cfw_dns_control_policies":                          {Tok: dataSourceType(volcengineCloudFirewallMod, "CfwDnsControlPolicies")},
			"volcengine_cfw_address_books":                                 {Tok: dataSourceType(volcengineCloudFirewallMod, "CfwAddressBooks")},
			"volcengine_cfw_vpc_firewall_acl_rules":                        {Tok: dataSourceType(volcengineCloudFirewallMod, "CfwVpcFirewallAclRules")},
			"volcengine_cfw_nat_firewall_control_policies":                 {Tok: dataSourceType(volcengineCloudFirewallMod, "CfwNatFirewallControlPolicies")},
			"volcengine_cfw_control_policies":                              {Tok: dataSourceType(volcengineCloudFirewallMod, "CfwControlPolicies")},
			"volcengine_ecs_invocations":                                   {Tok: dataSourceType(volcengineEcsMod, "Invocations")},
			"volcengine_image_share_permissions":                           {Tok: dataSourceType(volcengineEcsMod, "ImageSharePermissions")},
			"volcengine_ecs_instance_types":                                {Tok: dataSourceType(volcengineEcsMod, "InstanceTypes")},
			"volcengine_ecs_hpc_clusters":                                  {Tok: dataSourceType(volcengineEcsMod, "HpcClusters")},
			"volcengine_ecs_launch_templates":                              {Tok: dataSourceType(volcengineEcsMod, "LaunchTemplates")},
			"volcengine_ecs_available_resources":                           {Tok: dataSourceType(volcengineEcsMod, "AvailableResources")},
			"volcengine_ecs_invocation_results":                            {Tok: dataSourceType(volcengineEcsMod, "InvocationResults")},
			"volcengine_regions":                                           {Tok: dataSourceType(volcengineEcsMod, "Regions")},
			"volcengine_ecs_commands":                                      {Tok: dataSourceType(volcengineEcsMod, "Commands")},
			"volcengine_images":                                            {Tok: dataSourceType(volcengineEcsMod, "Images")},
			"volcengine_ecs_key_pairs":                                     {Tok: dataSourceType(volcengineEcsMod, "KeyPairs")},
			"volcengine_ecs_instances":                                     {Tok: dataSourceType(volcengineEcsMod, "Instances")},
			"volcengine_ecs_deployment_sets":                               {Tok: dataSourceType(volcengineEcsMod, "DeploymentSets")},
			"volcengine_zones":                                             {Tok: dataSourceType(volcengineEcsMod, "Zones")},
			"volcengine_rabbitmq_instance_plugins":                         {Tok: dataSourceType(volcengineRabbitmqMod, "InstancePlugins")},
			"volcengine_rabbitmq_zones":                                    {Tok: dataSourceType(volcengineRabbitmqMod, "Zones")},
			"volcengine_rabbitmq_regions":                                  {Tok: dataSourceType(volcengineRabbitmqMod, "Regions")},
			"volcengine_rabbitmq_instances":                                {Tok: dataSourceType(volcengineRabbitmqMod, "Instances")},
			"volcengine_organization_units":                                {Tok: dataSourceType(volcengineOrganizationMod, "Units")},
			"volcengine_organization_accounts":                             {Tok: dataSourceType(volcengineOrganizationMod, "Accounts")},
			"volcengine_organizations":                                     {Tok: dataSourceType(volcengineOrganizationMod, "Organizations")},
			"volcengine_organization_service_control_policies":             {Tok: dataSourceType(volcengineOrganizationMod, "ServiceControlPolicies")},
			"volcengine_tos_bucket_inventories":                            {Tok: dataSourceType(volcengineTosMod, "BucketInventories")},
			"volcengine_tos_buckets":                                       {Tok: dataSourceType(volcengineTosMod, "Buckets")},
			"volcengine_tos_objects":                                       {Tok: dataSourceType(volcengineTosMod, "BucketObjects")},
			"volcengine_apig_upstreams":                                    {Tok: dataSourceType(volcengineApigMod, "Upstreams")},
			"volcengine_apig_upstream_sources":                             {Tok: dataSourceType(volcengineApigMod, "UpstreamSources")},
			"volcengine_apig_upstream_versions":                            {Tok: dataSourceType(volcengineApigMod, "UpstreamVersions")},
			"volcengine_apig_gateway_services":                             {Tok: dataSourceType(volcengineApigMod, "GatewayServices")},
			"volcengine_apig_gateways":                                     {Tok: dataSourceType(volcengineApigMod, "Gateways")},
			"volcengine_apig_routes":                                       {Tok: dataSourceType(volcengineApigMod, "Routes")},
			"volcengine_apig_custom_domains":                               {Tok: dataSourceType(volcengineApigMod, "CustomDomains")},
			"volcengine_rds_instances_v2":                                  {Tok: dataSourceType(volcengineRdsV2Mod, "RdsInstancesV2")},
			"volcengine_kafka_sasl_users":                                  {Tok: dataSourceType(volcengineKafkaMod, "SaslUsers")},
			"volcengine_kafka_allow_lists":                                 {Tok: dataSourceType(volcengineKafkaMod, "AllowLists")},
			"volcengine_kafka_groups":                                      {Tok: dataSourceType(volcengineKafkaMod, "Groups")},
			"volcengine_kafka_consumed_partitions":                         {Tok: dataSourceType(volcengineKafkaMod, "ConsumedPartitions")},
			"volcengine_kafka_zones":                                       {Tok: dataSourceType(volcengineKafkaMod, "Zones")},
			"volcengine_kafka_topics":                                      {Tok: dataSourceType(volcengineKafkaMod, "Topics")},
			"volcengine_kafka_instances":                                   {Tok: dataSourceType(volcengineKafkaMod, "Instances")},
			"volcengine_kafka_consumed_topics":                             {Tok: dataSourceType(volcengineKafkaMod, "ConsumedTopics")},
			"volcengine_kafka_regions":                                     {Tok: dataSourceType(volcengineKafkaMod, "Regions")},
			"volcengine_kafka_topic_partitions":                            {Tok: dataSourceType(volcengineKafkaMod, "TopicPartitions")},
			"volcengine_private_zone_records":                              {Tok: dataSourceType(volcenginePrivateZoneMod, "Records")},
			"volcengine_private_zone_resolver_rules":                       {Tok: dataSourceType(volcenginePrivateZoneMod, "ResolverRules")},
			"volcengine_private_zone_resolver_endpoints":                   {Tok: dataSourceType(volcenginePrivateZoneMod, "ResolverEndpoints")},
			"volcengine_private_zones":                                     {Tok: dataSourceType(volcenginePrivateZoneMod, "PrivateZones")},
			"volcengine_private_zone_record_sets":                          {Tok: dataSourceType(volcenginePrivateZoneMod, "RecordSets")},
			"volcengine_vefaas_releases":                                   {Tok: dataSourceType(volcengineVefaasMod, "Releases")},
			"volcengine_vefaas_kafka_triggers":                             {Tok: dataSourceType(volcengineVefaasMod, "KafkaTriggers")},
			"volcengine_vefaas_functions":                                  {Tok: dataSourceType(volcengineVefaasMod, "Functions")},
			"volcengine_vefaas_timers":                                     {Tok: dataSourceType(volcengineVefaasMod, "Timers")},
			"volcengine_waf_custom_bots":                                   {Tok: dataSourceType(volcengineWafMod, "CustomBots")},
			"volcengine_waf_acl_rules":                                     {Tok: dataSourceType(volcengineWafMod, "AclRules")},
			"volcengine_waf_system_bots":                                   {Tok: dataSourceType(volcengineWafMod, "SystemBots")},
			"volcengine_waf_service_certificates":                          {Tok: dataSourceType(volcengineWafMod, "ServiceCertificates")},
			"volcengine_waf_custom_pages":                                  {Tok: dataSourceType(volcengineWafMod, "CustomPages")},
			"volcengine_waf_domains":                                       {Tok: dataSourceType(volcengineWafMod, "Domains")},
			"volcengine_waf_host_groups":                                   {Tok: dataSourceType(volcengineWafMod, "HostGroups")},
			"volcengine_waf_prohibitions":                                  {Tok: dataSourceType(volcengineWafMod, "Prohibitions")},
			"volcengine_waf_bot_analyse_protect_rules":                     {Tok: dataSourceType(volcengineWafMod, "BotAnalyseProtectRules")},
			"volcengine_waf_cc_rules":                                      {Tok: dataSourceType(volcengineWafMod, "CcRules")},
			"volcengine_waf_ip_groups":                                     {Tok: dataSourceType(volcengineWafMod, "IpGroups")},
			"volcengine_vepfs_filesets":                                    {Tok: dataSourceType(volcengineVepfsMod, "Filesets")},
			"volcengine_vepfs_file_systems":                                {Tok: dataSourceType(volcengineVepfsMod, "FileSystems")},
			"volcengine_vepfs_mount_services":                              {Tok: dataSourceType(volcengineVepfsMod, "MountServices")},
			"volcengine_scaling_configurations":                            {Tok: dataSourceType(volcengineAutoscalingMod, "ScalingConfigurations")},
			"volcengine_scaling_lifecycle_hooks":                           {Tok: dataSourceType(volcengineAutoscalingMod, "ScalingLifecycleHooks")},
			"volcengine_scaling_groups":                                    {Tok: dataSourceType(volcengineAutoscalingMod, "ScalingGroups")},
			"volcengine_scaling_policies":                                  {Tok: dataSourceType(volcengineAutoscalingMod, "ScalingPolicies")},
			"volcengine_scaling_instances":                                 {Tok: dataSourceType(volcengineAutoscalingMod, "ScalingInstances")},
			"volcengine_scaling_activities":                                {Tok: dataSourceType(volcengineAutoscalingMod, "ScalingActivities")},
			"volcengine_bandwidth_packages":                                {Tok: dataSourceType(volcengineBandwidthPackageMod, "BandwidthPackages")},
			"volcengine_veecp_edge_nodes":                                  {Tok: dataSourceType(volcengineVeecpMod, "EdgeNodes")},
			"volcengine_veecp_node_pools":                                  {Tok: dataSourceType(volcengineVeecpMod, "NodePools")},
			"volcengine_veecp_edge_node_pools":                             {Tok: dataSourceType(volcengineVeecpMod, "EdgeNodePools")},
			"volcengine_veecp_support_addons":                              {Tok: dataSourceType(volcengineVeecpMod, "SupportAddons")},
			"volcengine_veecp_support_resource_types":                      {Tok: dataSourceType(volcengineVeecpMod, "SupportResourceTypes")},
			"volcengine_veecp_batch_edge_machines":                         {Tok: dataSourceType(volcengineVeecpMod, "BatchEdgeMachines")},
			"volcengine_veecp_addons":                                      {Tok: dataSourceType(volcengineVeecpMod, "Addons")},
			"volcengine_veecp_clusters":                                    {Tok: dataSourceType(volcengineVeecpMod, "Clusters")},
			"volcengine_veenedge_cloud_servers":                            {Tok: dataSourceType(volcengineVeenedgeMod, "CloudServers")},
			"volcengine_veenedge_instance_types":                           {Tok: dataSourceType(volcengineVeenedgeMod, "InstanceTypes")},
			"volcengine_veenedge_available_resources":                      {Tok: dataSourceType(volcengineVeenedgeMod, "AvailableResources")},
			"volcengine_veenedge_instances":                                {Tok: dataSourceType(volcengineVeenedgeMod, "Instances")},
			"volcengine_veenedge_vpcs":                                     {Tok: dataSourceType(volcengineVeenedgeMod, "Vpcs")},
			"volcengine_eip_addresses":                                     {Tok: dataSourceType(volcengineEipMod, "Addresses")},
			"volcengine_vedb_mysql_endpoints":                              {Tok: dataSourceType(volcengineVedbMysqlMod, "Endpoints")},
			"volcengine_vedb_mysql_backups":                                {Tok: dataSourceType(volcengineVedbMysqlMod, "Backups")},
			"volcengine_vedb_mysql_allowlists":                             {Tok: dataSourceType(volcengineVedbMysqlMod, "Allowlists")},
			"volcengine_vedb_mysql_databases":                              {Tok: dataSourceType(volcengineVedbMysqlMod, "Databases")},
			"volcengine_vedb_mysql_instances":                              {Tok: dataSourceType(volcengineVedbMysqlMod, "Instances")},
			"volcengine_vedb_mysql_accounts":                               {Tok: dataSourceType(volcengineVedbMysqlMod, "Accounts")},
			"volcengine_cloud_identity_permission_sets":                    {Tok: dataSourceType(volcengineCloudIdentityMod, "PermissionSets")},
			"volcengine_cloud_identity_permission_set_provisionings":       {Tok: dataSourceType(volcengineCloudIdentityMod, "PermissionSetProvisionings")},
			"volcengine_cloud_identity_user_provisionings":                 {Tok: dataSourceType(volcengineCloudIdentityMod, "UserProvisionings")},
			"volcengine_cloud_identity_users":                              {Tok: dataSourceType(volcengineCloudIdentityMod, "Users")},
			"volcengine_cloud_identity_groups":                             {Tok: dataSourceType(volcengineCloudIdentityMod, "Groups")},
			"volcengine_cloud_identity_permission_set_assignments":         {Tok: dataSourceType(volcengineCloudIdentityMod, "PermissionSetAssignments")},
			"volcengine_cloudfs_quotas":                                    {Tok: dataSourceType(volcengineCloudfsMod, "Quotas")},
			"volcengine_cloudfs_ns_quotas":                                 {Tok: dataSourceType(volcengineCloudfsMod, "NsQuotas")},
			"volcengine_cloudfs_namespaces":                                {Tok: dataSourceType(volcengineCloudfsMod, "Namespaces")},
			"volcengine_cloudfs_file_systems":                              {Tok: dataSourceType(volcengineCloudfsMod, "FileSystems")},
			"volcengine_cloudfs_accesses":                                  {Tok: dataSourceType(volcengineCloudfsMod, "Accesses")},
			"volcengine_ebs_auto_snapshot_policies":                        {Tok: dataSourceType(volcengineEbsMod, "AutoSnapshotPolicies")},
			"volcengine_ebs_max_extra_performances":                        {Tok: dataSourceType(volcengineEbsMod, "MaxExtraPerformances")},
			"volcengine_ebs_snapshot_groups":                               {Tok: dataSourceType(volcengineEbsMod, "SnapshotGroups")},
			"volcengine_volumes":                                           {Tok: dataSourceType(volcengineEbsMod, "Volumes")},
			"volcengine_ebs_snapshots":                                     {Tok: dataSourceType(volcengineEbsMod, "Snapshots")},
			"volcengine_redis_pitr_time_windows":                           {Tok: dataSourceType(volcengineRedisMod, "PitrTimeWindows")},
			"volcengine_redis_planned_events":                              {Tok: dataSourceType(volcengineRedisMod, "PlannedEvents")},
			"volcengine_redis_accounts":                                    {Tok: dataSourceType(volcengineRedisMod, "Accounts")},
			"volcengine_redis_backups":                                     {Tok: dataSourceType(volcengineRedisMod, "Backups")},
			"volcengine_redis_zones":                                       {Tok: dataSourceType(volcengineRedisMod, "Zones")},
			"volcengine_redis_instance_specs":                              {Tok: dataSourceType(volcengineRedisMod, "InstanceSpecs")},
			"volcengine_redis_parameter_groups":                            {Tok: dataSourceType(volcengineRedisMod, "ParameterGroups")},
			"volcengine_redis_big_keys":                                    {Tok: dataSourceType(volcengineRedisMod, "BigKeys")},
			"volcengine_redis_regions":                                     {Tok: dataSourceType(volcengineRedisMod, "Regions")},
			"volcengine_redis_allow_lists":                                 {Tok: dataSourceType(volcengineRedisMod, "AllowLists")},
			"volcengine_redis_instances":                                   {Tok: dataSourceType(volcengineRedisMod, "Instances")},
			"volcengine_rds_postgresql_schemas":                            {Tok: dataSourceType(volcengineRdsPostgresqlMod, "Schemas")},
			"volcengine_rds_postgresql_databases":                          {Tok: dataSourceType(volcengineRdsPostgresqlMod, "Databases")},
			"volcengine_rds_postgresql_accounts":                           {Tok: dataSourceType(volcengineRdsPostgresqlMod, "Accounts")},
			"volcengine_rds_postgresql_allowlists":                         {Tok: dataSourceType(volcengineRdsPostgresqlMod, "Allowlists")},
			"volcengine_rds_postgresql_instances":                          {Tok: dataSourceType(volcengineRdsPostgresqlMod, "Instances")},
			"volcengine_escloud_instances":                                 {Tok: dataSourceType(volcengineEscloudMod, "Instances")},
			"volcengine_escloud_regions":                                   {Tok: dataSourceType(volcengineEscloudMod, "Regions")},
			"volcengine_escloud_zones":                                     {Tok: dataSourceType(volcengineEscloudMod, "Zones")},
			"volcengine_cr_endpoints":                                      {Tok: dataSourceType(volcengineCrMod, "Endpoints")},
			"volcengine_cr_tags":                                           {Tok: dataSourceType(volcengineCrMod, "Tags")},
			"volcengine_cr_registries":                                     {Tok: dataSourceType(volcengineCrMod, "Registries")},
			"volcengine_cr_repositories":                                   {Tok: dataSourceType(volcengineCrMod, "Repositories")},
			"volcengine_cr_authorization_tokens":                           {Tok: dataSourceType(volcengineCrMod, "AuthorizationTokens")},
			"volcengine_cr_namespaces":                                     {Tok: dataSourceType(volcengineCrMod, "Namespaces")},
			"volcengine_cr_vpc_endpoints":                                  {Tok: dataSourceType(volcengineCrMod, "VpcEndpoints")},
			"volcengine_vmp_notify_group_policies":                         {Tok: dataSourceType(volcengineVmpMod, "NotifyGroupPolicies")},
			"volcengine_vmp_contact_groups":                                {Tok: dataSourceType(volcengineVmpMod, "ContactGroups")},
			"volcengine_vmp_rule_files":                                    {Tok: dataSourceType(volcengineVmpMod, "RuleFiles")},
			"volcengine_vmp_notify_templates":                              {Tok: dataSourceType(volcengineVmpMod, "NotifyTemplates")},
			"volcengine_vmp_rules":                                         {Tok: dataSourceType(volcengineVmpMod, "Rules")},
			"volcengine_vmp_alerts":                                        {Tok: dataSourceType(volcengineVmpMod, "Alerts")},
			"volcengine_vmp_alerting_rules":                                {Tok: dataSourceType(volcengineVmpMod, "AlertingRules")},
			"volcengine_vmp_workspaces":                                    {Tok: dataSourceType(volcengineVmpMod, "Workspaces")},
			"volcengine_vmp_contacts":                                      {Tok: dataSourceType(volcengineVmpMod, "Contacts")},
			"volcengine_vmp_notify_policies":                               {Tok: dataSourceType(volcengineVmpMod, "NotifyPolicies")},
			"volcengine_vmp_instance_types":                                {Tok: dataSourceType(volcengineVmpMod, "InstanceTypes")},
			"volcengine_escloud_node_available_specs":                      {Tok: dataSourceType(volcengineEscloudV2Mod, "EscloudNodeAvailableSpecs")},
			"volcengine_escloud_zones_v2":                                  {Tok: dataSourceType(volcengineEscloudV2Mod, "EscloudZonesV2")},
			"volcengine_escloud_instances_v2":                              {Tok: dataSourceType(volcengineEscloudV2Mod, "EscloudInstancesV2")},
			"volcengine_tls_alarms":                                        {Tok: dataSourceType(volcengineTlsMod, "Alarms")},
			"volcengine_tls_projects":                                      {Tok: dataSourceType(volcengineTlsMod, "Projects")},
			"volcengine_tls_topics":                                        {Tok: dataSourceType(volcengineTlsMod, "Topics")},
			"volcengine_tls_rule_appliers":                                 {Tok: dataSourceType(volcengineTlsMod, "RuleAppliers")},
			"volcengine_tls_etl_tasks":                                     {Tok: dataSourceType(volcengineTlsMod, "EtlTasks")},
			"volcengine_tls_rules":                                         {Tok: dataSourceType(volcengineTlsMod, "Rules")},
			"volcengine_tls_host_groups":                                   {Tok: dataSourceType(volcengineTlsMod, "HostGroups")},
			"volcengine_tls_shippers":                                      {Tok: dataSourceType(volcengineTlsMod, "Shippers")},
			"volcengine_tls_indexes":                                       {Tok: dataSourceType(volcengineTlsMod, "Indexes")},
			"volcengine_tls_consumer_groups":                               {Tok: dataSourceType(volcengineTlsMod, "ConsumerGroups")},
			"volcengine_tls_hosts":                                         {Tok: dataSourceType(volcengineTlsMod, "Hosts")},
			"volcengine_tls_schedule_sql_tasks":                            {Tok: dataSourceType(volcengineTlsMod, "ScheduleSqlTasks")},
			"volcengine_tls_alarm_notify_groups":                           {Tok: dataSourceType(volcengineTlsMod, "AlarmNotifyGroups")},
			"volcengine_tls_kafka_consumers":                               {Tok: dataSourceType(volcengineTlsMod, "KafkaConsumers")},
			"volcengine_tls_import_tasks":                                  {Tok: dataSourceType(volcengineTlsMod, "ImportTasks")},
			"volcengine_tls_shards":                                        {Tok: dataSourceType(volcengineTlsMod, "Shards")},
			"volcengine_bioos_clusters":                                    {Tok: dataSourceType(volcengineBioosMod, "Clusters")},
			"volcengine_bioos_workspaces":                                  {Tok: dataSourceType(volcengineBioosMod, "Workspaces")},
			"volcengine_rds_mysql_tasks":                                   {Tok: dataSourceType(volcengineRdsMysqlMod, "Tasks")},
			"volcengine_rds_mysql_instance_specs":                          {Tok: dataSourceType(volcengineRdsMysqlMod, "InstanceSpecs")},
			"volcengine_rds_mysql_planned_events":                          {Tok: dataSourceType(volcengineRdsMysqlMod, "PlannedEvents")},
			"volcengine_rds_mysql_zones":                                   {Tok: dataSourceType(volcengineRdsMysqlMod, "Zones")},
			"volcengine_rds_mysql_accounts":                                {Tok: dataSourceType(volcengineRdsMysqlMod, "Accounts")},
			"volcengine_rds_mysql_backups":                                 {Tok: dataSourceType(volcengineRdsMysqlMod, "Backups")},
			"volcengine_rds_mysql_instances":                               {Tok: dataSourceType(volcengineRdsMysqlMod, "Instances")},
			"volcengine_rds_mysql_regions":                                 {Tok: dataSourceType(volcengineRdsMysqlMod, "Regions")},
			"volcengine_rds_mysql_parameter_templates":                     {Tok: dataSourceType(volcengineRdsMysqlMod, "ParameterTemplates")},
			"volcengine_rds_mysql_allowlists":                              {Tok: dataSourceType(volcengineRdsMysqlMod, "Allowlists")},
			"volcengine_rds_mysql_endpoints":                               {Tok: dataSourceType(volcengineRdsMysqlMod, "Endpoints")},
			"volcengine_rds_mysql_account_table_column_infos":              {Tok: dataSourceType(volcengineRdsMysqlMod, "AccountTableColumnInfos")},
			"volcengine_rds_mysql_databases":                               {Tok: dataSourceType(volcengineRdsMysqlMod, "Databases")},
			"volcengine_rds_mssql_instances":                               {Tok: dataSourceType(volcengineRdsMssqlMod, "Instances")},
			"volcengine_rds_mssql_backups":                                 {Tok: dataSourceType(volcengineRdsMssqlMod, "Backups")},
			"volcengine_rds_mssql_regions":                                 {Tok: dataSourceType(volcengineRdsMssqlMod, "Regions")},
			"volcengine_rds_mssql_zones":                                   {Tok: dataSourceType(volcengineRdsMssqlMod, "Zones")},
			"volcengine_vpn_connections":                                   {Tok: dataSourceType(volcengineVpnMod, "Connections")},
			"volcengine_vpn_gateway_routes":                                {Tok: dataSourceType(volcengineVpnMod, "GatewayRoutes")},
			"volcengine_ssl_vpn_client_certs":                              {Tok: dataSourceType(volcengineVpnMod, "SslVpnClientCerts")},
			"volcengine_vpn_gateways":                                      {Tok: dataSourceType(volcengineVpnMod, "Gateways")},
			"volcengine_ssl_vpn_servers":                                   {Tok: dataSourceType(volcengineVpnMod, "SslVpnServers")},
			"volcengine_customer_gateways":                                 {Tok: dataSourceType(volcengineVpnMod, "CustomerGateways")},
			"volcengine_cen_route_entries":                                 {Tok: dataSourceType(volcengineCenMod, "RouteEntries")},
			"volcengine_cen_bandwidth_packages":                            {Tok: dataSourceType(volcengineCenMod, "BandwidthPackages")},
			"volcengine_cen_attach_instances":                              {Tok: dataSourceType(volcengineCenMod, "AttachInstances")},
			"volcengine_cen_service_route_entries":                         {Tok: dataSourceType(volcengineCenMod, "ServiceRouteEntries")},
			"volcengine_cen_grant_instances":                               {Tok: dataSourceType(volcengineCenMod, "GrantInstances")},
			"volcengine_cens":                                              {Tok: dataSourceType(volcengineCenMod, "Cens")},
			"volcengine_cen_inter_region_bandwidths":                       {Tok: dataSourceType(volcengineCenMod, "InterRegionBandwidths")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			//Overlay: &tfbridge.OverlayInfo{},
			PackageName: "@volcengine/pulumi",
		},
		Python: &tfbridge.PythonInfo{
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
			PackageName: "pulumi_volcengine",
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/volcengine/pulumi-%[1]s/sdk/", volcenginePkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				volcenginePkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
			//RootNamespace: "Volcengine.Pulumi",
		},
	}

	prov.SetAutonaming(255, "-")

	// rename data source
	prov.RenameDataSource("volcengine_traffic_mirror_filter_rules", dataSourceType(volcengineVpcMod, "TrafficMirrorFilterRules"),
		dataSourceType(volcengineVpcMod, "getTrafficMirrorFilterRules"), volcengineVpcMod, volcengineVpcMod, nil)
	prov.RenameDataSource("volcengine_vpc_ipv6_gateways", dataSourceType(volcengineVpcMod, "Ipv6Gateways"),
		dataSourceType(volcengineVpcMod, "getIpv6Gateways"), volcengineVpcMod, volcengineVpcMod, nil)
	prov.RenameDataSource("volcengine_vpc_ipv6_addresses", dataSourceType(volcengineVpcMod, "Ipv6Addresses"),
		dataSourceType(volcengineVpcMod, "getIpv6Addresses"), volcengineVpcMod, volcengineVpcMod, nil)
	prov.RenameDataSource("volcengine_flow_logs", dataSourceType(volcengineVpcMod, "FlowLogs"),
		dataSourceType(volcengineVpcMod, "getFlowLogs"), volcengineVpcMod, volcengineVpcMod, nil)
	prov.RenameDataSource("volcengine_route_entries", dataSourceType(volcengineVpcMod, "RouteEntries"),
		dataSourceType(volcengineVpcMod, "getRouteEntries"), volcengineVpcMod, volcengineVpcMod, nil)
	prov.RenameDataSource("volcengine_network_acls", dataSourceType(volcengineVpcMod, "NetworkAcls"),
		dataSourceType(volcengineVpcMod, "getNetworkAcls"), volcengineVpcMod, volcengineVpcMod, nil)
	prov.RenameDataSource("volcengine_security_group_rules", dataSourceType(volcengineVpcMod, "SecurityGroupRules"),
		dataSourceType(volcengineVpcMod, "getSecurityGroupRules"), volcengineVpcMod, volcengineVpcMod, nil)
	prov.RenameDataSource("volcengine_subnets", dataSourceType(volcengineVpcMod, "Subnets"),
		dataSourceType(volcengineVpcMod, "getSubnets"), volcengineVpcMod, volcengineVpcMod, nil)
	prov.RenameDataSource("volcengine_traffic_mirror_filters", dataSourceType(volcengineVpcMod, "TrafficMirrorFilters"),
		dataSourceType(volcengineVpcMod, "getTrafficMirrorFilters"), volcengineVpcMod, volcengineVpcMod, nil)
	prov.RenameDataSource("volcengine_vpcs", dataSourceType(volcengineVpcMod, "Vpcs"),
		dataSourceType(volcengineVpcMod, "getVpcs"), volcengineVpcMod, volcengineVpcMod, nil)
	prov.RenameDataSource("volcengine_security_groups", dataSourceType(volcengineVpcMod, "SecurityGroups"),
		dataSourceType(volcengineVpcMod, "getSecurityGroups"), volcengineVpcMod, volcengineVpcMod, nil)
	prov.RenameDataSource("volcengine_traffic_mirror_targets", dataSourceType(volcengineVpcMod, "TrafficMirrorTargets"),
		dataSourceType(volcengineVpcMod, "getTrafficMirrorTargets"), volcengineVpcMod, volcengineVpcMod, nil)
	prov.RenameDataSource("volcengine_vpc_prefix_lists", dataSourceType(volcengineVpcMod, "PrefixLists"),
		dataSourceType(volcengineVpcMod, "getPrefixLists"), volcengineVpcMod, volcengineVpcMod, nil)
	prov.RenameDataSource("volcengine_vpc_ipv6_address_bandwidths", dataSourceType(volcengineVpcMod, "Ipv6AddressBandwidths"),
		dataSourceType(volcengineVpcMod, "getIpv6AddressBandwidths"), volcengineVpcMod, volcengineVpcMod, nil)
	prov.RenameDataSource("volcengine_route_tables", dataSourceType(volcengineVpcMod, "RouteTables"),
		dataSourceType(volcengineVpcMod, "getRouteTables"), volcengineVpcMod, volcengineVpcMod, nil)
	prov.RenameDataSource("volcengine_traffic_mirror_sessions", dataSourceType(volcengineVpcMod, "TrafficMirrorSessions"),
		dataSourceType(volcengineVpcMod, "getTrafficMirrorSessions"), volcengineVpcMod, volcengineVpcMod, nil)
	prov.RenameDataSource("volcengine_network_interfaces", dataSourceType(volcengineVpcMod, "NetworkInterfaces"),
		dataSourceType(volcengineVpcMod, "getNetworkInterfaces"), volcengineVpcMod, volcengineVpcMod, nil)
	prov.RenameDataSource("volcengine_ha_vips", dataSourceType(volcengineVpcMod, "HaVips"),
		dataSourceType(volcengineVpcMod, "getHaVips"), volcengineVpcMod, volcengineVpcMod, nil)
	prov.RenameDataSource("volcengine_nas_mount_points", dataSourceType(volcengineNasMod, "MountPoints"),
		dataSourceType(volcengineNasMod, "getMountPoints"), volcengineNasMod, volcengineNasMod, nil)
	prov.RenameDataSource("volcengine_nas_permission_groups", dataSourceType(volcengineNasMod, "PermissionGroups"),
		dataSourceType(volcengineNasMod, "getPermissionGroups"), volcengineNasMod, volcengineNasMod, nil)
	prov.RenameDataSource("volcengine_nas_zones", dataSourceType(volcengineNasMod, "Zones"),
		dataSourceType(volcengineNasMod, "getZones"), volcengineNasMod, volcengineNasMod, nil)
	prov.RenameDataSource("volcengine_nas_regions", dataSourceType(volcengineNasMod, "Regions"),
		dataSourceType(volcengineNasMod, "getRegions"), volcengineNasMod, volcengineNasMod, nil)
	prov.RenameDataSource("volcengine_nas_auto_snapshot_policies", dataSourceType(volcengineNasMod, "AutoSnapshotPolicies"),
		dataSourceType(volcengineNasMod, "getAutoSnapshotPolicies"), volcengineNasMod, volcengineNasMod, nil)
	prov.RenameDataSource("volcengine_nas_file_systems", dataSourceType(volcengineNasMod, "FileSystems"),
		dataSourceType(volcengineNasMod, "getFileSystems"), volcengineNasMod, volcengineNasMod, nil)
	prov.RenameDataSource("volcengine_nas_snapshots", dataSourceType(volcengineNasMod, "Snapshots"),
		dataSourceType(volcengineNasMod, "getSnapshots"), volcengineNasMod, volcengineNasMod, nil)
	prov.RenameDataSource("volcengine_rds_ip_lists", dataSourceType(volcengineRdsMod, "IpLists"),
		dataSourceType(volcengineRdsMod, "getIpLists"), volcengineRdsMod, volcengineRdsMod, nil)
	prov.RenameDataSource("volcengine_rds_databases", dataSourceType(volcengineRdsMod, "Databases"),
		dataSourceType(volcengineRdsMod, "getDatabases"), volcengineRdsMod, volcengineRdsMod, nil)
	prov.RenameDataSource("volcengine_rds_accounts", dataSourceType(volcengineRdsMod, "Accounts"),
		dataSourceType(volcengineRdsMod, "getAccounts"), volcengineRdsMod, volcengineRdsMod, nil)
	prov.RenameDataSource("volcengine_rds_parameter_templates", dataSourceType(volcengineRdsMod, "ParameterTemplates"),
		dataSourceType(volcengineRdsMod, "getParameterTemplates"), volcengineRdsMod, volcengineRdsMod, nil)
	prov.RenameDataSource("volcengine_rds_instances", dataSourceType(volcengineRdsMod, "Instances"),
		dataSourceType(volcengineRdsMod, "getInstances"), volcengineRdsMod, volcengineRdsMod, nil)
	prov.RenameDataSource("volcengine_listeners", dataSourceType(volcengineClbMod, "Listeners"),
		dataSourceType(volcengineClbMod, "getListeners"), volcengineClbMod, volcengineClbMod, nil)
	prov.RenameDataSource("volcengine_acls", dataSourceType(volcengineClbMod, "Acls"),
		dataSourceType(volcengineClbMod, "getAcls"), volcengineClbMod, volcengineClbMod, nil)
	prov.RenameDataSource("volcengine_clb_rules", dataSourceType(volcengineClbMod, "Rules"),
		dataSourceType(volcengineClbMod, "getRules"), volcengineClbMod, volcengineClbMod, nil)
	prov.RenameDataSource("volcengine_clbs", dataSourceType(volcengineClbMod, "Clbs"),
		dataSourceType(volcengineClbMod, "getClbs"), volcengineClbMod, volcengineClbMod, nil)
	prov.RenameDataSource("volcengine_server_groups", dataSourceType(volcengineClbMod, "ServerGroups"),
		dataSourceType(volcengineClbMod, "getServerGroups"), volcengineClbMod, volcengineClbMod, nil)
	prov.RenameDataSource("volcengine_server_group_servers", dataSourceType(volcengineClbMod, "ServerGroupServers"),
		dataSourceType(volcengineClbMod, "getServerGroupServers"), volcengineClbMod, volcengineClbMod, nil)
	prov.RenameDataSource("volcengine_clb_zones", dataSourceType(volcengineClbMod, "Zones"),
		dataSourceType(volcengineClbMod, "getZones"), volcengineClbMod, volcengineClbMod, nil)
	prov.RenameDataSource("volcengine_certificates", dataSourceType(volcengineClbMod, "Certificates"),
		dataSourceType(volcengineClbMod, "getCertificates"), volcengineClbMod, volcengineClbMod, nil)
	prov.RenameDataSource("volcengine_financial_relations", dataSourceType(volcengineFinancialRelationMod, "FinancialRelations"),
		dataSourceType(volcengineFinancialRelationMod, "getFinancialRelations"), volcengineFinancialRelationMod, volcengineFinancialRelationMod, nil)
	prov.RenameDataSource("volcengine_snat_entries", dataSourceType(volcengineNatMod, "SnatEntries"),
		dataSourceType(volcengineNatMod, "getSnatEntries"), volcengineNatMod, volcengineNatMod, nil)
	prov.RenameDataSource("volcengine_nat_gateways", dataSourceType(volcengineNatMod, "Gateways"),
		dataSourceType(volcengineNatMod, "getGateways"), volcengineNatMod, volcengineNatMod, nil)
	prov.RenameDataSource("volcengine_dnat_entries", dataSourceType(volcengineNatMod, "DnatEntries"),
		dataSourceType(volcengineNatMod, "getDnatEntries"), volcengineNatMod, volcengineNatMod, nil)
	prov.RenameDataSource("volcengine_direct_connect_gateways", dataSourceType(volcengineDirectConnectMod, "Gateways"),
		dataSourceType(volcengineDirectConnectMod, "getGateways"), volcengineDirectConnectMod, volcengineDirectConnectMod, nil)
	prov.RenameDataSource("volcengine_direct_connect_gateway_routes", dataSourceType(volcengineDirectConnectMod, "GatewayRoutes"),
		dataSourceType(volcengineDirectConnectMod, "getGatewayRoutes"), volcengineDirectConnectMod, volcengineDirectConnectMod, nil)
	prov.RenameDataSource("volcengine_direct_connect_connections", dataSourceType(volcengineDirectConnectMod, "Connections"),
		dataSourceType(volcengineDirectConnectMod, "getConnections"), volcengineDirectConnectMod, volcengineDirectConnectMod, nil)
	prov.RenameDataSource("volcengine_direct_connect_virtual_interfaces", dataSourceType(volcengineDirectConnectMod, "VirtualInterfaces"),
		dataSourceType(volcengineDirectConnectMod, "getVirtualInterfaces"), volcengineDirectConnectMod, volcengineDirectConnectMod, nil)
	prov.RenameDataSource("volcengine_direct_connect_bgp_peers", dataSourceType(volcengineDirectConnectMod, "BgpPeers"),
		dataSourceType(volcengineDirectConnectMod, "getBgpPeers"), volcengineDirectConnectMod, volcengineDirectConnectMod, nil)
	prov.RenameDataSource("volcengine_vke_node_pools", dataSourceType(volcengineVkeMod, "NodePools"),
		dataSourceType(volcengineVkeMod, "getNodePools"), volcengineVkeMod, volcengineVkeMod, nil)
	prov.RenameDataSource("volcengine_vke_kubeconfigs", dataSourceType(volcengineVkeMod, "Kubeconfigs"),
		dataSourceType(volcengineVkeMod, "getKubeconfigs"), volcengineVkeMod, volcengineVkeMod, nil)
	prov.RenameDataSource("volcengine_vke_nodes", dataSourceType(volcengineVkeMod, "Nodes"),
		dataSourceType(volcengineVkeMod, "getNodes"), volcengineVkeMod, volcengineVkeMod, nil)
	prov.RenameDataSource("volcengine_vke_clusters", dataSourceType(volcengineVkeMod, "Clusters"),
		dataSourceType(volcengineVkeMod, "getClusters"), volcengineVkeMod, volcengineVkeMod, nil)
	prov.RenameDataSource("volcengine_vke_support_addons", dataSourceType(volcengineVkeMod, "SupportAddons"),
		dataSourceType(volcengineVkeMod, "getSupportAddons"), volcengineVkeMod, volcengineVkeMod, nil)
	prov.RenameDataSource("volcengine_vke_addons", dataSourceType(volcengineVkeMod, "Addons"),
		dataSourceType(volcengineVkeMod, "getAddons"), volcengineVkeMod, volcengineVkeMod, nil)
	prov.RenameDataSource("volcengine_vke_permissions", dataSourceType(volcengineVkeMod, "Permissions"),
		dataSourceType(volcengineVkeMod, "getPermissions"), volcengineVkeMod, volcengineVkeMod, nil)
	prov.RenameDataSource("volcengine_vke_support_resource_types", dataSourceType(volcengineVkeMod, "SupportResourceTypes"),
		dataSourceType(volcengineVkeMod, "getSupportResourceTypes"), volcengineVkeMod, volcengineVkeMod, nil)
	prov.RenameDataSource("volcengine_kms_keyrings", dataSourceType(volcengineKmsMod, "Keyrings"),
		dataSourceType(volcengineKmsMod, "getKeyrings"), volcengineKmsMod, volcengineKmsMod, nil)
	prov.RenameDataSource("volcengine_kms_keys", dataSourceType(volcengineKmsMod, "Keys"),
		dataSourceType(volcengineKmsMod, "getKeys"), volcengineKmsMod, volcengineKmsMod, nil)
	prov.RenameDataSource("volcengine_kms_secrets", dataSourceType(volcengineKmsMod, "Secrets"),
		dataSourceType(volcengineKmsMod, "getSecrets"), volcengineKmsMod, volcengineKmsMod, nil)
	prov.RenameDataSource("volcengine_cloud_monitor_event_rules", dataSourceType(volcengineCloudMonitorMod, "EventRules"),
		dataSourceType(volcengineCloudMonitorMod, "getEventRules"), volcengineCloudMonitorMod, volcengineCloudMonitorMod, nil)
	prov.RenameDataSource("volcengine_cloud_monitor_contact_groups", dataSourceType(volcengineCloudMonitorMod, "ContactGroups"),
		dataSourceType(volcengineCloudMonitorMod, "getContactGroups"), volcengineCloudMonitorMod, volcengineCloudMonitorMod, nil)
	prov.RenameDataSource("volcengine_cloud_monitor_contacts", dataSourceType(volcengineCloudMonitorMod, "Contacts"),
		dataSourceType(volcengineCloudMonitorMod, "getContacts"), volcengineCloudMonitorMod, volcengineCloudMonitorMod, nil)
	prov.RenameDataSource("volcengine_cloud_monitor_rules", dataSourceType(volcengineCloudMonitorMod, "Rules"),
		dataSourceType(volcengineCloudMonitorMod, "getRules"), volcengineCloudMonitorMod, volcengineCloudMonitorMod, nil)
	prov.RenameDataSource("volcengine_iam_user_group_policy_attachments", dataSourceType(volcengineIamMod, "UserGroupPolicyAttachments"),
		dataSourceType(volcengineIamMod, "getUserGroupPolicyAttachments"), volcengineIamMod, volcengineIamMod, nil)
	prov.RenameDataSource("volcengine_iam_roles", dataSourceType(volcengineIamMod, "Roles"),
		dataSourceType(volcengineIamMod, "getRoles"), volcengineIamMod, volcengineIamMod, nil)
	prov.RenameDataSource("volcengine_iam_access_keys", dataSourceType(volcengineIamMod, "AccessKeys"),
		dataSourceType(volcengineIamMod, "getAccessKeys"), volcengineIamMod, volcengineIamMod, nil)
	prov.RenameDataSource("volcengine_iam_saml_providers", dataSourceType(volcengineIamMod, "SamlProviders"),
		dataSourceType(volcengineIamMod, "getSamlProviders"), volcengineIamMod, volcengineIamMod, nil)
	prov.RenameDataSource("volcengine_iam_policies", dataSourceType(volcengineIamMod, "Policies"),
		dataSourceType(volcengineIamMod, "getPolicies"), volcengineIamMod, volcengineIamMod, nil)
	prov.RenameDataSource("volcengine_iam_user_groups", dataSourceType(volcengineIamMod, "UserGroups"),
		dataSourceType(volcengineIamMod, "getUserGroups"), volcengineIamMod, volcengineIamMod, nil)
	prov.RenameDataSource("volcengine_iam_users", dataSourceType(volcengineIamMod, "Users"),
		dataSourceType(volcengineIamMod, "getUsers"), volcengineIamMod, volcengineIamMod, nil)
	prov.RenameDataSource("volcengine_privatelink_vpc_endpoint_services", dataSourceType(volcenginePrivatelinkMod, "VpcEndpointServices"),
		dataSourceType(volcenginePrivatelinkMod, "getVpcEndpointServices"), volcenginePrivatelinkMod, volcenginePrivatelinkMod, nil)
	prov.RenameDataSource("volcengine_privatelink_vpc_endpoint_zones", dataSourceType(volcenginePrivatelinkMod, "VpcEndpointZones"),
		dataSourceType(volcenginePrivatelinkMod, "getVpcEndpointZones"), volcenginePrivatelinkMod, volcenginePrivatelinkMod, nil)
	prov.RenameDataSource("volcengine_privatelink_vpc_endpoints", dataSourceType(volcenginePrivatelinkMod, "VpcEndpoints"),
		dataSourceType(volcenginePrivatelinkMod, "getVpcEndpoints"), volcenginePrivatelinkMod, volcenginePrivatelinkMod, nil)
	prov.RenameDataSource("volcengine_privatelink_vpc_endpoint_service_permissions", dataSourceType(volcenginePrivatelinkMod, "VpcEndpointServicePermissions"),
		dataSourceType(volcenginePrivatelinkMod, "getVpcEndpointServicePermissions"), volcenginePrivatelinkMod, volcenginePrivatelinkMod, nil)
	prov.RenameDataSource("volcengine_privatelink_vpc_endpoint_connections", dataSourceType(volcenginePrivatelinkMod, "VpcEndpointConnections"),
		dataSourceType(volcenginePrivatelinkMod, "getVpcEndpointConnections"), volcenginePrivatelinkMod, volcenginePrivatelinkMod, nil)
	prov.RenameDataSource("volcengine_rocketmq_topics", dataSourceType(volcengineRocketmqMod, "Topics"),
		dataSourceType(volcengineRocketmqMod, "getTopics"), volcengineRocketmqMod, volcengineRocketmqMod, nil)
	prov.RenameDataSource("volcengine_rocketmq_instances", dataSourceType(volcengineRocketmqMod, "Instances"),
		dataSourceType(volcengineRocketmqMod, "getInstances"), volcengineRocketmqMod, volcengineRocketmqMod, nil)
	prov.RenameDataSource("volcengine_rocketmq_groups", dataSourceType(volcengineRocketmqMod, "Groups"),
		dataSourceType(volcengineRocketmqMod, "getGroups"), volcengineRocketmqMod, volcengineRocketmqMod, nil)
	prov.RenameDataSource("volcengine_rocketmq_access_keys", dataSourceType(volcengineRocketmqMod, "AccessKeys"),
		dataSourceType(volcengineRocketmqMod, "getAccessKeys"), volcengineRocketmqMod, volcengineRocketmqMod, nil)
	prov.RenameDataSource("volcengine_rocketmq_allow_lists", dataSourceType(volcengineRocketmqMod, "AllowLists"),
		dataSourceType(volcengineRocketmqMod, "getAllowLists"), volcengineRocketmqMod, volcengineRocketmqMod, nil)
	prov.RenameDataSource("volcengine_transit_router_route_entries", dataSourceType(volcengineTransitRouterMod, "RouteEntries"),
		dataSourceType(volcengineTransitRouterMod, "getRouteEntries"), volcengineTransitRouterMod, volcengineTransitRouterMod, nil)
	prov.RenameDataSource("volcengine_transit_router_route_tables", dataSourceType(volcengineTransitRouterMod, "RouteTables"),
		dataSourceType(volcengineTransitRouterMod, "getRouteTables"), volcengineTransitRouterMod, volcengineTransitRouterMod, nil)
	prov.RenameDataSource("volcengine_transit_routers", dataSourceType(volcengineTransitRouterMod, "TransitRouters"),
		dataSourceType(volcengineTransitRouterMod, "getTransitRouters"), volcengineTransitRouterMod, volcengineTransitRouterMod, nil)
	prov.RenameDataSource("volcengine_transit_router_bandwidth_packages", dataSourceType(volcengineTransitRouterMod, "BandwidthPackages"),
		dataSourceType(volcengineTransitRouterMod, "getBandwidthPackages"), volcengineTransitRouterMod, volcengineTransitRouterMod, nil)
	prov.RenameDataSource("volcengine_transit_router_route_table_associations", dataSourceType(volcengineTransitRouterMod, "RouteTableAssociations"),
		dataSourceType(volcengineTransitRouterMod, "getRouteTableAssociations"), volcengineTransitRouterMod, volcengineTransitRouterMod, nil)
	prov.RenameDataSource("volcengine_transit_router_route_table_propagations", dataSourceType(volcengineTransitRouterMod, "RouteTablePropagations"),
		dataSourceType(volcengineTransitRouterMod, "getRouteTablePropagations"), volcengineTransitRouterMod, volcengineTransitRouterMod, nil)
	prov.RenameDataSource("volcengine_transit_router_grant_rules", dataSourceType(volcengineTransitRouterMod, "GrantRules"),
		dataSourceType(volcengineTransitRouterMod, "getGrantRules"), volcengineTransitRouterMod, volcengineTransitRouterMod, nil)
	prov.RenameDataSource("volcengine_transit_router_vpc_attachments", dataSourceType(volcengineTransitRouterMod, "VpcAttachments"),
		dataSourceType(volcengineTransitRouterMod, "getVpcAttachments"), volcengineTransitRouterMod, volcengineTransitRouterMod, nil)
	prov.RenameDataSource("volcengine_transit_router_direct_connect_gateway_attachments", dataSourceType(volcengineTransitRouterMod, "DirectConnectGatewayAttachments"),
		dataSourceType(volcengineTransitRouterMod, "getDirectConnectGatewayAttachments"), volcengineTransitRouterMod, volcengineTransitRouterMod, nil)
	prov.RenameDataSource("volcengine_transit_router_peer_attachments", dataSourceType(volcengineTransitRouterMod, "PeerAttachments"),
		dataSourceType(volcengineTransitRouterMod, "getPeerAttachments"), volcengineTransitRouterMod, volcengineTransitRouterMod, nil)
	prov.RenameDataSource("volcengine_transit_router_vpn_attachments", dataSourceType(volcengineTransitRouterMod, "VpnAttachments"),
		dataSourceType(volcengineTransitRouterMod, "getVpnAttachments"), volcengineTransitRouterMod, volcengineTransitRouterMod, nil)
	prov.RenameDataSource("volcengine_albs", dataSourceType(volcengineAlbMod, "Albs"),
		dataSourceType(volcengineAlbMod, "getAlbs"), volcengineAlbMod, volcengineAlbMod, nil)
	prov.RenameDataSource("volcengine_alb_listener_domain_extensions", dataSourceType(volcengineAlbMod, "ListenerDomainExtensions"),
		dataSourceType(volcengineAlbMod, "getListenerDomainExtensions"), volcengineAlbMod, volcengineAlbMod, nil)
	prov.RenameDataSource("volcengine_alb_health_check_templates", dataSourceType(volcengineAlbMod, "HealthCheckTemplates"),
		dataSourceType(volcengineAlbMod, "getHealthCheckTemplates"), volcengineAlbMod, volcengineAlbMod, nil)
	prov.RenameDataSource("volcengine_alb_certificates", dataSourceType(volcengineAlbMod, "Certificates"),
		dataSourceType(volcengineAlbMod, "getCertificates"), volcengineAlbMod, volcengineAlbMod, nil)
	prov.RenameDataSource("volcengine_alb_ca_certificates", dataSourceType(volcengineAlbMod, "CaCertificates"),
		dataSourceType(volcengineAlbMod, "getCaCertificates"), volcengineAlbMod, volcengineAlbMod, nil)
	prov.RenameDataSource("volcengine_alb_rules", dataSourceType(volcengineAlbMod, "Rules"),
		dataSourceType(volcengineAlbMod, "getRules"), volcengineAlbMod, volcengineAlbMod, nil)
	prov.RenameDataSource("volcengine_alb_acls", dataSourceType(volcengineAlbMod, "Acls"),
		dataSourceType(volcengineAlbMod, "getAcls"), volcengineAlbMod, volcengineAlbMod, nil)
	prov.RenameDataSource("volcengine_alb_listeners", dataSourceType(volcengineAlbMod, "Listeners"),
		dataSourceType(volcengineAlbMod, "getListeners"), volcengineAlbMod, volcengineAlbMod, nil)
	prov.RenameDataSource("volcengine_alb_server_group_servers", dataSourceType(volcengineAlbMod, "ServerGroupServers"),
		dataSourceType(volcengineAlbMod, "getServerGroupServers"), volcengineAlbMod, volcengineAlbMod, nil)
	prov.RenameDataSource("volcengine_alb_customized_cfgs", dataSourceType(volcengineAlbMod, "CustomizedCfgs"),
		dataSourceType(volcengineAlbMod, "getCustomizedCfgs"), volcengineAlbMod, volcengineAlbMod, nil)
	prov.RenameDataSource("volcengine_alb_zones", dataSourceType(volcengineAlbMod, "Zones"),
		dataSourceType(volcengineAlbMod, "getZones"), volcengineAlbMod, volcengineAlbMod, nil)
	prov.RenameDataSource("volcengine_alb_server_groups", dataSourceType(volcengineAlbMod, "ServerGroups"),
		dataSourceType(volcengineAlbMod, "getServerGroups"), volcengineAlbMod, volcengineAlbMod, nil)
	prov.RenameDataSource("volcengine_dns_records", dataSourceType(volcengineDnsMod, "Records"),
		dataSourceType(volcengineDnsMod, "getRecords"), volcengineDnsMod, volcengineDnsMod, nil)
	prov.RenameDataSource("volcengine_dns_zones", dataSourceType(volcengineDnsMod, "Zones"),
		dataSourceType(volcengineDnsMod, "getZones"), volcengineDnsMod, volcengineDnsMod, nil)
	prov.RenameDataSource("volcengine_dns_backups", dataSourceType(volcengineDnsMod, "Backups"),
		dataSourceType(volcengineDnsMod, "getBackups"), volcengineDnsMod, volcengineDnsMod, nil)
	prov.RenameDataSource("volcengine_dns_record_sets", dataSourceType(volcengineDnsMod, "RecordSets"),
		dataSourceType(volcengineDnsMod, "getRecordSets"), volcengineDnsMod, volcengineDnsMod, nil)
	prov.RenameDataSource("volcengine_mongodb_zones", dataSourceType(volcengineMongodbMod, "Zones"),
		dataSourceType(volcengineMongodbMod, "getZones"), volcengineMongodbMod, volcengineMongodbMod, nil)
	prov.RenameDataSource("volcengine_mongodb_instances", dataSourceType(volcengineMongodbMod, "Instances"),
		dataSourceType(volcengineMongodbMod, "getInstances"), volcengineMongodbMod, volcengineMongodbMod, nil)
	prov.RenameDataSource("volcengine_mongodb_ssl_states", dataSourceType(volcengineMongodbMod, "SslStates"),
		dataSourceType(volcengineMongodbMod, "getSslStates"), volcengineMongodbMod, volcengineMongodbMod, nil)
	prov.RenameDataSource("volcengine_mongodb_endpoints", dataSourceType(volcengineMongodbMod, "Endpoints"),
		dataSourceType(volcengineMongodbMod, "getEndpoints"), volcengineMongodbMod, volcengineMongodbMod, nil)
	prov.RenameDataSource("volcengine_mongodb_instance_parameters", dataSourceType(volcengineMongodbMod, "InstanceParameters"),
		dataSourceType(volcengineMongodbMod, "getInstanceParameters"), volcengineMongodbMod, volcengineMongodbMod, nil)
	prov.RenameDataSource("volcengine_mongodb_instance_parameter_logs", dataSourceType(volcengineMongodbMod, "InstanceParameterLogs"),
		dataSourceType(volcengineMongodbMod, "getInstanceParameterLogs"), volcengineMongodbMod, volcengineMongodbMod, nil)
	prov.RenameDataSource("volcengine_mongodb_specs", dataSourceType(volcengineMongodbMod, "Specs"),
		dataSourceType(volcengineMongodbMod, "getSpecs"), volcengineMongodbMod, volcengineMongodbMod, nil)
	prov.RenameDataSource("volcengine_mongodb_regions", dataSourceType(volcengineMongodbMod, "Regions"),
		dataSourceType(volcengineMongodbMod, "getRegions"), volcengineMongodbMod, volcengineMongodbMod, nil)
	prov.RenameDataSource("volcengine_mongodb_allow_lists", dataSourceType(volcengineMongodbMod, "MongoAllowLists"),
		dataSourceType(volcengineMongodbMod, "getMongoAllowLists"), volcengineMongodbMod, volcengineMongodbMod, nil)
	prov.RenameDataSource("volcengine_mongodb_accounts", dataSourceType(volcengineMongodbMod, "Accounts"),
		dataSourceType(volcengineMongodbMod, "getAccounts"), volcengineMongodbMod, volcengineMongodbMod, nil)
	prov.RenameDataSource("volcengine_cdn_certificates", dataSourceType(volcengineCdnMod, "Certificates"),
		dataSourceType(volcengineCdnMod, "getCertificates"), volcengineCdnMod, volcengineCdnMod, nil)
	prov.RenameDataSource("volcengine_cdn_shared_configs", dataSourceType(volcengineCdnMod, "SharedConfigs"),
		dataSourceType(volcengineCdnMod, "getSharedConfigs"), volcengineCdnMod, volcengineCdnMod, nil)
	prov.RenameDataSource("volcengine_cdn_domains", dataSourceType(volcengineCdnMod, "Domains"),
		dataSourceType(volcengineCdnMod, "getDomains"), volcengineCdnMod, volcengineCdnMod, nil)
	prov.RenameDataSource("volcengine_cdn_configs", dataSourceType(volcengineCdnMod, "Configs"),
		dataSourceType(volcengineCdnMod, "getConfigs"), volcengineCdnMod, volcengineCdnMod, nil)
	prov.RenameDataSource("volcengine_cfw_dns_control_policies", dataSourceType(volcengineCloudFirewallMod, "CfwDnsControlPolicies"),
		dataSourceType(volcengineCloudFirewallMod, "getCfwDnsControlPolicies"), volcengineCloudFirewallMod, volcengineCloudFirewallMod, nil)
	prov.RenameDataSource("volcengine_cfw_address_books", dataSourceType(volcengineCloudFirewallMod, "CfwAddressBooks"),
		dataSourceType(volcengineCloudFirewallMod, "getCfwAddressBooks"), volcengineCloudFirewallMod, volcengineCloudFirewallMod, nil)
	prov.RenameDataSource("volcengine_cfw_vpc_firewall_acl_rules", dataSourceType(volcengineCloudFirewallMod, "CfwVpcFirewallAclRules"),
		dataSourceType(volcengineCloudFirewallMod, "getCfwVpcFirewallAclRules"), volcengineCloudFirewallMod, volcengineCloudFirewallMod, nil)
	prov.RenameDataSource("volcengine_cfw_nat_firewall_control_policies", dataSourceType(volcengineCloudFirewallMod, "CfwNatFirewallControlPolicies"),
		dataSourceType(volcengineCloudFirewallMod, "getCfwNatFirewallControlPolicies"), volcengineCloudFirewallMod, volcengineCloudFirewallMod, nil)
	prov.RenameDataSource("volcengine_cfw_control_policies", dataSourceType(volcengineCloudFirewallMod, "CfwControlPolicies"),
		dataSourceType(volcengineCloudFirewallMod, "getCfwControlPolicies"), volcengineCloudFirewallMod, volcengineCloudFirewallMod, nil)
	prov.RenameDataSource("volcengine_ecs_invocations", dataSourceType(volcengineEcsMod, "Invocations"),
		dataSourceType(volcengineEcsMod, "getInvocations"), volcengineEcsMod, volcengineEcsMod, nil)
	prov.RenameDataSource("volcengine_image_share_permissions", dataSourceType(volcengineEcsMod, "ImageSharePermissions"),
		dataSourceType(volcengineEcsMod, "getImageSharePermissions"), volcengineEcsMod, volcengineEcsMod, nil)
	prov.RenameDataSource("volcengine_ecs_instance_types", dataSourceType(volcengineEcsMod, "InstanceTypes"),
		dataSourceType(volcengineEcsMod, "getInstanceTypes"), volcengineEcsMod, volcengineEcsMod, nil)
	prov.RenameDataSource("volcengine_ecs_hpc_clusters", dataSourceType(volcengineEcsMod, "HpcClusters"),
		dataSourceType(volcengineEcsMod, "getHpcClusters"), volcengineEcsMod, volcengineEcsMod, nil)
	prov.RenameDataSource("volcengine_ecs_launch_templates", dataSourceType(volcengineEcsMod, "LaunchTemplates"),
		dataSourceType(volcengineEcsMod, "getLaunchTemplates"), volcengineEcsMod, volcengineEcsMod, nil)
	prov.RenameDataSource("volcengine_ecs_available_resources", dataSourceType(volcengineEcsMod, "AvailableResources"),
		dataSourceType(volcengineEcsMod, "getAvailableResources"), volcengineEcsMod, volcengineEcsMod, nil)
	prov.RenameDataSource("volcengine_ecs_invocation_results", dataSourceType(volcengineEcsMod, "InvocationResults"),
		dataSourceType(volcengineEcsMod, "getInvocationResults"), volcengineEcsMod, volcengineEcsMod, nil)
	prov.RenameDataSource("volcengine_regions", dataSourceType(volcengineEcsMod, "Regions"),
		dataSourceType(volcengineEcsMod, "getRegions"), volcengineEcsMod, volcengineEcsMod, nil)
	prov.RenameDataSource("volcengine_ecs_commands", dataSourceType(volcengineEcsMod, "Commands"),
		dataSourceType(volcengineEcsMod, "getCommands"), volcengineEcsMod, volcengineEcsMod, nil)
	prov.RenameDataSource("volcengine_images", dataSourceType(volcengineEcsMod, "Images"),
		dataSourceType(volcengineEcsMod, "getImages"), volcengineEcsMod, volcengineEcsMod, nil)
	prov.RenameDataSource("volcengine_ecs_key_pairs", dataSourceType(volcengineEcsMod, "KeyPairs"),
		dataSourceType(volcengineEcsMod, "getKeyPairs"), volcengineEcsMod, volcengineEcsMod, nil)
	prov.RenameDataSource("volcengine_ecs_instances", dataSourceType(volcengineEcsMod, "Instances"),
		dataSourceType(volcengineEcsMod, "getInstances"), volcengineEcsMod, volcengineEcsMod, nil)
	prov.RenameDataSource("volcengine_ecs_deployment_sets", dataSourceType(volcengineEcsMod, "DeploymentSets"),
		dataSourceType(volcengineEcsMod, "getDeploymentSets"), volcengineEcsMod, volcengineEcsMod, nil)
	prov.RenameDataSource("volcengine_zones", dataSourceType(volcengineEcsMod, "Zones"),
		dataSourceType(volcengineEcsMod, "getZones"), volcengineEcsMod, volcengineEcsMod, nil)
	prov.RenameDataSource("volcengine_rabbitmq_instance_plugins", dataSourceType(volcengineRabbitmqMod, "InstancePlugins"),
		dataSourceType(volcengineRabbitmqMod, "getInstancePlugins"), volcengineRabbitmqMod, volcengineRabbitmqMod, nil)
	prov.RenameDataSource("volcengine_rabbitmq_zones", dataSourceType(volcengineRabbitmqMod, "Zones"),
		dataSourceType(volcengineRabbitmqMod, "getZones"), volcengineRabbitmqMod, volcengineRabbitmqMod, nil)
	prov.RenameDataSource("volcengine_rabbitmq_regions", dataSourceType(volcengineRabbitmqMod, "Regions"),
		dataSourceType(volcengineRabbitmqMod, "getRegions"), volcengineRabbitmqMod, volcengineRabbitmqMod, nil)
	prov.RenameDataSource("volcengine_rabbitmq_instances", dataSourceType(volcengineRabbitmqMod, "Instances"),
		dataSourceType(volcengineRabbitmqMod, "getInstances"), volcengineRabbitmqMod, volcengineRabbitmqMod, nil)
	prov.RenameDataSource("volcengine_organization_units", dataSourceType(volcengineOrganizationMod, "Units"),
		dataSourceType(volcengineOrganizationMod, "getUnits"), volcengineOrganizationMod, volcengineOrganizationMod, nil)
	prov.RenameDataSource("volcengine_organization_accounts", dataSourceType(volcengineOrganizationMod, "Accounts"),
		dataSourceType(volcengineOrganizationMod, "getAccounts"), volcengineOrganizationMod, volcengineOrganizationMod, nil)
	prov.RenameDataSource("volcengine_organizations", dataSourceType(volcengineOrganizationMod, "Organizations"),
		dataSourceType(volcengineOrganizationMod, "getOrganizations"), volcengineOrganizationMod, volcengineOrganizationMod, nil)
	prov.RenameDataSource("volcengine_organization_service_control_policies", dataSourceType(volcengineOrganizationMod, "ServiceControlPolicies"),
		dataSourceType(volcengineOrganizationMod, "getServiceControlPolicies"), volcengineOrganizationMod, volcengineOrganizationMod, nil)
	prov.RenameDataSource("volcengine_tos_bucket_inventories", dataSourceType(volcengineTosMod, "BucketInventories"),
		dataSourceType(volcengineTosMod, "getBucketInventories"), volcengineTosMod, volcengineTosMod, nil)
	prov.RenameDataSource("volcengine_tos_buckets", dataSourceType(volcengineTosMod, "Buckets"),
		dataSourceType(volcengineTosMod, "getBuckets"), volcengineTosMod, volcengineTosMod, nil)
	prov.RenameDataSource("volcengine_tos_objects", dataSourceType(volcengineTosMod, "BucketObjects"),
		dataSourceType(volcengineTosMod, "getBucketObjects"), volcengineTosMod, volcengineTosMod, nil)
	prov.RenameDataSource("volcengine_apig_upstreams", dataSourceType(volcengineApigMod, "Upstreams"),
		dataSourceType(volcengineApigMod, "getUpstreams"), volcengineApigMod, volcengineApigMod, nil)
	prov.RenameDataSource("volcengine_apig_upstream_sources", dataSourceType(volcengineApigMod, "UpstreamSources"),
		dataSourceType(volcengineApigMod, "getUpstreamSources"), volcengineApigMod, volcengineApigMod, nil)
	prov.RenameDataSource("volcengine_apig_upstream_versions", dataSourceType(volcengineApigMod, "UpstreamVersions"),
		dataSourceType(volcengineApigMod, "getUpstreamVersions"), volcengineApigMod, volcengineApigMod, nil)
	prov.RenameDataSource("volcengine_apig_gateway_services", dataSourceType(volcengineApigMod, "GatewayServices"),
		dataSourceType(volcengineApigMod, "getGatewayServices"), volcengineApigMod, volcengineApigMod, nil)
	prov.RenameDataSource("volcengine_apig_gateways", dataSourceType(volcengineApigMod, "Gateways"),
		dataSourceType(volcengineApigMod, "getGateways"), volcengineApigMod, volcengineApigMod, nil)
	prov.RenameDataSource("volcengine_apig_routes", dataSourceType(volcengineApigMod, "Routes"),
		dataSourceType(volcengineApigMod, "getRoutes"), volcengineApigMod, volcengineApigMod, nil)
	prov.RenameDataSource("volcengine_apig_custom_domains", dataSourceType(volcengineApigMod, "CustomDomains"),
		dataSourceType(volcengineApigMod, "getCustomDomains"), volcengineApigMod, volcengineApigMod, nil)
	prov.RenameDataSource("volcengine_rds_instances_v2", dataSourceType(volcengineRdsV2Mod, "RdsInstancesV2"),
		dataSourceType(volcengineRdsV2Mod, "getRdsInstancesV2"), volcengineRdsV2Mod, volcengineRdsV2Mod, nil)
	prov.RenameDataSource("volcengine_kafka_sasl_users", dataSourceType(volcengineKafkaMod, "SaslUsers"),
		dataSourceType(volcengineKafkaMod, "getSaslUsers"), volcengineKafkaMod, volcengineKafkaMod, nil)
	prov.RenameDataSource("volcengine_kafka_allow_lists", dataSourceType(volcengineKafkaMod, "AllowLists"),
		dataSourceType(volcengineKafkaMod, "getAllowLists"), volcengineKafkaMod, volcengineKafkaMod, nil)
	prov.RenameDataSource("volcengine_kafka_groups", dataSourceType(volcengineKafkaMod, "Groups"),
		dataSourceType(volcengineKafkaMod, "getGroups"), volcengineKafkaMod, volcengineKafkaMod, nil)
	prov.RenameDataSource("volcengine_kafka_consumed_partitions", dataSourceType(volcengineKafkaMod, "ConsumedPartitions"),
		dataSourceType(volcengineKafkaMod, "getConsumedPartitions"), volcengineKafkaMod, volcengineKafkaMod, nil)
	prov.RenameDataSource("volcengine_kafka_zones", dataSourceType(volcengineKafkaMod, "Zones"),
		dataSourceType(volcengineKafkaMod, "getZones"), volcengineKafkaMod, volcengineKafkaMod, nil)
	prov.RenameDataSource("volcengine_kafka_topics", dataSourceType(volcengineKafkaMod, "Topics"),
		dataSourceType(volcengineKafkaMod, "getTopics"), volcengineKafkaMod, volcengineKafkaMod, nil)
	prov.RenameDataSource("volcengine_kafka_instances", dataSourceType(volcengineKafkaMod, "Instances"),
		dataSourceType(volcengineKafkaMod, "getInstances"), volcengineKafkaMod, volcengineKafkaMod, nil)
	prov.RenameDataSource("volcengine_kafka_consumed_topics", dataSourceType(volcengineKafkaMod, "ConsumedTopics"),
		dataSourceType(volcengineKafkaMod, "getConsumedTopics"), volcengineKafkaMod, volcengineKafkaMod, nil)
	prov.RenameDataSource("volcengine_kafka_regions", dataSourceType(volcengineKafkaMod, "Regions"),
		dataSourceType(volcengineKafkaMod, "getRegions"), volcengineKafkaMod, volcengineKafkaMod, nil)
	prov.RenameDataSource("volcengine_kafka_topic_partitions", dataSourceType(volcengineKafkaMod, "TopicPartitions"),
		dataSourceType(volcengineKafkaMod, "getTopicPartitions"), volcengineKafkaMod, volcengineKafkaMod, nil)
	prov.RenameDataSource("volcengine_private_zone_records", dataSourceType(volcenginePrivateZoneMod, "Records"),
		dataSourceType(volcenginePrivateZoneMod, "getRecords"), volcenginePrivateZoneMod, volcenginePrivateZoneMod, nil)
	prov.RenameDataSource("volcengine_private_zone_resolver_rules", dataSourceType(volcenginePrivateZoneMod, "ResolverRules"),
		dataSourceType(volcenginePrivateZoneMod, "getResolverRules"), volcenginePrivateZoneMod, volcenginePrivateZoneMod, nil)
	prov.RenameDataSource("volcengine_private_zone_resolver_endpoints", dataSourceType(volcenginePrivateZoneMod, "ResolverEndpoints"),
		dataSourceType(volcenginePrivateZoneMod, "getResolverEndpoints"), volcenginePrivateZoneMod, volcenginePrivateZoneMod, nil)
	prov.RenameDataSource("volcengine_private_zones", dataSourceType(volcenginePrivateZoneMod, "PrivateZones"),
		dataSourceType(volcenginePrivateZoneMod, "getPrivateZones"), volcenginePrivateZoneMod, volcenginePrivateZoneMod, nil)
	prov.RenameDataSource("volcengine_private_zone_record_sets", dataSourceType(volcenginePrivateZoneMod, "RecordSets"),
		dataSourceType(volcenginePrivateZoneMod, "getRecordSets"), volcenginePrivateZoneMod, volcenginePrivateZoneMod, nil)
	prov.RenameDataSource("volcengine_vefaas_releases", dataSourceType(volcengineVefaasMod, "Releases"),
		dataSourceType(volcengineVefaasMod, "getReleases"), volcengineVefaasMod, volcengineVefaasMod, nil)
	prov.RenameDataSource("volcengine_vefaas_kafka_triggers", dataSourceType(volcengineVefaasMod, "KafkaTriggers"),
		dataSourceType(volcengineVefaasMod, "getKafkaTriggers"), volcengineVefaasMod, volcengineVefaasMod, nil)
	prov.RenameDataSource("volcengine_vefaas_functions", dataSourceType(volcengineVefaasMod, "Functions"),
		dataSourceType(volcengineVefaasMod, "getFunctions"), volcengineVefaasMod, volcengineVefaasMod, nil)
	prov.RenameDataSource("volcengine_vefaas_timers", dataSourceType(volcengineVefaasMod, "Timers"),
		dataSourceType(volcengineVefaasMod, "getTimers"), volcengineVefaasMod, volcengineVefaasMod, nil)
	prov.RenameDataSource("volcengine_waf_custom_bots", dataSourceType(volcengineWafMod, "CustomBots"),
		dataSourceType(volcengineWafMod, "getCustomBots"), volcengineWafMod, volcengineWafMod, nil)
	prov.RenameDataSource("volcengine_waf_acl_rules", dataSourceType(volcengineWafMod, "AclRules"),
		dataSourceType(volcengineWafMod, "getAclRules"), volcengineWafMod, volcengineWafMod, nil)
	prov.RenameDataSource("volcengine_waf_system_bots", dataSourceType(volcengineWafMod, "SystemBots"),
		dataSourceType(volcengineWafMod, "getSystemBots"), volcengineWafMod, volcengineWafMod, nil)
	prov.RenameDataSource("volcengine_waf_service_certificates", dataSourceType(volcengineWafMod, "ServiceCertificates"),
		dataSourceType(volcengineWafMod, "getServiceCertificates"), volcengineWafMod, volcengineWafMod, nil)
	prov.RenameDataSource("volcengine_waf_custom_pages", dataSourceType(volcengineWafMod, "CustomPages"),
		dataSourceType(volcengineWafMod, "getCustomPages"), volcengineWafMod, volcengineWafMod, nil)
	prov.RenameDataSource("volcengine_waf_domains", dataSourceType(volcengineWafMod, "Domains"),
		dataSourceType(volcengineWafMod, "getDomains"), volcengineWafMod, volcengineWafMod, nil)
	prov.RenameDataSource("volcengine_waf_host_groups", dataSourceType(volcengineWafMod, "HostGroups"),
		dataSourceType(volcengineWafMod, "getHostGroups"), volcengineWafMod, volcengineWafMod, nil)
	prov.RenameDataSource("volcengine_waf_prohibitions", dataSourceType(volcengineWafMod, "Prohibitions"),
		dataSourceType(volcengineWafMod, "getProhibitions"), volcengineWafMod, volcengineWafMod, nil)
	prov.RenameDataSource("volcengine_waf_bot_analyse_protect_rules", dataSourceType(volcengineWafMod, "BotAnalyseProtectRules"),
		dataSourceType(volcengineWafMod, "getBotAnalyseProtectRules"), volcengineWafMod, volcengineWafMod, nil)
	prov.RenameDataSource("volcengine_waf_cc_rules", dataSourceType(volcengineWafMod, "CcRules"),
		dataSourceType(volcengineWafMod, "getCcRules"), volcengineWafMod, volcengineWafMod, nil)
	prov.RenameDataSource("volcengine_waf_ip_groups", dataSourceType(volcengineWafMod, "IpGroups"),
		dataSourceType(volcengineWafMod, "getIpGroups"), volcengineWafMod, volcengineWafMod, nil)
	prov.RenameDataSource("volcengine_vepfs_filesets", dataSourceType(volcengineVepfsMod, "Filesets"),
		dataSourceType(volcengineVepfsMod, "getFilesets"), volcengineVepfsMod, volcengineVepfsMod, nil)
	prov.RenameDataSource("volcengine_vepfs_file_systems", dataSourceType(volcengineVepfsMod, "FileSystems"),
		dataSourceType(volcengineVepfsMod, "getFileSystems"), volcengineVepfsMod, volcengineVepfsMod, nil)
	prov.RenameDataSource("volcengine_vepfs_mount_services", dataSourceType(volcengineVepfsMod, "MountServices"),
		dataSourceType(volcengineVepfsMod, "getMountServices"), volcengineVepfsMod, volcengineVepfsMod, nil)
	prov.RenameDataSource("volcengine_scaling_configurations", dataSourceType(volcengineAutoscalingMod, "ScalingConfigurations"),
		dataSourceType(volcengineAutoscalingMod, "getScalingConfigurations"), volcengineAutoscalingMod, volcengineAutoscalingMod, nil)
	prov.RenameDataSource("volcengine_scaling_lifecycle_hooks", dataSourceType(volcengineAutoscalingMod, "ScalingLifecycleHooks"),
		dataSourceType(volcengineAutoscalingMod, "getScalingLifecycleHooks"), volcengineAutoscalingMod, volcengineAutoscalingMod, nil)
	prov.RenameDataSource("volcengine_scaling_groups", dataSourceType(volcengineAutoscalingMod, "ScalingGroups"),
		dataSourceType(volcengineAutoscalingMod, "getScalingGroups"), volcengineAutoscalingMod, volcengineAutoscalingMod, nil)
	prov.RenameDataSource("volcengine_scaling_policies", dataSourceType(volcengineAutoscalingMod, "ScalingPolicies"),
		dataSourceType(volcengineAutoscalingMod, "getScalingPolicies"), volcengineAutoscalingMod, volcengineAutoscalingMod, nil)
	prov.RenameDataSource("volcengine_scaling_instances", dataSourceType(volcengineAutoscalingMod, "ScalingInstances"),
		dataSourceType(volcengineAutoscalingMod, "getScalingInstances"), volcengineAutoscalingMod, volcengineAutoscalingMod, nil)
	prov.RenameDataSource("volcengine_scaling_activities", dataSourceType(volcengineAutoscalingMod, "ScalingActivities"),
		dataSourceType(volcengineAutoscalingMod, "getScalingActivities"), volcengineAutoscalingMod, volcengineAutoscalingMod, nil)
	prov.RenameDataSource("volcengine_bandwidth_packages", dataSourceType(volcengineBandwidthPackageMod, "BandwidthPackages"),
		dataSourceType(volcengineBandwidthPackageMod, "getBandwidthPackages"), volcengineBandwidthPackageMod, volcengineBandwidthPackageMod, nil)
	prov.RenameDataSource("volcengine_veecp_edge_nodes", dataSourceType(volcengineVeecpMod, "EdgeNodes"),
		dataSourceType(volcengineVeecpMod, "getEdgeNodes"), volcengineVeecpMod, volcengineVeecpMod, nil)
	prov.RenameDataSource("volcengine_veecp_node_pools", dataSourceType(volcengineVeecpMod, "NodePools"),
		dataSourceType(volcengineVeecpMod, "getNodePools"), volcengineVeecpMod, volcengineVeecpMod, nil)
	prov.RenameDataSource("volcengine_veecp_edge_node_pools", dataSourceType(volcengineVeecpMod, "EdgeNodePools"),
		dataSourceType(volcengineVeecpMod, "getEdgeNodePools"), volcengineVeecpMod, volcengineVeecpMod, nil)
	prov.RenameDataSource("volcengine_veecp_support_addons", dataSourceType(volcengineVeecpMod, "SupportAddons"),
		dataSourceType(volcengineVeecpMod, "getSupportAddons"), volcengineVeecpMod, volcengineVeecpMod, nil)
	prov.RenameDataSource("volcengine_veecp_support_resource_types", dataSourceType(volcengineVeecpMod, "SupportResourceTypes"),
		dataSourceType(volcengineVeecpMod, "getSupportResourceTypes"), volcengineVeecpMod, volcengineVeecpMod, nil)
	prov.RenameDataSource("volcengine_veecp_batch_edge_machines", dataSourceType(volcengineVeecpMod, "BatchEdgeMachines"),
		dataSourceType(volcengineVeecpMod, "getBatchEdgeMachines"), volcengineVeecpMod, volcengineVeecpMod, nil)
	prov.RenameDataSource("volcengine_veecp_addons", dataSourceType(volcengineVeecpMod, "Addons"),
		dataSourceType(volcengineVeecpMod, "getAddons"), volcengineVeecpMod, volcengineVeecpMod, nil)
	prov.RenameDataSource("volcengine_veecp_clusters", dataSourceType(volcengineVeecpMod, "Clusters"),
		dataSourceType(volcengineVeecpMod, "getClusters"), volcengineVeecpMod, volcengineVeecpMod, nil)
	prov.RenameDataSource("volcengine_veenedge_cloud_servers", dataSourceType(volcengineVeenedgeMod, "CloudServers"),
		dataSourceType(volcengineVeenedgeMod, "getCloudServers"), volcengineVeenedgeMod, volcengineVeenedgeMod, nil)
	prov.RenameDataSource("volcengine_veenedge_instance_types", dataSourceType(volcengineVeenedgeMod, "InstanceTypes"),
		dataSourceType(volcengineVeenedgeMod, "getInstanceTypes"), volcengineVeenedgeMod, volcengineVeenedgeMod, nil)
	prov.RenameDataSource("volcengine_veenedge_available_resources", dataSourceType(volcengineVeenedgeMod, "AvailableResources"),
		dataSourceType(volcengineVeenedgeMod, "getAvailableResources"), volcengineVeenedgeMod, volcengineVeenedgeMod, nil)
	prov.RenameDataSource("volcengine_veenedge_instances", dataSourceType(volcengineVeenedgeMod, "Instances"),
		dataSourceType(volcengineVeenedgeMod, "getInstances"), volcengineVeenedgeMod, volcengineVeenedgeMod, nil)
	prov.RenameDataSource("volcengine_veenedge_vpcs", dataSourceType(volcengineVeenedgeMod, "Vpcs"),
		dataSourceType(volcengineVeenedgeMod, "getVpcs"), volcengineVeenedgeMod, volcengineVeenedgeMod, nil)
	prov.RenameDataSource("volcengine_eip_addresses", dataSourceType(volcengineEipMod, "Addresses"),
		dataSourceType(volcengineEipMod, "getAddresses"), volcengineEipMod, volcengineEipMod, nil)
	prov.RenameDataSource("volcengine_vedb_mysql_endpoints", dataSourceType(volcengineVedbMysqlMod, "Endpoints"),
		dataSourceType(volcengineVedbMysqlMod, "getEndpoints"), volcengineVedbMysqlMod, volcengineVedbMysqlMod, nil)
	prov.RenameDataSource("volcengine_vedb_mysql_backups", dataSourceType(volcengineVedbMysqlMod, "Backups"),
		dataSourceType(volcengineVedbMysqlMod, "getBackups"), volcengineVedbMysqlMod, volcengineVedbMysqlMod, nil)
	prov.RenameDataSource("volcengine_vedb_mysql_allowlists", dataSourceType(volcengineVedbMysqlMod, "Allowlists"),
		dataSourceType(volcengineVedbMysqlMod, "getAllowlists"), volcengineVedbMysqlMod, volcengineVedbMysqlMod, nil)
	prov.RenameDataSource("volcengine_vedb_mysql_databases", dataSourceType(volcengineVedbMysqlMod, "Databases"),
		dataSourceType(volcengineVedbMysqlMod, "getDatabases"), volcengineVedbMysqlMod, volcengineVedbMysqlMod, nil)
	prov.RenameDataSource("volcengine_vedb_mysql_instances", dataSourceType(volcengineVedbMysqlMod, "Instances"),
		dataSourceType(volcengineVedbMysqlMod, "getInstances"), volcengineVedbMysqlMod, volcengineVedbMysqlMod, nil)
	prov.RenameDataSource("volcengine_vedb_mysql_accounts", dataSourceType(volcengineVedbMysqlMod, "Accounts"),
		dataSourceType(volcengineVedbMysqlMod, "getAccounts"), volcengineVedbMysqlMod, volcengineVedbMysqlMod, nil)
	prov.RenameDataSource("volcengine_cloud_identity_permission_sets", dataSourceType(volcengineCloudIdentityMod, "PermissionSets"),
		dataSourceType(volcengineCloudIdentityMod, "getPermissionSets"), volcengineCloudIdentityMod, volcengineCloudIdentityMod, nil)
	prov.RenameDataSource("volcengine_cloud_identity_permission_set_provisionings", dataSourceType(volcengineCloudIdentityMod, "PermissionSetProvisionings"),
		dataSourceType(volcengineCloudIdentityMod, "getPermissionSetProvisionings"), volcengineCloudIdentityMod, volcengineCloudIdentityMod, nil)
	prov.RenameDataSource("volcengine_cloud_identity_user_provisionings", dataSourceType(volcengineCloudIdentityMod, "UserProvisionings"),
		dataSourceType(volcengineCloudIdentityMod, "getUserProvisionings"), volcengineCloudIdentityMod, volcengineCloudIdentityMod, nil)
	prov.RenameDataSource("volcengine_cloud_identity_users", dataSourceType(volcengineCloudIdentityMod, "Users"),
		dataSourceType(volcengineCloudIdentityMod, "getUsers"), volcengineCloudIdentityMod, volcengineCloudIdentityMod, nil)
	prov.RenameDataSource("volcengine_cloud_identity_groups", dataSourceType(volcengineCloudIdentityMod, "Groups"),
		dataSourceType(volcengineCloudIdentityMod, "getGroups"), volcengineCloudIdentityMod, volcengineCloudIdentityMod, nil)
	prov.RenameDataSource("volcengine_cloud_identity_permission_set_assignments", dataSourceType(volcengineCloudIdentityMod, "PermissionSetAssignments"),
		dataSourceType(volcengineCloudIdentityMod, "getPermissionSetAssignments"), volcengineCloudIdentityMod, volcengineCloudIdentityMod, nil)
	prov.RenameDataSource("volcengine_cloudfs_quotas", dataSourceType(volcengineCloudfsMod, "Quotas"),
		dataSourceType(volcengineCloudfsMod, "getQuotas"), volcengineCloudfsMod, volcengineCloudfsMod, nil)
	prov.RenameDataSource("volcengine_cloudfs_ns_quotas", dataSourceType(volcengineCloudfsMod, "NsQuotas"),
		dataSourceType(volcengineCloudfsMod, "getNsQuotas"), volcengineCloudfsMod, volcengineCloudfsMod, nil)
	prov.RenameDataSource("volcengine_cloudfs_namespaces", dataSourceType(volcengineCloudfsMod, "Namespaces"),
		dataSourceType(volcengineCloudfsMod, "getNamespaces"), volcengineCloudfsMod, volcengineCloudfsMod, nil)
	prov.RenameDataSource("volcengine_cloudfs_file_systems", dataSourceType(volcengineCloudfsMod, "FileSystems"),
		dataSourceType(volcengineCloudfsMod, "getFileSystems"), volcengineCloudfsMod, volcengineCloudfsMod, nil)
	prov.RenameDataSource("volcengine_cloudfs_accesses", dataSourceType(volcengineCloudfsMod, "Accesses"),
		dataSourceType(volcengineCloudfsMod, "getAccesses"), volcengineCloudfsMod, volcengineCloudfsMod, nil)
	prov.RenameDataSource("volcengine_ebs_auto_snapshot_policies", dataSourceType(volcengineEbsMod, "AutoSnapshotPolicies"),
		dataSourceType(volcengineEbsMod, "getAutoSnapshotPolicies"), volcengineEbsMod, volcengineEbsMod, nil)
	prov.RenameDataSource("volcengine_ebs_max_extra_performances", dataSourceType(volcengineEbsMod, "MaxExtraPerformances"),
		dataSourceType(volcengineEbsMod, "getMaxExtraPerformances"), volcengineEbsMod, volcengineEbsMod, nil)
	prov.RenameDataSource("volcengine_ebs_snapshot_groups", dataSourceType(volcengineEbsMod, "SnapshotGroups"),
		dataSourceType(volcengineEbsMod, "getSnapshotGroups"), volcengineEbsMod, volcengineEbsMod, nil)
	prov.RenameDataSource("volcengine_volumes", dataSourceType(volcengineEbsMod, "Volumes"),
		dataSourceType(volcengineEbsMod, "getVolumes"), volcengineEbsMod, volcengineEbsMod, nil)
	prov.RenameDataSource("volcengine_ebs_snapshots", dataSourceType(volcengineEbsMod, "Snapshots"),
		dataSourceType(volcengineEbsMod, "getSnapshots"), volcengineEbsMod, volcengineEbsMod, nil)
	prov.RenameDataSource("volcengine_redis_pitr_time_windows", dataSourceType(volcengineRedisMod, "PitrTimeWindows"),
		dataSourceType(volcengineRedisMod, "getPitrTimeWindows"), volcengineRedisMod, volcengineRedisMod, nil)
	prov.RenameDataSource("volcengine_redis_planned_events", dataSourceType(volcengineRedisMod, "PlannedEvents"),
		dataSourceType(volcengineRedisMod, "getPlannedEvents"), volcengineRedisMod, volcengineRedisMod, nil)
	prov.RenameDataSource("volcengine_redis_accounts", dataSourceType(volcengineRedisMod, "Accounts"),
		dataSourceType(volcengineRedisMod, "getAccounts"), volcengineRedisMod, volcengineRedisMod, nil)
	prov.RenameDataSource("volcengine_redis_backups", dataSourceType(volcengineRedisMod, "Backups"),
		dataSourceType(volcengineRedisMod, "getBackups"), volcengineRedisMod, volcengineRedisMod, nil)
	prov.RenameDataSource("volcengine_redis_zones", dataSourceType(volcengineRedisMod, "Zones"),
		dataSourceType(volcengineRedisMod, "getZones"), volcengineRedisMod, volcengineRedisMod, nil)
	prov.RenameDataSource("volcengine_redis_instance_specs", dataSourceType(volcengineRedisMod, "InstanceSpecs"),
		dataSourceType(volcengineRedisMod, "getInstanceSpecs"), volcengineRedisMod, volcengineRedisMod, nil)
	prov.RenameDataSource("volcengine_redis_parameter_groups", dataSourceType(volcengineRedisMod, "ParameterGroups"),
		dataSourceType(volcengineRedisMod, "getParameterGroups"), volcengineRedisMod, volcengineRedisMod, nil)
	prov.RenameDataSource("volcengine_redis_big_keys", dataSourceType(volcengineRedisMod, "BigKeys"),
		dataSourceType(volcengineRedisMod, "getBigKeys"), volcengineRedisMod, volcengineRedisMod, nil)
	prov.RenameDataSource("volcengine_redis_regions", dataSourceType(volcengineRedisMod, "Regions"),
		dataSourceType(volcengineRedisMod, "getRegions"), volcengineRedisMod, volcengineRedisMod, nil)
	prov.RenameDataSource("volcengine_redis_allow_lists", dataSourceType(volcengineRedisMod, "AllowLists"),
		dataSourceType(volcengineRedisMod, "getAllowLists"), volcengineRedisMod, volcengineRedisMod, nil)
	prov.RenameDataSource("volcengine_redis_instances", dataSourceType(volcengineRedisMod, "Instances"),
		dataSourceType(volcengineRedisMod, "getInstances"), volcengineRedisMod, volcengineRedisMod, nil)
	prov.RenameDataSource("volcengine_rds_postgresql_schemas", dataSourceType(volcengineRdsPostgresqlMod, "Schemas"),
		dataSourceType(volcengineRdsPostgresqlMod, "getSchemas"), volcengineRdsPostgresqlMod, volcengineRdsPostgresqlMod, nil)
	prov.RenameDataSource("volcengine_rds_postgresql_databases", dataSourceType(volcengineRdsPostgresqlMod, "Databases"),
		dataSourceType(volcengineRdsPostgresqlMod, "getDatabases"), volcengineRdsPostgresqlMod, volcengineRdsPostgresqlMod, nil)
	prov.RenameDataSource("volcengine_rds_postgresql_accounts", dataSourceType(volcengineRdsPostgresqlMod, "Accounts"),
		dataSourceType(volcengineRdsPostgresqlMod, "getAccounts"), volcengineRdsPostgresqlMod, volcengineRdsPostgresqlMod, nil)
	prov.RenameDataSource("volcengine_rds_postgresql_allowlists", dataSourceType(volcengineRdsPostgresqlMod, "Allowlists"),
		dataSourceType(volcengineRdsPostgresqlMod, "getAllowlists"), volcengineRdsPostgresqlMod, volcengineRdsPostgresqlMod, nil)
	prov.RenameDataSource("volcengine_rds_postgresql_instances", dataSourceType(volcengineRdsPostgresqlMod, "Instances"),
		dataSourceType(volcengineRdsPostgresqlMod, "getInstances"), volcengineRdsPostgresqlMod, volcengineRdsPostgresqlMod, nil)
	prov.RenameDataSource("volcengine_escloud_instances", dataSourceType(volcengineEscloudMod, "Instances"),
		dataSourceType(volcengineEscloudMod, "getInstances"), volcengineEscloudMod, volcengineEscloudMod, nil)
	prov.RenameDataSource("volcengine_escloud_regions", dataSourceType(volcengineEscloudMod, "Regions"),
		dataSourceType(volcengineEscloudMod, "getRegions"), volcengineEscloudMod, volcengineEscloudMod, nil)
	prov.RenameDataSource("volcengine_escloud_zones", dataSourceType(volcengineEscloudMod, "Zones"),
		dataSourceType(volcengineEscloudMod, "getZones"), volcengineEscloudMod, volcengineEscloudMod, nil)
	prov.RenameDataSource("volcengine_cr_endpoints", dataSourceType(volcengineCrMod, "Endpoints"),
		dataSourceType(volcengineCrMod, "getEndpoints"), volcengineCrMod, volcengineCrMod, nil)
	prov.RenameDataSource("volcengine_cr_tags", dataSourceType(volcengineCrMod, "Tags"),
		dataSourceType(volcengineCrMod, "getTags"), volcengineCrMod, volcengineCrMod, nil)
	prov.RenameDataSource("volcengine_cr_registries", dataSourceType(volcengineCrMod, "Registries"),
		dataSourceType(volcengineCrMod, "getRegistries"), volcengineCrMod, volcengineCrMod, nil)
	prov.RenameDataSource("volcengine_cr_repositories", dataSourceType(volcengineCrMod, "Repositories"),
		dataSourceType(volcengineCrMod, "getRepositories"), volcengineCrMod, volcengineCrMod, nil)
	prov.RenameDataSource("volcengine_cr_authorization_tokens", dataSourceType(volcengineCrMod, "AuthorizationTokens"),
		dataSourceType(volcengineCrMod, "getAuthorizationTokens"), volcengineCrMod, volcengineCrMod, nil)
	prov.RenameDataSource("volcengine_cr_namespaces", dataSourceType(volcengineCrMod, "Namespaces"),
		dataSourceType(volcengineCrMod, "getNamespaces"), volcengineCrMod, volcengineCrMod, nil)
	prov.RenameDataSource("volcengine_cr_vpc_endpoints", dataSourceType(volcengineCrMod, "VpcEndpoints"),
		dataSourceType(volcengineCrMod, "getVpcEndpoints"), volcengineCrMod, volcengineCrMod, nil)
	prov.RenameDataSource("volcengine_vmp_notify_group_policies", dataSourceType(volcengineVmpMod, "NotifyGroupPolicies"),
		dataSourceType(volcengineVmpMod, "getNotifyGroupPolicies"), volcengineVmpMod, volcengineVmpMod, nil)
	prov.RenameDataSource("volcengine_vmp_contact_groups", dataSourceType(volcengineVmpMod, "ContactGroups"),
		dataSourceType(volcengineVmpMod, "getContactGroups"), volcengineVmpMod, volcengineVmpMod, nil)
	prov.RenameDataSource("volcengine_vmp_rule_files", dataSourceType(volcengineVmpMod, "RuleFiles"),
		dataSourceType(volcengineVmpMod, "getRuleFiles"), volcengineVmpMod, volcengineVmpMod, nil)
	prov.RenameDataSource("volcengine_vmp_notify_templates", dataSourceType(volcengineVmpMod, "NotifyTemplates"),
		dataSourceType(volcengineVmpMod, "getNotifyTemplates"), volcengineVmpMod, volcengineVmpMod, nil)
	prov.RenameDataSource("volcengine_vmp_rules", dataSourceType(volcengineVmpMod, "Rules"),
		dataSourceType(volcengineVmpMod, "getRules"), volcengineVmpMod, volcengineVmpMod, nil)
	prov.RenameDataSource("volcengine_vmp_alerts", dataSourceType(volcengineVmpMod, "Alerts"),
		dataSourceType(volcengineVmpMod, "getAlerts"), volcengineVmpMod, volcengineVmpMod, nil)
	prov.RenameDataSource("volcengine_vmp_alerting_rules", dataSourceType(volcengineVmpMod, "AlertingRules"),
		dataSourceType(volcengineVmpMod, "getAlertingRules"), volcengineVmpMod, volcengineVmpMod, nil)
	prov.RenameDataSource("volcengine_vmp_workspaces", dataSourceType(volcengineVmpMod, "Workspaces"),
		dataSourceType(volcengineVmpMod, "getWorkspaces"), volcengineVmpMod, volcengineVmpMod, nil)
	prov.RenameDataSource("volcengine_vmp_contacts", dataSourceType(volcengineVmpMod, "Contacts"),
		dataSourceType(volcengineVmpMod, "getContacts"), volcengineVmpMod, volcengineVmpMod, nil)
	prov.RenameDataSource("volcengine_vmp_notify_policies", dataSourceType(volcengineVmpMod, "NotifyPolicies"),
		dataSourceType(volcengineVmpMod, "getNotifyPolicies"), volcengineVmpMod, volcengineVmpMod, nil)
	prov.RenameDataSource("volcengine_vmp_instance_types", dataSourceType(volcengineVmpMod, "InstanceTypes"),
		dataSourceType(volcengineVmpMod, "getInstanceTypes"), volcengineVmpMod, volcengineVmpMod, nil)
	prov.RenameDataSource("volcengine_escloud_node_available_specs", dataSourceType(volcengineEscloudV2Mod, "EscloudNodeAvailableSpecs"),
		dataSourceType(volcengineEscloudV2Mod, "getEscloudNodeAvailableSpecs"), volcengineEscloudV2Mod, volcengineEscloudV2Mod, nil)
	prov.RenameDataSource("volcengine_escloud_zones_v2", dataSourceType(volcengineEscloudV2Mod, "EscloudZonesV2"),
		dataSourceType(volcengineEscloudV2Mod, "getEscloudZonesV2"), volcengineEscloudV2Mod, volcengineEscloudV2Mod, nil)
	prov.RenameDataSource("volcengine_escloud_instances_v2", dataSourceType(volcengineEscloudV2Mod, "EscloudInstancesV2"),
		dataSourceType(volcengineEscloudV2Mod, "getEscloudInstancesV2"), volcengineEscloudV2Mod, volcengineEscloudV2Mod, nil)
	prov.RenameDataSource("volcengine_tls_alarms", dataSourceType(volcengineTlsMod, "Alarms"),
		dataSourceType(volcengineTlsMod, "getAlarms"), volcengineTlsMod, volcengineTlsMod, nil)
	prov.RenameDataSource("volcengine_tls_projects", dataSourceType(volcengineTlsMod, "Projects"),
		dataSourceType(volcengineTlsMod, "getProjects"), volcengineTlsMod, volcengineTlsMod, nil)
	prov.RenameDataSource("volcengine_tls_topics", dataSourceType(volcengineTlsMod, "Topics"),
		dataSourceType(volcengineTlsMod, "getTopics"), volcengineTlsMod, volcengineTlsMod, nil)
	prov.RenameDataSource("volcengine_tls_rule_appliers", dataSourceType(volcengineTlsMod, "RuleAppliers"),
		dataSourceType(volcengineTlsMod, "getRuleAppliers"), volcengineTlsMod, volcengineTlsMod, nil)
	prov.RenameDataSource("volcengine_tls_etl_tasks", dataSourceType(volcengineTlsMod, "EtlTasks"),
		dataSourceType(volcengineTlsMod, "getEtlTasks"), volcengineTlsMod, volcengineTlsMod, nil)
	prov.RenameDataSource("volcengine_tls_rules", dataSourceType(volcengineTlsMod, "Rules"),
		dataSourceType(volcengineTlsMod, "getRules"), volcengineTlsMod, volcengineTlsMod, nil)
	prov.RenameDataSource("volcengine_tls_host_groups", dataSourceType(volcengineTlsMod, "HostGroups"),
		dataSourceType(volcengineTlsMod, "getHostGroups"), volcengineTlsMod, volcengineTlsMod, nil)
	prov.RenameDataSource("volcengine_tls_shippers", dataSourceType(volcengineTlsMod, "Shippers"),
		dataSourceType(volcengineTlsMod, "getShippers"), volcengineTlsMod, volcengineTlsMod, nil)
	prov.RenameDataSource("volcengine_tls_indexes", dataSourceType(volcengineTlsMod, "Indexes"),
		dataSourceType(volcengineTlsMod, "getIndexes"), volcengineTlsMod, volcengineTlsMod, nil)
	prov.RenameDataSource("volcengine_tls_consumer_groups", dataSourceType(volcengineTlsMod, "ConsumerGroups"),
		dataSourceType(volcengineTlsMod, "getConsumerGroups"), volcengineTlsMod, volcengineTlsMod, nil)
	prov.RenameDataSource("volcengine_tls_hosts", dataSourceType(volcengineTlsMod, "Hosts"),
		dataSourceType(volcengineTlsMod, "getHosts"), volcengineTlsMod, volcengineTlsMod, nil)
	prov.RenameDataSource("volcengine_tls_schedule_sql_tasks", dataSourceType(volcengineTlsMod, "ScheduleSqlTasks"),
		dataSourceType(volcengineTlsMod, "getScheduleSqlTasks"), volcengineTlsMod, volcengineTlsMod, nil)
	prov.RenameDataSource("volcengine_tls_alarm_notify_groups", dataSourceType(volcengineTlsMod, "AlarmNotifyGroups"),
		dataSourceType(volcengineTlsMod, "getAlarmNotifyGroups"), volcengineTlsMod, volcengineTlsMod, nil)
	prov.RenameDataSource("volcengine_tls_kafka_consumers", dataSourceType(volcengineTlsMod, "KafkaConsumers"),
		dataSourceType(volcengineTlsMod, "getKafkaConsumers"), volcengineTlsMod, volcengineTlsMod, nil)
	prov.RenameDataSource("volcengine_tls_import_tasks", dataSourceType(volcengineTlsMod, "ImportTasks"),
		dataSourceType(volcengineTlsMod, "getImportTasks"), volcengineTlsMod, volcengineTlsMod, nil)
	prov.RenameDataSource("volcengine_tls_shards", dataSourceType(volcengineTlsMod, "Shards"),
		dataSourceType(volcengineTlsMod, "getShards"), volcengineTlsMod, volcengineTlsMod, nil)
	prov.RenameDataSource("volcengine_bioos_clusters", dataSourceType(volcengineBioosMod, "Clusters"),
		dataSourceType(volcengineBioosMod, "getClusters"), volcengineBioosMod, volcengineBioosMod, nil)
	prov.RenameDataSource("volcengine_bioos_workspaces", dataSourceType(volcengineBioosMod, "Workspaces"),
		dataSourceType(volcengineBioosMod, "getWorkspaces"), volcengineBioosMod, volcengineBioosMod, nil)
	prov.RenameDataSource("volcengine_rds_mysql_tasks", dataSourceType(volcengineRdsMysqlMod, "Tasks"),
		dataSourceType(volcengineRdsMysqlMod, "getTasks"), volcengineRdsMysqlMod, volcengineRdsMysqlMod, nil)
	prov.RenameDataSource("volcengine_rds_mysql_instance_specs", dataSourceType(volcengineRdsMysqlMod, "InstanceSpecs"),
		dataSourceType(volcengineRdsMysqlMod, "getInstanceSpecs"), volcengineRdsMysqlMod, volcengineRdsMysqlMod, nil)
	prov.RenameDataSource("volcengine_rds_mysql_planned_events", dataSourceType(volcengineRdsMysqlMod, "PlannedEvents"),
		dataSourceType(volcengineRdsMysqlMod, "getPlannedEvents"), volcengineRdsMysqlMod, volcengineRdsMysqlMod, nil)
	prov.RenameDataSource("volcengine_rds_mysql_zones", dataSourceType(volcengineRdsMysqlMod, "Zones"),
		dataSourceType(volcengineRdsMysqlMod, "getZones"), volcengineRdsMysqlMod, volcengineRdsMysqlMod, nil)
	prov.RenameDataSource("volcengine_rds_mysql_accounts", dataSourceType(volcengineRdsMysqlMod, "Accounts"),
		dataSourceType(volcengineRdsMysqlMod, "getAccounts"), volcengineRdsMysqlMod, volcengineRdsMysqlMod, nil)
	prov.RenameDataSource("volcengine_rds_mysql_backups", dataSourceType(volcengineRdsMysqlMod, "Backups"),
		dataSourceType(volcengineRdsMysqlMod, "getBackups"), volcengineRdsMysqlMod, volcengineRdsMysqlMod, nil)
	prov.RenameDataSource("volcengine_rds_mysql_instances", dataSourceType(volcengineRdsMysqlMod, "Instances"),
		dataSourceType(volcengineRdsMysqlMod, "getInstances"), volcengineRdsMysqlMod, volcengineRdsMysqlMod, nil)
	prov.RenameDataSource("volcengine_rds_mysql_regions", dataSourceType(volcengineRdsMysqlMod, "Regions"),
		dataSourceType(volcengineRdsMysqlMod, "getRegions"), volcengineRdsMysqlMod, volcengineRdsMysqlMod, nil)
	prov.RenameDataSource("volcengine_rds_mysql_parameter_templates", dataSourceType(volcengineRdsMysqlMod, "ParameterTemplates"),
		dataSourceType(volcengineRdsMysqlMod, "getParameterTemplates"), volcengineRdsMysqlMod, volcengineRdsMysqlMod, nil)
	prov.RenameDataSource("volcengine_rds_mysql_allowlists", dataSourceType(volcengineRdsMysqlMod, "Allowlists"),
		dataSourceType(volcengineRdsMysqlMod, "getAllowlists"), volcengineRdsMysqlMod, volcengineRdsMysqlMod, nil)
	prov.RenameDataSource("volcengine_rds_mysql_endpoints", dataSourceType(volcengineRdsMysqlMod, "Endpoints"),
		dataSourceType(volcengineRdsMysqlMod, "getEndpoints"), volcengineRdsMysqlMod, volcengineRdsMysqlMod, nil)
	prov.RenameDataSource("volcengine_rds_mysql_account_table_column_infos", dataSourceType(volcengineRdsMysqlMod, "AccountTableColumnInfos"),
		dataSourceType(volcengineRdsMysqlMod, "getAccountTableColumnInfos"), volcengineRdsMysqlMod, volcengineRdsMysqlMod, nil)
	prov.RenameDataSource("volcengine_rds_mysql_databases", dataSourceType(volcengineRdsMysqlMod, "Databases"),
		dataSourceType(volcengineRdsMysqlMod, "getDatabases"), volcengineRdsMysqlMod, volcengineRdsMysqlMod, nil)
	prov.RenameDataSource("volcengine_rds_mssql_instances", dataSourceType(volcengineRdsMssqlMod, "Instances"),
		dataSourceType(volcengineRdsMssqlMod, "getInstances"), volcengineRdsMssqlMod, volcengineRdsMssqlMod, nil)
	prov.RenameDataSource("volcengine_rds_mssql_backups", dataSourceType(volcengineRdsMssqlMod, "Backups"),
		dataSourceType(volcengineRdsMssqlMod, "getBackups"), volcengineRdsMssqlMod, volcengineRdsMssqlMod, nil)
	prov.RenameDataSource("volcengine_rds_mssql_regions", dataSourceType(volcengineRdsMssqlMod, "Regions"),
		dataSourceType(volcengineRdsMssqlMod, "getRegions"), volcengineRdsMssqlMod, volcengineRdsMssqlMod, nil)
	prov.RenameDataSource("volcengine_rds_mssql_zones", dataSourceType(volcengineRdsMssqlMod, "Zones"),
		dataSourceType(volcengineRdsMssqlMod, "getZones"), volcengineRdsMssqlMod, volcengineRdsMssqlMod, nil)
	prov.RenameDataSource("volcengine_vpn_connections", dataSourceType(volcengineVpnMod, "Connections"),
		dataSourceType(volcengineVpnMod, "getConnections"), volcengineVpnMod, volcengineVpnMod, nil)
	prov.RenameDataSource("volcengine_vpn_gateway_routes", dataSourceType(volcengineVpnMod, "GatewayRoutes"),
		dataSourceType(volcengineVpnMod, "getGatewayRoutes"), volcengineVpnMod, volcengineVpnMod, nil)
	prov.RenameDataSource("volcengine_ssl_vpn_client_certs", dataSourceType(volcengineVpnMod, "SslVpnClientCerts"),
		dataSourceType(volcengineVpnMod, "getSslVpnClientCerts"), volcengineVpnMod, volcengineVpnMod, nil)
	prov.RenameDataSource("volcengine_vpn_gateways", dataSourceType(volcengineVpnMod, "Gateways"),
		dataSourceType(volcengineVpnMod, "getGateways"), volcengineVpnMod, volcengineVpnMod, nil)
	prov.RenameDataSource("volcengine_ssl_vpn_servers", dataSourceType(volcengineVpnMod, "SslVpnServers"),
		dataSourceType(volcengineVpnMod, "getSslVpnServers"), volcengineVpnMod, volcengineVpnMod, nil)
	prov.RenameDataSource("volcengine_customer_gateways", dataSourceType(volcengineVpnMod, "CustomerGateways"),
		dataSourceType(volcengineVpnMod, "getCustomerGateways"), volcengineVpnMod, volcengineVpnMod, nil)
	prov.RenameDataSource("volcengine_cen_route_entries", dataSourceType(volcengineCenMod, "RouteEntries"),
		dataSourceType(volcengineCenMod, "getRouteEntries"), volcengineCenMod, volcengineCenMod, nil)
	prov.RenameDataSource("volcengine_cen_bandwidth_packages", dataSourceType(volcengineCenMod, "BandwidthPackages"),
		dataSourceType(volcengineCenMod, "getBandwidthPackages"), volcengineCenMod, volcengineCenMod, nil)
	prov.RenameDataSource("volcengine_cen_attach_instances", dataSourceType(volcengineCenMod, "AttachInstances"),
		dataSourceType(volcengineCenMod, "getAttachInstances"), volcengineCenMod, volcengineCenMod, nil)
	prov.RenameDataSource("volcengine_cen_service_route_entries", dataSourceType(volcengineCenMod, "ServiceRouteEntries"),
		dataSourceType(volcengineCenMod, "getServiceRouteEntries"), volcengineCenMod, volcengineCenMod, nil)
	prov.RenameDataSource("volcengine_cen_grant_instances", dataSourceType(volcengineCenMod, "GrantInstances"),
		dataSourceType(volcengineCenMod, "getGrantInstances"), volcengineCenMod, volcengineCenMod, nil)
	prov.RenameDataSource("volcengine_cens", dataSourceType(volcengineCenMod, "Cens"),
		dataSourceType(volcengineCenMod, "getCens"), volcengineCenMod, volcengineCenMod, nil)
	prov.RenameDataSource("volcengine_cen_inter_region_bandwidths", dataSourceType(volcengineCenMod, "InterRegionBandwidths"),
		dataSourceType(volcengineCenMod, "getInterRegionBandwidths"), volcengineCenMod, volcengineCenMod, nil)

	return prov
}
