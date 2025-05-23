// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package veecp

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Use this data source to query detailed information of veecp node pools
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
//			fooSecurityGroup, err := vpc.NewSecurityGroup(ctx, "fooSecurityGroup", &vpc.SecurityGroupArgs{
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
//			fooNodePool, err := veecp.NewNodePool(ctx, "fooNodePool", &veecp.NodePoolArgs{
//				ClusterId:   fooCluster.ID(),
//				ClientToken: pulumi.String("FGAHIxa23412FGAIOHioj"),
//				AutoScaling: &veecp.NodePoolAutoScalingArgs{
//					Enabled:         pulumi.Bool(true),
//					MinReplicas:     pulumi.Int(0),
//					MaxReplicas:     pulumi.Int(5),
//					DesiredReplicas: pulumi.Int(0),
//					Priority:        pulumi.Int(5),
//					SubnetPolicy:    pulumi.String("ZoneBalance"),
//				},
//				NodeConfig: &veecp.NodePoolNodeConfigArgs{
//					InstanceTypeIds: pulumi.StringArray{
//						pulumi.String("ecs.c1ie.xlarge"),
//					},
//					SubnetIds: pulumi.StringArray{
//						fooSubnet.ID(),
//					},
//					ImageId: pulumi.String(""),
//					SystemVolume: &veecp.NodePoolNodeConfigSystemVolumeArgs{
//						Type: pulumi.String("ESSD_PL0"),
//						Size: pulumi.Int(80),
//					},
//					DataVolumes: veecp.NodePoolNodeConfigDataVolumeArray{
//						&veecp.NodePoolNodeConfigDataVolumeArgs{
//							Type:       pulumi.String("ESSD_PL0"),
//							Size:       pulumi.Int(80),
//							MountPoint: pulumi.String("/tf1"),
//						},
//						&veecp.NodePoolNodeConfigDataVolumeArgs{
//							Type:       pulumi.String("ESSD_PL0"),
//							Size:       pulumi.Int(60),
//							MountPoint: pulumi.String("/tf2"),
//						},
//					},
//					InitializeScript: pulumi.String("ZWNobyBoZWxsbyB0ZXJyYWZvcm0h"),
//					Security: &veecp.NodePoolNodeConfigSecurityArgs{
//						Login: &veecp.NodePoolNodeConfigSecurityLoginArgs{
//							Password: pulumi.String("UHdkMTIzNDU2"),
//						},
//						SecurityStrategies: pulumi.StringArray{
//							pulumi.String("Hids"),
//						},
//						SecurityGroupIds: pulumi.StringArray{
//							fooSecurityGroup.ID(),
//						},
//					},
//					AdditionalContainerStorageEnabled: pulumi.Bool(false),
//					InstanceChargeType:                pulumi.String("PostPaid"),
//					NamePrefix:                        pulumi.String("acc-test"),
//					EcsTags: veecp.NodePoolNodeConfigEcsTagArray{
//						&veecp.NodePoolNodeConfigEcsTagArgs{
//							Key:   pulumi.String("ecs_k1"),
//							Value: pulumi.String("ecs_v1"),
//						},
//					},
//				},
//				KubernetesConfig: &veecp.NodePoolKubernetesConfigArgs{
//					Labels: veecp.NodePoolKubernetesConfigLabelArray{
//						&veecp.NodePoolKubernetesConfigLabelArgs{
//							Key:   pulumi.String("label1"),
//							Value: pulumi.String("value1"),
//						},
//					},
//					Taints: veecp.NodePoolKubernetesConfigTaintArray{
//						&veecp.NodePoolKubernetesConfigTaintArgs{
//							Key:    pulumi.String("taint-key/node-type"),
//							Value:  pulumi.String("taint-value"),
//							Effect: pulumi.String("NoSchedule"),
//						},
//					},
//					Cordon: pulumi.Bool(true),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_ = veecp.GetNodePoolsOutput(ctx, veecp.GetNodePoolsOutputArgs{
//				Ids: pulumi.StringArray{
//					fooNodePool.ID(),
//				},
//			}, nil)
//			return nil
//		})
//	}
//
// ```
//
// Deprecated: volcengine.veecp.NodePools has been deprecated in favor of volcengine.veecp.getNodePools
func NodePools(ctx *pulumi.Context, args *NodePoolsArgs, opts ...pulumi.InvokeOption) (*NodePoolsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv NodePoolsResult
	err := ctx.Invoke("volcengine:veecp/nodePools:NodePools", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking NodePools.
type NodePoolsArgs struct {
	// Is enabled of AutoScaling.
	AutoScalingEnabled *bool `pulumi:"autoScalingEnabled"`
	// The ClusterId of NodePool.
	ClusterId *string `pulumi:"clusterId"`
	// The ClusterIds of NodePool IDs.
	ClusterIds []string `pulumi:"clusterIds"`
	// The ClientToken when successfully created.
	CreateClientToken *string `pulumi:"createClientToken"`
	// The IDs of NodePool.
	Ids []string `pulumi:"ids"`
	// The Name of NodePool.
	Name *string `pulumi:"name"`
	// A Name Regex of Resource.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// The Status of NodePool.
	Statuses []NodePoolsStatus `pulumi:"statuses"`
	// The ClientToken when last update was successful.
	UpdateClientToken *string `pulumi:"updateClientToken"`
}

// A collection of values returned by NodePools.
type NodePoolsResult struct {
	AutoScalingEnabled *bool `pulumi:"autoScalingEnabled"`
	// The ClusterId of NodePool.
	ClusterId  *string  `pulumi:"clusterId"`
	ClusterIds []string `pulumi:"clusterIds"`
	// The ClientToken when successfully created.
	CreateClientToken *string `pulumi:"createClientToken"`
	// The provider-assigned unique ID for this managed resource.
	Id  string   `pulumi:"id"`
	Ids []string `pulumi:"ids"`
	// The Name of NodePool.
	Name      *string `pulumi:"name"`
	NameRegex *string `pulumi:"nameRegex"`
	// The collection of NodePools query.
	NodePools  []NodePoolsNodePool `pulumi:"nodePools"`
	OutputFile *string             `pulumi:"outputFile"`
	Statuses   []NodePoolsStatus   `pulumi:"statuses"`
	// The total count of query.
	TotalCount int `pulumi:"totalCount"`
	// The ClientToken when last update was successful.
	UpdateClientToken *string `pulumi:"updateClientToken"`
}

func NodePoolsOutput(ctx *pulumi.Context, args NodePoolsOutputArgs, opts ...pulumi.InvokeOption) NodePoolsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (NodePoolsResult, error) {
			args := v.(NodePoolsArgs)
			r, err := NodePools(ctx, &args, opts...)
			var s NodePoolsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(NodePoolsResultOutput)
}

// A collection of arguments for invoking NodePools.
type NodePoolsOutputArgs struct {
	// Is enabled of AutoScaling.
	AutoScalingEnabled pulumi.BoolPtrInput `pulumi:"autoScalingEnabled"`
	// The ClusterId of NodePool.
	ClusterId pulumi.StringPtrInput `pulumi:"clusterId"`
	// The ClusterIds of NodePool IDs.
	ClusterIds pulumi.StringArrayInput `pulumi:"clusterIds"`
	// The ClientToken when successfully created.
	CreateClientToken pulumi.StringPtrInput `pulumi:"createClientToken"`
	// The IDs of NodePool.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// The Name of NodePool.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// A Name Regex of Resource.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The Status of NodePool.
	Statuses NodePoolsStatusArrayInput `pulumi:"statuses"`
	// The ClientToken when last update was successful.
	UpdateClientToken pulumi.StringPtrInput `pulumi:"updateClientToken"`
}

func (NodePoolsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NodePoolsArgs)(nil)).Elem()
}

// A collection of values returned by NodePools.
type NodePoolsResultOutput struct{ *pulumi.OutputState }

func (NodePoolsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NodePoolsResult)(nil)).Elem()
}

func (o NodePoolsResultOutput) ToNodePoolsResultOutput() NodePoolsResultOutput {
	return o
}

func (o NodePoolsResultOutput) ToNodePoolsResultOutputWithContext(ctx context.Context) NodePoolsResultOutput {
	return o
}

func (o NodePoolsResultOutput) AutoScalingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v NodePoolsResult) *bool { return v.AutoScalingEnabled }).(pulumi.BoolPtrOutput)
}

// The ClusterId of NodePool.
func (o NodePoolsResultOutput) ClusterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NodePoolsResult) *string { return v.ClusterId }).(pulumi.StringPtrOutput)
}

func (o NodePoolsResultOutput) ClusterIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NodePoolsResult) []string { return v.ClusterIds }).(pulumi.StringArrayOutput)
}

// The ClientToken when successfully created.
func (o NodePoolsResultOutput) CreateClientToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NodePoolsResult) *string { return v.CreateClientToken }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o NodePoolsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v NodePoolsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o NodePoolsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NodePoolsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

// The Name of NodePool.
func (o NodePoolsResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NodePoolsResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o NodePoolsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NodePoolsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

// The collection of NodePools query.
func (o NodePoolsResultOutput) NodePools() NodePoolsNodePoolArrayOutput {
	return o.ApplyT(func(v NodePoolsResult) []NodePoolsNodePool { return v.NodePools }).(NodePoolsNodePoolArrayOutput)
}

func (o NodePoolsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NodePoolsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o NodePoolsResultOutput) Statuses() NodePoolsStatusArrayOutput {
	return o.ApplyT(func(v NodePoolsResult) []NodePoolsStatus { return v.Statuses }).(NodePoolsStatusArrayOutput)
}

// The total count of query.
func (o NodePoolsResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v NodePoolsResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

// The ClientToken when last update was successful.
func (o NodePoolsResultOutput) UpdateClientToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NodePoolsResult) *string { return v.UpdateClientToken }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(NodePoolsResultOutput{})
}
