// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package veecp

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Use this data source to query detailed information of veecp clusters
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/ecs"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/veecp"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/vpc"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			fooZones, err := ecs.GetZones(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			fooVpc, err := vpc.NewVpc(ctx, "fooVpc", &vpc.VpcArgs{
//				VpcName:   pulumi.String("acc-test-project1"),
//				CidrBlock: pulumi.String("172.16.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			fooSubnet, err := vpc.NewSubnet(ctx, "fooSubnet", &vpc.SubnetArgs{
//				SubnetName: pulumi.String("acc-subnet-test-2"),
//				CidrBlock:  pulumi.String("172.16.0.0/24"),
//				ZoneId:     pulumi.String(fooZones.Zones[0].Id),
//				VpcId:      fooVpc.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = vpc.NewSecurityGroup(ctx, "fooSecurityGroup", &vpc.SecurityGroupArgs{
//				VpcId:             fooVpc.ID(),
//				SecurityGroupName: pulumi.String("acc-test-security-group2"),
//			})
//			if err != nil {
//				return err
//			}
//			fooCluster, err := veecp.NewCluster(ctx, "fooCluster", &veecp.ClusterArgs{
//				Description:             pulumi.String("created by terraform"),
//				DeleteProtectionEnabled: pulumi.Bool(false),
//				Profile:                 pulumi.String("Edge"),
//				ClusterConfig: &veecp.ClusterClusterConfigArgs{
//					SubnetIds: pulumi.StringArray{
//						fooSubnet.ID(),
//					},
//					ApiServerPublicAccessEnabled: pulumi.Bool(true),
//					ApiServerPublicAccessConfig: &veecp.ClusterClusterConfigApiServerPublicAccessConfigArgs{
//						PublicAccessNetworkConfig: &veecp.ClusterClusterConfigApiServerPublicAccessConfigPublicAccessNetworkConfigArgs{
//							BillingType: pulumi.String("PostPaidByBandwidth"),
//							Bandwidth:   pulumi.Int(1),
//						},
//					},
//					ResourcePublicAccessDefaultEnabled: pulumi.Bool(true),
//				},
//				PodsConfig: &veecp.ClusterPodsConfigArgs{
//					PodNetworkMode: pulumi.String("Flannel"),
//					FlannelConfig: &veecp.ClusterPodsConfigFlannelConfigArgs{
//						PodCidrs: pulumi.StringArray{
//							pulumi.String("172.22.224.0/20"),
//						},
//						MaxPodsPerNode: pulumi.Int(64),
//					},
//				},
//				ServicesConfig: &veecp.ClusterServicesConfigArgs{
//					ServiceCidrsv4s: pulumi.StringArray{
//						pulumi.String("172.30.0.0/18"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_ = veecp.GetClustersOutput(ctx, veecp.GetClustersOutputArgs{
//				Ids: pulumi.StringArray{
//					fooCluster.ID(),
//				},
//			}, nil)
//			return nil
//		})
//	}
//
// ```
//
// Deprecated: volcengine.veecp.Clusters has been deprecated in favor of volcengine.veecp.getClusters
func Clusters(ctx *pulumi.Context, args *ClustersArgs, opts ...pulumi.InvokeOption) (*ClustersResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ClustersResult
	err := ctx.Invoke("volcengine:veecp/clusters:Clusters", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking Clusters.
type ClustersArgs struct {
	// ClientToken when the cluster is created successfully. ClientToken is a string that guarantees the idempotency of the request. This string is passed in by the caller.
	CreateClientToken *string `pulumi:"createClientToken"`
	// Cluster deletion protection. Values: true: Enable deletion protection. false: Disable deletion protection.
	DeleteProtectionEnabled *bool `pulumi:"deleteProtectionEnabled"`
	// Whether to enable the edge tunnel. The value is `true` or `false`.
	EdgeTunnelEnabled *bool `pulumi:"edgeTunnelEnabled"`
	// Cluster ID. Supports exact matching. A maximum of 100 array elements can be filled in at a time. Note: When this parameter is an empty array, filtering is based on all clusters in the specified region under the account.
	Ids []string `pulumi:"ids"`
	// Cluster name.
	Name *string `pulumi:"name"`
	// A Name Regex of Cluster.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// The container network model of the cluster, the value is `Flannel` or `VpcCniShared`. Flannel: Flannel network model, an independent Underlay container network solution, combined with the global routing capability of VPC, to achieve a high-performance network experience for the cluster. VpcCniShared: VPC-CNI network model, an Underlay container network solution based on the ENI of the private network elastic network card, with high network communication performance.
	PodsConfigPodNetworkMode *string `pulumi:"podsConfigPodNetworkMode"`
	// Filter by cluster scenario: Cloud: non-edge cluster; Edge: edge cluster.
	Profiles []string `pulumi:"profiles"`
	// Array of cluster states to filter. (The elements of the array are logically ORed. A maximum of 15 state array elements can be filled at a time).
	Statuses []ClustersStatus `pulumi:"statuses"`
	// The ClientToken when the last cluster update succeeded. ClientToken is a string that guarantees the idempotency of the request. This string is passed in by the caller.
	UpdateClientToken *string `pulumi:"updateClientToken"`
}

// A collection of values returned by Clusters.
type ClustersResult struct {
	// The collection of query.
	Clusters []ClustersCluster `pulumi:"clusters"`
	// ClientToken when creation is successful. ClientToken is a string that guarantees request idempotency. This string is passed in by the caller.
	CreateClientToken *string `pulumi:"createClientToken"`
	// The delete protection of the cluster, the value is `true` or `false`.
	DeleteProtectionEnabled *bool `pulumi:"deleteProtectionEnabled"`
	EdgeTunnelEnabled       *bool `pulumi:"edgeTunnelEnabled"`
	// The provider-assigned unique ID for this managed resource.
	Id  string   `pulumi:"id"`
	Ids []string `pulumi:"ids"`
	// Cluster name.
	Name                     *string          `pulumi:"name"`
	NameRegex                *string          `pulumi:"nameRegex"`
	OutputFile               *string          `pulumi:"outputFile"`
	PodsConfigPodNetworkMode *string          `pulumi:"podsConfigPodNetworkMode"`
	Profiles                 []string         `pulumi:"profiles"`
	Statuses                 []ClustersStatus `pulumi:"statuses"`
	// The total count of query.
	TotalCount int `pulumi:"totalCount"`
	// ClientToken when the last update was successful. ClientToken is a string that guarantees request idempotency. This string is passed in by the caller.
	UpdateClientToken *string `pulumi:"updateClientToken"`
}

func ClustersOutput(ctx *pulumi.Context, args ClustersOutputArgs, opts ...pulumi.InvokeOption) ClustersResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ClustersResult, error) {
			args := v.(ClustersArgs)
			r, err := Clusters(ctx, &args, opts...)
			var s ClustersResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ClustersResultOutput)
}

// A collection of arguments for invoking Clusters.
type ClustersOutputArgs struct {
	// ClientToken when the cluster is created successfully. ClientToken is a string that guarantees the idempotency of the request. This string is passed in by the caller.
	CreateClientToken pulumi.StringPtrInput `pulumi:"createClientToken"`
	// Cluster deletion protection. Values: true: Enable deletion protection. false: Disable deletion protection.
	DeleteProtectionEnabled pulumi.BoolPtrInput `pulumi:"deleteProtectionEnabled"`
	// Whether to enable the edge tunnel. The value is `true` or `false`.
	EdgeTunnelEnabled pulumi.BoolPtrInput `pulumi:"edgeTunnelEnabled"`
	// Cluster ID. Supports exact matching. A maximum of 100 array elements can be filled in at a time. Note: When this parameter is an empty array, filtering is based on all clusters in the specified region under the account.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// Cluster name.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// A Name Regex of Cluster.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The container network model of the cluster, the value is `Flannel` or `VpcCniShared`. Flannel: Flannel network model, an independent Underlay container network solution, combined with the global routing capability of VPC, to achieve a high-performance network experience for the cluster. VpcCniShared: VPC-CNI network model, an Underlay container network solution based on the ENI of the private network elastic network card, with high network communication performance.
	PodsConfigPodNetworkMode pulumi.StringPtrInput `pulumi:"podsConfigPodNetworkMode"`
	// Filter by cluster scenario: Cloud: non-edge cluster; Edge: edge cluster.
	Profiles pulumi.StringArrayInput `pulumi:"profiles"`
	// Array of cluster states to filter. (The elements of the array are logically ORed. A maximum of 15 state array elements can be filled at a time).
	Statuses ClustersStatusArrayInput `pulumi:"statuses"`
	// The ClientToken when the last cluster update succeeded. ClientToken is a string that guarantees the idempotency of the request. This string is passed in by the caller.
	UpdateClientToken pulumi.StringPtrInput `pulumi:"updateClientToken"`
}

func (ClustersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClustersArgs)(nil)).Elem()
}

// A collection of values returned by Clusters.
type ClustersResultOutput struct{ *pulumi.OutputState }

func (ClustersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClustersResult)(nil)).Elem()
}

func (o ClustersResultOutput) ToClustersResultOutput() ClustersResultOutput {
	return o
}

func (o ClustersResultOutput) ToClustersResultOutputWithContext(ctx context.Context) ClustersResultOutput {
	return o
}

// The collection of query.
func (o ClustersResultOutput) Clusters() ClustersClusterArrayOutput {
	return o.ApplyT(func(v ClustersResult) []ClustersCluster { return v.Clusters }).(ClustersClusterArrayOutput)
}

// ClientToken when creation is successful. ClientToken is a string that guarantees request idempotency. This string is passed in by the caller.
func (o ClustersResultOutput) CreateClientToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClustersResult) *string { return v.CreateClientToken }).(pulumi.StringPtrOutput)
}

// The delete protection of the cluster, the value is `true` or `false`.
func (o ClustersResultOutput) DeleteProtectionEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ClustersResult) *bool { return v.DeleteProtectionEnabled }).(pulumi.BoolPtrOutput)
}

func (o ClustersResultOutput) EdgeTunnelEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ClustersResult) *bool { return v.EdgeTunnelEnabled }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o ClustersResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ClustersResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o ClustersResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ClustersResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

// Cluster name.
func (o ClustersResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClustersResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ClustersResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClustersResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o ClustersResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClustersResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o ClustersResultOutput) PodsConfigPodNetworkMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClustersResult) *string { return v.PodsConfigPodNetworkMode }).(pulumi.StringPtrOutput)
}

func (o ClustersResultOutput) Profiles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ClustersResult) []string { return v.Profiles }).(pulumi.StringArrayOutput)
}

func (o ClustersResultOutput) Statuses() ClustersStatusArrayOutput {
	return o.ApplyT(func(v ClustersResult) []ClustersStatus { return v.Statuses }).(ClustersStatusArrayOutput)
}

// The total count of query.
func (o ClustersResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v ClustersResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

// ClientToken when the last update was successful. ClientToken is a string that guarantees request idempotency. This string is passed in by the caller.
func (o ClustersResultOutput) UpdateClientToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClustersResult) *string { return v.UpdateClientToken }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ClustersResultOutput{})
}
