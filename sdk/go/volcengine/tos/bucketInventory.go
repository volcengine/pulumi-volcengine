// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tos

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Provides a resource to manage tos bucket inventory
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/tos"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := tos.NewBucketInventory(ctx, "foo", &tos.BucketInventoryArgs{
//				BucketName: pulumi.String("terraform-demo"),
//				Destination: &tos.BucketInventoryDestinationArgs{
//					TosBucketDestination: &tos.BucketInventoryDestinationTosBucketDestinationArgs{
//						AccountId: pulumi.String("21000*****"),
//						Bucket:    pulumi.String("terraform-demo"),
//						Format:    pulumi.String("CSV"),
//						Prefix:    pulumi.String("tf-test-prefix"),
//						Role:      pulumi.String("TosArchiveTOSInventory"),
//					},
//				},
//				Filter: &tos.BucketInventoryFilterArgs{
//					Prefix: pulumi.String("test-tf"),
//				},
//				IncludedObjectVersions: pulumi.String("All"),
//				InventoryId:            pulumi.String("acc-test-inventory"),
//				IsEnabled:              pulumi.Bool(true),
//				OptionalFields: &tos.BucketInventoryOptionalFieldsArgs{
//					Fields: pulumi.StringArray{
//						pulumi.String("Size"),
//						pulumi.String("StorageClass"),
//						pulumi.String("CRC64"),
//					},
//				},
//				Schedule: &tos.BucketInventoryScheduleArgs{
//					Frequency: pulumi.String("Weekly"),
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
// TosBucketInventory can be imported using the bucket_name:inventory_id, e.g.
//
// ```sh
// $ pulumi import volcengine:tos/bucketInventory:BucketInventory default resource_id
// ```
type BucketInventory struct {
	pulumi.CustomResourceState

	// The name of the bucket.
	BucketName pulumi.StringOutput `pulumi:"bucketName"`
	// The destination information of the bucket inventory.
	Destination BucketInventoryDestinationOutput `pulumi:"destination"`
	// The filter of the bucket inventory.
	Filter BucketInventoryFilterPtrOutput `pulumi:"filter"`
	// The export version of object. Valid values: `All`, `Current`.
	IncludedObjectVersions pulumi.StringOutput `pulumi:"includedObjectVersions"`
	// The name of the bucket inventory.
	InventoryId pulumi.StringOutput `pulumi:"inventoryId"`
	// Whether to enable the bucket inventory.
	IsEnabled pulumi.BoolOutput `pulumi:"isEnabled"`
	// The information exported from the bucket inventory.
	OptionalFields BucketInventoryOptionalFieldsPtrOutput `pulumi:"optionalFields"`
	// The export schedule of the bucket inventory.
	Schedule BucketInventoryScheduleOutput `pulumi:"schedule"`
}

// NewBucketInventory registers a new resource with the given unique name, arguments, and options.
func NewBucketInventory(ctx *pulumi.Context,
	name string, args *BucketInventoryArgs, opts ...pulumi.ResourceOption) (*BucketInventory, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BucketName == nil {
		return nil, errors.New("invalid value for required argument 'BucketName'")
	}
	if args.Destination == nil {
		return nil, errors.New("invalid value for required argument 'Destination'")
	}
	if args.IncludedObjectVersions == nil {
		return nil, errors.New("invalid value for required argument 'IncludedObjectVersions'")
	}
	if args.InventoryId == nil {
		return nil, errors.New("invalid value for required argument 'InventoryId'")
	}
	if args.IsEnabled == nil {
		return nil, errors.New("invalid value for required argument 'IsEnabled'")
	}
	if args.Schedule == nil {
		return nil, errors.New("invalid value for required argument 'Schedule'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource BucketInventory
	err := ctx.RegisterResource("volcengine:tos/bucketInventory:BucketInventory", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBucketInventory gets an existing BucketInventory resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBucketInventory(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BucketInventoryState, opts ...pulumi.ResourceOption) (*BucketInventory, error) {
	var resource BucketInventory
	err := ctx.ReadResource("volcengine:tos/bucketInventory:BucketInventory", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BucketInventory resources.
type bucketInventoryState struct {
	// The name of the bucket.
	BucketName *string `pulumi:"bucketName"`
	// The destination information of the bucket inventory.
	Destination *BucketInventoryDestination `pulumi:"destination"`
	// The filter of the bucket inventory.
	Filter *BucketInventoryFilter `pulumi:"filter"`
	// The export version of object. Valid values: `All`, `Current`.
	IncludedObjectVersions *string `pulumi:"includedObjectVersions"`
	// The name of the bucket inventory.
	InventoryId *string `pulumi:"inventoryId"`
	// Whether to enable the bucket inventory.
	IsEnabled *bool `pulumi:"isEnabled"`
	// The information exported from the bucket inventory.
	OptionalFields *BucketInventoryOptionalFields `pulumi:"optionalFields"`
	// The export schedule of the bucket inventory.
	Schedule *BucketInventorySchedule `pulumi:"schedule"`
}

type BucketInventoryState struct {
	// The name of the bucket.
	BucketName pulumi.StringPtrInput
	// The destination information of the bucket inventory.
	Destination BucketInventoryDestinationPtrInput
	// The filter of the bucket inventory.
	Filter BucketInventoryFilterPtrInput
	// The export version of object. Valid values: `All`, `Current`.
	IncludedObjectVersions pulumi.StringPtrInput
	// The name of the bucket inventory.
	InventoryId pulumi.StringPtrInput
	// Whether to enable the bucket inventory.
	IsEnabled pulumi.BoolPtrInput
	// The information exported from the bucket inventory.
	OptionalFields BucketInventoryOptionalFieldsPtrInput
	// The export schedule of the bucket inventory.
	Schedule BucketInventorySchedulePtrInput
}

func (BucketInventoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketInventoryState)(nil)).Elem()
}

type bucketInventoryArgs struct {
	// The name of the bucket.
	BucketName string `pulumi:"bucketName"`
	// The destination information of the bucket inventory.
	Destination BucketInventoryDestination `pulumi:"destination"`
	// The filter of the bucket inventory.
	Filter *BucketInventoryFilter `pulumi:"filter"`
	// The export version of object. Valid values: `All`, `Current`.
	IncludedObjectVersions string `pulumi:"includedObjectVersions"`
	// The name of the bucket inventory.
	InventoryId string `pulumi:"inventoryId"`
	// Whether to enable the bucket inventory.
	IsEnabled bool `pulumi:"isEnabled"`
	// The information exported from the bucket inventory.
	OptionalFields *BucketInventoryOptionalFields `pulumi:"optionalFields"`
	// The export schedule of the bucket inventory.
	Schedule BucketInventorySchedule `pulumi:"schedule"`
}

// The set of arguments for constructing a BucketInventory resource.
type BucketInventoryArgs struct {
	// The name of the bucket.
	BucketName pulumi.StringInput
	// The destination information of the bucket inventory.
	Destination BucketInventoryDestinationInput
	// The filter of the bucket inventory.
	Filter BucketInventoryFilterPtrInput
	// The export version of object. Valid values: `All`, `Current`.
	IncludedObjectVersions pulumi.StringInput
	// The name of the bucket inventory.
	InventoryId pulumi.StringInput
	// Whether to enable the bucket inventory.
	IsEnabled pulumi.BoolInput
	// The information exported from the bucket inventory.
	OptionalFields BucketInventoryOptionalFieldsPtrInput
	// The export schedule of the bucket inventory.
	Schedule BucketInventoryScheduleInput
}

func (BucketInventoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketInventoryArgs)(nil)).Elem()
}

type BucketInventoryInput interface {
	pulumi.Input

	ToBucketInventoryOutput() BucketInventoryOutput
	ToBucketInventoryOutputWithContext(ctx context.Context) BucketInventoryOutput
}

func (*BucketInventory) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketInventory)(nil)).Elem()
}

func (i *BucketInventory) ToBucketInventoryOutput() BucketInventoryOutput {
	return i.ToBucketInventoryOutputWithContext(context.Background())
}

func (i *BucketInventory) ToBucketInventoryOutputWithContext(ctx context.Context) BucketInventoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketInventoryOutput)
}

// BucketInventoryArrayInput is an input type that accepts BucketInventoryArray and BucketInventoryArrayOutput values.
// You can construct a concrete instance of `BucketInventoryArrayInput` via:
//
//	BucketInventoryArray{ BucketInventoryArgs{...} }
type BucketInventoryArrayInput interface {
	pulumi.Input

	ToBucketInventoryArrayOutput() BucketInventoryArrayOutput
	ToBucketInventoryArrayOutputWithContext(context.Context) BucketInventoryArrayOutput
}

type BucketInventoryArray []BucketInventoryInput

func (BucketInventoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BucketInventory)(nil)).Elem()
}

func (i BucketInventoryArray) ToBucketInventoryArrayOutput() BucketInventoryArrayOutput {
	return i.ToBucketInventoryArrayOutputWithContext(context.Background())
}

func (i BucketInventoryArray) ToBucketInventoryArrayOutputWithContext(ctx context.Context) BucketInventoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketInventoryArrayOutput)
}

// BucketInventoryMapInput is an input type that accepts BucketInventoryMap and BucketInventoryMapOutput values.
// You can construct a concrete instance of `BucketInventoryMapInput` via:
//
//	BucketInventoryMap{ "key": BucketInventoryArgs{...} }
type BucketInventoryMapInput interface {
	pulumi.Input

	ToBucketInventoryMapOutput() BucketInventoryMapOutput
	ToBucketInventoryMapOutputWithContext(context.Context) BucketInventoryMapOutput
}

type BucketInventoryMap map[string]BucketInventoryInput

func (BucketInventoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BucketInventory)(nil)).Elem()
}

func (i BucketInventoryMap) ToBucketInventoryMapOutput() BucketInventoryMapOutput {
	return i.ToBucketInventoryMapOutputWithContext(context.Background())
}

func (i BucketInventoryMap) ToBucketInventoryMapOutputWithContext(ctx context.Context) BucketInventoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketInventoryMapOutput)
}

type BucketInventoryOutput struct{ *pulumi.OutputState }

func (BucketInventoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketInventory)(nil)).Elem()
}

func (o BucketInventoryOutput) ToBucketInventoryOutput() BucketInventoryOutput {
	return o
}

func (o BucketInventoryOutput) ToBucketInventoryOutputWithContext(ctx context.Context) BucketInventoryOutput {
	return o
}

// The name of the bucket.
func (o BucketInventoryOutput) BucketName() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketInventory) pulumi.StringOutput { return v.BucketName }).(pulumi.StringOutput)
}

// The destination information of the bucket inventory.
func (o BucketInventoryOutput) Destination() BucketInventoryDestinationOutput {
	return o.ApplyT(func(v *BucketInventory) BucketInventoryDestinationOutput { return v.Destination }).(BucketInventoryDestinationOutput)
}

// The filter of the bucket inventory.
func (o BucketInventoryOutput) Filter() BucketInventoryFilterPtrOutput {
	return o.ApplyT(func(v *BucketInventory) BucketInventoryFilterPtrOutput { return v.Filter }).(BucketInventoryFilterPtrOutput)
}

// The export version of object. Valid values: `All`, `Current`.
func (o BucketInventoryOutput) IncludedObjectVersions() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketInventory) pulumi.StringOutput { return v.IncludedObjectVersions }).(pulumi.StringOutput)
}

// The name of the bucket inventory.
func (o BucketInventoryOutput) InventoryId() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketInventory) pulumi.StringOutput { return v.InventoryId }).(pulumi.StringOutput)
}

// Whether to enable the bucket inventory.
func (o BucketInventoryOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *BucketInventory) pulumi.BoolOutput { return v.IsEnabled }).(pulumi.BoolOutput)
}

// The information exported from the bucket inventory.
func (o BucketInventoryOutput) OptionalFields() BucketInventoryOptionalFieldsPtrOutput {
	return o.ApplyT(func(v *BucketInventory) BucketInventoryOptionalFieldsPtrOutput { return v.OptionalFields }).(BucketInventoryOptionalFieldsPtrOutput)
}

// The export schedule of the bucket inventory.
func (o BucketInventoryOutput) Schedule() BucketInventoryScheduleOutput {
	return o.ApplyT(func(v *BucketInventory) BucketInventoryScheduleOutput { return v.Schedule }).(BucketInventoryScheduleOutput)
}

type BucketInventoryArrayOutput struct{ *pulumi.OutputState }

func (BucketInventoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BucketInventory)(nil)).Elem()
}

func (o BucketInventoryArrayOutput) ToBucketInventoryArrayOutput() BucketInventoryArrayOutput {
	return o
}

func (o BucketInventoryArrayOutput) ToBucketInventoryArrayOutputWithContext(ctx context.Context) BucketInventoryArrayOutput {
	return o
}

func (o BucketInventoryArrayOutput) Index(i pulumi.IntInput) BucketInventoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BucketInventory {
		return vs[0].([]*BucketInventory)[vs[1].(int)]
	}).(BucketInventoryOutput)
}

type BucketInventoryMapOutput struct{ *pulumi.OutputState }

func (BucketInventoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BucketInventory)(nil)).Elem()
}

func (o BucketInventoryMapOutput) ToBucketInventoryMapOutput() BucketInventoryMapOutput {
	return o
}

func (o BucketInventoryMapOutput) ToBucketInventoryMapOutputWithContext(ctx context.Context) BucketInventoryMapOutput {
	return o
}

func (o BucketInventoryMapOutput) MapIndex(k pulumi.StringInput) BucketInventoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BucketInventory {
		return vs[0].(map[string]*BucketInventory)[vs[1].(string)]
	}).(BucketInventoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BucketInventoryInput)(nil)).Elem(), &BucketInventory{})
	pulumi.RegisterInputType(reflect.TypeOf((*BucketInventoryArrayInput)(nil)).Elem(), BucketInventoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BucketInventoryMapInput)(nil)).Elem(), BucketInventoryMap{})
	pulumi.RegisterOutputType(BucketInventoryOutput{})
	pulumi.RegisterOutputType(BucketInventoryArrayOutput{})
	pulumi.RegisterOutputType(BucketInventoryMapOutput{})
}
