// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ebs

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Provides a resource to manage ebs auto snapshot policy
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/ebs"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ebs.NewAutoSnapshotPolicy(ctx, "foo", &ebs.AutoSnapshotPolicyArgs{
//				AutoSnapshotPolicyName: pulumi.String("acc-test-auto-snapshot-policy"),
//				ProjectName:            pulumi.String("default"),
//				RepeatWeekdays: pulumi.StringArray{
//					pulumi.String("2"),
//					pulumi.String("6"),
//				},
//				RetentionDays: -1,
//				Tags: ebs.AutoSnapshotPolicyTagArray{
//					&ebs.AutoSnapshotPolicyTagArgs{
//						Key:   pulumi.String("k1"),
//						Value: pulumi.String("v1"),
//					},
//				},
//				TimePoints: pulumi.StringArray{
//					pulumi.String("1"),
//					pulumi.String("5"),
//					pulumi.String("9"),
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
// EbsAutoSnapshotPolicy can be imported using the id, e.g.
//
// ```sh
// $ pulumi import volcengine:ebs/autoSnapshotPolicy:AutoSnapshotPolicy default resource_id
// ```
type AutoSnapshotPolicy struct {
	pulumi.CustomResourceState

	// The name of the auto snapshot policy.
	AutoSnapshotPolicyName pulumi.StringOutput `pulumi:"autoSnapshotPolicyName"`
	// The creation time of the auto snapshot policy.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The project name of the auto snapshot policy.
	ProjectName pulumi.StringOutput `pulumi:"projectName"`
	// Create snapshots repeatedly on a daily basis, with intervals of a certain number of days between each snapshot. The value range is `1-30`. Only one of `repeat_weekdays, repeatDays` can be specified.
	RepeatDays pulumi.IntPtrOutput `pulumi:"repeatDays"`
	// The date of creating snapshot repeatedly by week. The value range is `1-7`, for example, 1 represents Monday. Only one of `repeat_weekdays, repeatDays` can be specified.
	RepeatWeekdays pulumi.StringArrayOutput `pulumi:"repeatWeekdays"`
	// The retention days of the auto snapshot. Valid values: -1 and 1~65536. `-1` means permanently preserving the snapshot.
	RetentionDays pulumi.IntOutput `pulumi:"retentionDays"`
	// The status of the auto snapshot policy.
	Status pulumi.StringOutput `pulumi:"status"`
	// Tags.
	Tags AutoSnapshotPolicyTagArrayOutput `pulumi:"tags"`
	// The creation time points of the auto snapshot policy. The value range is `0~23`, representing a total of 24 time points from 00:00 to 23:00, for example, 1 represents 01:00.
	TimePoints pulumi.StringArrayOutput `pulumi:"timePoints"`
	// The updated time of the auto snapshot policy.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// The number of volumes associated with the auto snapshot policy.
	VolumeNums pulumi.IntOutput `pulumi:"volumeNums"`
}

// NewAutoSnapshotPolicy registers a new resource with the given unique name, arguments, and options.
func NewAutoSnapshotPolicy(ctx *pulumi.Context,
	name string, args *AutoSnapshotPolicyArgs, opts ...pulumi.ResourceOption) (*AutoSnapshotPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AutoSnapshotPolicyName == nil {
		return nil, errors.New("invalid value for required argument 'AutoSnapshotPolicyName'")
	}
	if args.RetentionDays == nil {
		return nil, errors.New("invalid value for required argument 'RetentionDays'")
	}
	if args.TimePoints == nil {
		return nil, errors.New("invalid value for required argument 'TimePoints'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AutoSnapshotPolicy
	err := ctx.RegisterResource("volcengine:ebs/autoSnapshotPolicy:AutoSnapshotPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAutoSnapshotPolicy gets an existing AutoSnapshotPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAutoSnapshotPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AutoSnapshotPolicyState, opts ...pulumi.ResourceOption) (*AutoSnapshotPolicy, error) {
	var resource AutoSnapshotPolicy
	err := ctx.ReadResource("volcengine:ebs/autoSnapshotPolicy:AutoSnapshotPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AutoSnapshotPolicy resources.
type autoSnapshotPolicyState struct {
	// The name of the auto snapshot policy.
	AutoSnapshotPolicyName *string `pulumi:"autoSnapshotPolicyName"`
	// The creation time of the auto snapshot policy.
	CreatedAt *string `pulumi:"createdAt"`
	// The project name of the auto snapshot policy.
	ProjectName *string `pulumi:"projectName"`
	// Create snapshots repeatedly on a daily basis, with intervals of a certain number of days between each snapshot. The value range is `1-30`. Only one of `repeat_weekdays, repeatDays` can be specified.
	RepeatDays *int `pulumi:"repeatDays"`
	// The date of creating snapshot repeatedly by week. The value range is `1-7`, for example, 1 represents Monday. Only one of `repeat_weekdays, repeatDays` can be specified.
	RepeatWeekdays []string `pulumi:"repeatWeekdays"`
	// The retention days of the auto snapshot. Valid values: -1 and 1~65536. `-1` means permanently preserving the snapshot.
	RetentionDays *int `pulumi:"retentionDays"`
	// The status of the auto snapshot policy.
	Status *string `pulumi:"status"`
	// Tags.
	Tags []AutoSnapshotPolicyTag `pulumi:"tags"`
	// The creation time points of the auto snapshot policy. The value range is `0~23`, representing a total of 24 time points from 00:00 to 23:00, for example, 1 represents 01:00.
	TimePoints []string `pulumi:"timePoints"`
	// The updated time of the auto snapshot policy.
	UpdatedAt *string `pulumi:"updatedAt"`
	// The number of volumes associated with the auto snapshot policy.
	VolumeNums *int `pulumi:"volumeNums"`
}

type AutoSnapshotPolicyState struct {
	// The name of the auto snapshot policy.
	AutoSnapshotPolicyName pulumi.StringPtrInput
	// The creation time of the auto snapshot policy.
	CreatedAt pulumi.StringPtrInput
	// The project name of the auto snapshot policy.
	ProjectName pulumi.StringPtrInput
	// Create snapshots repeatedly on a daily basis, with intervals of a certain number of days between each snapshot. The value range is `1-30`. Only one of `repeat_weekdays, repeatDays` can be specified.
	RepeatDays pulumi.IntPtrInput
	// The date of creating snapshot repeatedly by week. The value range is `1-7`, for example, 1 represents Monday. Only one of `repeat_weekdays, repeatDays` can be specified.
	RepeatWeekdays pulumi.StringArrayInput
	// The retention days of the auto snapshot. Valid values: -1 and 1~65536. `-1` means permanently preserving the snapshot.
	RetentionDays pulumi.IntPtrInput
	// The status of the auto snapshot policy.
	Status pulumi.StringPtrInput
	// Tags.
	Tags AutoSnapshotPolicyTagArrayInput
	// The creation time points of the auto snapshot policy. The value range is `0~23`, representing a total of 24 time points from 00:00 to 23:00, for example, 1 represents 01:00.
	TimePoints pulumi.StringArrayInput
	// The updated time of the auto snapshot policy.
	UpdatedAt pulumi.StringPtrInput
	// The number of volumes associated with the auto snapshot policy.
	VolumeNums pulumi.IntPtrInput
}

func (AutoSnapshotPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*autoSnapshotPolicyState)(nil)).Elem()
}

type autoSnapshotPolicyArgs struct {
	// The name of the auto snapshot policy.
	AutoSnapshotPolicyName string `pulumi:"autoSnapshotPolicyName"`
	// The project name of the auto snapshot policy.
	ProjectName *string `pulumi:"projectName"`
	// Create snapshots repeatedly on a daily basis, with intervals of a certain number of days between each snapshot. The value range is `1-30`. Only one of `repeat_weekdays, repeatDays` can be specified.
	RepeatDays *int `pulumi:"repeatDays"`
	// The date of creating snapshot repeatedly by week. The value range is `1-7`, for example, 1 represents Monday. Only one of `repeat_weekdays, repeatDays` can be specified.
	RepeatWeekdays []string `pulumi:"repeatWeekdays"`
	// The retention days of the auto snapshot. Valid values: -1 and 1~65536. `-1` means permanently preserving the snapshot.
	RetentionDays int `pulumi:"retentionDays"`
	// Tags.
	Tags []AutoSnapshotPolicyTag `pulumi:"tags"`
	// The creation time points of the auto snapshot policy. The value range is `0~23`, representing a total of 24 time points from 00:00 to 23:00, for example, 1 represents 01:00.
	TimePoints []string `pulumi:"timePoints"`
}

// The set of arguments for constructing a AutoSnapshotPolicy resource.
type AutoSnapshotPolicyArgs struct {
	// The name of the auto snapshot policy.
	AutoSnapshotPolicyName pulumi.StringInput
	// The project name of the auto snapshot policy.
	ProjectName pulumi.StringPtrInput
	// Create snapshots repeatedly on a daily basis, with intervals of a certain number of days between each snapshot. The value range is `1-30`. Only one of `repeat_weekdays, repeatDays` can be specified.
	RepeatDays pulumi.IntPtrInput
	// The date of creating snapshot repeatedly by week. The value range is `1-7`, for example, 1 represents Monday. Only one of `repeat_weekdays, repeatDays` can be specified.
	RepeatWeekdays pulumi.StringArrayInput
	// The retention days of the auto snapshot. Valid values: -1 and 1~65536. `-1` means permanently preserving the snapshot.
	RetentionDays pulumi.IntInput
	// Tags.
	Tags AutoSnapshotPolicyTagArrayInput
	// The creation time points of the auto snapshot policy. The value range is `0~23`, representing a total of 24 time points from 00:00 to 23:00, for example, 1 represents 01:00.
	TimePoints pulumi.StringArrayInput
}

func (AutoSnapshotPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*autoSnapshotPolicyArgs)(nil)).Elem()
}

type AutoSnapshotPolicyInput interface {
	pulumi.Input

	ToAutoSnapshotPolicyOutput() AutoSnapshotPolicyOutput
	ToAutoSnapshotPolicyOutputWithContext(ctx context.Context) AutoSnapshotPolicyOutput
}

func (*AutoSnapshotPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoSnapshotPolicy)(nil)).Elem()
}

func (i *AutoSnapshotPolicy) ToAutoSnapshotPolicyOutput() AutoSnapshotPolicyOutput {
	return i.ToAutoSnapshotPolicyOutputWithContext(context.Background())
}

func (i *AutoSnapshotPolicy) ToAutoSnapshotPolicyOutputWithContext(ctx context.Context) AutoSnapshotPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoSnapshotPolicyOutput)
}

// AutoSnapshotPolicyArrayInput is an input type that accepts AutoSnapshotPolicyArray and AutoSnapshotPolicyArrayOutput values.
// You can construct a concrete instance of `AutoSnapshotPolicyArrayInput` via:
//
//	AutoSnapshotPolicyArray{ AutoSnapshotPolicyArgs{...} }
type AutoSnapshotPolicyArrayInput interface {
	pulumi.Input

	ToAutoSnapshotPolicyArrayOutput() AutoSnapshotPolicyArrayOutput
	ToAutoSnapshotPolicyArrayOutputWithContext(context.Context) AutoSnapshotPolicyArrayOutput
}

type AutoSnapshotPolicyArray []AutoSnapshotPolicyInput

func (AutoSnapshotPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AutoSnapshotPolicy)(nil)).Elem()
}

func (i AutoSnapshotPolicyArray) ToAutoSnapshotPolicyArrayOutput() AutoSnapshotPolicyArrayOutput {
	return i.ToAutoSnapshotPolicyArrayOutputWithContext(context.Background())
}

func (i AutoSnapshotPolicyArray) ToAutoSnapshotPolicyArrayOutputWithContext(ctx context.Context) AutoSnapshotPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoSnapshotPolicyArrayOutput)
}

// AutoSnapshotPolicyMapInput is an input type that accepts AutoSnapshotPolicyMap and AutoSnapshotPolicyMapOutput values.
// You can construct a concrete instance of `AutoSnapshotPolicyMapInput` via:
//
//	AutoSnapshotPolicyMap{ "key": AutoSnapshotPolicyArgs{...} }
type AutoSnapshotPolicyMapInput interface {
	pulumi.Input

	ToAutoSnapshotPolicyMapOutput() AutoSnapshotPolicyMapOutput
	ToAutoSnapshotPolicyMapOutputWithContext(context.Context) AutoSnapshotPolicyMapOutput
}

type AutoSnapshotPolicyMap map[string]AutoSnapshotPolicyInput

func (AutoSnapshotPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AutoSnapshotPolicy)(nil)).Elem()
}

func (i AutoSnapshotPolicyMap) ToAutoSnapshotPolicyMapOutput() AutoSnapshotPolicyMapOutput {
	return i.ToAutoSnapshotPolicyMapOutputWithContext(context.Background())
}

func (i AutoSnapshotPolicyMap) ToAutoSnapshotPolicyMapOutputWithContext(ctx context.Context) AutoSnapshotPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoSnapshotPolicyMapOutput)
}

type AutoSnapshotPolicyOutput struct{ *pulumi.OutputState }

func (AutoSnapshotPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoSnapshotPolicy)(nil)).Elem()
}

func (o AutoSnapshotPolicyOutput) ToAutoSnapshotPolicyOutput() AutoSnapshotPolicyOutput {
	return o
}

func (o AutoSnapshotPolicyOutput) ToAutoSnapshotPolicyOutputWithContext(ctx context.Context) AutoSnapshotPolicyOutput {
	return o
}

// The name of the auto snapshot policy.
func (o AutoSnapshotPolicyOutput) AutoSnapshotPolicyName() pulumi.StringOutput {
	return o.ApplyT(func(v *AutoSnapshotPolicy) pulumi.StringOutput { return v.AutoSnapshotPolicyName }).(pulumi.StringOutput)
}

// The creation time of the auto snapshot policy.
func (o AutoSnapshotPolicyOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *AutoSnapshotPolicy) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The project name of the auto snapshot policy.
func (o AutoSnapshotPolicyOutput) ProjectName() pulumi.StringOutput {
	return o.ApplyT(func(v *AutoSnapshotPolicy) pulumi.StringOutput { return v.ProjectName }).(pulumi.StringOutput)
}

// Create snapshots repeatedly on a daily basis, with intervals of a certain number of days between each snapshot. The value range is `1-30`. Only one of `repeat_weekdays, repeatDays` can be specified.
func (o AutoSnapshotPolicyOutput) RepeatDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AutoSnapshotPolicy) pulumi.IntPtrOutput { return v.RepeatDays }).(pulumi.IntPtrOutput)
}

// The date of creating snapshot repeatedly by week. The value range is `1-7`, for example, 1 represents Monday. Only one of `repeat_weekdays, repeatDays` can be specified.
func (o AutoSnapshotPolicyOutput) RepeatWeekdays() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AutoSnapshotPolicy) pulumi.StringArrayOutput { return v.RepeatWeekdays }).(pulumi.StringArrayOutput)
}

// The retention days of the auto snapshot. Valid values: -1 and 1~65536. `-1` means permanently preserving the snapshot.
func (o AutoSnapshotPolicyOutput) RetentionDays() pulumi.IntOutput {
	return o.ApplyT(func(v *AutoSnapshotPolicy) pulumi.IntOutput { return v.RetentionDays }).(pulumi.IntOutput)
}

// The status of the auto snapshot policy.
func (o AutoSnapshotPolicyOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *AutoSnapshotPolicy) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Tags.
func (o AutoSnapshotPolicyOutput) Tags() AutoSnapshotPolicyTagArrayOutput {
	return o.ApplyT(func(v *AutoSnapshotPolicy) AutoSnapshotPolicyTagArrayOutput { return v.Tags }).(AutoSnapshotPolicyTagArrayOutput)
}

// The creation time points of the auto snapshot policy. The value range is `0~23`, representing a total of 24 time points from 00:00 to 23:00, for example, 1 represents 01:00.
func (o AutoSnapshotPolicyOutput) TimePoints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AutoSnapshotPolicy) pulumi.StringArrayOutput { return v.TimePoints }).(pulumi.StringArrayOutput)
}

// The updated time of the auto snapshot policy.
func (o AutoSnapshotPolicyOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *AutoSnapshotPolicy) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

// The number of volumes associated with the auto snapshot policy.
func (o AutoSnapshotPolicyOutput) VolumeNums() pulumi.IntOutput {
	return o.ApplyT(func(v *AutoSnapshotPolicy) pulumi.IntOutput { return v.VolumeNums }).(pulumi.IntOutput)
}

type AutoSnapshotPolicyArrayOutput struct{ *pulumi.OutputState }

func (AutoSnapshotPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AutoSnapshotPolicy)(nil)).Elem()
}

func (o AutoSnapshotPolicyArrayOutput) ToAutoSnapshotPolicyArrayOutput() AutoSnapshotPolicyArrayOutput {
	return o
}

func (o AutoSnapshotPolicyArrayOutput) ToAutoSnapshotPolicyArrayOutputWithContext(ctx context.Context) AutoSnapshotPolicyArrayOutput {
	return o
}

func (o AutoSnapshotPolicyArrayOutput) Index(i pulumi.IntInput) AutoSnapshotPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AutoSnapshotPolicy {
		return vs[0].([]*AutoSnapshotPolicy)[vs[1].(int)]
	}).(AutoSnapshotPolicyOutput)
}

type AutoSnapshotPolicyMapOutput struct{ *pulumi.OutputState }

func (AutoSnapshotPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AutoSnapshotPolicy)(nil)).Elem()
}

func (o AutoSnapshotPolicyMapOutput) ToAutoSnapshotPolicyMapOutput() AutoSnapshotPolicyMapOutput {
	return o
}

func (o AutoSnapshotPolicyMapOutput) ToAutoSnapshotPolicyMapOutputWithContext(ctx context.Context) AutoSnapshotPolicyMapOutput {
	return o
}

func (o AutoSnapshotPolicyMapOutput) MapIndex(k pulumi.StringInput) AutoSnapshotPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AutoSnapshotPolicy {
		return vs[0].(map[string]*AutoSnapshotPolicy)[vs[1].(string)]
	}).(AutoSnapshotPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AutoSnapshotPolicyInput)(nil)).Elem(), &AutoSnapshotPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*AutoSnapshotPolicyArrayInput)(nil)).Elem(), AutoSnapshotPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AutoSnapshotPolicyMapInput)(nil)).Elem(), AutoSnapshotPolicyMap{})
	pulumi.RegisterOutputType(AutoSnapshotPolicyOutput{})
	pulumi.RegisterOutputType(AutoSnapshotPolicyArrayOutput{})
	pulumi.RegisterOutputType(AutoSnapshotPolicyMapOutput{})
}