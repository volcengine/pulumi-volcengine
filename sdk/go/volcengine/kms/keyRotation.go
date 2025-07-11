// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kms

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Provides a resource to manage kms key rotation
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/kms"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := kms.NewKeyRotation(ctx, "foo", &kms.KeyRotationArgs{
//				KeyId: pulumi.String("m_cn-guilin-boe_63c08fe9-42e8-4c10-a09e-8e8e6xxxxxx"),
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
// KmsKeyRotation can be imported using the id, e.g.
//
// ```sh
// $ pulumi import volcengine:kms/keyRotation:KeyRotation default resource_id
// ```
//
// or
//
// ```sh
// $ pulumi import volcengine:kms/keyRotation:KeyRotation default key_name:keyring_name
// ```
type KeyRotation struct {
	pulumi.CustomResourceState

	// The id of the CMK.
	KeyId pulumi.StringOutput `pulumi:"keyId"`
	// The name of the CMK.
	KeyName pulumi.StringOutput `pulumi:"keyName"`
	// The name of the keyring.
	KeyringName pulumi.StringPtrOutput `pulumi:"keyringName"`
	// The state of the key rotation.
	RotationState pulumi.StringOutput `pulumi:"rotationState"`
}

// NewKeyRotation registers a new resource with the given unique name, arguments, and options.
func NewKeyRotation(ctx *pulumi.Context,
	name string, args *KeyRotationArgs, opts ...pulumi.ResourceOption) (*KeyRotation, error) {
	if args == nil {
		args = &KeyRotationArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource KeyRotation
	err := ctx.RegisterResource("volcengine:kms/keyRotation:KeyRotation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKeyRotation gets an existing KeyRotation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKeyRotation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KeyRotationState, opts ...pulumi.ResourceOption) (*KeyRotation, error) {
	var resource KeyRotation
	err := ctx.ReadResource("volcengine:kms/keyRotation:KeyRotation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KeyRotation resources.
type keyRotationState struct {
	// The id of the CMK.
	KeyId *string `pulumi:"keyId"`
	// The name of the CMK.
	KeyName *string `pulumi:"keyName"`
	// The name of the keyring.
	KeyringName *string `pulumi:"keyringName"`
	// The state of the key rotation.
	RotationState *string `pulumi:"rotationState"`
}

type KeyRotationState struct {
	// The id of the CMK.
	KeyId pulumi.StringPtrInput
	// The name of the CMK.
	KeyName pulumi.StringPtrInput
	// The name of the keyring.
	KeyringName pulumi.StringPtrInput
	// The state of the key rotation.
	RotationState pulumi.StringPtrInput
}

func (KeyRotationState) ElementType() reflect.Type {
	return reflect.TypeOf((*keyRotationState)(nil)).Elem()
}

type keyRotationArgs struct {
	// The id of the CMK.
	KeyId *string `pulumi:"keyId"`
	// The name of the CMK.
	KeyName *string `pulumi:"keyName"`
	// The name of the keyring.
	KeyringName *string `pulumi:"keyringName"`
}

// The set of arguments for constructing a KeyRotation resource.
type KeyRotationArgs struct {
	// The id of the CMK.
	KeyId pulumi.StringPtrInput
	// The name of the CMK.
	KeyName pulumi.StringPtrInput
	// The name of the keyring.
	KeyringName pulumi.StringPtrInput
}

func (KeyRotationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*keyRotationArgs)(nil)).Elem()
}

type KeyRotationInput interface {
	pulumi.Input

	ToKeyRotationOutput() KeyRotationOutput
	ToKeyRotationOutputWithContext(ctx context.Context) KeyRotationOutput
}

func (*KeyRotation) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyRotation)(nil)).Elem()
}

func (i *KeyRotation) ToKeyRotationOutput() KeyRotationOutput {
	return i.ToKeyRotationOutputWithContext(context.Background())
}

func (i *KeyRotation) ToKeyRotationOutputWithContext(ctx context.Context) KeyRotationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyRotationOutput)
}

// KeyRotationArrayInput is an input type that accepts KeyRotationArray and KeyRotationArrayOutput values.
// You can construct a concrete instance of `KeyRotationArrayInput` via:
//
//	KeyRotationArray{ KeyRotationArgs{...} }
type KeyRotationArrayInput interface {
	pulumi.Input

	ToKeyRotationArrayOutput() KeyRotationArrayOutput
	ToKeyRotationArrayOutputWithContext(context.Context) KeyRotationArrayOutput
}

type KeyRotationArray []KeyRotationInput

func (KeyRotationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KeyRotation)(nil)).Elem()
}

func (i KeyRotationArray) ToKeyRotationArrayOutput() KeyRotationArrayOutput {
	return i.ToKeyRotationArrayOutputWithContext(context.Background())
}

func (i KeyRotationArray) ToKeyRotationArrayOutputWithContext(ctx context.Context) KeyRotationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyRotationArrayOutput)
}

// KeyRotationMapInput is an input type that accepts KeyRotationMap and KeyRotationMapOutput values.
// You can construct a concrete instance of `KeyRotationMapInput` via:
//
//	KeyRotationMap{ "key": KeyRotationArgs{...} }
type KeyRotationMapInput interface {
	pulumi.Input

	ToKeyRotationMapOutput() KeyRotationMapOutput
	ToKeyRotationMapOutputWithContext(context.Context) KeyRotationMapOutput
}

type KeyRotationMap map[string]KeyRotationInput

func (KeyRotationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KeyRotation)(nil)).Elem()
}

func (i KeyRotationMap) ToKeyRotationMapOutput() KeyRotationMapOutput {
	return i.ToKeyRotationMapOutputWithContext(context.Background())
}

func (i KeyRotationMap) ToKeyRotationMapOutputWithContext(ctx context.Context) KeyRotationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyRotationMapOutput)
}

type KeyRotationOutput struct{ *pulumi.OutputState }

func (KeyRotationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyRotation)(nil)).Elem()
}

func (o KeyRotationOutput) ToKeyRotationOutput() KeyRotationOutput {
	return o
}

func (o KeyRotationOutput) ToKeyRotationOutputWithContext(ctx context.Context) KeyRotationOutput {
	return o
}

// The id of the CMK.
func (o KeyRotationOutput) KeyId() pulumi.StringOutput {
	return o.ApplyT(func(v *KeyRotation) pulumi.StringOutput { return v.KeyId }).(pulumi.StringOutput)
}

// The name of the CMK.
func (o KeyRotationOutput) KeyName() pulumi.StringOutput {
	return o.ApplyT(func(v *KeyRotation) pulumi.StringOutput { return v.KeyName }).(pulumi.StringOutput)
}

// The name of the keyring.
func (o KeyRotationOutput) KeyringName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyRotation) pulumi.StringPtrOutput { return v.KeyringName }).(pulumi.StringPtrOutput)
}

// The state of the key rotation.
func (o KeyRotationOutput) RotationState() pulumi.StringOutput {
	return o.ApplyT(func(v *KeyRotation) pulumi.StringOutput { return v.RotationState }).(pulumi.StringOutput)
}

type KeyRotationArrayOutput struct{ *pulumi.OutputState }

func (KeyRotationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KeyRotation)(nil)).Elem()
}

func (o KeyRotationArrayOutput) ToKeyRotationArrayOutput() KeyRotationArrayOutput {
	return o
}

func (o KeyRotationArrayOutput) ToKeyRotationArrayOutputWithContext(ctx context.Context) KeyRotationArrayOutput {
	return o
}

func (o KeyRotationArrayOutput) Index(i pulumi.IntInput) KeyRotationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *KeyRotation {
		return vs[0].([]*KeyRotation)[vs[1].(int)]
	}).(KeyRotationOutput)
}

type KeyRotationMapOutput struct{ *pulumi.OutputState }

func (KeyRotationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KeyRotation)(nil)).Elem()
}

func (o KeyRotationMapOutput) ToKeyRotationMapOutput() KeyRotationMapOutput {
	return o
}

func (o KeyRotationMapOutput) ToKeyRotationMapOutputWithContext(ctx context.Context) KeyRotationMapOutput {
	return o
}

func (o KeyRotationMapOutput) MapIndex(k pulumi.StringInput) KeyRotationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *KeyRotation {
		return vs[0].(map[string]*KeyRotation)[vs[1].(string)]
	}).(KeyRotationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KeyRotationInput)(nil)).Elem(), &KeyRotation{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyRotationArrayInput)(nil)).Elem(), KeyRotationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyRotationMapInput)(nil)).Elem(), KeyRotationMap{})
	pulumi.RegisterOutputType(KeyRotationOutput{})
	pulumi.RegisterOutputType(KeyRotationArrayOutput{})
	pulumi.RegisterOutputType(KeyRotationMapOutput{})
}
