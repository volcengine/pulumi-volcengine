// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of ecs instances
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-volcengine/sdk/go/volcengine/ecs"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ecs.Instances(ctx, &ecs.InstancesArgs{
//				Ids: []string{
//					"i-ebgy6xmgjve0384ncgsc",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func Instances(ctx *pulumi.Context, args *InstancesArgs, opts ...pulumi.InvokeOption) (*InstancesResult, error) {
	var rv InstancesResult
	err := ctx.Invoke("volcengine:ecs/instances:Instances", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking Instances.
type InstancesArgs struct {
	// A list of DeploymentSet IDs.
	DeploymentSetIds []string `pulumi:"deploymentSetIds"`
	// The hpc cluster ID of ECS instance.
	HpcClusterId *string `pulumi:"hpcClusterId"`
	// A list of ECS instance IDs.
	Ids []string `pulumi:"ids"`
	// The charge type of ECS instance.
	InstanceChargeType *string `pulumi:"instanceChargeType"`
	// The key pair name of ECS instance.
	KeyPairName *string `pulumi:"keyPairName"`
	// A Name Regex of ECS instance.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// The primary ip address of ECS instance.
	PrimaryIpAddress *string `pulumi:"primaryIpAddress"`
	// The ProjectName of ECS instance.
	ProjectName *string `pulumi:"projectName"`
	// The status of ECS instance.
	Status *string `pulumi:"status"`
	// Tags.
	Tags []InstancesTag `pulumi:"tags"`
	// The VPC ID of ECS instance.
	VpcId *string `pulumi:"vpcId"`
	// The available zone ID of ECS instance.
	ZoneId *string `pulumi:"zoneId"`
}

// A collection of values returned by Instances.
type InstancesResult struct {
	DeploymentSetIds []string `pulumi:"deploymentSetIds"`
	HpcClusterId     *string  `pulumi:"hpcClusterId"`
	// The provider-assigned unique ID for this managed resource.
	Id  string   `pulumi:"id"`
	Ids []string `pulumi:"ids"`
	// The charge type of ECS instance.
	InstanceChargeType *string `pulumi:"instanceChargeType"`
	// The collection of ECS instance query.
	Instances []InstancesInstance `pulumi:"instances"`
	// The ssh key name of ECS instance.
	KeyPairName *string `pulumi:"keyPairName"`
	NameRegex   *string `pulumi:"nameRegex"`
	OutputFile  *string `pulumi:"outputFile"`
	// The private ip address of networkInterface.
	PrimaryIpAddress *string `pulumi:"primaryIpAddress"`
	// The ProjectName of ECS instance.
	ProjectName *string `pulumi:"projectName"`
	// The status of ECS instance.
	Status *string `pulumi:"status"`
	// Tags.
	Tags []InstancesTag `pulumi:"tags"`
	// The total count of ECS instance query.
	TotalCount int `pulumi:"totalCount"`
	// The VPC ID of ECS instance.
	VpcId *string `pulumi:"vpcId"`
	// The available zone ID of ECS instance.
	ZoneId *string `pulumi:"zoneId"`
}

func InstancesOutput(ctx *pulumi.Context, args InstancesOutputArgs, opts ...pulumi.InvokeOption) InstancesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (InstancesResult, error) {
			args := v.(InstancesArgs)
			r, err := Instances(ctx, &args, opts...)
			var s InstancesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(InstancesResultOutput)
}

// A collection of arguments for invoking Instances.
type InstancesOutputArgs struct {
	// A list of DeploymentSet IDs.
	DeploymentSetIds pulumi.StringArrayInput `pulumi:"deploymentSetIds"`
	// The hpc cluster ID of ECS instance.
	HpcClusterId pulumi.StringPtrInput `pulumi:"hpcClusterId"`
	// A list of ECS instance IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// The charge type of ECS instance.
	InstanceChargeType pulumi.StringPtrInput `pulumi:"instanceChargeType"`
	// The key pair name of ECS instance.
	KeyPairName pulumi.StringPtrInput `pulumi:"keyPairName"`
	// A Name Regex of ECS instance.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The primary ip address of ECS instance.
	PrimaryIpAddress pulumi.StringPtrInput `pulumi:"primaryIpAddress"`
	// The ProjectName of ECS instance.
	ProjectName pulumi.StringPtrInput `pulumi:"projectName"`
	// The status of ECS instance.
	Status pulumi.StringPtrInput `pulumi:"status"`
	// Tags.
	Tags InstancesTagArrayInput `pulumi:"tags"`
	// The VPC ID of ECS instance.
	VpcId pulumi.StringPtrInput `pulumi:"vpcId"`
	// The available zone ID of ECS instance.
	ZoneId pulumi.StringPtrInput `pulumi:"zoneId"`
}

func (InstancesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InstancesArgs)(nil)).Elem()
}

// A collection of values returned by Instances.
type InstancesResultOutput struct{ *pulumi.OutputState }

func (InstancesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstancesResult)(nil)).Elem()
}

func (o InstancesResultOutput) ToInstancesResultOutput() InstancesResultOutput {
	return o
}

func (o InstancesResultOutput) ToInstancesResultOutputWithContext(ctx context.Context) InstancesResultOutput {
	return o
}

func (o InstancesResultOutput) DeploymentSetIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v InstancesResult) []string { return v.DeploymentSetIds }).(pulumi.StringArrayOutput)
}

func (o InstancesResultOutput) HpcClusterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstancesResult) *string { return v.HpcClusterId }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o InstancesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o InstancesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v InstancesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

// The charge type of ECS instance.
func (o InstancesResultOutput) InstanceChargeType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstancesResult) *string { return v.InstanceChargeType }).(pulumi.StringPtrOutput)
}

// The collection of ECS instance query.
func (o InstancesResultOutput) Instances() InstancesInstanceArrayOutput {
	return o.ApplyT(func(v InstancesResult) []InstancesInstance { return v.Instances }).(InstancesInstanceArrayOutput)
}

// The ssh key name of ECS instance.
func (o InstancesResultOutput) KeyPairName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstancesResult) *string { return v.KeyPairName }).(pulumi.StringPtrOutput)
}

func (o InstancesResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstancesResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o InstancesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstancesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The private ip address of networkInterface.
func (o InstancesResultOutput) PrimaryIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstancesResult) *string { return v.PrimaryIpAddress }).(pulumi.StringPtrOutput)
}

// The ProjectName of ECS instance.
func (o InstancesResultOutput) ProjectName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstancesResult) *string { return v.ProjectName }).(pulumi.StringPtrOutput)
}

// The status of ECS instance.
func (o InstancesResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstancesResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

// Tags.
func (o InstancesResultOutput) Tags() InstancesTagArrayOutput {
	return o.ApplyT(func(v InstancesResult) []InstancesTag { return v.Tags }).(InstancesTagArrayOutput)
}

// The total count of ECS instance query.
func (o InstancesResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v InstancesResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

// The VPC ID of ECS instance.
func (o InstancesResultOutput) VpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstancesResult) *string { return v.VpcId }).(pulumi.StringPtrOutput)
}

// The available zone ID of ECS instance.
func (o InstancesResultOutput) ZoneId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstancesResult) *string { return v.ZoneId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(InstancesResultOutput{})
}