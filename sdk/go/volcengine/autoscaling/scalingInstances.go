// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package autoscaling

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Use this data source to query detailed information of scaling instances
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/autoscaling"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/ecs"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/vpc"
//
// )
// func main() {
// pulumi.Run(func(ctx *pulumi.Context) error {
// fooZones, err := ecs.GetZones(ctx, nil, nil);
// if err != nil {
// return err
// }
// fooVpc, err := vpc.NewVpc(ctx, "fooVpc", &vpc.VpcArgs{
// VpcName: pulumi.String("acc-test-vpc"),
// CidrBlock: pulumi.String("172.16.0.0/16"),
// })
// if err != nil {
// return err
// }
// fooSubnet, err := vpc.NewSubnet(ctx, "fooSubnet", &vpc.SubnetArgs{
// SubnetName: pulumi.String("acc-test-subnet"),
// CidrBlock: pulumi.String("172.16.0.0/24"),
// ZoneId: pulumi.String(fooZones.Zones[0].Id),
// VpcId: fooVpc.ID(),
// })
// if err != nil {
// return err
// }
// fooSecurityGroup, err := vpc.NewSecurityGroup(ctx, "fooSecurityGroup", &vpc.SecurityGroupArgs{
// SecurityGroupName: pulumi.String("acc-test-security-group"),
// VpcId: fooVpc.ID(),
// })
// if err != nil {
// return err
// }
// fooImages, err := ecs.GetImages(ctx, &ecs.GetImagesArgs{
// OsType: pulumi.StringRef("Linux"),
// Visibility: pulumi.StringRef("public"),
// InstanceTypeId: pulumi.StringRef("ecs.g1.large"),
// }, nil);
// if err != nil {
// return err
// }
// fooKeyPair, err := ecs.NewKeyPair(ctx, "fooKeyPair", &ecs.KeyPairArgs{
// Description: pulumi.String("acc-test-2"),
// KeyPairName: pulumi.String("acc-test-key-pair-name"),
// })
// if err != nil {
// return err
// }
// fooLaunchTemplate, err := ecs.NewLaunchTemplate(ctx, "fooLaunchTemplate", &ecs.LaunchTemplateArgs{
// Description: pulumi.String("acc-test-desc"),
// EipBandwidth: pulumi.Int(200),
// EipBillingType: pulumi.String("PostPaidByBandwidth"),
// EipIsp: pulumi.String("BGP"),
// HostName: pulumi.String("acc-hostname"),
// ImageId: pulumi.String(fooImages.Images[0].ImageId),
// InstanceChargeType: pulumi.String("PostPaid"),
// InstanceName: pulumi.String("acc-instance-name"),
// InstanceTypeId: pulumi.String("ecs.g1.large"),
// KeyPairName: fooKeyPair.KeyPairName,
// LaunchTemplateName: pulumi.String("acc-test-template"),
// NetworkInterfaces: ecs.LaunchTemplateNetworkInterfaceArray{
// &ecs.LaunchTemplateNetworkInterfaceArgs{
// SubnetId: fooSubnet.ID(),
// SecurityGroupIds: pulumi.StringArray{
// fooSecurityGroup.ID(),
// },
// },
// },
// Volumes: ecs.LaunchTemplateVolumeArray{
// &ecs.LaunchTemplateVolumeArgs{
// VolumeType: pulumi.String("ESSD_PL0"),
// Size: pulumi.Int(50),
// DeleteWithInstance: pulumi.Bool(true),
// },
// },
// })
// if err != nil {
// return err
// }
// fooScalingGroup, err := autoscaling.NewScalingGroup(ctx, "fooScalingGroup", &autoscaling.ScalingGroupArgs{
// ScalingGroupName: pulumi.String("acc-test-scaling-group"),
// SubnetIds: pulumi.StringArray{
// fooSubnet.ID(),
// },
// MultiAzPolicy: pulumi.String("BALANCE"),
// DesireInstanceNumber: -1,
// MinInstanceNumber: pulumi.Int(0),
// MaxInstanceNumber: pulumi.Int(10),
// InstanceTerminatePolicy: pulumi.String("OldestInstance"),
// DefaultCooldown: pulumi.Int(10),
// LaunchTemplateId: fooLaunchTemplate.ID(),
// LaunchTemplateVersion: pulumi.String("Default"),
// })
// if err != nil {
// return err
// }
// fooScalingGroupEnabler, err := autoscaling.NewScalingGroupEnabler(ctx, "fooScalingGroupEnabler", &autoscaling.ScalingGroupEnablerArgs{
// ScalingGroupId: fooScalingGroup.ID(),
// })
// if err != nil {
// return err
// }
// var fooInstance []*ecs.Instance
//
//	for index := 0; index < 3; index++ {
//	    key0 := index
//	    val0 := index
//
// __res, err := ecs.NewInstance(ctx, fmt.Sprintf("fooInstance-%v", key0), &ecs.InstanceArgs{
// InstanceName: pulumi.String(fmt.Sprintf("acc-test-ecs-%v", val0)),
// Description: pulumi.String("acc-test"),
// HostName: pulumi.String("tf-acc-test"),
// ImageId: pulumi.String(fooImages.Images[0].ImageId),
// InstanceType: pulumi.String("ecs.g1.large"),
// Password: pulumi.String("93f0cb0614Aab12"),
// InstanceChargeType: pulumi.String("PostPaid"),
// SystemVolumeType: pulumi.String("ESSD_PL0"),
// SystemVolumeSize: pulumi.Int(40),
// SubnetId: fooSubnet.ID(),
// SecurityGroupIds: pulumi.StringArray{
// fooSecurityGroup.ID(),
// },
// })
// if err != nil {
// return err
// }
// fooInstance = append(fooInstance, __res)
// }
// var fooScalingInstanceAttachment []*autoscaling.ScalingInstanceAttachment
//
//	for index := 0; index < len(fooInstance); index++ {
//	    key0 := index
//	    val0 := index
//
// __res, err := autoscaling.NewScalingInstanceAttachment(ctx, fmt.Sprintf("fooScalingInstanceAttachment-%v", key0), &autoscaling.ScalingInstanceAttachmentArgs{
// InstanceId: fooInstance[val0].ID(),
// ScalingGroupId: fooScalingGroup.ID(),
// Entrusted: pulumi.Bool(true),
// }, pulumi.DependsOn([]pulumi.Resource{
// fooScalingGroupEnabler,
// }))
// if err != nil {
// return err
// }
// fooScalingInstanceAttachment = append(fooScalingInstanceAttachment, __res)
// }
// _ = autoscaling.GetScalingInstancesOutput(ctx, autoscaling.GetScalingInstancesOutputArgs{
// ScalingGroupId: fooScalingGroup.ID(),
// Ids: %!v(PANIC=Format method: fatal: A failure has occurred: unlowered splat expression @ #-functions-volcengine:autoscaling-scalingInstances:ScalingInstances.pp:93,22-64),
// }, nil);
// return nil
// })
// }
// ```
//
// Deprecated: volcengine.autoscaling.ScalingInstances has been deprecated in favor of volcengine.autoscaling.getScalingInstances
func ScalingInstances(ctx *pulumi.Context, args *ScalingInstancesArgs, opts ...pulumi.InvokeOption) (*ScalingInstancesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ScalingInstancesResult
	err := ctx.Invoke("volcengine:autoscaling/scalingInstances:ScalingInstances", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking ScalingInstances.
type ScalingInstancesArgs struct {
	// The creation type of the instances. Valid values: AutoCreated, Attached.
	CreationType *string `pulumi:"creationType"`
	// A list of instance ids.
	Ids []string `pulumi:"ids"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// The id of the scaling configuration id.
	ScalingConfigurationId *string `pulumi:"scalingConfigurationId"`
	// The id of the scaling group.
	ScalingGroupId string `pulumi:"scalingGroupId"`
	// The status of instances. Valid values: Init, Pending, Pending:Wait, InService, Error, Removing, Removing:Wait, Stopped, Protected.
	Status *string `pulumi:"status"`
}

// A collection of values returned by ScalingInstances.
type ScalingInstancesResult struct {
	// The creation type of the instance. Valid values: AutoCreated, Attached.
	CreationType *string `pulumi:"creationType"`
	// The provider-assigned unique ID for this managed resource.
	Id         string   `pulumi:"id"`
	Ids        []string `pulumi:"ids"`
	OutputFile *string  `pulumi:"outputFile"`
	// The id of the scaling configuration.
	ScalingConfigurationId *string `pulumi:"scalingConfigurationId"`
	// The id of the scaling group.
	ScalingGroupId string `pulumi:"scalingGroupId"`
	// The collection of scaling instances.
	ScalingInstances []ScalingInstancesScalingInstance `pulumi:"scalingInstances"`
	// The status of instances.
	Status *string `pulumi:"status"`
	// The total count of scaling instances query.
	TotalCount int `pulumi:"totalCount"`
}

func ScalingInstancesOutput(ctx *pulumi.Context, args ScalingInstancesOutputArgs, opts ...pulumi.InvokeOption) ScalingInstancesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ScalingInstancesResult, error) {
			args := v.(ScalingInstancesArgs)
			r, err := ScalingInstances(ctx, &args, opts...)
			var s ScalingInstancesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ScalingInstancesResultOutput)
}

// A collection of arguments for invoking ScalingInstances.
type ScalingInstancesOutputArgs struct {
	// The creation type of the instances. Valid values: AutoCreated, Attached.
	CreationType pulumi.StringPtrInput `pulumi:"creationType"`
	// A list of instance ids.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The id of the scaling configuration id.
	ScalingConfigurationId pulumi.StringPtrInput `pulumi:"scalingConfigurationId"`
	// The id of the scaling group.
	ScalingGroupId pulumi.StringInput `pulumi:"scalingGroupId"`
	// The status of instances. Valid values: Init, Pending, Pending:Wait, InService, Error, Removing, Removing:Wait, Stopped, Protected.
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (ScalingInstancesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalingInstancesArgs)(nil)).Elem()
}

// A collection of values returned by ScalingInstances.
type ScalingInstancesResultOutput struct{ *pulumi.OutputState }

func (ScalingInstancesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalingInstancesResult)(nil)).Elem()
}

func (o ScalingInstancesResultOutput) ToScalingInstancesResultOutput() ScalingInstancesResultOutput {
	return o
}

func (o ScalingInstancesResultOutput) ToScalingInstancesResultOutputWithContext(ctx context.Context) ScalingInstancesResultOutput {
	return o
}

// The creation type of the instance. Valid values: AutoCreated, Attached.
func (o ScalingInstancesResultOutput) CreationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalingInstancesResult) *string { return v.CreationType }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o ScalingInstancesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ScalingInstancesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o ScalingInstancesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ScalingInstancesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o ScalingInstancesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalingInstancesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The id of the scaling configuration.
func (o ScalingInstancesResultOutput) ScalingConfigurationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalingInstancesResult) *string { return v.ScalingConfigurationId }).(pulumi.StringPtrOutput)
}

// The id of the scaling group.
func (o ScalingInstancesResultOutput) ScalingGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v ScalingInstancesResult) string { return v.ScalingGroupId }).(pulumi.StringOutput)
}

// The collection of scaling instances.
func (o ScalingInstancesResultOutput) ScalingInstances() ScalingInstancesScalingInstanceArrayOutput {
	return o.ApplyT(func(v ScalingInstancesResult) []ScalingInstancesScalingInstance { return v.ScalingInstances }).(ScalingInstancesScalingInstanceArrayOutput)
}

// The status of instances.
func (o ScalingInstancesResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalingInstancesResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

// The total count of scaling instances query.
func (o ScalingInstancesResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v ScalingInstancesResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(ScalingInstancesResultOutput{})
}
