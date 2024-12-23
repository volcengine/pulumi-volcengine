// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package alb

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Provides a resource to manage alb certificate
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/alb"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := alb.NewCertificate(ctx, "foo", &alb.CertificateArgs{
//				Description: pulumi.String("test123"),
//				PrivateKey:  pulumi.String("private key"),
//				PublicKey:   pulumi.String("public key"),
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
// Certificate can be imported using the id, e.g.
//
// ```sh
// $ pulumi import volcengine:alb/certificate:Certificate default cert-2fe5k****c16o5oxruvtk3qf5
// ```
type Certificate struct {
	pulumi.CustomResourceState

	// The name of the Certificate.
	CertificateName pulumi.StringOutput `pulumi:"certificateName"`
	// The type of the Certificate.
	CertificateType pulumi.StringOutput `pulumi:"certificateType"`
	// The create time of the Certificate.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The description of the Certificate.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The domain name of the Certificate.
	DomainName pulumi.StringOutput `pulumi:"domainName"`
	// The expire time of the Certificate.
	ExpiredAt pulumi.StringOutput `pulumi:"expiredAt"`
	// The ID list of the Listener.
	Listeners pulumi.StringArrayOutput `pulumi:"listeners"`
	// The private key of the Certificate.
	PrivateKey pulumi.StringOutput `pulumi:"privateKey"`
	// The project name of the Certificate.
	ProjectName pulumi.StringOutput `pulumi:"projectName"`
	// The public key of the Certificate.
	PublicKey pulumi.StringOutput `pulumi:"publicKey"`
	// The status of the Certificate.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewCertificate registers a new resource with the given unique name, arguments, and options.
func NewCertificate(ctx *pulumi.Context,
	name string, args *CertificateArgs, opts ...pulumi.ResourceOption) (*Certificate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PrivateKey == nil {
		return nil, errors.New("invalid value for required argument 'PrivateKey'")
	}
	if args.PublicKey == nil {
		return nil, errors.New("invalid value for required argument 'PublicKey'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Certificate
	err := ctx.RegisterResource("volcengine:alb/certificate:Certificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCertificate gets an existing Certificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CertificateState, opts ...pulumi.ResourceOption) (*Certificate, error) {
	var resource Certificate
	err := ctx.ReadResource("volcengine:alb/certificate:Certificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Certificate resources.
type certificateState struct {
	// The name of the Certificate.
	CertificateName *string `pulumi:"certificateName"`
	// The type of the Certificate.
	CertificateType *string `pulumi:"certificateType"`
	// The create time of the Certificate.
	CreateTime *string `pulumi:"createTime"`
	// The description of the Certificate.
	Description *string `pulumi:"description"`
	// The domain name of the Certificate.
	DomainName *string `pulumi:"domainName"`
	// The expire time of the Certificate.
	ExpiredAt *string `pulumi:"expiredAt"`
	// The ID list of the Listener.
	Listeners []string `pulumi:"listeners"`
	// The private key of the Certificate.
	PrivateKey *string `pulumi:"privateKey"`
	// The project name of the Certificate.
	ProjectName *string `pulumi:"projectName"`
	// The public key of the Certificate.
	PublicKey *string `pulumi:"publicKey"`
	// The status of the Certificate.
	Status *string `pulumi:"status"`
}

type CertificateState struct {
	// The name of the Certificate.
	CertificateName pulumi.StringPtrInput
	// The type of the Certificate.
	CertificateType pulumi.StringPtrInput
	// The create time of the Certificate.
	CreateTime pulumi.StringPtrInput
	// The description of the Certificate.
	Description pulumi.StringPtrInput
	// The domain name of the Certificate.
	DomainName pulumi.StringPtrInput
	// The expire time of the Certificate.
	ExpiredAt pulumi.StringPtrInput
	// The ID list of the Listener.
	Listeners pulumi.StringArrayInput
	// The private key of the Certificate.
	PrivateKey pulumi.StringPtrInput
	// The project name of the Certificate.
	ProjectName pulumi.StringPtrInput
	// The public key of the Certificate.
	PublicKey pulumi.StringPtrInput
	// The status of the Certificate.
	Status pulumi.StringPtrInput
}

func (CertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateState)(nil)).Elem()
}

type certificateArgs struct {
	// The name of the Certificate.
	CertificateName *string `pulumi:"certificateName"`
	// The description of the Certificate.
	Description *string `pulumi:"description"`
	// The private key of the Certificate.
	PrivateKey string `pulumi:"privateKey"`
	// The project name of the Certificate.
	ProjectName *string `pulumi:"projectName"`
	// The public key of the Certificate.
	PublicKey string `pulumi:"publicKey"`
}

// The set of arguments for constructing a Certificate resource.
type CertificateArgs struct {
	// The name of the Certificate.
	CertificateName pulumi.StringPtrInput
	// The description of the Certificate.
	Description pulumi.StringPtrInput
	// The private key of the Certificate.
	PrivateKey pulumi.StringInput
	// The project name of the Certificate.
	ProjectName pulumi.StringPtrInput
	// The public key of the Certificate.
	PublicKey pulumi.StringInput
}

func (CertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateArgs)(nil)).Elem()
}

type CertificateInput interface {
	pulumi.Input

	ToCertificateOutput() CertificateOutput
	ToCertificateOutputWithContext(ctx context.Context) CertificateOutput
}

func (*Certificate) ElementType() reflect.Type {
	return reflect.TypeOf((**Certificate)(nil)).Elem()
}

func (i *Certificate) ToCertificateOutput() CertificateOutput {
	return i.ToCertificateOutputWithContext(context.Background())
}

func (i *Certificate) ToCertificateOutputWithContext(ctx context.Context) CertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateOutput)
}

// CertificateArrayInput is an input type that accepts CertificateArray and CertificateArrayOutput values.
// You can construct a concrete instance of `CertificateArrayInput` via:
//
//	CertificateArray{ CertificateArgs{...} }
type CertificateArrayInput interface {
	pulumi.Input

	ToCertificateArrayOutput() CertificateArrayOutput
	ToCertificateArrayOutputWithContext(context.Context) CertificateArrayOutput
}

type CertificateArray []CertificateInput

func (CertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Certificate)(nil)).Elem()
}

func (i CertificateArray) ToCertificateArrayOutput() CertificateArrayOutput {
	return i.ToCertificateArrayOutputWithContext(context.Background())
}

func (i CertificateArray) ToCertificateArrayOutputWithContext(ctx context.Context) CertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateArrayOutput)
}

// CertificateMapInput is an input type that accepts CertificateMap and CertificateMapOutput values.
// You can construct a concrete instance of `CertificateMapInput` via:
//
//	CertificateMap{ "key": CertificateArgs{...} }
type CertificateMapInput interface {
	pulumi.Input

	ToCertificateMapOutput() CertificateMapOutput
	ToCertificateMapOutputWithContext(context.Context) CertificateMapOutput
}

type CertificateMap map[string]CertificateInput

func (CertificateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Certificate)(nil)).Elem()
}

func (i CertificateMap) ToCertificateMapOutput() CertificateMapOutput {
	return i.ToCertificateMapOutputWithContext(context.Background())
}

func (i CertificateMap) ToCertificateMapOutputWithContext(ctx context.Context) CertificateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateMapOutput)
}

type CertificateOutput struct{ *pulumi.OutputState }

func (CertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Certificate)(nil)).Elem()
}

func (o CertificateOutput) ToCertificateOutput() CertificateOutput {
	return o
}

func (o CertificateOutput) ToCertificateOutputWithContext(ctx context.Context) CertificateOutput {
	return o
}

// The name of the Certificate.
func (o CertificateOutput) CertificateName() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.CertificateName }).(pulumi.StringOutput)
}

// The type of the Certificate.
func (o CertificateOutput) CertificateType() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.CertificateType }).(pulumi.StringOutput)
}

// The create time of the Certificate.
func (o CertificateOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// The description of the Certificate.
func (o CertificateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The domain name of the Certificate.
func (o CertificateOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.DomainName }).(pulumi.StringOutput)
}

// The expire time of the Certificate.
func (o CertificateOutput) ExpiredAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.ExpiredAt }).(pulumi.StringOutput)
}

// The ID list of the Listener.
func (o CertificateOutput) Listeners() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringArrayOutput { return v.Listeners }).(pulumi.StringArrayOutput)
}

// The private key of the Certificate.
func (o CertificateOutput) PrivateKey() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.PrivateKey }).(pulumi.StringOutput)
}

// The project name of the Certificate.
func (o CertificateOutput) ProjectName() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.ProjectName }).(pulumi.StringOutput)
}

// The public key of the Certificate.
func (o CertificateOutput) PublicKey() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.PublicKey }).(pulumi.StringOutput)
}

// The status of the Certificate.
func (o CertificateOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type CertificateArrayOutput struct{ *pulumi.OutputState }

func (CertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Certificate)(nil)).Elem()
}

func (o CertificateArrayOutput) ToCertificateArrayOutput() CertificateArrayOutput {
	return o
}

func (o CertificateArrayOutput) ToCertificateArrayOutputWithContext(ctx context.Context) CertificateArrayOutput {
	return o
}

func (o CertificateArrayOutput) Index(i pulumi.IntInput) CertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Certificate {
		return vs[0].([]*Certificate)[vs[1].(int)]
	}).(CertificateOutput)
}

type CertificateMapOutput struct{ *pulumi.OutputState }

func (CertificateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Certificate)(nil)).Elem()
}

func (o CertificateMapOutput) ToCertificateMapOutput() CertificateMapOutput {
	return o
}

func (o CertificateMapOutput) ToCertificateMapOutputWithContext(ctx context.Context) CertificateMapOutput {
	return o
}

func (o CertificateMapOutput) MapIndex(k pulumi.StringInput) CertificateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Certificate {
		return vs[0].(map[string]*Certificate)[vs[1].(string)]
	}).(CertificateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateInput)(nil)).Elem(), &Certificate{})
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateArrayInput)(nil)).Elem(), CertificateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateMapInput)(nil)).Elem(), CertificateMap{})
	pulumi.RegisterOutputType(CertificateOutput{})
	pulumi.RegisterOutputType(CertificateArrayOutput{})
	pulumi.RegisterOutputType(CertificateMapOutput{})
}
