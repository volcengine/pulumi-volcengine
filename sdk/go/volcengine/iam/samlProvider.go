// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Provides a resource to manage iam saml provider
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/iam"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := iam.NewSamlProvider(ctx, "foo", &iam.SamlProviderArgs{
//				EncodedSamlMetadataDocument: pulumi.String("your document"),
//				SamlProviderName:            pulumi.String("terraform"),
//				SsoType:                     pulumi.Int(2),
//				Status:                      pulumi.Int(1),
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
// IamSamlProvider can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import volcengine:iam/samlProvider:SamlProvider default SAMLProviderName
//
// ```
type SamlProvider struct {
	pulumi.CustomResourceState

	// Identity provider creation time, such as 20150123T123318Z.
	CreateDate pulumi.StringOutput `pulumi:"createDate"`
	// The description of the SAML provider.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Metadata document, encoded in Base64.
	EncodedSamlMetadataDocument pulumi.StringOutput `pulumi:"encodedSamlMetadataDocument"`
	// The name of the SAML provider.
	SamlProviderName pulumi.StringOutput `pulumi:"samlProviderName"`
	// SSO types, 1. Role-based SSO, 2. User-based SSO.
	SsoType pulumi.IntOutput `pulumi:"ssoType"`
	// User SSO status, 1. Enabled, 2. Disable other console login methods after enabling, 3. Disabled, is a required field when creating user SSO.
	Status pulumi.IntPtrOutput `pulumi:"status"`
	// The format for the resource name of an identity provider is trn:iam::${accountID}:saml-provider/{$SAMLProviderName}.
	Trn pulumi.StringOutput `pulumi:"trn"`
	// Identity provider update time, such as: 20150123T123318Z.
	UpdateDate pulumi.StringOutput `pulumi:"updateDate"`
}

// NewSamlProvider registers a new resource with the given unique name, arguments, and options.
func NewSamlProvider(ctx *pulumi.Context,
	name string, args *SamlProviderArgs, opts ...pulumi.ResourceOption) (*SamlProvider, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EncodedSamlMetadataDocument == nil {
		return nil, errors.New("invalid value for required argument 'EncodedSamlMetadataDocument'")
	}
	if args.SamlProviderName == nil {
		return nil, errors.New("invalid value for required argument 'SamlProviderName'")
	}
	if args.SsoType == nil {
		return nil, errors.New("invalid value for required argument 'SsoType'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SamlProvider
	err := ctx.RegisterResource("volcengine:iam/samlProvider:SamlProvider", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSamlProvider gets an existing SamlProvider resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSamlProvider(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SamlProviderState, opts ...pulumi.ResourceOption) (*SamlProvider, error) {
	var resource SamlProvider
	err := ctx.ReadResource("volcengine:iam/samlProvider:SamlProvider", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SamlProvider resources.
type samlProviderState struct {
	// Identity provider creation time, such as 20150123T123318Z.
	CreateDate *string `pulumi:"createDate"`
	// The description of the SAML provider.
	Description *string `pulumi:"description"`
	// Metadata document, encoded in Base64.
	EncodedSamlMetadataDocument *string `pulumi:"encodedSamlMetadataDocument"`
	// The name of the SAML provider.
	SamlProviderName *string `pulumi:"samlProviderName"`
	// SSO types, 1. Role-based SSO, 2. User-based SSO.
	SsoType *int `pulumi:"ssoType"`
	// User SSO status, 1. Enabled, 2. Disable other console login methods after enabling, 3. Disabled, is a required field when creating user SSO.
	Status *int `pulumi:"status"`
	// The format for the resource name of an identity provider is trn:iam::${accountID}:saml-provider/{$SAMLProviderName}.
	Trn *string `pulumi:"trn"`
	// Identity provider update time, such as: 20150123T123318Z.
	UpdateDate *string `pulumi:"updateDate"`
}

type SamlProviderState struct {
	// Identity provider creation time, such as 20150123T123318Z.
	CreateDate pulumi.StringPtrInput
	// The description of the SAML provider.
	Description pulumi.StringPtrInput
	// Metadata document, encoded in Base64.
	EncodedSamlMetadataDocument pulumi.StringPtrInput
	// The name of the SAML provider.
	SamlProviderName pulumi.StringPtrInput
	// SSO types, 1. Role-based SSO, 2. User-based SSO.
	SsoType pulumi.IntPtrInput
	// User SSO status, 1. Enabled, 2. Disable other console login methods after enabling, 3. Disabled, is a required field when creating user SSO.
	Status pulumi.IntPtrInput
	// The format for the resource name of an identity provider is trn:iam::${accountID}:saml-provider/{$SAMLProviderName}.
	Trn pulumi.StringPtrInput
	// Identity provider update time, such as: 20150123T123318Z.
	UpdateDate pulumi.StringPtrInput
}

func (SamlProviderState) ElementType() reflect.Type {
	return reflect.TypeOf((*samlProviderState)(nil)).Elem()
}

type samlProviderArgs struct {
	// The description of the SAML provider.
	Description *string `pulumi:"description"`
	// Metadata document, encoded in Base64.
	EncodedSamlMetadataDocument string `pulumi:"encodedSamlMetadataDocument"`
	// The name of the SAML provider.
	SamlProviderName string `pulumi:"samlProviderName"`
	// SSO types, 1. Role-based SSO, 2. User-based SSO.
	SsoType int `pulumi:"ssoType"`
	// User SSO status, 1. Enabled, 2. Disable other console login methods after enabling, 3. Disabled, is a required field when creating user SSO.
	Status *int `pulumi:"status"`
}

// The set of arguments for constructing a SamlProvider resource.
type SamlProviderArgs struct {
	// The description of the SAML provider.
	Description pulumi.StringPtrInput
	// Metadata document, encoded in Base64.
	EncodedSamlMetadataDocument pulumi.StringInput
	// The name of the SAML provider.
	SamlProviderName pulumi.StringInput
	// SSO types, 1. Role-based SSO, 2. User-based SSO.
	SsoType pulumi.IntInput
	// User SSO status, 1. Enabled, 2. Disable other console login methods after enabling, 3. Disabled, is a required field when creating user SSO.
	Status pulumi.IntPtrInput
}

func (SamlProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*samlProviderArgs)(nil)).Elem()
}

type SamlProviderInput interface {
	pulumi.Input

	ToSamlProviderOutput() SamlProviderOutput
	ToSamlProviderOutputWithContext(ctx context.Context) SamlProviderOutput
}

func (*SamlProvider) ElementType() reflect.Type {
	return reflect.TypeOf((**SamlProvider)(nil)).Elem()
}

func (i *SamlProvider) ToSamlProviderOutput() SamlProviderOutput {
	return i.ToSamlProviderOutputWithContext(context.Background())
}

func (i *SamlProvider) ToSamlProviderOutputWithContext(ctx context.Context) SamlProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SamlProviderOutput)
}

// SamlProviderArrayInput is an input type that accepts SamlProviderArray and SamlProviderArrayOutput values.
// You can construct a concrete instance of `SamlProviderArrayInput` via:
//
//	SamlProviderArray{ SamlProviderArgs{...} }
type SamlProviderArrayInput interface {
	pulumi.Input

	ToSamlProviderArrayOutput() SamlProviderArrayOutput
	ToSamlProviderArrayOutputWithContext(context.Context) SamlProviderArrayOutput
}

type SamlProviderArray []SamlProviderInput

func (SamlProviderArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SamlProvider)(nil)).Elem()
}

func (i SamlProviderArray) ToSamlProviderArrayOutput() SamlProviderArrayOutput {
	return i.ToSamlProviderArrayOutputWithContext(context.Background())
}

func (i SamlProviderArray) ToSamlProviderArrayOutputWithContext(ctx context.Context) SamlProviderArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SamlProviderArrayOutput)
}

// SamlProviderMapInput is an input type that accepts SamlProviderMap and SamlProviderMapOutput values.
// You can construct a concrete instance of `SamlProviderMapInput` via:
//
//	SamlProviderMap{ "key": SamlProviderArgs{...} }
type SamlProviderMapInput interface {
	pulumi.Input

	ToSamlProviderMapOutput() SamlProviderMapOutput
	ToSamlProviderMapOutputWithContext(context.Context) SamlProviderMapOutput
}

type SamlProviderMap map[string]SamlProviderInput

func (SamlProviderMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SamlProvider)(nil)).Elem()
}

func (i SamlProviderMap) ToSamlProviderMapOutput() SamlProviderMapOutput {
	return i.ToSamlProviderMapOutputWithContext(context.Background())
}

func (i SamlProviderMap) ToSamlProviderMapOutputWithContext(ctx context.Context) SamlProviderMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SamlProviderMapOutput)
}

type SamlProviderOutput struct{ *pulumi.OutputState }

func (SamlProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SamlProvider)(nil)).Elem()
}

func (o SamlProviderOutput) ToSamlProviderOutput() SamlProviderOutput {
	return o
}

func (o SamlProviderOutput) ToSamlProviderOutputWithContext(ctx context.Context) SamlProviderOutput {
	return o
}

// Identity provider creation time, such as 20150123T123318Z.
func (o SamlProviderOutput) CreateDate() pulumi.StringOutput {
	return o.ApplyT(func(v *SamlProvider) pulumi.StringOutput { return v.CreateDate }).(pulumi.StringOutput)
}

// The description of the SAML provider.
func (o SamlProviderOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SamlProvider) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Metadata document, encoded in Base64.
func (o SamlProviderOutput) EncodedSamlMetadataDocument() pulumi.StringOutput {
	return o.ApplyT(func(v *SamlProvider) pulumi.StringOutput { return v.EncodedSamlMetadataDocument }).(pulumi.StringOutput)
}

// The name of the SAML provider.
func (o SamlProviderOutput) SamlProviderName() pulumi.StringOutput {
	return o.ApplyT(func(v *SamlProvider) pulumi.StringOutput { return v.SamlProviderName }).(pulumi.StringOutput)
}

// SSO types, 1. Role-based SSO, 2. User-based SSO.
func (o SamlProviderOutput) SsoType() pulumi.IntOutput {
	return o.ApplyT(func(v *SamlProvider) pulumi.IntOutput { return v.SsoType }).(pulumi.IntOutput)
}

// User SSO status, 1. Enabled, 2. Disable other console login methods after enabling, 3. Disabled, is a required field when creating user SSO.
func (o SamlProviderOutput) Status() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SamlProvider) pulumi.IntPtrOutput { return v.Status }).(pulumi.IntPtrOutput)
}

// The format for the resource name of an identity provider is trn:iam::${accountID}:saml-provider/{$SAMLProviderName}.
func (o SamlProviderOutput) Trn() pulumi.StringOutput {
	return o.ApplyT(func(v *SamlProvider) pulumi.StringOutput { return v.Trn }).(pulumi.StringOutput)
}

// Identity provider update time, such as: 20150123T123318Z.
func (o SamlProviderOutput) UpdateDate() pulumi.StringOutput {
	return o.ApplyT(func(v *SamlProvider) pulumi.StringOutput { return v.UpdateDate }).(pulumi.StringOutput)
}

type SamlProviderArrayOutput struct{ *pulumi.OutputState }

func (SamlProviderArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SamlProvider)(nil)).Elem()
}

func (o SamlProviderArrayOutput) ToSamlProviderArrayOutput() SamlProviderArrayOutput {
	return o
}

func (o SamlProviderArrayOutput) ToSamlProviderArrayOutputWithContext(ctx context.Context) SamlProviderArrayOutput {
	return o
}

func (o SamlProviderArrayOutput) Index(i pulumi.IntInput) SamlProviderOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SamlProvider {
		return vs[0].([]*SamlProvider)[vs[1].(int)]
	}).(SamlProviderOutput)
}

type SamlProviderMapOutput struct{ *pulumi.OutputState }

func (SamlProviderMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SamlProvider)(nil)).Elem()
}

func (o SamlProviderMapOutput) ToSamlProviderMapOutput() SamlProviderMapOutput {
	return o
}

func (o SamlProviderMapOutput) ToSamlProviderMapOutputWithContext(ctx context.Context) SamlProviderMapOutput {
	return o
}

func (o SamlProviderMapOutput) MapIndex(k pulumi.StringInput) SamlProviderOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SamlProvider {
		return vs[0].(map[string]*SamlProvider)[vs[1].(string)]
	}).(SamlProviderOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SamlProviderInput)(nil)).Elem(), &SamlProvider{})
	pulumi.RegisterInputType(reflect.TypeOf((*SamlProviderArrayInput)(nil)).Elem(), SamlProviderArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SamlProviderMapInput)(nil)).Elem(), SamlProviderMap{})
	pulumi.RegisterOutputType(SamlProviderOutput{})
	pulumi.RegisterOutputType(SamlProviderArrayOutput{})
	pulumi.RegisterOutputType(SamlProviderMapOutput{})
}