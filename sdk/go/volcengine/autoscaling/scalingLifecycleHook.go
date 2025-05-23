// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package autoscaling

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Provides a resource to manage scaling lifecycle hook
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/autoscaling"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/ecs"
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
//				VpcName:   pulumi.String("acc-test-vpc"),
//				CidrBlock: pulumi.String("172.16.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			fooSubnet, err := vpc.NewSubnet(ctx, "fooSubnet", &vpc.SubnetArgs{
//				SubnetName: pulumi.String("acc-test-subnet"),
//				CidrBlock:  pulumi.String("172.16.0.0/24"),
//				ZoneId:     pulumi.String(fooZones.Zones[0].Id),
//				VpcId:      fooVpc.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ecs.NewCommand(ctx, "fooCommand", &ecs.CommandArgs{
//				Description:    pulumi.String("tf"),
//				WorkingDir:     pulumi.String("/home"),
//				Username:       pulumi.String("root"),
//				Timeout:        pulumi.Int(100),
//				CommandContent: pulumi.String("IyEvYmluL2Jhc2gKCgplY2hvICJvcGVyYXRpb24gc3VjY2VzcyEi"),
//			})
//			if err != nil {
//				return err
//			}
//			fooScalingGroup, err := autoscaling.NewScalingGroup(ctx, "fooScalingGroup", &autoscaling.ScalingGroupArgs{
//				ScalingGroupName: pulumi.String("acc-test-scaling-group-lifecycle"),
//				SubnetIds: pulumi.StringArray{
//					fooSubnet.ID(),
//				},
//				MultiAzPolicy:           pulumi.String("BALANCE"),
//				DesireInstanceNumber:    pulumi.Int(0),
//				MinInstanceNumber:       pulumi.Int(0),
//				MaxInstanceNumber:       pulumi.Int(1),
//				InstanceTerminatePolicy: pulumi.String("OldestInstance"),
//				DefaultCooldown:         pulumi.Int(10),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = autoscaling.NewScalingLifecycleHook(ctx, "fooScalingLifecycleHook", &autoscaling.ScalingLifecycleHookArgs{
//				LifecycleHookName:    pulumi.String("acc-test-lifecycle"),
//				LifecycleHookPolicy:  pulumi.String("ROLLBACK"),
//				LifecycleHookTimeout: pulumi.Int(300),
//				LifecycleHookType:    pulumi.String("SCALE_OUT"),
//				ScalingGroupId:       fooScalingGroup.ID(),
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
// ScalingLifecycleHook can be imported using the ScalingGroupId:LifecycleHookId, e.g.
//
// ```sh
// $ pulumi import volcengine:autoscaling/scalingLifecycleHook:ScalingLifecycleHook default scg-yblfbfhy7agh9zn72iaz:sgh-ybqholahe4gso0ee88sd
// ```
type ScalingLifecycleHook struct {
	pulumi.CustomResourceState

	// Batch job command.
	LifecycleCommand ScalingLifecycleHookLifecycleCommandPtrOutput `pulumi:"lifecycleCommand"`
	// The id of the lifecycle hook.
	LifecycleHookId pulumi.StringOutput `pulumi:"lifecycleHookId"`
	// The name of the lifecycle hook.
	LifecycleHookName pulumi.StringOutput `pulumi:"lifecycleHookName"`
	// The policy of the lifecycle hook. Valid values: CONTINUE, REJECT, ROLLBACK.
	LifecycleHookPolicy pulumi.StringOutput `pulumi:"lifecycleHookPolicy"`
	// The timeout of the lifecycle hook.
	LifecycleHookTimeout pulumi.IntOutput `pulumi:"lifecycleHookTimeout"`
	// The type of the lifecycle hook. Valid values: SCALE_IN, SCALE_OUT.
	LifecycleHookType pulumi.StringOutput `pulumi:"lifecycleHookType"`
	// The id of the scaling group.
	ScalingGroupId pulumi.StringOutput `pulumi:"scalingGroupId"`
}

// NewScalingLifecycleHook registers a new resource with the given unique name, arguments, and options.
func NewScalingLifecycleHook(ctx *pulumi.Context,
	name string, args *ScalingLifecycleHookArgs, opts ...pulumi.ResourceOption) (*ScalingLifecycleHook, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LifecycleHookName == nil {
		return nil, errors.New("invalid value for required argument 'LifecycleHookName'")
	}
	if args.LifecycleHookPolicy == nil {
		return nil, errors.New("invalid value for required argument 'LifecycleHookPolicy'")
	}
	if args.LifecycleHookTimeout == nil {
		return nil, errors.New("invalid value for required argument 'LifecycleHookTimeout'")
	}
	if args.LifecycleHookType == nil {
		return nil, errors.New("invalid value for required argument 'LifecycleHookType'")
	}
	if args.ScalingGroupId == nil {
		return nil, errors.New("invalid value for required argument 'ScalingGroupId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ScalingLifecycleHook
	err := ctx.RegisterResource("volcengine:autoscaling/scalingLifecycleHook:ScalingLifecycleHook", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetScalingLifecycleHook gets an existing ScalingLifecycleHook resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetScalingLifecycleHook(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScalingLifecycleHookState, opts ...pulumi.ResourceOption) (*ScalingLifecycleHook, error) {
	var resource ScalingLifecycleHook
	err := ctx.ReadResource("volcengine:autoscaling/scalingLifecycleHook:ScalingLifecycleHook", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ScalingLifecycleHook resources.
type scalingLifecycleHookState struct {
	// Batch job command.
	LifecycleCommand *ScalingLifecycleHookLifecycleCommand `pulumi:"lifecycleCommand"`
	// The id of the lifecycle hook.
	LifecycleHookId *string `pulumi:"lifecycleHookId"`
	// The name of the lifecycle hook.
	LifecycleHookName *string `pulumi:"lifecycleHookName"`
	// The policy of the lifecycle hook. Valid values: CONTINUE, REJECT, ROLLBACK.
	LifecycleHookPolicy *string `pulumi:"lifecycleHookPolicy"`
	// The timeout of the lifecycle hook.
	LifecycleHookTimeout *int `pulumi:"lifecycleHookTimeout"`
	// The type of the lifecycle hook. Valid values: SCALE_IN, SCALE_OUT.
	LifecycleHookType *string `pulumi:"lifecycleHookType"`
	// The id of the scaling group.
	ScalingGroupId *string `pulumi:"scalingGroupId"`
}

type ScalingLifecycleHookState struct {
	// Batch job command.
	LifecycleCommand ScalingLifecycleHookLifecycleCommandPtrInput
	// The id of the lifecycle hook.
	LifecycleHookId pulumi.StringPtrInput
	// The name of the lifecycle hook.
	LifecycleHookName pulumi.StringPtrInput
	// The policy of the lifecycle hook. Valid values: CONTINUE, REJECT, ROLLBACK.
	LifecycleHookPolicy pulumi.StringPtrInput
	// The timeout of the lifecycle hook.
	LifecycleHookTimeout pulumi.IntPtrInput
	// The type of the lifecycle hook. Valid values: SCALE_IN, SCALE_OUT.
	LifecycleHookType pulumi.StringPtrInput
	// The id of the scaling group.
	ScalingGroupId pulumi.StringPtrInput
}

func (ScalingLifecycleHookState) ElementType() reflect.Type {
	return reflect.TypeOf((*scalingLifecycleHookState)(nil)).Elem()
}

type scalingLifecycleHookArgs struct {
	// Batch job command.
	LifecycleCommand *ScalingLifecycleHookLifecycleCommand `pulumi:"lifecycleCommand"`
	// The name of the lifecycle hook.
	LifecycleHookName string `pulumi:"lifecycleHookName"`
	// The policy of the lifecycle hook. Valid values: CONTINUE, REJECT, ROLLBACK.
	LifecycleHookPolicy string `pulumi:"lifecycleHookPolicy"`
	// The timeout of the lifecycle hook.
	LifecycleHookTimeout int `pulumi:"lifecycleHookTimeout"`
	// The type of the lifecycle hook. Valid values: SCALE_IN, SCALE_OUT.
	LifecycleHookType string `pulumi:"lifecycleHookType"`
	// The id of the scaling group.
	ScalingGroupId string `pulumi:"scalingGroupId"`
}

// The set of arguments for constructing a ScalingLifecycleHook resource.
type ScalingLifecycleHookArgs struct {
	// Batch job command.
	LifecycleCommand ScalingLifecycleHookLifecycleCommandPtrInput
	// The name of the lifecycle hook.
	LifecycleHookName pulumi.StringInput
	// The policy of the lifecycle hook. Valid values: CONTINUE, REJECT, ROLLBACK.
	LifecycleHookPolicy pulumi.StringInput
	// The timeout of the lifecycle hook.
	LifecycleHookTimeout pulumi.IntInput
	// The type of the lifecycle hook. Valid values: SCALE_IN, SCALE_OUT.
	LifecycleHookType pulumi.StringInput
	// The id of the scaling group.
	ScalingGroupId pulumi.StringInput
}

func (ScalingLifecycleHookArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scalingLifecycleHookArgs)(nil)).Elem()
}

type ScalingLifecycleHookInput interface {
	pulumi.Input

	ToScalingLifecycleHookOutput() ScalingLifecycleHookOutput
	ToScalingLifecycleHookOutputWithContext(ctx context.Context) ScalingLifecycleHookOutput
}

func (*ScalingLifecycleHook) ElementType() reflect.Type {
	return reflect.TypeOf((**ScalingLifecycleHook)(nil)).Elem()
}

func (i *ScalingLifecycleHook) ToScalingLifecycleHookOutput() ScalingLifecycleHookOutput {
	return i.ToScalingLifecycleHookOutputWithContext(context.Background())
}

func (i *ScalingLifecycleHook) ToScalingLifecycleHookOutputWithContext(ctx context.Context) ScalingLifecycleHookOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalingLifecycleHookOutput)
}

// ScalingLifecycleHookArrayInput is an input type that accepts ScalingLifecycleHookArray and ScalingLifecycleHookArrayOutput values.
// You can construct a concrete instance of `ScalingLifecycleHookArrayInput` via:
//
//	ScalingLifecycleHookArray{ ScalingLifecycleHookArgs{...} }
type ScalingLifecycleHookArrayInput interface {
	pulumi.Input

	ToScalingLifecycleHookArrayOutput() ScalingLifecycleHookArrayOutput
	ToScalingLifecycleHookArrayOutputWithContext(context.Context) ScalingLifecycleHookArrayOutput
}

type ScalingLifecycleHookArray []ScalingLifecycleHookInput

func (ScalingLifecycleHookArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ScalingLifecycleHook)(nil)).Elem()
}

func (i ScalingLifecycleHookArray) ToScalingLifecycleHookArrayOutput() ScalingLifecycleHookArrayOutput {
	return i.ToScalingLifecycleHookArrayOutputWithContext(context.Background())
}

func (i ScalingLifecycleHookArray) ToScalingLifecycleHookArrayOutputWithContext(ctx context.Context) ScalingLifecycleHookArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalingLifecycleHookArrayOutput)
}

// ScalingLifecycleHookMapInput is an input type that accepts ScalingLifecycleHookMap and ScalingLifecycleHookMapOutput values.
// You can construct a concrete instance of `ScalingLifecycleHookMapInput` via:
//
//	ScalingLifecycleHookMap{ "key": ScalingLifecycleHookArgs{...} }
type ScalingLifecycleHookMapInput interface {
	pulumi.Input

	ToScalingLifecycleHookMapOutput() ScalingLifecycleHookMapOutput
	ToScalingLifecycleHookMapOutputWithContext(context.Context) ScalingLifecycleHookMapOutput
}

type ScalingLifecycleHookMap map[string]ScalingLifecycleHookInput

func (ScalingLifecycleHookMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ScalingLifecycleHook)(nil)).Elem()
}

func (i ScalingLifecycleHookMap) ToScalingLifecycleHookMapOutput() ScalingLifecycleHookMapOutput {
	return i.ToScalingLifecycleHookMapOutputWithContext(context.Background())
}

func (i ScalingLifecycleHookMap) ToScalingLifecycleHookMapOutputWithContext(ctx context.Context) ScalingLifecycleHookMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalingLifecycleHookMapOutput)
}

type ScalingLifecycleHookOutput struct{ *pulumi.OutputState }

func (ScalingLifecycleHookOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScalingLifecycleHook)(nil)).Elem()
}

func (o ScalingLifecycleHookOutput) ToScalingLifecycleHookOutput() ScalingLifecycleHookOutput {
	return o
}

func (o ScalingLifecycleHookOutput) ToScalingLifecycleHookOutputWithContext(ctx context.Context) ScalingLifecycleHookOutput {
	return o
}

// Batch job command.
func (o ScalingLifecycleHookOutput) LifecycleCommand() ScalingLifecycleHookLifecycleCommandPtrOutput {
	return o.ApplyT(func(v *ScalingLifecycleHook) ScalingLifecycleHookLifecycleCommandPtrOutput { return v.LifecycleCommand }).(ScalingLifecycleHookLifecycleCommandPtrOutput)
}

// The id of the lifecycle hook.
func (o ScalingLifecycleHookOutput) LifecycleHookId() pulumi.StringOutput {
	return o.ApplyT(func(v *ScalingLifecycleHook) pulumi.StringOutput { return v.LifecycleHookId }).(pulumi.StringOutput)
}

// The name of the lifecycle hook.
func (o ScalingLifecycleHookOutput) LifecycleHookName() pulumi.StringOutput {
	return o.ApplyT(func(v *ScalingLifecycleHook) pulumi.StringOutput { return v.LifecycleHookName }).(pulumi.StringOutput)
}

// The policy of the lifecycle hook. Valid values: CONTINUE, REJECT, ROLLBACK.
func (o ScalingLifecycleHookOutput) LifecycleHookPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v *ScalingLifecycleHook) pulumi.StringOutput { return v.LifecycleHookPolicy }).(pulumi.StringOutput)
}

// The timeout of the lifecycle hook.
func (o ScalingLifecycleHookOutput) LifecycleHookTimeout() pulumi.IntOutput {
	return o.ApplyT(func(v *ScalingLifecycleHook) pulumi.IntOutput { return v.LifecycleHookTimeout }).(pulumi.IntOutput)
}

// The type of the lifecycle hook. Valid values: SCALE_IN, SCALE_OUT.
func (o ScalingLifecycleHookOutput) LifecycleHookType() pulumi.StringOutput {
	return o.ApplyT(func(v *ScalingLifecycleHook) pulumi.StringOutput { return v.LifecycleHookType }).(pulumi.StringOutput)
}

// The id of the scaling group.
func (o ScalingLifecycleHookOutput) ScalingGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *ScalingLifecycleHook) pulumi.StringOutput { return v.ScalingGroupId }).(pulumi.StringOutput)
}

type ScalingLifecycleHookArrayOutput struct{ *pulumi.OutputState }

func (ScalingLifecycleHookArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ScalingLifecycleHook)(nil)).Elem()
}

func (o ScalingLifecycleHookArrayOutput) ToScalingLifecycleHookArrayOutput() ScalingLifecycleHookArrayOutput {
	return o
}

func (o ScalingLifecycleHookArrayOutput) ToScalingLifecycleHookArrayOutputWithContext(ctx context.Context) ScalingLifecycleHookArrayOutput {
	return o
}

func (o ScalingLifecycleHookArrayOutput) Index(i pulumi.IntInput) ScalingLifecycleHookOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ScalingLifecycleHook {
		return vs[0].([]*ScalingLifecycleHook)[vs[1].(int)]
	}).(ScalingLifecycleHookOutput)
}

type ScalingLifecycleHookMapOutput struct{ *pulumi.OutputState }

func (ScalingLifecycleHookMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ScalingLifecycleHook)(nil)).Elem()
}

func (o ScalingLifecycleHookMapOutput) ToScalingLifecycleHookMapOutput() ScalingLifecycleHookMapOutput {
	return o
}

func (o ScalingLifecycleHookMapOutput) ToScalingLifecycleHookMapOutputWithContext(ctx context.Context) ScalingLifecycleHookMapOutput {
	return o
}

func (o ScalingLifecycleHookMapOutput) MapIndex(k pulumi.StringInput) ScalingLifecycleHookOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ScalingLifecycleHook {
		return vs[0].(map[string]*ScalingLifecycleHook)[vs[1].(string)]
	}).(ScalingLifecycleHookOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ScalingLifecycleHookInput)(nil)).Elem(), &ScalingLifecycleHook{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScalingLifecycleHookArrayInput)(nil)).Elem(), ScalingLifecycleHookArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScalingLifecycleHookMapInput)(nil)).Elem(), ScalingLifecycleHookMap{})
	pulumi.RegisterOutputType(ScalingLifecycleHookOutput{})
	pulumi.RegisterOutputType(ScalingLifecycleHookArrayOutput{})
	pulumi.RegisterOutputType(ScalingLifecycleHookMapOutput{})
}
