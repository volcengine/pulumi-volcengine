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
	volcengineEbsMod               = "ebs"
	volcengineKafkaMod             = "kafka"
	volcengineMongodbMod           = "mongodb"
	volcengineOrganizationMod      = "organization"
	volcengineCdnMod               = "cdn"
	volcengineCloudIdentityMod     = "cloud_identity"
	volcengineCloudMonitorMod      = "cloud_monitor"
	volcengineCloudfsMod           = "cloudfs"
	volcengineRdsMssqlMod          = "rds_mssql"
	volcengineRdsV2Mod             = "rds_v2"
	volcengineVeenedgeMod          = "veenedge"
	volcengineTosMod               = "tos"
	volcengineBioosMod             = "bioos"
	volcengineCrMod                = "cr"
	volcenginePrivatelinkMod       = "privatelink"
	volcengineRdsMysqlMod          = "rds_mysql"
	volcengineBandwidthPackageMod  = "bandwidth_package"
	volcengineFinancialRelationMod = "financial_relation"
	volcengineIamMod               = "iam"
	volcengineNasMod               = "nas"
	volcengineNatMod               = "nat"
	volcengineTlsMod               = "tls"
	volcengineClbMod               = "clb"
	volcengineDirectConnectMod     = "direct_connect"
	volcengineEcsMod               = "ecs"
	volcengineEscloudMod           = "escloud"
	volcengineCenMod               = "cen"
	volcengineRdsPostgresqlMod     = "rds_postgresql"
	volcengineAlbMod               = "alb"
	volcengineAutoscalingMod       = "autoscaling"
	volcengineVkeMod               = "vke"
	volcengineVpnMod               = "vpn"
	volcengineVpcMod               = "vpc"
	volcengineEipMod               = "eip"
	volcengineRdsMod               = "rds"
	volcengineRedisMod             = "redis"
	volcengineTransitRouterMod     = "transit_router"
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
			"volcengine_cr_tag":                                           {Tok: resourceType(volcengineCrMod, "Tag")},
			"volcengine_cr_endpoint":                                      {Tok: resourceType(volcengineCrMod, "Endpoint")},
			"volcengine_cr_vpc_endpoint":                                  {Tok: resourceType(volcengineCrMod, "VpcEndpoint")},
			"volcengine_cr_repository":                                    {Tok: resourceType(volcengineCrMod, "Repository")},
			"volcengine_cr_registry":                                      {Tok: resourceType(volcengineCrMod, "Registry")},
			"volcengine_cr_registry_state":                                {Tok: resourceType(volcengineCrMod, "State")},
			"volcengine_cr_namespace":                                     {Tok: resourceType(volcengineCrMod, "Namespace")},
			"volcengine_ssl_vpn_server":                                   {Tok: resourceType(volcengineVpnMod, "SslVpnServer")},
			"volcengine_vpn_gateway_route":                                {Tok: resourceType(volcengineVpnMod, "GatewayRoute")},
			"volcengine_vpn_connection":                                   {Tok: resourceType(volcengineVpnMod, "Connection")},
			"volcengine_customer_gateway":                                 {Tok: resourceType(volcengineVpnMod, "CustomerGateway")},
			"volcengine_ssl_vpn_client_cert":                              {Tok: resourceType(volcengineVpnMod, "SslVpnClientCert")},
			"volcengine_vpn_gateway":                                      {Tok: resourceType(volcengineVpnMod, "Gateway")},
			"volcengine_rds_instance_v2":                                  {Tok: resourceType(volcengineRdsV2Mod, "RdsInstanceV2")},
			"volcengine_cloud_identity_user_attachment":                   {Tok: resourceType(volcengineCloudIdentityMod, "UserAttachment")},
			"volcengine_cloud_identity_user_provisioning":                 {Tok: resourceType(volcengineCloudIdentityMod, "UserProvisioning")},
			"volcengine_cloud_identity_user":                              {Tok: resourceType(volcengineCloudIdentityMod, "User")},
			"volcengine_cloud_identity_permission_set_provisioning":       {Tok: resourceType(volcengineCloudIdentityMod, "PermissionSetProvisioning")},
			"volcengine_cloud_identity_permission_set_assignment":         {Tok: resourceType(volcengineCloudIdentityMod, "PermissionSetAssignment")},
			"volcengine_cloud_identity_permission_set":                    {Tok: resourceType(volcengineCloudIdentityMod, "PermissionSet")},
			"volcengine_cloud_identity_group":                             {Tok: resourceType(volcengineCloudIdentityMod, "Group")},
			"volcengine_cloudfs_access":                                   {Tok: resourceType(volcengineCloudfsMod, "Access")},
			"volcengine_cloudfs_namespace":                                {Tok: resourceType(volcengineCloudfsMod, "Namespace")},
			"volcengine_cloudfs_file_system":                              {Tok: resourceType(volcengineCloudfsMod, "FileSystem")},
			"volcengine_veenedge_cloud_server":                            {Tok: resourceType(volcengineVeenedgeMod, "CloudServer")},
			"volcengine_veenedge_vpc":                                     {Tok: resourceType(volcengineVeenedgeMod, "Vpc")},
			"volcengine_veenedge_instance":                                {Tok: resourceType(volcengineVeenedgeMod, "Instance")},
			"volcengine_cloud_monitor_contact":                            {Tok: resourceType(volcengineCloudMonitorMod, "Contact")},
			"volcengine_cloud_monitor_rule":                               {Tok: resourceType(volcengineCloudMonitorMod, "Rule")},
			"volcengine_cloud_monitor_contact_group":                      {Tok: resourceType(volcengineCloudMonitorMod, "ContactGroup")},
			"volcengine_cloud_monitor_event_rule":                         {Tok: resourceType(volcengineCloudMonitorMod, "EventRule")},
			"volcengine_network_acl":                                      {Tok: resourceType(volcengineVpcMod, "NetworkAcl")},
			"volcengine_route_table":                                      {Tok: resourceType(volcengineVpcMod, "RouteTable")},
			"volcengine_security_group":                                   {Tok: resourceType(volcengineVpcMod, "SecurityGroup")},
			"volcengine_ha_vip_associate":                                 {Tok: resourceType(volcengineVpcMod, "HaVipAssociate")},
			"volcengine_vpc":                                              {Tok: resourceType(volcengineVpcMod, "Vpc")},
			"volcengine_route_table_associate":                            {Tok: resourceType(volcengineVpcMod, "RouteTableAssociate")},
			"volcengine_ha_vip":                                           {Tok: resourceType(volcengineVpcMod, "HaVip")},
			"volcengine_subnet":                                           {Tok: resourceType(volcengineVpcMod, "Subnet")},
			"volcengine_vpc_ipv6_gateway":                                 {Tok: resourceType(volcengineVpcMod, "Ipv6Gateway")},
			"volcengine_network_interface_attach":                         {Tok: resourceType(volcengineVpcMod, "NetworkInterfaceAttach")},
			"volcengine_security_group_rule":                              {Tok: resourceType(volcengineVpcMod, "SecurityGroupRule")},
			"volcengine_vpc_prefix_list":                                  {Tok: resourceType(volcengineVpcMod, "PrefixList")},
			"volcengine_network_interface":                                {Tok: resourceType(volcengineVpcMod, "NetworkInterface")},
			"volcengine_vpc_ipv6_address_bandwidth":                       {Tok: resourceType(volcengineVpcMod, "Ipv6AddressBandwidth")},
			"volcengine_network_acl_associate":                            {Tok: resourceType(volcengineVpcMod, "NetworkAclAssociate")},
			"volcengine_route_entry":                                      {Tok: resourceType(volcengineVpcMod, "RouteEntry")},
			"volcengine_tls_rule":                                         {Tok: resourceType(volcengineTlsMod, "Rule")},
			"volcengine_tls_alarm":                                        {Tok: resourceType(volcengineTlsMod, "Alarm")},
			"volcengine_tls_alarm_notify_group":                           {Tok: resourceType(volcengineTlsMod, "AlarmNotifyGroup")},
			"volcengine_tls_project":                                      {Tok: resourceType(volcengineTlsMod, "Project")},
			"volcengine_tls_topic":                                        {Tok: resourceType(volcengineTlsMod, "Topic")},
			"volcengine_tls_host":                                         {Tok: resourceType(volcengineTlsMod, "Host")},
			"volcengine_tls_index":                                        {Tok: resourceType(volcengineTlsMod, "Index")},
			"volcengine_tls_host_group":                                   {Tok: resourceType(volcengineTlsMod, "HostGroup")},
			"volcengine_tls_kafka_consumer":                               {Tok: resourceType(volcengineTlsMod, "KafkaConsumer")},
			"volcengine_tls_rule_applier":                                 {Tok: resourceType(volcengineTlsMod, "RuleApplier")},
			"volcengine_cdn_shared_config":                                {Tok: resourceType(volcengineCdnMod, "SharedConfig")},
			"volcengine_cdn_certificate":                                  {Tok: resourceType(volcengineCdnMod, "CdnCertificate")},
			"volcengine_cdn_domain":                                       {Tok: resourceType(volcengineCdnMod, "CdnDomain")},
			"volcengine_escloud_instance":                                 {Tok: resourceType(volcengineEscloudMod, "Instance")},
			"volcengine_organization_unit":                                {Tok: resourceType(volcengineOrganizationMod, "Unit")},
			"volcengine_organization_service_control_policy_enabler":      {Tok: resourceType(volcengineOrganizationMod, "ServiceControlPolicyEnabler")},
			"volcengine_organization_account":                             {Tok: resourceType(volcengineOrganizationMod, "Account")},
			"volcengine_organization_service_control_policy_attachment":   {Tok: resourceType(volcengineOrganizationMod, "ServiceControlPolicyAttachment")},
			"volcengine_organization_service_control_policy":              {Tok: resourceType(volcengineOrganizationMod, "ServiceControlPolicy")},
			"volcengine_organization":                                     {Tok: resourceType(volcengineOrganizationMod, "Organization")},
			"volcengine_privatelink_vpc_endpoint":                         {Tok: resourceType(volcenginePrivatelinkMod, "VpcEndpoint")},
			"volcengine_privatelink_vpc_endpoint_connection":              {Tok: resourceType(volcenginePrivatelinkMod, "VpcEndpointConnection")},
			"volcengine_privatelink_vpc_endpoint_service":                 {Tok: resourceType(volcenginePrivatelinkMod, "VpcEndpointService")},
			"volcengine_privatelink_security_group":                       {Tok: resourceType(volcenginePrivatelinkMod, "SecurityGroup")},
			"volcengine_privatelink_vpc_endpoint_service_resource":        {Tok: resourceType(volcenginePrivatelinkMod, "VpcEndpointServiceResource")},
			"volcengine_privatelink_vpc_endpoint_zone":                    {Tok: resourceType(volcenginePrivatelinkMod, "VpcEndpointZone")},
			"volcengine_privatelink_vpc_endpoint_service_permission":      {Tok: resourceType(volcenginePrivatelinkMod, "VpcEndpointServicePermission")},
			"volcengine_volume_attach":                                    {Tok: resourceType(volcengineEbsMod, "VolumeAttach")},
			"volcengine_volume":                                           {Tok: resourceType(volcengineEbsMod, "Volume")},
			"volcengine_rds_postgresql_database":                          {Tok: resourceType(volcengineRdsPostgresqlMod, "Database")},
			"volcengine_rds_postgresql_account":                           {Tok: resourceType(volcengineRdsPostgresqlMod, "Account")},
			"volcengine_rds_postgresql_schema":                            {Tok: resourceType(volcengineRdsPostgresqlMod, "Schema")},
			"volcengine_rds_postgresql_instance":                          {Tok: resourceType(volcengineRdsPostgresqlMod, "Instance")},
			"volcengine_rds_postgresql_instance_readonly_node":            {Tok: resourceType(volcengineRdsPostgresqlMod, "InstanceReadonlyNode")},
			"volcengine_mongodb_allow_list":                               {Tok: resourceType(volcengineMongodbMod, "MongoAllowList")},
			"volcengine_mongodb_ssl_state":                                {Tok: resourceType(volcengineMongodbMod, "SslState")},
			"volcengine_mongodb_instance_parameter":                       {Tok: resourceType(volcengineMongodbMod, "InstanceParameter")},
			"volcengine_mongodb_instance":                                 {Tok: resourceType(volcengineMongodbMod, "Instance")},
			"volcengine_mongodb_endpoint":                                 {Tok: resourceType(volcengineMongodbMod, "Endpoint")},
			"volcengine_mongodb_allow_list_associate":                     {Tok: resourceType(volcengineMongodbMod, "MongoAllowListAssociate")},
			"volcengine_redis_allow_list_associate":                       {Tok: resourceType(volcengineRedisMod, "AllowListAssociate")},
			"volcengine_redis_endpoint":                                   {Tok: resourceType(volcengineRedisMod, "Endpoint")},
			"volcengine_redis_account":                                    {Tok: resourceType(volcengineRedisMod, "Account")},
			"volcengine_redis_backup_restore":                             {Tok: resourceType(volcengineRedisMod, "BackupRestore")},
			"volcengine_redis_instance":                                   {Tok: resourceType(volcengineRedisMod, "Instance")},
			"volcengine_redis_allow_list":                                 {Tok: resourceType(volcengineRedisMod, "AllowList")},
			"volcengine_redis_instance_state":                             {Tok: resourceType(volcengineRedisMod, "State")},
			"volcengine_redis_continuous_backup":                          {Tok: resourceType(volcengineRedisMod, "ContinuousBackup")},
			"volcengine_redis_backup":                                     {Tok: resourceType(volcengineRedisMod, "Backup")},
			"volcengine_kafka_group":                                      {Tok: resourceType(volcengineKafkaMod, "Group")},
			"volcengine_kafka_sasl_user":                                  {Tok: resourceType(volcengineKafkaMod, "SaslUser")},
			"volcengine_kafka_public_address":                             {Tok: resourceType(volcengineKafkaMod, "PublicAddress")},
			"volcengine_kafka_instance":                                   {Tok: resourceType(volcengineKafkaMod, "Instance")},
			"volcengine_kafka_topic":                                      {Tok: resourceType(volcengineKafkaMod, "Topic")},
			"volcengine_tos_object":                                       {Tok: resourceType(volcengineTosMod, "BucketObject")},
			"volcengine_tos_bucket_policy":                                {Tok: resourceType(volcengineTosMod, "BucketPolicy")},
			"volcengine_tos_bucket":                                       {Tok: resourceType(volcengineTosMod, "Bucket")},
			"volcengine_cen_route_entry":                                  {Tok: resourceType(volcengineCenMod, "RouteEntry")},
			"volcengine_cen_bandwidth_package":                            {Tok: resourceType(volcengineCenMod, "BandwidthPackage")},
			"volcengine_cen":                                              {Tok: resourceType(volcengineCenMod, "Cen")},
			"volcengine_cen_bandwidth_package_associate":                  {Tok: resourceType(volcengineCenMod, "BandwidthPackageAssociate")},
			"volcengine_cen_inter_region_bandwidth":                       {Tok: resourceType(volcengineCenMod, "InterRegionBandwidth")},
			"volcengine_cen_service_route_entry":                          {Tok: resourceType(volcengineCenMod, "ServiceRouteEntry")},
			"volcengine_cen_attach_instance":                              {Tok: resourceType(volcengineCenMod, "AttachInstance")},
			"volcengine_cen_grant_instance":                               {Tok: resourceType(volcengineCenMod, "GrantInstance")},
			"volcengine_direct_connect_bgp_peer":                          {Tok: resourceType(volcengineDirectConnectMod, "BgpPeer")},
			"volcengine_direct_connect_gateway":                           {Tok: resourceType(volcengineDirectConnectMod, "Gateway")},
			"volcengine_direct_connect_connection":                        {Tok: resourceType(volcengineDirectConnectMod, "Connection")},
			"volcengine_direct_connect_gateway_route":                     {Tok: resourceType(volcengineDirectConnectMod, "GatewayRoute")},
			"volcengine_direct_connect_virtual_interface":                 {Tok: resourceType(volcengineDirectConnectMod, "VirtualInterface")},
			"volcengine_financial_relation":                               {Tok: resourceType(volcengineFinancialRelationMod, "FinancialRelation")},
			"volcengine_nat_gateway":                                      {Tok: resourceType(volcengineNatMod, "Gateway")},
			"volcengine_dnat_entry":                                       {Tok: resourceType(volcengineNatMod, "DnatEntry")},
			"volcengine_snat_entry":                                       {Tok: resourceType(volcengineNatMod, "SnatEntry")},
			"volcengine_transit_router_route_table":                       {Tok: resourceType(volcengineTransitRouterMod, "RouteTable")},
			"volcengine_transit_router_route_table_propagation":           {Tok: resourceType(volcengineTransitRouterMod, "RouteTablePropagation")},
			"volcengine_transit_router_vpn_attachment":                    {Tok: resourceType(volcengineTransitRouterMod, "VpnAttachment")},
			"volcengine_transit_router_route_entry":                       {Tok: resourceType(volcengineTransitRouterMod, "RouteEntry")},
			"volcengine_transit_router":                                   {Tok: resourceType(volcengineTransitRouterMod, "TransitRouter")},
			"volcengine_transit_router_bandwidth_package":                 {Tok: resourceType(volcengineTransitRouterMod, "BandwidthPackage")},
			"volcengine_transit_router_vpc_attachment":                    {Tok: resourceType(volcengineTransitRouterMod, "VpcAttachment")},
			"volcengine_transit_router_direct_connect_gateway_attachment": {Tok: resourceType(volcengineTransitRouterMod, "DirectConnectGatewayAttachment")},
			"volcengine_transit_router_peer_attachment":                   {Tok: resourceType(volcengineTransitRouterMod, "PeerAttachment")},
			"volcengine_transit_router_shared_transit_router_state":       {Tok: resourceType(volcengineTransitRouterMod, "SharedTransitRouterState")},
			"volcengine_transit_router_grant_rule":                        {Tok: resourceType(volcengineTransitRouterMod, "GrantRule")},
			"volcengine_transit_router_route_table_association":           {Tok: resourceType(volcengineTransitRouterMod, "RouteTableAssociation")},
			"volcengine_alb_listener":                                     {Tok: resourceType(volcengineAlbMod, "Listener")},
			"volcengine_alb_server_group_server":                          {Tok: resourceType(volcengineAlbMod, "ServerGroupServer")},
			"volcengine_alb_acl":                                          {Tok: resourceType(volcengineAlbMod, "Acl")},
			"volcengine_alb_customized_cfg":                               {Tok: resourceType(volcengineAlbMod, "CustomizedCfg")},
			"volcengine_alb_listener_domain_extension":                    {Tok: resourceType(volcengineAlbMod, "ListenerDomainExtension")},
			"volcengine_alb_server_group":                                 {Tok: resourceType(volcengineAlbMod, "ServerGroup")},
			"volcengine_alb_rule":                                         {Tok: resourceType(volcengineAlbMod, "Rule")},
			"volcengine_alb":                                              {Tok: resourceType(volcengineAlbMod, "Alb")},
			"volcengine_alb_certificate":                                  {Tok: resourceType(volcengineAlbMod, "Certificate")},
			"volcengine_alb_health_check_template":                        {Tok: resourceType(volcengineAlbMod, "HealthCheckTemplate")},
			"volcengine_alb_ca_certificate":                               {Tok: resourceType(volcengineAlbMod, "CACertificate")},
			"volcengine_rds_mssql_instance":                               {Tok: resourceType(volcengineRdsMssqlMod, "Instance")},
			"volcengine_rds_mssql_backup":                                 {Tok: resourceType(volcengineRdsMssqlMod, "Backup")},
			"volcengine_server_group_server":                              {Tok: resourceType(volcengineClbMod, "ServerGroupServer")},
			"volcengine_clb":                                              {Tok: resourceType(volcengineClbMod, "Clb")},
			"volcengine_server_group":                                     {Tok: resourceType(volcengineClbMod, "ServerGroup")},
			"volcengine_certificate":                                      {Tok: resourceType(volcengineClbMod, "Certificate")},
			"volcengine_listener":                                         {Tok: resourceType(volcengineClbMod, "Listener")},
			"volcengine_clb_rule":                                         {Tok: resourceType(volcengineClbMod, "Rule")},
			"volcengine_acl_entry":                                        {Tok: resourceType(volcengineClbMod, "AclEntry")},
			"volcengine_acl":                                              {Tok: resourceType(volcengineClbMod, "Acl")},
			"volcengine_bandwidth_package":                                {Tok: resourceType(volcengineBandwidthPackageMod, "BandwidthPackage")},
			"volcengine_bandwidth_package_attachment":                     {Tok: resourceType(volcengineBandwidthPackageMod, "Attachment")},
			"volcengine_bioos_cluster_bind":                               {Tok: resourceType(volcengineBioosMod, "ClusterBind")},
			"volcengine_bioos_cluster":                                    {Tok: resourceType(volcengineBioosMod, "Cluster")},
			"volcengine_bioos_workspace":                                  {Tok: resourceType(volcengineBioosMod, "Workspace")},
			"volcengine_scaling_lifecycle_hook":                           {Tok: resourceType(volcengineAutoscalingMod, "ScalingLifecycleHook")},
			"volcengine_scaling_configuration":                            {Tok: resourceType(volcengineAutoscalingMod, "ScalingConfiguration")},
			"volcengine_scaling_group_enabler":                            {Tok: resourceType(volcengineAutoscalingMod, "ScalingGroupEnabler")},
			"volcengine_scaling_configuration_attachment":                 {Tok: resourceType(volcengineAutoscalingMod, "ScalingConfigurationAttachment")},
			"volcengine_scaling_group":                                    {Tok: resourceType(volcengineAutoscalingMod, "ScalingGroup")},
			"volcengine_scaling_instance_attachment":                      {Tok: resourceType(volcengineAutoscalingMod, "ScalingInstanceAttachment")},
			"volcengine_scaling_policy":                                   {Tok: resourceType(volcengineAutoscalingMod, "ScalingPolicy")},
			"volcengine_eip_address":                                      {Tok: resourceType(volcengineEipMod, "Address")},
			"volcengine_eip_associate":                                    {Tok: resourceType(volcengineEipMod, "Associate")},
			"volcengine_ecs_instance":                                     {Tok: resourceType(volcengineEcsMod, "Instance")},
			"volcengine_ecs_deployment_set_associate":                     {Tok: resourceType(volcengineEcsMod, "DeploymentSetAssociate")},
			"volcengine_ecs_key_pair_associate":                           {Tok: resourceType(volcengineEcsMod, "KeyPairAssociate")},
			"volcengine_ecs_command":                                      {Tok: resourceType(volcengineEcsMod, "Command")},
			"volcengine_ecs_deployment_set":                               {Tok: resourceType(volcengineEcsMod, "DeploymentSet")},
			"volcengine_ecs_instance_state":                               {Tok: resourceType(volcengineEcsMod, "State")},
			"volcengine_ecs_key_pair":                                     {Tok: resourceType(volcengineEcsMod, "KeyPair")},
			"volcengine_iam_role_attachment":                              {Tok: resourceType(volcengineEcsMod, "IamRoleAttachment")},
			"volcengine_ecs_launch_template":                              {Tok: resourceType(volcengineEcsMod, "LaunchTemplate")},
			"volcengine_ecs_invocation":                                   {Tok: resourceType(volcengineEcsMod, "Invocation")},
			"volcengine_rds_mysql_instance":                               {Tok: resourceType(volcengineRdsMysqlMod, "Instance")},
			"volcengine_rds_mysql_instance_readonly_node":                 {Tok: resourceType(volcengineRdsMysqlMod, "InstanceReadonlyNode")},
			"volcengine_rds_mysql_account":                                {Tok: resourceType(volcengineRdsMysqlMod, "Account")},
			"volcengine_rds_mysql_allowlist":                              {Tok: resourceType(volcengineRdsMysqlMod, "Allowlist")},
			"volcengine_rds_mysql_database":                               {Tok: resourceType(volcengineRdsMysqlMod, "Database")},
			"volcengine_rds_mysql_allowlist_associate":                    {Tok: resourceType(volcengineRdsMysqlMod, "AllowlistAssociate")},
			"volcengine_iam_role":                                         {Tok: resourceType(volcengineIamMod, "Role")},
			"volcengine_iam_policy":                                       {Tok: resourceType(volcengineIamMod, "Policy")},
			"volcengine_iam_access_key":                                   {Tok: resourceType(volcengineIamMod, "AccessKey")},
			"volcengine_iam_user_policy_attachment":                       {Tok: resourceType(volcengineIamMod, "UserPolicyAttachment")},
			"volcengine_iam_user_group_attachment":                        {Tok: resourceType(volcengineIamMod, "UserGroupAttachment")},
			"volcengine_iam_saml_provider":                                {Tok: resourceType(volcengineIamMod, "SamlProvider")},
			"volcengine_iam_role_policy_attachment":                       {Tok: resourceType(volcengineIamMod, "RolePolicyAttachment")},
			"volcengine_iam_user":                                         {Tok: resourceType(volcengineIamMod, "User")},
			"volcengine_iam_login_profile":                                {Tok: resourceType(volcengineIamMod, "LoginProfile")},
			"volcengine_iam_user_group":                                   {Tok: resourceType(volcengineIamMod, "UserGroup")},
			"volcengine_iam_user_group_policy_attachment":                 {Tok: resourceType(volcengineIamMod, "UserGroupPolicyAttachment")},
			"volcengine_rds_parameter_template":                           {Tok: resourceType(volcengineRdsMod, "ParameterTemplate")},
			"volcengine_rds_ip_list":                                      {Tok: resourceType(volcengineRdsMod, "IpList")},
			"volcengine_rds_instance":                                     {Tok: resourceType(volcengineRdsMod, "Instance")},
			"volcengine_rds_account_privilege":                            {Tok: resourceType(volcengineRdsMod, "AccountPrivilege")},
			"volcengine_rds_account":                                      {Tok: resourceType(volcengineRdsMod, "Account")},
			"volcengine_rds_database":                                     {Tok: resourceType(volcengineRdsMod, "Database")},
			"volcengine_vke_node":                                         {Tok: resourceType(volcengineVkeMod, "Node")},
			"volcengine_vke_addon":                                        {Tok: resourceType(volcengineVkeMod, "Addon")},
			"volcengine_vke_default_node_pool":                            {Tok: resourceType(volcengineVkeMod, "DefaultNodePool")},
			"volcengine_vke_node_pool":                                    {Tok: resourceType(volcengineVkeMod, "NodePool")},
			"volcengine_vke_default_node_pool_batch_attach":               {Tok: resourceType(volcengineVkeMod, "DefaultNodePoolBatchAttach")},
			"volcengine_vke_cluster":                                      {Tok: resourceType(volcengineVkeMod, "Cluster")},
			"volcengine_vke_kubeconfig":                                   {Tok: resourceType(volcengineVkeMod, "Kubeconfig")},
			"volcengine_nas_mount_point":                                  {Tok: resourceType(volcengineNasMod, "MountPoint")},
			"volcengine_nas_snapshot":                                     {Tok: resourceType(volcengineNasMod, "Snapshot")},
			"volcengine_nas_permission_group":                             {Tok: resourceType(volcengineNasMod, "PermissionGroup")},
			"volcengine_nas_file_system":                                  {Tok: resourceType(volcengineNasMod, "FileSystem")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"volcengine_cr_tags":                                           {Tok: dataSourceType(volcengineCrMod, "Tags")},
			"volcengine_cr_endpoints":                                      {Tok: dataSourceType(volcengineCrMod, "Endpoints")},
			"volcengine_cr_vpc_endpoints":                                  {Tok: dataSourceType(volcengineCrMod, "VpcEndpoints")},
			"volcengine_cr_authorization_tokens":                           {Tok: dataSourceType(volcengineCrMod, "AuthorizationTokens")},
			"volcengine_cr_repositories":                                   {Tok: dataSourceType(volcengineCrMod, "Repositories")},
			"volcengine_cr_registries":                                     {Tok: dataSourceType(volcengineCrMod, "Registries")},
			"volcengine_cr_namespaces":                                     {Tok: dataSourceType(volcengineCrMod, "Namespaces")},
			"volcengine_ssl_vpn_servers":                                   {Tok: dataSourceType(volcengineVpnMod, "SslVpnServers")},
			"volcengine_vpn_gateway_routes":                                {Tok: dataSourceType(volcengineVpnMod, "GatewayRoutes")},
			"volcengine_vpn_connections":                                   {Tok: dataSourceType(volcengineVpnMod, "Connections")},
			"volcengine_customer_gateways":                                 {Tok: dataSourceType(volcengineVpnMod, "CustomerGateways")},
			"volcengine_ssl_vpn_client_certs":                              {Tok: dataSourceType(volcengineVpnMod, "SslVpnClientCerts")},
			"volcengine_vpn_gateways":                                      {Tok: dataSourceType(volcengineVpnMod, "Gateways")},
			"volcengine_rds_instances_v2":                                  {Tok: dataSourceType(volcengineRdsV2Mod, "RdsInstancesV2")},
			"volcengine_cloud_identity_user_provisionings":                 {Tok: dataSourceType(volcengineCloudIdentityMod, "UserProvisionings")},
			"volcengine_cloud_identity_users":                              {Tok: dataSourceType(volcengineCloudIdentityMod, "Users")},
			"volcengine_cloud_identity_permission_set_provisionings":       {Tok: dataSourceType(volcengineCloudIdentityMod, "PermissionSetProvisionings")},
			"volcengine_cloud_identity_permission_set_assignments":         {Tok: dataSourceType(volcengineCloudIdentityMod, "PermissionSetAssignments")},
			"volcengine_cloud_identity_permission_sets":                    {Tok: dataSourceType(volcengineCloudIdentityMod, "PermissionSets")},
			"volcengine_cloud_identity_groups":                             {Tok: dataSourceType(volcengineCloudIdentityMod, "Groups")},
			"volcengine_cloudfs_accesses":                                  {Tok: dataSourceType(volcengineCloudfsMod, "Accesses")},
			"volcengine_cloudfs_namespaces":                                {Tok: dataSourceType(volcengineCloudfsMod, "Namespaces")},
			"volcengine_cloudfs_file_systems":                              {Tok: dataSourceType(volcengineCloudfsMod, "FileSystems")},
			"volcengine_cloudfs_quotas":                                    {Tok: dataSourceType(volcengineCloudfsMod, "Quotas")},
			"volcengine_cloudfs_ns_quotas":                                 {Tok: dataSourceType(volcengineCloudfsMod, "NsQuotas")},
			"volcengine_veenedge_available_resources":                      {Tok: dataSourceType(volcengineVeenedgeMod, "AvailableResources")},
			"volcengine_veenedge_cloud_servers":                            {Tok: dataSourceType(volcengineVeenedgeMod, "CloudServers")},
			"volcengine_veenedge_vpcs":                                     {Tok: dataSourceType(volcengineVeenedgeMod, "Vpcs")},
			"volcengine_veenedge_instances":                                {Tok: dataSourceType(volcengineVeenedgeMod, "Instances")},
			"volcengine_veenedge_instance_types":                           {Tok: dataSourceType(volcengineVeenedgeMod, "InstanceTypes")},
			"volcengine_cloud_monitor_contacts":                            {Tok: dataSourceType(volcengineCloudMonitorMod, "Contacts")},
			"volcengine_cloud_monitor_rules":                               {Tok: dataSourceType(volcengineCloudMonitorMod, "Rules")},
			"volcengine_cloud_monitor_contact_groups":                      {Tok: dataSourceType(volcengineCloudMonitorMod, "ContactGroups")},
			"volcengine_cloud_monitor_event_rules":                         {Tok: dataSourceType(volcengineCloudMonitorMod, "EventRules")},
			"volcengine_network_acls":                                      {Tok: dataSourceType(volcengineVpcMod, "NetworkAcls")},
			"volcengine_route_tables":                                      {Tok: dataSourceType(volcengineVpcMod, "RouteTables")},
			"volcengine_security_groups":                                   {Tok: dataSourceType(volcengineVpcMod, "SecurityGroups")},
			"volcengine_vpc_ipv6_addresses":                                {Tok: dataSourceType(volcengineVpcMod, "Ipv6Addresses")},
			"volcengine_vpcs":                                              {Tok: dataSourceType(volcengineVpcMod, "Vpcs")},
			"volcengine_ha_vips":                                           {Tok: dataSourceType(volcengineVpcMod, "HaVips")},
			"volcengine_subnets":                                           {Tok: dataSourceType(volcengineVpcMod, "Subnets")},
			"volcengine_vpc_ipv6_gateways":                                 {Tok: dataSourceType(volcengineVpcMod, "Ipv6Gateways")},
			"volcengine_security_group_rules":                              {Tok: dataSourceType(volcengineVpcMod, "SecurityGroupRules")},
			"volcengine_vpc_prefix_lists":                                  {Tok: dataSourceType(volcengineVpcMod, "PrefixLists")},
			"volcengine_network_interfaces":                                {Tok: dataSourceType(volcengineVpcMod, "NetworkInterfaces")},
			"volcengine_vpc_ipv6_address_bandwidths":                       {Tok: dataSourceType(volcengineVpcMod, "Ipv6AddressBandwidths")},
			"volcengine_route_entries":                                     {Tok: dataSourceType(volcengineVpcMod, "RouteEntries")},
			"volcengine_tls_rules":                                         {Tok: dataSourceType(volcengineTlsMod, "Rules")},
			"volcengine_tls_alarms":                                        {Tok: dataSourceType(volcengineTlsMod, "Alarms")},
			"volcengine_tls_alarm_notify_groups":                           {Tok: dataSourceType(volcengineTlsMod, "AlarmNotifyGroups")},
			"volcengine_tls_projects":                                      {Tok: dataSourceType(volcengineTlsMod, "Projects")},
			"volcengine_tls_topics":                                        {Tok: dataSourceType(volcengineTlsMod, "Topics")},
			"volcengine_tls_hosts":                                         {Tok: dataSourceType(volcengineTlsMod, "Hosts")},
			"volcengine_tls_indexes":                                       {Tok: dataSourceType(volcengineTlsMod, "Indexes")},
			"volcengine_tls_host_groups":                                   {Tok: dataSourceType(volcengineTlsMod, "HostGroups")},
			"volcengine_tls_kafka_consumers":                               {Tok: dataSourceType(volcengineTlsMod, "KafkaConsumers")},
			"volcengine_tls_rule_appliers":                                 {Tok: dataSourceType(volcengineTlsMod, "RuleAppliers")},
			"volcengine_tls_shards":                                        {Tok: dataSourceType(volcengineTlsMod, "Shards")},
			"volcengine_cdn_configs":                                       {Tok: dataSourceType(volcengineCdnMod, "Configs")},
			"volcengine_cdn_shared_configs":                                {Tok: dataSourceType(volcengineCdnMod, "SharedConfigs")},
			"volcengine_cdn_certificates":                                  {Tok: dataSourceType(volcengineCdnMod, "Certificates")},
			"volcengine_cdn_domains":                                       {Tok: dataSourceType(volcengineCdnMod, "Domains")},
			"volcengine_escloud_zones":                                     {Tok: dataSourceType(volcengineEscloudMod, "Zones")},
			"volcengine_escloud_instances":                                 {Tok: dataSourceType(volcengineEscloudMod, "Instances")},
			"volcengine_escloud_regions":                                   {Tok: dataSourceType(volcengineEscloudMod, "Regions")},
			"volcengine_organization_units":                                {Tok: dataSourceType(volcengineOrganizationMod, "Units")},
			"volcengine_organization_accounts":                             {Tok: dataSourceType(volcengineOrganizationMod, "Accounts")},
			"volcengine_organization_service_control_policies":             {Tok: dataSourceType(volcengineOrganizationMod, "ServiceControlPolicies")},
			"volcengine_organizations":                                     {Tok: dataSourceType(volcengineOrganizationMod, "Organizations")},
			"volcengine_privatelink_vpc_endpoints":                         {Tok: dataSourceType(volcenginePrivatelinkMod, "VpcEndpoints")},
			"volcengine_privatelink_vpc_endpoint_connections":              {Tok: dataSourceType(volcenginePrivatelinkMod, "VpcEndpointConnections")},
			"volcengine_privatelink_vpc_endpoint_services":                 {Tok: dataSourceType(volcenginePrivatelinkMod, "VpcEndpointServices")},
			"volcengine_privatelink_vpc_endpoint_zones":                    {Tok: dataSourceType(volcenginePrivatelinkMod, "VpcEndpointZones")},
			"volcengine_privatelink_vpc_endpoint_service_permissions":      {Tok: dataSourceType(volcenginePrivatelinkMod, "VpcEndpointServicePermissions")},
			"volcengine_volumes":                                           {Tok: dataSourceType(volcengineEbsMod, "Volumes")},
			"volcengine_rds_postgresql_databases":                          {Tok: dataSourceType(volcengineRdsPostgresqlMod, "Databases")},
			"volcengine_rds_postgresql_accounts":                           {Tok: dataSourceType(volcengineRdsPostgresqlMod, "Accounts")},
			"volcengine_rds_postgresql_schemas":                            {Tok: dataSourceType(volcengineRdsPostgresqlMod, "Schemas")},
			"volcengine_rds_postgresql_instances":                          {Tok: dataSourceType(volcengineRdsPostgresqlMod, "Instances")},
			"volcengine_mongodb_allow_lists":                               {Tok: dataSourceType(volcengineMongodbMod, "MongoAllowLists")},
			"volcengine_mongodb_ssl_states":                                {Tok: dataSourceType(volcengineMongodbMod, "SslStates")},
			"volcengine_mongodb_instance_parameters":                       {Tok: dataSourceType(volcengineMongodbMod, "InstanceParameters")},
			"volcengine_mongodb_zones":                                     {Tok: dataSourceType(volcengineMongodbMod, "Zones")},
			"volcengine_mongodb_regions":                                   {Tok: dataSourceType(volcengineMongodbMod, "Regions")},
			"volcengine_mongodb_accounts":                                  {Tok: dataSourceType(volcengineMongodbMod, "Accounts")},
			"volcengine_mongodb_instances":                                 {Tok: dataSourceType(volcengineMongodbMod, "Instances")},
			"volcengine_mongodb_endpoints":                                 {Tok: dataSourceType(volcengineMongodbMod, "Endpoints")},
			"volcengine_mongodb_instance_parameter_logs":                   {Tok: dataSourceType(volcengineMongodbMod, "InstanceParameterLogs")},
			"volcengine_mongodb_specs":                                     {Tok: dataSourceType(volcengineMongodbMod, "Specs")},
			"volcengine_redis_accounts":                                    {Tok: dataSourceType(volcengineRedisMod, "Accounts")},
			"volcengine_redis_pitr_time_windows":                           {Tok: dataSourceType(volcengineRedisMod, "PitrTimeWindows")},
			"volcengine_redis_zones":                                       {Tok: dataSourceType(volcengineRedisMod, "Zones")},
			"volcengine_redis_regions":                                     {Tok: dataSourceType(volcengineRedisMod, "Regions")},
			"volcengine_redis_instances":                                   {Tok: dataSourceType(volcengineRedisMod, "Instances")},
			"volcengine_redis_allow_lists":                                 {Tok: dataSourceType(volcengineRedisMod, "AllowLists")},
			"volcengine_redis_backups":                                     {Tok: dataSourceType(volcengineRedisMod, "Backups")},
			"volcengine_kafka_groups":                                      {Tok: dataSourceType(volcengineKafkaMod, "Groups")},
			"volcengine_kafka_regions":                                     {Tok: dataSourceType(volcengineKafkaMod, "Regions")},
			"volcengine_kafka_consumed_partitions":                         {Tok: dataSourceType(volcengineKafkaMod, "ConsumedPartitions")},
			"volcengine_kafka_sasl_users":                                  {Tok: dataSourceType(volcengineKafkaMod, "SaslUsers")},
			"volcengine_kafka_topic_partitions":                            {Tok: dataSourceType(volcengineKafkaMod, "TopicPartitions")},
			"volcengine_kafka_zones":                                       {Tok: dataSourceType(volcengineKafkaMod, "Zones")},
			"volcengine_kafka_consumed_topics":                             {Tok: dataSourceType(volcengineKafkaMod, "ConsumedTopics")},
			"volcengine_kafka_instances":                                   {Tok: dataSourceType(volcengineKafkaMod, "Instances")},
			"volcengine_kafka_topics":                                      {Tok: dataSourceType(volcengineKafkaMod, "Topics")},
			"volcengine_tos_objects":                                       {Tok: dataSourceType(volcengineTosMod, "BucketObjects")},
			"volcengine_tos_buckets":                                       {Tok: dataSourceType(volcengineTosMod, "Buckets")},
			"volcengine_cen_route_entries":                                 {Tok: dataSourceType(volcengineCenMod, "RouteEntries")},
			"volcengine_cen_bandwidth_packages":                            {Tok: dataSourceType(volcengineCenMod, "BandwidthPackages")},
			"volcengine_cens":                                              {Tok: dataSourceType(volcengineCenMod, "Cens")},
			"volcengine_cen_inter_region_bandwidths":                       {Tok: dataSourceType(volcengineCenMod, "InterRegionBandwidths")},
			"volcengine_cen_service_route_entries":                         {Tok: dataSourceType(volcengineCenMod, "ServiceRouteEntries")},
			"volcengine_cen_attach_instances":                              {Tok: dataSourceType(volcengineCenMod, "AttachInstances")},
			"volcengine_direct_connect_bgp_peers":                          {Tok: dataSourceType(volcengineDirectConnectMod, "BgpPeers")},
			"volcengine_direct_connect_gateways":                           {Tok: dataSourceType(volcengineDirectConnectMod, "Gateways")},
			"volcengine_direct_connect_connections":                        {Tok: dataSourceType(volcengineDirectConnectMod, "Connections")},
			"volcengine_direct_connect_gateway_routes":                     {Tok: dataSourceType(volcengineDirectConnectMod, "GatewayRoutes")},
			"volcengine_direct_connect_virtual_interfaces":                 {Tok: dataSourceType(volcengineDirectConnectMod, "VirtualInterfaces")},
			"volcengine_financial_relations":                               {Tok: dataSourceType(volcengineFinancialRelationMod, "FinancialRelations")},
			"volcengine_nat_gateways":                                      {Tok: dataSourceType(volcengineNatMod, "Gateways")},
			"volcengine_dnat_entries":                                      {Tok: dataSourceType(volcengineNatMod, "DnatEntries")},
			"volcengine_snat_entries":                                      {Tok: dataSourceType(volcengineNatMod, "SnatEntries")},
			"volcengine_transit_router_route_tables":                       {Tok: dataSourceType(volcengineTransitRouterMod, "RouteTables")},
			"volcengine_transit_router_route_table_propagations":           {Tok: dataSourceType(volcengineTransitRouterMod, "RouteTablePropagations")},
			"volcengine_transit_router_vpn_attachments":                    {Tok: dataSourceType(volcengineTransitRouterMod, "VpnAttachments")},
			"volcengine_transit_router_route_entries":                      {Tok: dataSourceType(volcengineTransitRouterMod, "RouteEntries")},
			"volcengine_transit_routers":                                   {Tok: dataSourceType(volcengineTransitRouterMod, "TransitRouters")},
			"volcengine_transit_router_bandwidth_packages":                 {Tok: dataSourceType(volcengineTransitRouterMod, "BandwidthPackages")},
			"volcengine_transit_router_vpc_attachments":                    {Tok: dataSourceType(volcengineTransitRouterMod, "VpcAttachments")},
			"volcengine_transit_router_direct_connect_gateway_attachments": {Tok: dataSourceType(volcengineTransitRouterMod, "DirectConnectGatewayAttachments")},
			"volcengine_transit_router_peer_attachments":                   {Tok: dataSourceType(volcengineTransitRouterMod, "PeerAttachments")},
			"volcengine_transit_router_grant_rules":                        {Tok: dataSourceType(volcengineTransitRouterMod, "GrantRules")},
			"volcengine_transit_router_route_table_associations":           {Tok: dataSourceType(volcengineTransitRouterMod, "RouteTableAssociations")},
			"volcengine_alb_listeners":                                     {Tok: dataSourceType(volcengineAlbMod, "Listeners")},
			"volcengine_alb_server_group_servers":                          {Tok: dataSourceType(volcengineAlbMod, "ServerGroupServers")},
			"volcengine_alb_acls":                                          {Tok: dataSourceType(volcengineAlbMod, "Acls")},
			"volcengine_alb_zones":                                         {Tok: dataSourceType(volcengineAlbMod, "Zones")},
			"volcengine_alb_customized_cfgs":                               {Tok: dataSourceType(volcengineAlbMod, "CustomizedCfgs")},
			"volcengine_alb_listener_domain_extensions":                    {Tok: dataSourceType(volcengineAlbMod, "ListenerDomainExtensions")},
			"volcengine_alb_server_groups":                                 {Tok: dataSourceType(volcengineAlbMod, "ServerGroups")},
			"volcengine_alb_rules":                                         {Tok: dataSourceType(volcengineAlbMod, "Rules")},
			"volcengine_albs":                                              {Tok: dataSourceType(volcengineAlbMod, "Albs")},
			"volcengine_alb_certificates":                                  {Tok: dataSourceType(volcengineAlbMod, "Certificates")},
			"volcengine_alb_health_check_templates":                        {Tok: dataSourceType(volcengineAlbMod, "HealthCheckTemplates")},
			"volcengine_alb_ca_certificates":                               {Tok: dataSourceType(volcengineAlbMod, "CaCertificates")},
			"volcengine_rds_mssql_instances":                               {Tok: dataSourceType(volcengineRdsMssqlMod, "Instances")},
			"volcengine_rds_mssql_backups":                                 {Tok: dataSourceType(volcengineRdsMssqlMod, "Backups")},
			"volcengine_rds_mssql_zones":                                   {Tok: dataSourceType(volcengineRdsMssqlMod, "Zones")},
			"volcengine_rds_mssql_regions":                                 {Tok: dataSourceType(volcengineRdsMssqlMod, "Regions")},
			"volcengine_server_group_servers":                              {Tok: dataSourceType(volcengineClbMod, "ServerGroupServers")},
			"volcengine_clbs":                                              {Tok: dataSourceType(volcengineClbMod, "Clbs")},
			"volcengine_server_groups":                                     {Tok: dataSourceType(volcengineClbMod, "ServerGroups")},
			"volcengine_certificates":                                      {Tok: dataSourceType(volcengineClbMod, "Certificates")},
			"volcengine_listeners":                                         {Tok: dataSourceType(volcengineClbMod, "Listeners")},
			"volcengine_clb_zones":                                         {Tok: dataSourceType(volcengineClbMod, "Zones")},
			"volcengine_clb_rules":                                         {Tok: dataSourceType(volcengineClbMod, "Rules")},
			"volcengine_acls":                                              {Tok: dataSourceType(volcengineClbMod, "Acls")},
			"volcengine_bandwidth_packages":                                {Tok: dataSourceType(volcengineBandwidthPackageMod, "BandwidthPackages")},
			"volcengine_bioos_clusters":                                    {Tok: dataSourceType(volcengineBioosMod, "Clusters")},
			"volcengine_bioos_workspaces":                                  {Tok: dataSourceType(volcengineBioosMod, "Workspaces")},
			"volcengine_scaling_lifecycle_hooks":                           {Tok: dataSourceType(volcengineAutoscalingMod, "ScalingLifecycleHooks")},
			"volcengine_scaling_configurations":                            {Tok: dataSourceType(volcengineAutoscalingMod, "ScalingConfigurations")},
			"volcengine_scaling_instances":                                 {Tok: dataSourceType(volcengineAutoscalingMod, "ScalingInstances")},
			"volcengine_scaling_groups":                                    {Tok: dataSourceType(volcengineAutoscalingMod, "ScalingGroups")},
			"volcengine_scaling_activities":                                {Tok: dataSourceType(volcengineAutoscalingMod, "ScalingActivities")},
			"volcengine_scaling_policies":                                  {Tok: dataSourceType(volcengineAutoscalingMod, "ScalingPolicies")},
			"volcengine_eip_addresses":                                     {Tok: dataSourceType(volcengineEipMod, "Addresses")},
			"volcengine_ecs_instances":                                     {Tok: dataSourceType(volcengineEcsMod, "Instances")},
			"volcengine_regions":                                           {Tok: dataSourceType(volcengineEcsMod, "Regions")},
			"volcengine_ecs_invocation_results":                            {Tok: dataSourceType(volcengineEcsMod, "InvocationResults")},
			"volcengine_zones":                                             {Tok: dataSourceType(volcengineEcsMod, "Zones")},
			"volcengine_ecs_commands":                                      {Tok: dataSourceType(volcengineEcsMod, "Commands")},
			"volcengine_ecs_deployment_sets":                               {Tok: dataSourceType(volcengineEcsMod, "DeploymentSets")},
			"volcengine_ecs_key_pairs":                                     {Tok: dataSourceType(volcengineEcsMod, "KeyPairs")},
			"volcengine_images":                                            {Tok: dataSourceType(volcengineEcsMod, "Images")},
			"volcengine_ecs_launch_templates":                              {Tok: dataSourceType(volcengineEcsMod, "LaunchTemplates")},
			"volcengine_ecs_invocations":                                   {Tok: dataSourceType(volcengineEcsMod, "Invocations")},
			"volcengine_ecs_available_resources":                           {Tok: dataSourceType(volcengineEcsMod, "AvailableResources")},
			"volcengine_ecs_instance_types":                                {Tok: dataSourceType(volcengineEcsMod, "InstanceTypes")},
			"volcengine_rds_mysql_instances":                               {Tok: dataSourceType(volcengineRdsMysqlMod, "Instances")},
			"volcengine_rds_mysql_accounts":                                {Tok: dataSourceType(volcengineRdsMysqlMod, "Accounts")},
			"volcengine_rds_mysql_allowlists":                              {Tok: dataSourceType(volcengineRdsMysqlMod, "Allowlists")},
			"volcengine_rds_mysql_databases":                               {Tok: dataSourceType(volcengineRdsMysqlMod, "Databases")},
			"volcengine_iam_roles":                                         {Tok: dataSourceType(volcengineIamMod, "Roles")},
			"volcengine_iam_policies":                                      {Tok: dataSourceType(volcengineIamMod, "Policies")},
			"volcengine_iam_access_keys":                                   {Tok: dataSourceType(volcengineIamMod, "AccessKeys")},
			"volcengine_iam_saml_providers":                                {Tok: dataSourceType(volcengineIamMod, "SamlProviders")},
			"volcengine_iam_users":                                         {Tok: dataSourceType(volcengineIamMod, "Users")},
			"volcengine_iam_user_groups":                                   {Tok: dataSourceType(volcengineIamMod, "UserGroups")},
			"volcengine_iam_user_group_policy_attachments":                 {Tok: dataSourceType(volcengineIamMod, "UserGroupPolicyAttachments")},
			"volcengine_rds_parameter_templates":                           {Tok: dataSourceType(volcengineRdsMod, "ParameterTemplates")},
			"volcengine_rds_ip_lists":                                      {Tok: dataSourceType(volcengineRdsMod, "IpLists")},
			"volcengine_rds_instances":                                     {Tok: dataSourceType(volcengineRdsMod, "Instances")},
			"volcengine_rds_accounts":                                      {Tok: dataSourceType(volcengineRdsMod, "Accounts")},
			"volcengine_rds_databases":                                     {Tok: dataSourceType(volcengineRdsMod, "Databases")},
			"volcengine_vke_nodes":                                         {Tok: dataSourceType(volcengineVkeMod, "Nodes")},
			"volcengine_vke_addons":                                        {Tok: dataSourceType(volcengineVkeMod, "Addons")},
			"volcengine_vke_support_addons":                                {Tok: dataSourceType(volcengineVkeMod, "SupportAddons")},
			"volcengine_vke_support_resource_types":                        {Tok: dataSourceType(volcengineVkeMod, "SupportResourceTypes")},
			"volcengine_vke_node_pools":                                    {Tok: dataSourceType(volcengineVkeMod, "NodePools")},
			"volcengine_vke_clusters":                                      {Tok: dataSourceType(volcengineVkeMod, "Clusters")},
			"volcengine_vke_kubeconfigs":                                   {Tok: dataSourceType(volcengineVkeMod, "Kubeconfigs")},
			"volcengine_nas_mount_points":                                  {Tok: dataSourceType(volcengineNasMod, "MountPoints")},
			"volcengine_nas_snapshots":                                     {Tok: dataSourceType(volcengineNasMod, "Snapshots")},
			"volcengine_nas_permission_groups":                             {Tok: dataSourceType(volcengineNasMod, "PermissionGroups")},
			"volcengine_nas_file_systems":                                  {Tok: dataSourceType(volcengineNasMod, "FileSystems")},
			"volcengine_nas_regions":                                       {Tok: dataSourceType(volcengineNasMod, "Regions")},
			"volcengine_nas_zones":                                         {Tok: dataSourceType(volcengineNasMod, "Zones")},
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

	return prov
}
