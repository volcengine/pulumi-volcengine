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
	"github.com/pulumi/pulumi-volcengine/provider/pkg/version"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	volcenginePkg = "volcengine"
	// modules:
	//mainMod          = "index" // the volcengine module
	volcengineVpcMod = "Vpc"
	volcengineClbMod = "Clb"
	volcengineEbsMod = "Ebs"
	volcengineEcsMod = "Ecs"
	volcengineEipMod = "Eip"
	volcengineVkeMod = "Vke"
	volcengineNatMod = "Nat"
	volcengineIamMod = "Iam"
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
		DisplayName: "",
		// The default publisher for all packages is Pulumi.
		// Change this to your personal name (or a company name) that you
		// would like to be shown in the Pulumi Registry if this package is published
		// there.
		Publisher: "Pulumi",
		// LogoURL is optional but useful to help identify your package in the Pulumi Registry
		// if this package is published there.
		//
		// You may host a logo on a domain you control or add an SVG logo for your package
		// in your repository and use the raw content URL for that file as your logo URL.
		LogoURL: "",
		// PluginDownloadURL is an optional URL used to download the Provider
		// for use in Pulumi programs
		// e.g https://github.com/org/pulumi-provider-name/releases/
		PluginDownloadURL: "",
		Description:       "A Pulumi package for creating and managing volcengine cloud resources.",
		// category/cloud tag helps with categorizing the package in the Pulumi Registry.
		// For all available categories, see `Keywords` in
		// https://www.pulumi.com/docs/guides/pulumi-packages/schema/#package.
		Keywords:   []string{"pulumi", "volcengine", "category/cloud"},
		License:    "Apache-2.0",
		Homepage:   "https://www.pulumi.com",
		Repository: "https://github.com/pulumi/pulumi-volcengine",
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
		},
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			//acl
			"volcengine_acl":       {Tok: resourceType(volcengineVpcMod, "Acl")},
			"volcengine_acl_entry": {Tok: resourceType(volcengineVpcMod, "AclEntry")},
			//certificate
			"volcengine_certificate": {Tok: resourceType(volcengineClbMod, "Certificate")},
			//clb
			"volcengine_clb":      {Tok: resourceType(volcengineClbMod, "Clb")},
			"volcengine_clb_rule": {Tok: resourceType(volcengineClbMod, "ClbRule")},
			//ecs
			"volcengine_ecs_instance":       {Tok: resourceType(volcengineEcsMod, "Instance")},
			"volcengine_ecs_instance_state": {Tok: resourceType(volcengineEcsMod, "State")},
			//eip
			"volcengine_eip_address":   {Tok: resourceType(volcengineEipMod, "Address")},
			"volcengine_eip_associate": {Tok: resourceType(volcengineEipMod, "Associate")},
			//listener
			"volcengine_listener": {Tok: resourceType(volcengineClbMod, "Listener")},
			//nat
			"volcengine_nat_gateway": {Tok: resourceType(volcengineNatMod, "Gateway")},
			"volcengine_snat_entry":  {Tok: resourceType(volcengineNatMod, "SnatEntry")},
			//network_interface
			"volcengine_network_interface":        {Tok: resourceType(volcengineVpcMod, "NetworkInterface")},
			"volcengine_network_interface_attach": {Tok: resourceType(volcengineVpcMod, "NetworkInterfaceAttach")},
			//route
			"volcengine_route_entry":           {Tok: resourceType(volcengineVpcMod, "RouteEntry")},
			"volcengine_route_table":           {Tok: resourceType(volcengineVpcMod, "RouteTable")},
			"volcengine_route_table_associate": {Tok: resourceType(volcengineVpcMod, "RouteTableAssociate")},
			//security_group
			"volcengine_security_group":      {Tok: resourceType(volcengineVpcMod, "SecurityGroup")},
			"volcengine_security_group_rule": {Tok: resourceType(volcengineVpcMod, "SecurityGroupRule")},
			//server_group
			"volcengine_server_group":        {Tok: resourceType(volcengineClbMod, "ServerGroup")},
			"volcengine_server_group_server": {Tok: resourceType(volcengineClbMod, "ServerGroupServer")},
			//subnet
			"volcengine_subnet": {Tok: resourceType(volcengineVpcMod, "Subnet")},
			//vke
			"volcengine_vke_cluster":   {Tok: resourceType(volcengineVkeMod, "Cluster")},
			"volcengine_vke_node":      {Tok: resourceType(volcengineVkeMod, "Node")},
			"volcengine_vke_node_pool": {Tok: resourceType(volcengineVkeMod, "NodePool")},
			//ebs
			"volcengine_volume":        {Tok: resourceType(volcengineEbsMod, "Volume")},
			"volcengine_volume_attach": {Tok: resourceType(volcengineEbsMod, "VolumeAttach")},
			//vpc
			"volcengine_vpc": {Tok: resourceType(volcengineVpcMod, "Vpc")},
			//iam
			"volcengine_iam_access_key":             {Tok: resourceType(volcengineIamMod, "AccessKey")},
			"volcengine_iam_login_profile":          {Tok: resourceType(volcengineIamMod, "LoginProfile")},
			"volcengine_iam_policy":                 {Tok: resourceType(volcengineIamMod, "Policy")},
			"volcengine_iam_role":                   {Tok: resourceType(volcengineIamMod, "Role")},
			"volcengine_iam_role_policy_attachment": {Tok: resourceType(volcengineIamMod, "RolePolicyAttachment")},
			"volcengine_iam_user":                   {Tok: resourceType(volcengineIamMod, "User")},
			"volcengine_iam_user_policy_attachment": {Tok: resourceType(volcengineIamMod, "UserPolicyAttachment")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			//acl
			"volcengine_acls": {Tok: dataSourceType(volcengineVpcMod, "Acls")},
			//certificates
			"volcengine_certificates": {Tok: dataSourceType(volcengineVpcMod, "Certificates")},
			//clb
			"volcengine_clb_rules": {Tok: dataSourceType(volcengineClbMod, "ClbRules")},
			"volcengine_clbs":      {Tok: dataSourceType(volcengineClbMod, "Clbs")},
			//ecs
			"volcengine_ecs_instances": {Tok: dataSourceType(volcengineEcsMod, "Instances")},
			//eip
			"volcengine_eip_addresses": {Tok: dataSourceType(volcengineEipMod, "Addresses")},
			//images
			"volcengine_images": {Tok: dataSourceType(volcengineEcsMod, "Images")},
			//listener
			"volcengine_listeners": {Tok: dataSourceType(volcengineClbMod, "Listeners")},
			//nat
			"volcengine_nat_gateways": {Tok: dataSourceType(volcengineVpcMod, "Gateways")},
			"volcengine_snat_entries": {Tok: dataSourceType(volcengineVpcMod, "SnatEntries")},
			//network_interface
			"volcengine_network_interfaces": {Tok: dataSourceType(volcengineVpcMod, "NetworkInterfaces")},
			//route
			"volcengine_route_entries": {Tok: dataSourceType(volcengineVpcMod, "RouteEntries")},
			"volcengine_route_tables":  {Tok: dataSourceType(volcengineVpcMod, "RouteTables")},
			//security_group
			"volcengine_security_groups": {Tok: dataSourceType(volcengineVpcMod, "SecurityGroups")},
			//server_group
			"volcengine_server_group_servers": {Tok: dataSourceType(volcengineClbMod, "ServerGroupServers")},
			"volcengine_server_groups":        {Tok: dataSourceType(volcengineClbMod, "ServerGroups")},
			//subnet
			"volcengine_subnets": {Tok: dataSourceType(volcengineVpcMod, "Subnets")},
			//vke
			"volcengine_vke_clusters":   {Tok: dataSourceType(volcengineVkeMod, "Clusters")},
			"volcengine_vke_node_pools": {Tok: dataSourceType(volcengineVkeMod, "NodePools")},
			"volcengine_vke_nodes":      {Tok: dataSourceType(volcengineVkeMod, "Nodes")},
			//ebs
			"volcengine_volumes": {Tok: dataSourceType(volcengineEbsMod, "Volumes")},
			//vpc
			"volcengine_vpcs":  {Tok: dataSourceType(volcengineVpcMod, "Vpcs")},
			"volcengine_zones": {Tok: dataSourceType(volcengineVpcMod, "Zones")},
			//iam
			"volcengine_iam_policies": {Tok: dataSourceType(volcengineIamMod, "Policies")},
			"volcengine_iam_roles":    {Tok: dataSourceType(volcengineIamMod, "Roles")},
			"volcengine_iam_users":    {Tok: dataSourceType(volcengineIamMod, "Users")},
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
		},
		Python: &tfbridge.PythonInfo{
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/pulumi/pulumi-%[1]s/sdk/", volcenginePkg),
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
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}
