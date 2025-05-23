// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cen

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Provides a resource to manage cen
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
//			_, err := cen.NewCen(ctx, "foo", &cen.CenArgs{
//				CenName:     pulumi.String("acc-test-cen"),
//				Description: pulumi.String("acc-test"),
//				ProjectName: pulumi.String("default"),
//				Tags: cen.CenTagArray{
//					&cen.CenTagArgs{
//						Key:   pulumi.String("k1"),
//						Value: pulumi.String("v1"),
//					},
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
// Cen can be imported using the id, e.g.
//
// ```sh
// $ pulumi import volcengine:cen/cen:Cen default cen-7qthudw0ll6jmc****
// ```
type Cen struct {
	pulumi.CustomResourceState

	// The account ID of the cen.
	AccountId pulumi.StringOutput `pulumi:"accountId"`
	// A list of bandwidth package IDs of the cen.
	CenBandwidthPackageIds pulumi.StringArrayOutput `pulumi:"cenBandwidthPackageIds"`
	// The ID of the cen.
	CenId pulumi.StringOutput `pulumi:"cenId"`
	// The name of the cen.
	CenName pulumi.StringOutput `pulumi:"cenName"`
	// The create time of the cen.
	CreationTime pulumi.StringOutput `pulumi:"creationTime"`
	// The description of the cen.
	Description pulumi.StringOutput `pulumi:"description"`
	// The ProjectName of the cen instance.
	ProjectName pulumi.StringOutput `pulumi:"projectName"`
	// The status of the cen.
	Status pulumi.StringOutput `pulumi:"status"`
	// Tags.
	Tags CenTagArrayOutput `pulumi:"tags"`
	// The update time of the cen.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewCen registers a new resource with the given unique name, arguments, and options.
func NewCen(ctx *pulumi.Context,
	name string, args *CenArgs, opts ...pulumi.ResourceOption) (*Cen, error) {
	if args == nil {
		args = &CenArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Cen
	err := ctx.RegisterResource("volcengine:cen/cen:Cen", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCen gets an existing Cen resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCen(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CenState, opts ...pulumi.ResourceOption) (*Cen, error) {
	var resource Cen
	err := ctx.ReadResource("volcengine:cen/cen:Cen", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Cen resources.
type cenState struct {
	// The account ID of the cen.
	AccountId *string `pulumi:"accountId"`
	// A list of bandwidth package IDs of the cen.
	CenBandwidthPackageIds []string `pulumi:"cenBandwidthPackageIds"`
	// The ID of the cen.
	CenId *string `pulumi:"cenId"`
	// The name of the cen.
	CenName *string `pulumi:"cenName"`
	// The create time of the cen.
	CreationTime *string `pulumi:"creationTime"`
	// The description of the cen.
	Description *string `pulumi:"description"`
	// The ProjectName of the cen instance.
	ProjectName *string `pulumi:"projectName"`
	// The status of the cen.
	Status *string `pulumi:"status"`
	// Tags.
	Tags []CenTag `pulumi:"tags"`
	// The update time of the cen.
	UpdateTime *string `pulumi:"updateTime"`
}

type CenState struct {
	// The account ID of the cen.
	AccountId pulumi.StringPtrInput
	// A list of bandwidth package IDs of the cen.
	CenBandwidthPackageIds pulumi.StringArrayInput
	// The ID of the cen.
	CenId pulumi.StringPtrInput
	// The name of the cen.
	CenName pulumi.StringPtrInput
	// The create time of the cen.
	CreationTime pulumi.StringPtrInput
	// The description of the cen.
	Description pulumi.StringPtrInput
	// The ProjectName of the cen instance.
	ProjectName pulumi.StringPtrInput
	// The status of the cen.
	Status pulumi.StringPtrInput
	// Tags.
	Tags CenTagArrayInput
	// The update time of the cen.
	UpdateTime pulumi.StringPtrInput
}

func (CenState) ElementType() reflect.Type {
	return reflect.TypeOf((*cenState)(nil)).Elem()
}

type cenArgs struct {
	// The name of the cen.
	CenName *string `pulumi:"cenName"`
	// The description of the cen.
	Description *string `pulumi:"description"`
	// The ProjectName of the cen instance.
	ProjectName *string `pulumi:"projectName"`
	// Tags.
	Tags []CenTag `pulumi:"tags"`
}

// The set of arguments for constructing a Cen resource.
type CenArgs struct {
	// The name of the cen.
	CenName pulumi.StringPtrInput
	// The description of the cen.
	Description pulumi.StringPtrInput
	// The ProjectName of the cen instance.
	ProjectName pulumi.StringPtrInput
	// Tags.
	Tags CenTagArrayInput
}

func (CenArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cenArgs)(nil)).Elem()
}

type CenInput interface {
	pulumi.Input

	ToCenOutput() CenOutput
	ToCenOutputWithContext(ctx context.Context) CenOutput
}

func (*Cen) ElementType() reflect.Type {
	return reflect.TypeOf((**Cen)(nil)).Elem()
}

func (i *Cen) ToCenOutput() CenOutput {
	return i.ToCenOutputWithContext(context.Background())
}

func (i *Cen) ToCenOutputWithContext(ctx context.Context) CenOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CenOutput)
}

// CenArrayInput is an input type that accepts CenArray and CenArrayOutput values.
// You can construct a concrete instance of `CenArrayInput` via:
//
//	CenArray{ CenArgs{...} }
type CenArrayInput interface {
	pulumi.Input

	ToCenArrayOutput() CenArrayOutput
	ToCenArrayOutputWithContext(context.Context) CenArrayOutput
}

type CenArray []CenInput

func (CenArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Cen)(nil)).Elem()
}

func (i CenArray) ToCenArrayOutput() CenArrayOutput {
	return i.ToCenArrayOutputWithContext(context.Background())
}

func (i CenArray) ToCenArrayOutputWithContext(ctx context.Context) CenArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CenArrayOutput)
}

// CenMapInput is an input type that accepts CenMap and CenMapOutput values.
// You can construct a concrete instance of `CenMapInput` via:
//
//	CenMap{ "key": CenArgs{...} }
type CenMapInput interface {
	pulumi.Input

	ToCenMapOutput() CenMapOutput
	ToCenMapOutputWithContext(context.Context) CenMapOutput
}

type CenMap map[string]CenInput

func (CenMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Cen)(nil)).Elem()
}

func (i CenMap) ToCenMapOutput() CenMapOutput {
	return i.ToCenMapOutputWithContext(context.Background())
}

func (i CenMap) ToCenMapOutputWithContext(ctx context.Context) CenMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CenMapOutput)
}

type CenOutput struct{ *pulumi.OutputState }

func (CenOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Cen)(nil)).Elem()
}

func (o CenOutput) ToCenOutput() CenOutput {
	return o
}

func (o CenOutput) ToCenOutputWithContext(ctx context.Context) CenOutput {
	return o
}

// The account ID of the cen.
func (o CenOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *Cen) pulumi.StringOutput { return v.AccountId }).(pulumi.StringOutput)
}

// A list of bandwidth package IDs of the cen.
func (o CenOutput) CenBandwidthPackageIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Cen) pulumi.StringArrayOutput { return v.CenBandwidthPackageIds }).(pulumi.StringArrayOutput)
}

// The ID of the cen.
func (o CenOutput) CenId() pulumi.StringOutput {
	return o.ApplyT(func(v *Cen) pulumi.StringOutput { return v.CenId }).(pulumi.StringOutput)
}

// The name of the cen.
func (o CenOutput) CenName() pulumi.StringOutput {
	return o.ApplyT(func(v *Cen) pulumi.StringOutput { return v.CenName }).(pulumi.StringOutput)
}

// The create time of the cen.
func (o CenOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Cen) pulumi.StringOutput { return v.CreationTime }).(pulumi.StringOutput)
}

// The description of the cen.
func (o CenOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Cen) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The ProjectName of the cen instance.
func (o CenOutput) ProjectName() pulumi.StringOutput {
	return o.ApplyT(func(v *Cen) pulumi.StringOutput { return v.ProjectName }).(pulumi.StringOutput)
}

// The status of the cen.
func (o CenOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Cen) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Tags.
func (o CenOutput) Tags() CenTagArrayOutput {
	return o.ApplyT(func(v *Cen) CenTagArrayOutput { return v.Tags }).(CenTagArrayOutput)
}

// The update time of the cen.
func (o CenOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Cen) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

type CenArrayOutput struct{ *pulumi.OutputState }

func (CenArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Cen)(nil)).Elem()
}

func (o CenArrayOutput) ToCenArrayOutput() CenArrayOutput {
	return o
}

func (o CenArrayOutput) ToCenArrayOutputWithContext(ctx context.Context) CenArrayOutput {
	return o
}

func (o CenArrayOutput) Index(i pulumi.IntInput) CenOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Cen {
		return vs[0].([]*Cen)[vs[1].(int)]
	}).(CenOutput)
}

type CenMapOutput struct{ *pulumi.OutputState }

func (CenMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Cen)(nil)).Elem()
}

func (o CenMapOutput) ToCenMapOutput() CenMapOutput {
	return o
}

func (o CenMapOutput) ToCenMapOutputWithContext(ctx context.Context) CenMapOutput {
	return o
}

func (o CenMapOutput) MapIndex(k pulumi.StringInput) CenOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Cen {
		return vs[0].(map[string]*Cen)[vs[1].(string)]
	}).(CenOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CenInput)(nil)).Elem(), &Cen{})
	pulumi.RegisterInputType(reflect.TypeOf((*CenArrayInput)(nil)).Elem(), CenArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CenMapInput)(nil)).Elem(), CenMap{})
	pulumi.RegisterOutputType(CenOutput{})
	pulumi.RegisterOutputType(CenArrayOutput{})
	pulumi.RegisterOutputType(CenMapOutput{})
}
