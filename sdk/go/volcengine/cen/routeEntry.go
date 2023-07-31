// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cen

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage cen route entry
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/cen"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cen.NewRouteEntry(ctx, "foo", &cen.RouteEntryArgs{
//				CenId:                pulumi.String("cen-12ar8uclj68sg17q7y20v9gil"),
//				DestinationCidrBlock: pulumi.String("192.168.0.0/24"),
//				InstanceId:           pulumi.String("vpc-im67wjcikxkw8gbssx8ufpj8"),
//				InstanceRegionId:     pulumi.String("cn-beijing"),
//				InstanceType:         pulumi.String("VPC"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = cen.NewRouteEntry(ctx, "foo1", &cen.RouteEntryArgs{
//				CenId:                pulumi.String("cen-12ar8uclj68sg17q7y20v9gil"),
//				DestinationCidrBlock: pulumi.String("192.168.17.0/24"),
//				InstanceId:           pulumi.String("vpc-im67wjcikxkw8gbssx8ufpj8"),
//				InstanceRegionId:     pulumi.String("cn-beijing"),
//				InstanceType:         pulumi.String("VPC"),
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
// CenRouteEntry can be imported using the CenId:DestinationCidrBlock:InstanceId:InstanceType:InstanceRegionId, e.g.
//
// ```sh
//
//	$ pulumi import volcengine:cen/routeEntry:RouteEntry default cen-2nim00ybaylts7trquyzt****:100.XX.XX.0/24:vpc-vtbnbb04qw3k2hgi12cv****:VPC:cn-beijing
//
// ```
type RouteEntry struct {
	pulumi.CustomResourceState

	// The AS path of the cen route entry.
	AsPaths pulumi.StringArrayOutput `pulumi:"asPaths"`
	// The cen ID of the cen route entry.
	CenId pulumi.StringOutput `pulumi:"cenId"`
	// The destination cidr block of the cen route entry.
	DestinationCidrBlock pulumi.StringOutput `pulumi:"destinationCidrBlock"`
	// The instance id of the next hop of the cen route entry.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The instance region id of the next hop of the cen route entry.
	InstanceRegionId pulumi.StringOutput `pulumi:"instanceRegionId"`
	// The instance type of the next hop of the cen route entry.
	InstanceType pulumi.StringPtrOutput `pulumi:"instanceType"`
	// The publish status of the cen route entry.
	PublishStatus pulumi.StringOutput `pulumi:"publishStatus"`
	// The status of the cen route entry.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewRouteEntry registers a new resource with the given unique name, arguments, and options.
func NewRouteEntry(ctx *pulumi.Context,
	name string, args *RouteEntryArgs, opts ...pulumi.ResourceOption) (*RouteEntry, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CenId == nil {
		return nil, errors.New("invalid value for required argument 'CenId'")
	}
	if args.DestinationCidrBlock == nil {
		return nil, errors.New("invalid value for required argument 'DestinationCidrBlock'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.InstanceRegionId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceRegionId'")
	}
	var resource RouteEntry
	err := ctx.RegisterResource("volcengine:cen/routeEntry:RouteEntry", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRouteEntry gets an existing RouteEntry resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRouteEntry(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouteEntryState, opts ...pulumi.ResourceOption) (*RouteEntry, error) {
	var resource RouteEntry
	err := ctx.ReadResource("volcengine:cen/routeEntry:RouteEntry", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RouteEntry resources.
type routeEntryState struct {
	// The AS path of the cen route entry.
	AsPaths []string `pulumi:"asPaths"`
	// The cen ID of the cen route entry.
	CenId *string `pulumi:"cenId"`
	// The destination cidr block of the cen route entry.
	DestinationCidrBlock *string `pulumi:"destinationCidrBlock"`
	// The instance id of the next hop of the cen route entry.
	InstanceId *string `pulumi:"instanceId"`
	// The instance region id of the next hop of the cen route entry.
	InstanceRegionId *string `pulumi:"instanceRegionId"`
	// The instance type of the next hop of the cen route entry.
	InstanceType *string `pulumi:"instanceType"`
	// The publish status of the cen route entry.
	PublishStatus *string `pulumi:"publishStatus"`
	// The status of the cen route entry.
	Status *string `pulumi:"status"`
}

type RouteEntryState struct {
	// The AS path of the cen route entry.
	AsPaths pulumi.StringArrayInput
	// The cen ID of the cen route entry.
	CenId pulumi.StringPtrInput
	// The destination cidr block of the cen route entry.
	DestinationCidrBlock pulumi.StringPtrInput
	// The instance id of the next hop of the cen route entry.
	InstanceId pulumi.StringPtrInput
	// The instance region id of the next hop of the cen route entry.
	InstanceRegionId pulumi.StringPtrInput
	// The instance type of the next hop of the cen route entry.
	InstanceType pulumi.StringPtrInput
	// The publish status of the cen route entry.
	PublishStatus pulumi.StringPtrInput
	// The status of the cen route entry.
	Status pulumi.StringPtrInput
}

func (RouteEntryState) ElementType() reflect.Type {
	return reflect.TypeOf((*routeEntryState)(nil)).Elem()
}

type routeEntryArgs struct {
	// The cen ID of the cen route entry.
	CenId string `pulumi:"cenId"`
	// The destination cidr block of the cen route entry.
	DestinationCidrBlock string `pulumi:"destinationCidrBlock"`
	// The instance id of the next hop of the cen route entry.
	InstanceId string `pulumi:"instanceId"`
	// The instance region id of the next hop of the cen route entry.
	InstanceRegionId string `pulumi:"instanceRegionId"`
	// The instance type of the next hop of the cen route entry.
	InstanceType *string `pulumi:"instanceType"`
}

// The set of arguments for constructing a RouteEntry resource.
type RouteEntryArgs struct {
	// The cen ID of the cen route entry.
	CenId pulumi.StringInput
	// The destination cidr block of the cen route entry.
	DestinationCidrBlock pulumi.StringInput
	// The instance id of the next hop of the cen route entry.
	InstanceId pulumi.StringInput
	// The instance region id of the next hop of the cen route entry.
	InstanceRegionId pulumi.StringInput
	// The instance type of the next hop of the cen route entry.
	InstanceType pulumi.StringPtrInput
}

func (RouteEntryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routeEntryArgs)(nil)).Elem()
}

type RouteEntryInput interface {
	pulumi.Input

	ToRouteEntryOutput() RouteEntryOutput
	ToRouteEntryOutputWithContext(ctx context.Context) RouteEntryOutput
}

func (*RouteEntry) ElementType() reflect.Type {
	return reflect.TypeOf((**RouteEntry)(nil)).Elem()
}

func (i *RouteEntry) ToRouteEntryOutput() RouteEntryOutput {
	return i.ToRouteEntryOutputWithContext(context.Background())
}

func (i *RouteEntry) ToRouteEntryOutputWithContext(ctx context.Context) RouteEntryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteEntryOutput)
}

// RouteEntryArrayInput is an input type that accepts RouteEntryArray and RouteEntryArrayOutput values.
// You can construct a concrete instance of `RouteEntryArrayInput` via:
//
//	RouteEntryArray{ RouteEntryArgs{...} }
type RouteEntryArrayInput interface {
	pulumi.Input

	ToRouteEntryArrayOutput() RouteEntryArrayOutput
	ToRouteEntryArrayOutputWithContext(context.Context) RouteEntryArrayOutput
}

type RouteEntryArray []RouteEntryInput

func (RouteEntryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RouteEntry)(nil)).Elem()
}

func (i RouteEntryArray) ToRouteEntryArrayOutput() RouteEntryArrayOutput {
	return i.ToRouteEntryArrayOutputWithContext(context.Background())
}

func (i RouteEntryArray) ToRouteEntryArrayOutputWithContext(ctx context.Context) RouteEntryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteEntryArrayOutput)
}

// RouteEntryMapInput is an input type that accepts RouteEntryMap and RouteEntryMapOutput values.
// You can construct a concrete instance of `RouteEntryMapInput` via:
//
//	RouteEntryMap{ "key": RouteEntryArgs{...} }
type RouteEntryMapInput interface {
	pulumi.Input

	ToRouteEntryMapOutput() RouteEntryMapOutput
	ToRouteEntryMapOutputWithContext(context.Context) RouteEntryMapOutput
}

type RouteEntryMap map[string]RouteEntryInput

func (RouteEntryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RouteEntry)(nil)).Elem()
}

func (i RouteEntryMap) ToRouteEntryMapOutput() RouteEntryMapOutput {
	return i.ToRouteEntryMapOutputWithContext(context.Background())
}

func (i RouteEntryMap) ToRouteEntryMapOutputWithContext(ctx context.Context) RouteEntryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteEntryMapOutput)
}

type RouteEntryOutput struct{ *pulumi.OutputState }

func (RouteEntryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RouteEntry)(nil)).Elem()
}

func (o RouteEntryOutput) ToRouteEntryOutput() RouteEntryOutput {
	return o
}

func (o RouteEntryOutput) ToRouteEntryOutputWithContext(ctx context.Context) RouteEntryOutput {
	return o
}

// The AS path of the cen route entry.
func (o RouteEntryOutput) AsPaths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RouteEntry) pulumi.StringArrayOutput { return v.AsPaths }).(pulumi.StringArrayOutput)
}

// The cen ID of the cen route entry.
func (o RouteEntryOutput) CenId() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteEntry) pulumi.StringOutput { return v.CenId }).(pulumi.StringOutput)
}

// The destination cidr block of the cen route entry.
func (o RouteEntryOutput) DestinationCidrBlock() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteEntry) pulumi.StringOutput { return v.DestinationCidrBlock }).(pulumi.StringOutput)
}

// The instance id of the next hop of the cen route entry.
func (o RouteEntryOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteEntry) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// The instance region id of the next hop of the cen route entry.
func (o RouteEntryOutput) InstanceRegionId() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteEntry) pulumi.StringOutput { return v.InstanceRegionId }).(pulumi.StringOutput)
}

// The instance type of the next hop of the cen route entry.
func (o RouteEntryOutput) InstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouteEntry) pulumi.StringPtrOutput { return v.InstanceType }).(pulumi.StringPtrOutput)
}

// The publish status of the cen route entry.
func (o RouteEntryOutput) PublishStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteEntry) pulumi.StringOutput { return v.PublishStatus }).(pulumi.StringOutput)
}

// The status of the cen route entry.
func (o RouteEntryOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteEntry) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type RouteEntryArrayOutput struct{ *pulumi.OutputState }

func (RouteEntryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RouteEntry)(nil)).Elem()
}

func (o RouteEntryArrayOutput) ToRouteEntryArrayOutput() RouteEntryArrayOutput {
	return o
}

func (o RouteEntryArrayOutput) ToRouteEntryArrayOutputWithContext(ctx context.Context) RouteEntryArrayOutput {
	return o
}

func (o RouteEntryArrayOutput) Index(i pulumi.IntInput) RouteEntryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RouteEntry {
		return vs[0].([]*RouteEntry)[vs[1].(int)]
	}).(RouteEntryOutput)
}

type RouteEntryMapOutput struct{ *pulumi.OutputState }

func (RouteEntryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RouteEntry)(nil)).Elem()
}

func (o RouteEntryMapOutput) ToRouteEntryMapOutput() RouteEntryMapOutput {
	return o
}

func (o RouteEntryMapOutput) ToRouteEntryMapOutputWithContext(ctx context.Context) RouteEntryMapOutput {
	return o
}

func (o RouteEntryMapOutput) MapIndex(k pulumi.StringInput) RouteEntryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RouteEntry {
		return vs[0].(map[string]*RouteEntry)[vs[1].(string)]
	}).(RouteEntryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RouteEntryInput)(nil)).Elem(), &RouteEntry{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouteEntryArrayInput)(nil)).Elem(), RouteEntryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouteEntryMapInput)(nil)).Elem(), RouteEntryMap{})
	pulumi.RegisterOutputType(RouteEntryOutput{})
	pulumi.RegisterOutputType(RouteEntryArrayOutput{})
	pulumi.RegisterOutputType(RouteEntryMapOutput{})
}