// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mongodb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage mongodb ssl state
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-volcengine/sdk/go/volcengine/mongodb"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := mongodb.NewSslState(ctx, "foo", &mongodb.SslStateArgs{
//				InstanceId: pulumi.String("mongo-replica-f16e9298b121"),
//				SslAction:  pulumi.String("Update"),
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
// mongodb ssl state can be imported using the ssl:instanceId, e.g.
//
// ```sh
//
//	$ pulumi import volcengine:mongodb/sslState:SslState default ssl:mongo-shard-d050db19xxx
//
// ```
//
//	Set `ssl_action` to `Update` will update ssl always when terraform apply.
type SslState struct {
	pulumi.CustomResourceState

	// The ID of mongodb instance.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Whetehr SSL is valid.
	IsValid pulumi.BoolOutput `pulumi:"isValid"`
	// The action of ssl, valid value contains `Update`. Set `ssl_action` to `Update` will update ssl always when terraform
	// apply.
	SslAction pulumi.StringPtrOutput `pulumi:"sslAction"`
	// Whether SSL is enabled.
	SslEnable pulumi.BoolOutput `pulumi:"sslEnable"`
	// The expire time of SSL.
	SslExpiredTime pulumi.StringOutput `pulumi:"sslExpiredTime"`
}

// NewSslState registers a new resource with the given unique name, arguments, and options.
func NewSslState(ctx *pulumi.Context,
	name string, args *SslStateArgs, opts ...pulumi.ResourceOption) (*SslState, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	var resource SslState
	err := ctx.RegisterResource("volcengine:mongodb/sslState:SslState", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSslState gets an existing SslState resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSslState(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SslStateState, opts ...pulumi.ResourceOption) (*SslState, error) {
	var resource SslState
	err := ctx.ReadResource("volcengine:mongodb/sslState:SslState", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SslState resources.
type sslStateState struct {
	// The ID of mongodb instance.
	InstanceId *string `pulumi:"instanceId"`
	// Whetehr SSL is valid.
	IsValid *bool `pulumi:"isValid"`
	// The action of ssl, valid value contains `Update`. Set `ssl_action` to `Update` will update ssl always when terraform
	// apply.
	SslAction *string `pulumi:"sslAction"`
	// Whether SSL is enabled.
	SslEnable *bool `pulumi:"sslEnable"`
	// The expire time of SSL.
	SslExpiredTime *string `pulumi:"sslExpiredTime"`
}

type SslStateState struct {
	// The ID of mongodb instance.
	InstanceId pulumi.StringPtrInput
	// Whetehr SSL is valid.
	IsValid pulumi.BoolPtrInput
	// The action of ssl, valid value contains `Update`. Set `ssl_action` to `Update` will update ssl always when terraform
	// apply.
	SslAction pulumi.StringPtrInput
	// Whether SSL is enabled.
	SslEnable pulumi.BoolPtrInput
	// The expire time of SSL.
	SslExpiredTime pulumi.StringPtrInput
}

func (SslStateState) ElementType() reflect.Type {
	return reflect.TypeOf((*sslStateState)(nil)).Elem()
}

type sslStateArgs struct {
	// The ID of mongodb instance.
	InstanceId string `pulumi:"instanceId"`
	// The action of ssl, valid value contains `Update`. Set `ssl_action` to `Update` will update ssl always when terraform
	// apply.
	SslAction *string `pulumi:"sslAction"`
}

// The set of arguments for constructing a SslState resource.
type SslStateArgs struct {
	// The ID of mongodb instance.
	InstanceId pulumi.StringInput
	// The action of ssl, valid value contains `Update`. Set `ssl_action` to `Update` will update ssl always when terraform
	// apply.
	SslAction pulumi.StringPtrInput
}

func (SslStateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sslStateArgs)(nil)).Elem()
}

type SslStateInput interface {
	pulumi.Input

	ToSslStateOutput() SslStateOutput
	ToSslStateOutputWithContext(ctx context.Context) SslStateOutput
}

func (*SslState) ElementType() reflect.Type {
	return reflect.TypeOf((**SslState)(nil)).Elem()
}

func (i *SslState) ToSslStateOutput() SslStateOutput {
	return i.ToSslStateOutputWithContext(context.Background())
}

func (i *SslState) ToSslStateOutputWithContext(ctx context.Context) SslStateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SslStateOutput)
}

// SslStateArrayInput is an input type that accepts SslStateArray and SslStateArrayOutput values.
// You can construct a concrete instance of `SslStateArrayInput` via:
//
//	SslStateArray{ SslStateArgs{...} }
type SslStateArrayInput interface {
	pulumi.Input

	ToSslStateArrayOutput() SslStateArrayOutput
	ToSslStateArrayOutputWithContext(context.Context) SslStateArrayOutput
}

type SslStateArray []SslStateInput

func (SslStateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SslState)(nil)).Elem()
}

func (i SslStateArray) ToSslStateArrayOutput() SslStateArrayOutput {
	return i.ToSslStateArrayOutputWithContext(context.Background())
}

func (i SslStateArray) ToSslStateArrayOutputWithContext(ctx context.Context) SslStateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SslStateArrayOutput)
}

// SslStateMapInput is an input type that accepts SslStateMap and SslStateMapOutput values.
// You can construct a concrete instance of `SslStateMapInput` via:
//
//	SslStateMap{ "key": SslStateArgs{...} }
type SslStateMapInput interface {
	pulumi.Input

	ToSslStateMapOutput() SslStateMapOutput
	ToSslStateMapOutputWithContext(context.Context) SslStateMapOutput
}

type SslStateMap map[string]SslStateInput

func (SslStateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SslState)(nil)).Elem()
}

func (i SslStateMap) ToSslStateMapOutput() SslStateMapOutput {
	return i.ToSslStateMapOutputWithContext(context.Background())
}

func (i SslStateMap) ToSslStateMapOutputWithContext(ctx context.Context) SslStateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SslStateMapOutput)
}

type SslStateOutput struct{ *pulumi.OutputState }

func (SslStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SslState)(nil)).Elem()
}

func (o SslStateOutput) ToSslStateOutput() SslStateOutput {
	return o
}

func (o SslStateOutput) ToSslStateOutputWithContext(ctx context.Context) SslStateOutput {
	return o
}

// The ID of mongodb instance.
func (o SslStateOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SslState) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Whetehr SSL is valid.
func (o SslStateOutput) IsValid() pulumi.BoolOutput {
	return o.ApplyT(func(v *SslState) pulumi.BoolOutput { return v.IsValid }).(pulumi.BoolOutput)
}

// The action of ssl, valid value contains `Update`. Set `ssl_action` to `Update` will update ssl always when terraform
// apply.
func (o SslStateOutput) SslAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SslState) pulumi.StringPtrOutput { return v.SslAction }).(pulumi.StringPtrOutput)
}

// Whether SSL is enabled.
func (o SslStateOutput) SslEnable() pulumi.BoolOutput {
	return o.ApplyT(func(v *SslState) pulumi.BoolOutput { return v.SslEnable }).(pulumi.BoolOutput)
}

// The expire time of SSL.
func (o SslStateOutput) SslExpiredTime() pulumi.StringOutput {
	return o.ApplyT(func(v *SslState) pulumi.StringOutput { return v.SslExpiredTime }).(pulumi.StringOutput)
}

type SslStateArrayOutput struct{ *pulumi.OutputState }

func (SslStateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SslState)(nil)).Elem()
}

func (o SslStateArrayOutput) ToSslStateArrayOutput() SslStateArrayOutput {
	return o
}

func (o SslStateArrayOutput) ToSslStateArrayOutputWithContext(ctx context.Context) SslStateArrayOutput {
	return o
}

func (o SslStateArrayOutput) Index(i pulumi.IntInput) SslStateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SslState {
		return vs[0].([]*SslState)[vs[1].(int)]
	}).(SslStateOutput)
}

type SslStateMapOutput struct{ *pulumi.OutputState }

func (SslStateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SslState)(nil)).Elem()
}

func (o SslStateMapOutput) ToSslStateMapOutput() SslStateMapOutput {
	return o
}

func (o SslStateMapOutput) ToSslStateMapOutputWithContext(ctx context.Context) SslStateMapOutput {
	return o
}

func (o SslStateMapOutput) MapIndex(k pulumi.StringInput) SslStateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SslState {
		return vs[0].(map[string]*SslState)[vs[1].(string)]
	}).(SslStateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SslStateInput)(nil)).Elem(), &SslState{})
	pulumi.RegisterInputType(reflect.TypeOf((*SslStateArrayInput)(nil)).Elem(), SslStateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SslStateMapInput)(nil)).Elem(), SslStateMap{})
	pulumi.RegisterOutputType(SslStateOutput{})
	pulumi.RegisterOutputType(SslStateArrayOutput{})
	pulumi.RegisterOutputType(SslStateMapOutput{})
}