// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package autoscaling

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage scaling group
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-volcengine/sdk/go/volcengine/autoscaling"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := autoscaling.NewScalingGroup(ctx, "foo", &autoscaling.ScalingGroupArgs{
//				DefaultCooldown:         pulumi.Int(10),
//				DesireInstanceNumber:    pulumi.Int(0),
//				InstanceTerminatePolicy: pulumi.String("OldestInstance"),
//				MaxInstanceNumber:       pulumi.Int(1),
//				MinInstanceNumber:       pulumi.Int(0),
//				MultiAzPolicy:           pulumi.String("BALANCE"),
//				ScalingGroupName:        pulumi.String("tf-test"),
//				SubnetIds: pulumi.StringArray{
//					pulumi.String("subnet-2ff1n75eyf08w59gp67qhnhqm"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// ScalingGroup can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import volcengine:autoscaling/scalingGroup:ScalingGroup default scg-mizl7m1kqccg5smt1bdpijuj
//
// ```
type ScalingGroup struct {
	pulumi.CustomResourceState

	// The scaling configuration id which used by the scaling group.
	ActiveScalingConfigurationId pulumi.StringOutput `pulumi:"activeScalingConfigurationId"`
	// The create time of the scaling group.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The list of db instance ids.
	DbInstanceIds pulumi.StringArrayOutput `pulumi:"dbInstanceIds"`
	// The default cooldown interval of the scaling group. Default value: 300.
	DefaultCooldown pulumi.IntOutput `pulumi:"defaultCooldown"`
	// The desire instance number of the scaling group.
	DesireInstanceNumber pulumi.IntOutput `pulumi:"desireInstanceNumber"`
	// The instance terminate policy of the scaling group. Valid values: OldestInstance, NewestInstance, OldestScalingConfigurationWithOldestInstance, OldestScalingConfigurationWithNewestInstance. Default value: OldestScalingConfigurationWithOldestInstance.
	InstanceTerminatePolicy pulumi.StringOutput `pulumi:"instanceTerminatePolicy"`
	// The ID of the launch template bound to the scaling group.
	LaunchTemplateId pulumi.StringPtrOutput `pulumi:"launchTemplateId"`
	// The version of the launch template bound to the scaling group.
	LaunchTemplateVersion pulumi.StringPtrOutput `pulumi:"launchTemplateVersion"`
	// The lifecycle state of the scaling group.
	LifecycleState pulumi.StringOutput `pulumi:"lifecycleState"`
	// The max instance number of the scaling group.
	MaxInstanceNumber pulumi.IntOutput `pulumi:"maxInstanceNumber"`
	// The min instance number of the scaling group.
	MinInstanceNumber pulumi.IntOutput `pulumi:"minInstanceNumber"`
	// The multi az policy of the scaling group. Valid values: PRIORITY, BALANCE. Default value: PRIORITY.
	MultiAzPolicy pulumi.StringOutput `pulumi:"multiAzPolicy"`
	// The id of the scaling group.
	ScalingGroupId pulumi.StringOutput `pulumi:"scalingGroupId"`
	// The name of the scaling group.
	ScalingGroupName pulumi.StringOutput `pulumi:"scalingGroupName"`
	// The load balancer server group attributes of the scaling group.
	ServerGroupAttributes ScalingGroupServerGroupAttributeArrayOutput `pulumi:"serverGroupAttributes"`
	// The list of the subnet id to which the ENI is connected.
	SubnetIds pulumi.StringArrayOutput `pulumi:"subnetIds"`
	// The total instance count of the scaling group.
	TotalInstanceCount pulumi.IntOutput `pulumi:"totalInstanceCount"`
	// The create time of the scaling group.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// The VPC id of the scaling group.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewScalingGroup registers a new resource with the given unique name, arguments, and options.
func NewScalingGroup(ctx *pulumi.Context,
	name string, args *ScalingGroupArgs, opts ...pulumi.ResourceOption) (*ScalingGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MaxInstanceNumber == nil {
		return nil, errors.New("invalid value for required argument 'MaxInstanceNumber'")
	}
	if args.MinInstanceNumber == nil {
		return nil, errors.New("invalid value for required argument 'MinInstanceNumber'")
	}
	if args.ScalingGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ScalingGroupName'")
	}
	if args.SubnetIds == nil {
		return nil, errors.New("invalid value for required argument 'SubnetIds'")
	}
	var resource ScalingGroup
	err := ctx.RegisterResource("volcengine:autoscaling/scalingGroup:ScalingGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetScalingGroup gets an existing ScalingGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetScalingGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScalingGroupState, opts ...pulumi.ResourceOption) (*ScalingGroup, error) {
	var resource ScalingGroup
	err := ctx.ReadResource("volcengine:autoscaling/scalingGroup:ScalingGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ScalingGroup resources.
type scalingGroupState struct {
	// The scaling configuration id which used by the scaling group.
	ActiveScalingConfigurationId *string `pulumi:"activeScalingConfigurationId"`
	// The create time of the scaling group.
	CreatedAt *string `pulumi:"createdAt"`
	// The list of db instance ids.
	DbInstanceIds []string `pulumi:"dbInstanceIds"`
	// The default cooldown interval of the scaling group. Default value: 300.
	DefaultCooldown *int `pulumi:"defaultCooldown"`
	// The desire instance number of the scaling group.
	DesireInstanceNumber *int `pulumi:"desireInstanceNumber"`
	// The instance terminate policy of the scaling group. Valid values: OldestInstance, NewestInstance, OldestScalingConfigurationWithOldestInstance, OldestScalingConfigurationWithNewestInstance. Default value: OldestScalingConfigurationWithOldestInstance.
	InstanceTerminatePolicy *string `pulumi:"instanceTerminatePolicy"`
	// The ID of the launch template bound to the scaling group.
	LaunchTemplateId *string `pulumi:"launchTemplateId"`
	// The version of the launch template bound to the scaling group.
	LaunchTemplateVersion *string `pulumi:"launchTemplateVersion"`
	// The lifecycle state of the scaling group.
	LifecycleState *string `pulumi:"lifecycleState"`
	// The max instance number of the scaling group.
	MaxInstanceNumber *int `pulumi:"maxInstanceNumber"`
	// The min instance number of the scaling group.
	MinInstanceNumber *int `pulumi:"minInstanceNumber"`
	// The multi az policy of the scaling group. Valid values: PRIORITY, BALANCE. Default value: PRIORITY.
	MultiAzPolicy *string `pulumi:"multiAzPolicy"`
	// The id of the scaling group.
	ScalingGroupId *string `pulumi:"scalingGroupId"`
	// The name of the scaling group.
	ScalingGroupName *string `pulumi:"scalingGroupName"`
	// The load balancer server group attributes of the scaling group.
	ServerGroupAttributes []ScalingGroupServerGroupAttribute `pulumi:"serverGroupAttributes"`
	// The list of the subnet id to which the ENI is connected.
	SubnetIds []string `pulumi:"subnetIds"`
	// The total instance count of the scaling group.
	TotalInstanceCount *int `pulumi:"totalInstanceCount"`
	// The create time of the scaling group.
	UpdatedAt *string `pulumi:"updatedAt"`
	// The VPC id of the scaling group.
	VpcId *string `pulumi:"vpcId"`
}

type ScalingGroupState struct {
	// The scaling configuration id which used by the scaling group.
	ActiveScalingConfigurationId pulumi.StringPtrInput
	// The create time of the scaling group.
	CreatedAt pulumi.StringPtrInput
	// The list of db instance ids.
	DbInstanceIds pulumi.StringArrayInput
	// The default cooldown interval of the scaling group. Default value: 300.
	DefaultCooldown pulumi.IntPtrInput
	// The desire instance number of the scaling group.
	DesireInstanceNumber pulumi.IntPtrInput
	// The instance terminate policy of the scaling group. Valid values: OldestInstance, NewestInstance, OldestScalingConfigurationWithOldestInstance, OldestScalingConfigurationWithNewestInstance. Default value: OldestScalingConfigurationWithOldestInstance.
	InstanceTerminatePolicy pulumi.StringPtrInput
	// The ID of the launch template bound to the scaling group.
	LaunchTemplateId pulumi.StringPtrInput
	// The version of the launch template bound to the scaling group.
	LaunchTemplateVersion pulumi.StringPtrInput
	// The lifecycle state of the scaling group.
	LifecycleState pulumi.StringPtrInput
	// The max instance number of the scaling group.
	MaxInstanceNumber pulumi.IntPtrInput
	// The min instance number of the scaling group.
	MinInstanceNumber pulumi.IntPtrInput
	// The multi az policy of the scaling group. Valid values: PRIORITY, BALANCE. Default value: PRIORITY.
	MultiAzPolicy pulumi.StringPtrInput
	// The id of the scaling group.
	ScalingGroupId pulumi.StringPtrInput
	// The name of the scaling group.
	ScalingGroupName pulumi.StringPtrInput
	// The load balancer server group attributes of the scaling group.
	ServerGroupAttributes ScalingGroupServerGroupAttributeArrayInput
	// The list of the subnet id to which the ENI is connected.
	SubnetIds pulumi.StringArrayInput
	// The total instance count of the scaling group.
	TotalInstanceCount pulumi.IntPtrInput
	// The create time of the scaling group.
	UpdatedAt pulumi.StringPtrInput
	// The VPC id of the scaling group.
	VpcId pulumi.StringPtrInput
}

func (ScalingGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*scalingGroupState)(nil)).Elem()
}

type scalingGroupArgs struct {
	// The default cooldown interval of the scaling group. Default value: 300.
	DefaultCooldown *int `pulumi:"defaultCooldown"`
	// The desire instance number of the scaling group.
	DesireInstanceNumber *int `pulumi:"desireInstanceNumber"`
	// The instance terminate policy of the scaling group. Valid values: OldestInstance, NewestInstance, OldestScalingConfigurationWithOldestInstance, OldestScalingConfigurationWithNewestInstance. Default value: OldestScalingConfigurationWithOldestInstance.
	InstanceTerminatePolicy *string `pulumi:"instanceTerminatePolicy"`
	// The ID of the launch template bound to the scaling group.
	LaunchTemplateId *string `pulumi:"launchTemplateId"`
	// The version of the launch template bound to the scaling group.
	LaunchTemplateVersion *string `pulumi:"launchTemplateVersion"`
	// The max instance number of the scaling group.
	MaxInstanceNumber int `pulumi:"maxInstanceNumber"`
	// The min instance number of the scaling group.
	MinInstanceNumber int `pulumi:"minInstanceNumber"`
	// The multi az policy of the scaling group. Valid values: PRIORITY, BALANCE. Default value: PRIORITY.
	MultiAzPolicy *string `pulumi:"multiAzPolicy"`
	// The name of the scaling group.
	ScalingGroupName string `pulumi:"scalingGroupName"`
	// The load balancer server group attributes of the scaling group.
	ServerGroupAttributes []ScalingGroupServerGroupAttribute `pulumi:"serverGroupAttributes"`
	// The list of the subnet id to which the ENI is connected.
	SubnetIds []string `pulumi:"subnetIds"`
}

// The set of arguments for constructing a ScalingGroup resource.
type ScalingGroupArgs struct {
	// The default cooldown interval of the scaling group. Default value: 300.
	DefaultCooldown pulumi.IntPtrInput
	// The desire instance number of the scaling group.
	DesireInstanceNumber pulumi.IntPtrInput
	// The instance terminate policy of the scaling group. Valid values: OldestInstance, NewestInstance, OldestScalingConfigurationWithOldestInstance, OldestScalingConfigurationWithNewestInstance. Default value: OldestScalingConfigurationWithOldestInstance.
	InstanceTerminatePolicy pulumi.StringPtrInput
	// The ID of the launch template bound to the scaling group.
	LaunchTemplateId pulumi.StringPtrInput
	// The version of the launch template bound to the scaling group.
	LaunchTemplateVersion pulumi.StringPtrInput
	// The max instance number of the scaling group.
	MaxInstanceNumber pulumi.IntInput
	// The min instance number of the scaling group.
	MinInstanceNumber pulumi.IntInput
	// The multi az policy of the scaling group. Valid values: PRIORITY, BALANCE. Default value: PRIORITY.
	MultiAzPolicy pulumi.StringPtrInput
	// The name of the scaling group.
	ScalingGroupName pulumi.StringInput
	// The load balancer server group attributes of the scaling group.
	ServerGroupAttributes ScalingGroupServerGroupAttributeArrayInput
	// The list of the subnet id to which the ENI is connected.
	SubnetIds pulumi.StringArrayInput
}

func (ScalingGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scalingGroupArgs)(nil)).Elem()
}

type ScalingGroupInput interface {
	pulumi.Input

	ToScalingGroupOutput() ScalingGroupOutput
	ToScalingGroupOutputWithContext(ctx context.Context) ScalingGroupOutput
}

func (*ScalingGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**ScalingGroup)(nil)).Elem()
}

func (i *ScalingGroup) ToScalingGroupOutput() ScalingGroupOutput {
	return i.ToScalingGroupOutputWithContext(context.Background())
}

func (i *ScalingGroup) ToScalingGroupOutputWithContext(ctx context.Context) ScalingGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalingGroupOutput)
}

// ScalingGroupArrayInput is an input type that accepts ScalingGroupArray and ScalingGroupArrayOutput values.
// You can construct a concrete instance of `ScalingGroupArrayInput` via:
//
//	ScalingGroupArray{ ScalingGroupArgs{...} }
type ScalingGroupArrayInput interface {
	pulumi.Input

	ToScalingGroupArrayOutput() ScalingGroupArrayOutput
	ToScalingGroupArrayOutputWithContext(context.Context) ScalingGroupArrayOutput
}

type ScalingGroupArray []ScalingGroupInput

func (ScalingGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ScalingGroup)(nil)).Elem()
}

func (i ScalingGroupArray) ToScalingGroupArrayOutput() ScalingGroupArrayOutput {
	return i.ToScalingGroupArrayOutputWithContext(context.Background())
}

func (i ScalingGroupArray) ToScalingGroupArrayOutputWithContext(ctx context.Context) ScalingGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalingGroupArrayOutput)
}

// ScalingGroupMapInput is an input type that accepts ScalingGroupMap and ScalingGroupMapOutput values.
// You can construct a concrete instance of `ScalingGroupMapInput` via:
//
//	ScalingGroupMap{ "key": ScalingGroupArgs{...} }
type ScalingGroupMapInput interface {
	pulumi.Input

	ToScalingGroupMapOutput() ScalingGroupMapOutput
	ToScalingGroupMapOutputWithContext(context.Context) ScalingGroupMapOutput
}

type ScalingGroupMap map[string]ScalingGroupInput

func (ScalingGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ScalingGroup)(nil)).Elem()
}

func (i ScalingGroupMap) ToScalingGroupMapOutput() ScalingGroupMapOutput {
	return i.ToScalingGroupMapOutputWithContext(context.Background())
}

func (i ScalingGroupMap) ToScalingGroupMapOutputWithContext(ctx context.Context) ScalingGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalingGroupMapOutput)
}

type ScalingGroupOutput struct{ *pulumi.OutputState }

func (ScalingGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScalingGroup)(nil)).Elem()
}

func (o ScalingGroupOutput) ToScalingGroupOutput() ScalingGroupOutput {
	return o
}

func (o ScalingGroupOutput) ToScalingGroupOutputWithContext(ctx context.Context) ScalingGroupOutput {
	return o
}

// The scaling configuration id which used by the scaling group.
func (o ScalingGroupOutput) ActiveScalingConfigurationId() pulumi.StringOutput {
	return o.ApplyT(func(v *ScalingGroup) pulumi.StringOutput { return v.ActiveScalingConfigurationId }).(pulumi.StringOutput)
}

// The create time of the scaling group.
func (o ScalingGroupOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ScalingGroup) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The list of db instance ids.
func (o ScalingGroupOutput) DbInstanceIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ScalingGroup) pulumi.StringArrayOutput { return v.DbInstanceIds }).(pulumi.StringArrayOutput)
}

// The default cooldown interval of the scaling group. Default value: 300.
func (o ScalingGroupOutput) DefaultCooldown() pulumi.IntOutput {
	return o.ApplyT(func(v *ScalingGroup) pulumi.IntOutput { return v.DefaultCooldown }).(pulumi.IntOutput)
}

// The desire instance number of the scaling group.
func (o ScalingGroupOutput) DesireInstanceNumber() pulumi.IntOutput {
	return o.ApplyT(func(v *ScalingGroup) pulumi.IntOutput { return v.DesireInstanceNumber }).(pulumi.IntOutput)
}

// The instance terminate policy of the scaling group. Valid values: OldestInstance, NewestInstance, OldestScalingConfigurationWithOldestInstance, OldestScalingConfigurationWithNewestInstance. Default value: OldestScalingConfigurationWithOldestInstance.
func (o ScalingGroupOutput) InstanceTerminatePolicy() pulumi.StringOutput {
	return o.ApplyT(func(v *ScalingGroup) pulumi.StringOutput { return v.InstanceTerminatePolicy }).(pulumi.StringOutput)
}

// The ID of the launch template bound to the scaling group.
func (o ScalingGroupOutput) LaunchTemplateId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingGroup) pulumi.StringPtrOutput { return v.LaunchTemplateId }).(pulumi.StringPtrOutput)
}

// The version of the launch template bound to the scaling group.
func (o ScalingGroupOutput) LaunchTemplateVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingGroup) pulumi.StringPtrOutput { return v.LaunchTemplateVersion }).(pulumi.StringPtrOutput)
}

// The lifecycle state of the scaling group.
func (o ScalingGroupOutput) LifecycleState() pulumi.StringOutput {
	return o.ApplyT(func(v *ScalingGroup) pulumi.StringOutput { return v.LifecycleState }).(pulumi.StringOutput)
}

// The max instance number of the scaling group.
func (o ScalingGroupOutput) MaxInstanceNumber() pulumi.IntOutput {
	return o.ApplyT(func(v *ScalingGroup) pulumi.IntOutput { return v.MaxInstanceNumber }).(pulumi.IntOutput)
}

// The min instance number of the scaling group.
func (o ScalingGroupOutput) MinInstanceNumber() pulumi.IntOutput {
	return o.ApplyT(func(v *ScalingGroup) pulumi.IntOutput { return v.MinInstanceNumber }).(pulumi.IntOutput)
}

// The multi az policy of the scaling group. Valid values: PRIORITY, BALANCE. Default value: PRIORITY.
func (o ScalingGroupOutput) MultiAzPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v *ScalingGroup) pulumi.StringOutput { return v.MultiAzPolicy }).(pulumi.StringOutput)
}

// The id of the scaling group.
func (o ScalingGroupOutput) ScalingGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *ScalingGroup) pulumi.StringOutput { return v.ScalingGroupId }).(pulumi.StringOutput)
}

// The name of the scaling group.
func (o ScalingGroupOutput) ScalingGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *ScalingGroup) pulumi.StringOutput { return v.ScalingGroupName }).(pulumi.StringOutput)
}

// The load balancer server group attributes of the scaling group.
func (o ScalingGroupOutput) ServerGroupAttributes() ScalingGroupServerGroupAttributeArrayOutput {
	return o.ApplyT(func(v *ScalingGroup) ScalingGroupServerGroupAttributeArrayOutput { return v.ServerGroupAttributes }).(ScalingGroupServerGroupAttributeArrayOutput)
}

// The list of the subnet id to which the ENI is connected.
func (o ScalingGroupOutput) SubnetIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ScalingGroup) pulumi.StringArrayOutput { return v.SubnetIds }).(pulumi.StringArrayOutput)
}

// The total instance count of the scaling group.
func (o ScalingGroupOutput) TotalInstanceCount() pulumi.IntOutput {
	return o.ApplyT(func(v *ScalingGroup) pulumi.IntOutput { return v.TotalInstanceCount }).(pulumi.IntOutput)
}

// The create time of the scaling group.
func (o ScalingGroupOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ScalingGroup) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

// The VPC id of the scaling group.
func (o ScalingGroupOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *ScalingGroup) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

type ScalingGroupArrayOutput struct{ *pulumi.OutputState }

func (ScalingGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ScalingGroup)(nil)).Elem()
}

func (o ScalingGroupArrayOutput) ToScalingGroupArrayOutput() ScalingGroupArrayOutput {
	return o
}

func (o ScalingGroupArrayOutput) ToScalingGroupArrayOutputWithContext(ctx context.Context) ScalingGroupArrayOutput {
	return o
}

func (o ScalingGroupArrayOutput) Index(i pulumi.IntInput) ScalingGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ScalingGroup {
		return vs[0].([]*ScalingGroup)[vs[1].(int)]
	}).(ScalingGroupOutput)
}

type ScalingGroupMapOutput struct{ *pulumi.OutputState }

func (ScalingGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ScalingGroup)(nil)).Elem()
}

func (o ScalingGroupMapOutput) ToScalingGroupMapOutput() ScalingGroupMapOutput {
	return o
}

func (o ScalingGroupMapOutput) ToScalingGroupMapOutputWithContext(ctx context.Context) ScalingGroupMapOutput {
	return o
}

func (o ScalingGroupMapOutput) MapIndex(k pulumi.StringInput) ScalingGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ScalingGroup {
		return vs[0].(map[string]*ScalingGroup)[vs[1].(string)]
	}).(ScalingGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ScalingGroupInput)(nil)).Elem(), &ScalingGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScalingGroupArrayInput)(nil)).Elem(), ScalingGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScalingGroupMapInput)(nil)).Elem(), ScalingGroupMap{})
	pulumi.RegisterOutputType(ScalingGroupOutput{})
	pulumi.RegisterOutputType(ScalingGroupArrayOutput{})
	pulumi.RegisterOutputType(ScalingGroupMapOutput{})
}