// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage ecs key pair associate
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
//			_, err := ecs.NewKeyPairAssociate(ctx, "default", &ecs.KeyPairAssociateArgs{
//				InstanceId: pulumi.String("i-ybskpw36rul8u1yekckh"),
//				KeyPairId:  pulumi.String("kp-ybvyy1e5msl8u258ovrv"),
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
// ECS key pair associate can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import volcengine:ecs/keyPairAssociate:KeyPairAssociate default kp-ybti5tkpkv2udbfolrft:i-mizl7m1kqccg5smt1bdpijuj
//
// ```
type KeyPairAssociate struct {
	pulumi.CustomResourceState

	// The ID of ECS Instance.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The ID of ECS KeyPair Associate.
	KeyPairId pulumi.StringOutput `pulumi:"keyPairId"`
}

// NewKeyPairAssociate registers a new resource with the given unique name, arguments, and options.
func NewKeyPairAssociate(ctx *pulumi.Context,
	name string, args *KeyPairAssociateArgs, opts ...pulumi.ResourceOption) (*KeyPairAssociate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.KeyPairId == nil {
		return nil, errors.New("invalid value for required argument 'KeyPairId'")
	}
	var resource KeyPairAssociate
	err := ctx.RegisterResource("volcengine:ecs/keyPairAssociate:KeyPairAssociate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKeyPairAssociate gets an existing KeyPairAssociate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKeyPairAssociate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KeyPairAssociateState, opts ...pulumi.ResourceOption) (*KeyPairAssociate, error) {
	var resource KeyPairAssociate
	err := ctx.ReadResource("volcengine:ecs/keyPairAssociate:KeyPairAssociate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KeyPairAssociate resources.
type keyPairAssociateState struct {
	// The ID of ECS Instance.
	InstanceId *string `pulumi:"instanceId"`
	// The ID of ECS KeyPair Associate.
	KeyPairId *string `pulumi:"keyPairId"`
}

type KeyPairAssociateState struct {
	// The ID of ECS Instance.
	InstanceId pulumi.StringPtrInput
	// The ID of ECS KeyPair Associate.
	KeyPairId pulumi.StringPtrInput
}

func (KeyPairAssociateState) ElementType() reflect.Type {
	return reflect.TypeOf((*keyPairAssociateState)(nil)).Elem()
}

type keyPairAssociateArgs struct {
	// The ID of ECS Instance.
	InstanceId string `pulumi:"instanceId"`
	// The ID of ECS KeyPair Associate.
	KeyPairId string `pulumi:"keyPairId"`
}

// The set of arguments for constructing a KeyPairAssociate resource.
type KeyPairAssociateArgs struct {
	// The ID of ECS Instance.
	InstanceId pulumi.StringInput
	// The ID of ECS KeyPair Associate.
	KeyPairId pulumi.StringInput
}

func (KeyPairAssociateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*keyPairAssociateArgs)(nil)).Elem()
}

type KeyPairAssociateInput interface {
	pulumi.Input

	ToKeyPairAssociateOutput() KeyPairAssociateOutput
	ToKeyPairAssociateOutputWithContext(ctx context.Context) KeyPairAssociateOutput
}

func (*KeyPairAssociate) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyPairAssociate)(nil)).Elem()
}

func (i *KeyPairAssociate) ToKeyPairAssociateOutput() KeyPairAssociateOutput {
	return i.ToKeyPairAssociateOutputWithContext(context.Background())
}

func (i *KeyPairAssociate) ToKeyPairAssociateOutputWithContext(ctx context.Context) KeyPairAssociateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyPairAssociateOutput)
}

// KeyPairAssociateArrayInput is an input type that accepts KeyPairAssociateArray and KeyPairAssociateArrayOutput values.
// You can construct a concrete instance of `KeyPairAssociateArrayInput` via:
//
//	KeyPairAssociateArray{ KeyPairAssociateArgs{...} }
type KeyPairAssociateArrayInput interface {
	pulumi.Input

	ToKeyPairAssociateArrayOutput() KeyPairAssociateArrayOutput
	ToKeyPairAssociateArrayOutputWithContext(context.Context) KeyPairAssociateArrayOutput
}

type KeyPairAssociateArray []KeyPairAssociateInput

func (KeyPairAssociateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KeyPairAssociate)(nil)).Elem()
}

func (i KeyPairAssociateArray) ToKeyPairAssociateArrayOutput() KeyPairAssociateArrayOutput {
	return i.ToKeyPairAssociateArrayOutputWithContext(context.Background())
}

func (i KeyPairAssociateArray) ToKeyPairAssociateArrayOutputWithContext(ctx context.Context) KeyPairAssociateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyPairAssociateArrayOutput)
}

// KeyPairAssociateMapInput is an input type that accepts KeyPairAssociateMap and KeyPairAssociateMapOutput values.
// You can construct a concrete instance of `KeyPairAssociateMapInput` via:
//
//	KeyPairAssociateMap{ "key": KeyPairAssociateArgs{...} }
type KeyPairAssociateMapInput interface {
	pulumi.Input

	ToKeyPairAssociateMapOutput() KeyPairAssociateMapOutput
	ToKeyPairAssociateMapOutputWithContext(context.Context) KeyPairAssociateMapOutput
}

type KeyPairAssociateMap map[string]KeyPairAssociateInput

func (KeyPairAssociateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KeyPairAssociate)(nil)).Elem()
}

func (i KeyPairAssociateMap) ToKeyPairAssociateMapOutput() KeyPairAssociateMapOutput {
	return i.ToKeyPairAssociateMapOutputWithContext(context.Background())
}

func (i KeyPairAssociateMap) ToKeyPairAssociateMapOutputWithContext(ctx context.Context) KeyPairAssociateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyPairAssociateMapOutput)
}

type KeyPairAssociateOutput struct{ *pulumi.OutputState }

func (KeyPairAssociateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyPairAssociate)(nil)).Elem()
}

func (o KeyPairAssociateOutput) ToKeyPairAssociateOutput() KeyPairAssociateOutput {
	return o
}

func (o KeyPairAssociateOutput) ToKeyPairAssociateOutputWithContext(ctx context.Context) KeyPairAssociateOutput {
	return o
}

// The ID of ECS Instance.
func (o KeyPairAssociateOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *KeyPairAssociate) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// The ID of ECS KeyPair Associate.
func (o KeyPairAssociateOutput) KeyPairId() pulumi.StringOutput {
	return o.ApplyT(func(v *KeyPairAssociate) pulumi.StringOutput { return v.KeyPairId }).(pulumi.StringOutput)
}

type KeyPairAssociateArrayOutput struct{ *pulumi.OutputState }

func (KeyPairAssociateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KeyPairAssociate)(nil)).Elem()
}

func (o KeyPairAssociateArrayOutput) ToKeyPairAssociateArrayOutput() KeyPairAssociateArrayOutput {
	return o
}

func (o KeyPairAssociateArrayOutput) ToKeyPairAssociateArrayOutputWithContext(ctx context.Context) KeyPairAssociateArrayOutput {
	return o
}

func (o KeyPairAssociateArrayOutput) Index(i pulumi.IntInput) KeyPairAssociateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *KeyPairAssociate {
		return vs[0].([]*KeyPairAssociate)[vs[1].(int)]
	}).(KeyPairAssociateOutput)
}

type KeyPairAssociateMapOutput struct{ *pulumi.OutputState }

func (KeyPairAssociateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KeyPairAssociate)(nil)).Elem()
}

func (o KeyPairAssociateMapOutput) ToKeyPairAssociateMapOutput() KeyPairAssociateMapOutput {
	return o
}

func (o KeyPairAssociateMapOutput) ToKeyPairAssociateMapOutputWithContext(ctx context.Context) KeyPairAssociateMapOutput {
	return o
}

func (o KeyPairAssociateMapOutput) MapIndex(k pulumi.StringInput) KeyPairAssociateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *KeyPairAssociate {
		return vs[0].(map[string]*KeyPairAssociate)[vs[1].(string)]
	}).(KeyPairAssociateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KeyPairAssociateInput)(nil)).Elem(), &KeyPairAssociate{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyPairAssociateArrayInput)(nil)).Elem(), KeyPairAssociateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyPairAssociateMapInput)(nil)).Elem(), KeyPairAssociateMap{})
	pulumi.RegisterOutputType(KeyPairAssociateOutput{})
	pulumi.RegisterOutputType(KeyPairAssociateArrayOutput{})
	pulumi.RegisterOutputType(KeyPairAssociateMapOutput{})
}