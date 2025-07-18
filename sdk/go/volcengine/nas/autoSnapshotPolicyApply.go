// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package nas

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Provides a resource to manage nas auto snapshot policy apply
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/nas"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			fooZones, err := nas.GetZones(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			fooFileSystem, err := nas.NewFileSystem(ctx, "fooFileSystem", &nas.FileSystemArgs{
//				FileSystemName: pulumi.String("acc-test-fs"),
//				Description:    pulumi.String("acc-test"),
//				ZoneId:         pulumi.String(fooZones.Zones[0].Id),
//				Capacity:       pulumi.Int(103),
//				ProjectName:    pulumi.String("default"),
//				Tags: nas.FileSystemTagArray{
//					&nas.FileSystemTagArgs{
//						Key:   pulumi.String("k1"),
//						Value: pulumi.String("v1"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			fooAutoSnapshotPolicy, err := nas.NewAutoSnapshotPolicy(ctx, "fooAutoSnapshotPolicy", &nas.AutoSnapshotPolicyArgs{
//				AutoSnapshotPolicyName: pulumi.String("acc-test-auto_snapshot_policy"),
//				RepeatWeekdays:         pulumi.String("1,3,5,7"),
//				TimePoints:             pulumi.String("0,7,17"),
//				RetentionDays:          pulumi.Int(20),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = nas.NewAutoSnapshotPolicyApply(ctx, "fooAutoSnapshotPolicyApply", &nas.AutoSnapshotPolicyApplyArgs{
//				FileSystemId:         fooFileSystem.ID(),
//				AutoSnapshotPolicyId: fooAutoSnapshotPolicy.ID(),
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
// NasAutoSnapshotPolicyApply can be imported using the auto_snapshot_policy_id:file_system_id, e.g.
//
// ```sh
// $ pulumi import volcengine:nas/autoSnapshotPolicyApply:AutoSnapshotPolicyApply default resource_id
// ```
type AutoSnapshotPolicyApply struct {
	pulumi.CustomResourceState

	// The id of auto snapshot policy.
	AutoSnapshotPolicyId pulumi.StringOutput `pulumi:"autoSnapshotPolicyId"`
	// The id of file system.
	FileSystemId pulumi.StringOutput `pulumi:"fileSystemId"`
}

// NewAutoSnapshotPolicyApply registers a new resource with the given unique name, arguments, and options.
func NewAutoSnapshotPolicyApply(ctx *pulumi.Context,
	name string, args *AutoSnapshotPolicyApplyArgs, opts ...pulumi.ResourceOption) (*AutoSnapshotPolicyApply, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AutoSnapshotPolicyId == nil {
		return nil, errors.New("invalid value for required argument 'AutoSnapshotPolicyId'")
	}
	if args.FileSystemId == nil {
		return nil, errors.New("invalid value for required argument 'FileSystemId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AutoSnapshotPolicyApply
	err := ctx.RegisterResource("volcengine:nas/autoSnapshotPolicyApply:AutoSnapshotPolicyApply", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAutoSnapshotPolicyApply gets an existing AutoSnapshotPolicyApply resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAutoSnapshotPolicyApply(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AutoSnapshotPolicyApplyState, opts ...pulumi.ResourceOption) (*AutoSnapshotPolicyApply, error) {
	var resource AutoSnapshotPolicyApply
	err := ctx.ReadResource("volcengine:nas/autoSnapshotPolicyApply:AutoSnapshotPolicyApply", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AutoSnapshotPolicyApply resources.
type autoSnapshotPolicyApplyState struct {
	// The id of auto snapshot policy.
	AutoSnapshotPolicyId *string `pulumi:"autoSnapshotPolicyId"`
	// The id of file system.
	FileSystemId *string `pulumi:"fileSystemId"`
}

type AutoSnapshotPolicyApplyState struct {
	// The id of auto snapshot policy.
	AutoSnapshotPolicyId pulumi.StringPtrInput
	// The id of file system.
	FileSystemId pulumi.StringPtrInput
}

func (AutoSnapshotPolicyApplyState) ElementType() reflect.Type {
	return reflect.TypeOf((*autoSnapshotPolicyApplyState)(nil)).Elem()
}

type autoSnapshotPolicyApplyArgs struct {
	// The id of auto snapshot policy.
	AutoSnapshotPolicyId string `pulumi:"autoSnapshotPolicyId"`
	// The id of file system.
	FileSystemId string `pulumi:"fileSystemId"`
}

// The set of arguments for constructing a AutoSnapshotPolicyApply resource.
type AutoSnapshotPolicyApplyArgs struct {
	// The id of auto snapshot policy.
	AutoSnapshotPolicyId pulumi.StringInput
	// The id of file system.
	FileSystemId pulumi.StringInput
}

func (AutoSnapshotPolicyApplyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*autoSnapshotPolicyApplyArgs)(nil)).Elem()
}

type AutoSnapshotPolicyApplyInput interface {
	pulumi.Input

	ToAutoSnapshotPolicyApplyOutput() AutoSnapshotPolicyApplyOutput
	ToAutoSnapshotPolicyApplyOutputWithContext(ctx context.Context) AutoSnapshotPolicyApplyOutput
}

func (*AutoSnapshotPolicyApply) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoSnapshotPolicyApply)(nil)).Elem()
}

func (i *AutoSnapshotPolicyApply) ToAutoSnapshotPolicyApplyOutput() AutoSnapshotPolicyApplyOutput {
	return i.ToAutoSnapshotPolicyApplyOutputWithContext(context.Background())
}

func (i *AutoSnapshotPolicyApply) ToAutoSnapshotPolicyApplyOutputWithContext(ctx context.Context) AutoSnapshotPolicyApplyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoSnapshotPolicyApplyOutput)
}

// AutoSnapshotPolicyApplyArrayInput is an input type that accepts AutoSnapshotPolicyApplyArray and AutoSnapshotPolicyApplyArrayOutput values.
// You can construct a concrete instance of `AutoSnapshotPolicyApplyArrayInput` via:
//
//	AutoSnapshotPolicyApplyArray{ AutoSnapshotPolicyApplyArgs{...} }
type AutoSnapshotPolicyApplyArrayInput interface {
	pulumi.Input

	ToAutoSnapshotPolicyApplyArrayOutput() AutoSnapshotPolicyApplyArrayOutput
	ToAutoSnapshotPolicyApplyArrayOutputWithContext(context.Context) AutoSnapshotPolicyApplyArrayOutput
}

type AutoSnapshotPolicyApplyArray []AutoSnapshotPolicyApplyInput

func (AutoSnapshotPolicyApplyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AutoSnapshotPolicyApply)(nil)).Elem()
}

func (i AutoSnapshotPolicyApplyArray) ToAutoSnapshotPolicyApplyArrayOutput() AutoSnapshotPolicyApplyArrayOutput {
	return i.ToAutoSnapshotPolicyApplyArrayOutputWithContext(context.Background())
}

func (i AutoSnapshotPolicyApplyArray) ToAutoSnapshotPolicyApplyArrayOutputWithContext(ctx context.Context) AutoSnapshotPolicyApplyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoSnapshotPolicyApplyArrayOutput)
}

// AutoSnapshotPolicyApplyMapInput is an input type that accepts AutoSnapshotPolicyApplyMap and AutoSnapshotPolicyApplyMapOutput values.
// You can construct a concrete instance of `AutoSnapshotPolicyApplyMapInput` via:
//
//	AutoSnapshotPolicyApplyMap{ "key": AutoSnapshotPolicyApplyArgs{...} }
type AutoSnapshotPolicyApplyMapInput interface {
	pulumi.Input

	ToAutoSnapshotPolicyApplyMapOutput() AutoSnapshotPolicyApplyMapOutput
	ToAutoSnapshotPolicyApplyMapOutputWithContext(context.Context) AutoSnapshotPolicyApplyMapOutput
}

type AutoSnapshotPolicyApplyMap map[string]AutoSnapshotPolicyApplyInput

func (AutoSnapshotPolicyApplyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AutoSnapshotPolicyApply)(nil)).Elem()
}

func (i AutoSnapshotPolicyApplyMap) ToAutoSnapshotPolicyApplyMapOutput() AutoSnapshotPolicyApplyMapOutput {
	return i.ToAutoSnapshotPolicyApplyMapOutputWithContext(context.Background())
}

func (i AutoSnapshotPolicyApplyMap) ToAutoSnapshotPolicyApplyMapOutputWithContext(ctx context.Context) AutoSnapshotPolicyApplyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoSnapshotPolicyApplyMapOutput)
}

type AutoSnapshotPolicyApplyOutput struct{ *pulumi.OutputState }

func (AutoSnapshotPolicyApplyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoSnapshotPolicyApply)(nil)).Elem()
}

func (o AutoSnapshotPolicyApplyOutput) ToAutoSnapshotPolicyApplyOutput() AutoSnapshotPolicyApplyOutput {
	return o
}

func (o AutoSnapshotPolicyApplyOutput) ToAutoSnapshotPolicyApplyOutputWithContext(ctx context.Context) AutoSnapshotPolicyApplyOutput {
	return o
}

// The id of auto snapshot policy.
func (o AutoSnapshotPolicyApplyOutput) AutoSnapshotPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *AutoSnapshotPolicyApply) pulumi.StringOutput { return v.AutoSnapshotPolicyId }).(pulumi.StringOutput)
}

// The id of file system.
func (o AutoSnapshotPolicyApplyOutput) FileSystemId() pulumi.StringOutput {
	return o.ApplyT(func(v *AutoSnapshotPolicyApply) pulumi.StringOutput { return v.FileSystemId }).(pulumi.StringOutput)
}

type AutoSnapshotPolicyApplyArrayOutput struct{ *pulumi.OutputState }

func (AutoSnapshotPolicyApplyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AutoSnapshotPolicyApply)(nil)).Elem()
}

func (o AutoSnapshotPolicyApplyArrayOutput) ToAutoSnapshotPolicyApplyArrayOutput() AutoSnapshotPolicyApplyArrayOutput {
	return o
}

func (o AutoSnapshotPolicyApplyArrayOutput) ToAutoSnapshotPolicyApplyArrayOutputWithContext(ctx context.Context) AutoSnapshotPolicyApplyArrayOutput {
	return o
}

func (o AutoSnapshotPolicyApplyArrayOutput) Index(i pulumi.IntInput) AutoSnapshotPolicyApplyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AutoSnapshotPolicyApply {
		return vs[0].([]*AutoSnapshotPolicyApply)[vs[1].(int)]
	}).(AutoSnapshotPolicyApplyOutput)
}

type AutoSnapshotPolicyApplyMapOutput struct{ *pulumi.OutputState }

func (AutoSnapshotPolicyApplyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AutoSnapshotPolicyApply)(nil)).Elem()
}

func (o AutoSnapshotPolicyApplyMapOutput) ToAutoSnapshotPolicyApplyMapOutput() AutoSnapshotPolicyApplyMapOutput {
	return o
}

func (o AutoSnapshotPolicyApplyMapOutput) ToAutoSnapshotPolicyApplyMapOutputWithContext(ctx context.Context) AutoSnapshotPolicyApplyMapOutput {
	return o
}

func (o AutoSnapshotPolicyApplyMapOutput) MapIndex(k pulumi.StringInput) AutoSnapshotPolicyApplyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AutoSnapshotPolicyApply {
		return vs[0].(map[string]*AutoSnapshotPolicyApply)[vs[1].(string)]
	}).(AutoSnapshotPolicyApplyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AutoSnapshotPolicyApplyInput)(nil)).Elem(), &AutoSnapshotPolicyApply{})
	pulumi.RegisterInputType(reflect.TypeOf((*AutoSnapshotPolicyApplyArrayInput)(nil)).Elem(), AutoSnapshotPolicyApplyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AutoSnapshotPolicyApplyMapInput)(nil)).Elem(), AutoSnapshotPolicyApplyMap{})
	pulumi.RegisterOutputType(AutoSnapshotPolicyApplyOutput{})
	pulumi.RegisterOutputType(AutoSnapshotPolicyApplyArrayOutput{})
	pulumi.RegisterOutputType(AutoSnapshotPolicyApplyMapOutput{})
}
