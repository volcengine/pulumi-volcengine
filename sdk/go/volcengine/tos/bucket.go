// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tos

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage tos bucket
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-volcengine/sdk/go/volcengine/tos"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := tos.NewBucket(ctx, "default", &tos.BucketArgs{
//				AccountAcls: tos.BucketAccountAclArray{
//					&tos.BucketAccountAclArgs{
//						AccountId:  pulumi.String("1"),
//						Permission: pulumi.String("READ"),
//					},
//					&tos.BucketAccountAclArgs{
//						AccountId:  pulumi.String("2001"),
//						Permission: pulumi.String("WRITE_ACP"),
//					},
//				},
//				BucketName:    pulumi.String("test-xym-1"),
//				EnableVersion: pulumi.Bool(true),
//				PublicAcl:     pulumi.String("private"),
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
// Tos Bucket can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import volcengine:tos/bucket:Bucket default bucketName
//
// ```
type Bucket struct {
	pulumi.CustomResourceState

	// The user set of grant full control.
	AccountAcls BucketAccountAclArrayOutput `pulumi:"accountAcls"`
	// The name of the bucket.
	BucketName pulumi.StringOutput `pulumi:"bucketName"`
	// The create date of the TOS bucket.
	CreationDate pulumi.StringOutput `pulumi:"creationDate"`
	// The flag of enable tos version.
	EnableVersion pulumi.BoolPtrOutput `pulumi:"enableVersion"`
	// The extranet endpoint of the TOS bucket.
	ExtranetEndpoint pulumi.StringOutput `pulumi:"extranetEndpoint"`
	// The intranet endpoint the TOS bucket.
	IntranetEndpoint pulumi.StringOutput `pulumi:"intranetEndpoint"`
	// The location of the TOS bucket.
	Location pulumi.StringOutput `pulumi:"location"`
	// The public acl control of object.Valid value is private|public-read|public-read-write|authenticated-read|bucket-owner-read.
	PublicAcl pulumi.StringPtrOutput `pulumi:"publicAcl"`
	// The storage type of the object.Valid value is STANDARD|IA|ARCHIVE_FR.Default is STANDARD.
	StorageClass pulumi.StringPtrOutput `pulumi:"storageClass"`
}

// NewBucket registers a new resource with the given unique name, arguments, and options.
func NewBucket(ctx *pulumi.Context,
	name string, args *BucketArgs, opts ...pulumi.ResourceOption) (*Bucket, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BucketName == nil {
		return nil, errors.New("invalid value for required argument 'BucketName'")
	}
	var resource Bucket
	err := ctx.RegisterResource("volcengine:tos/bucket:Bucket", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBucket gets an existing Bucket resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBucket(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BucketState, opts ...pulumi.ResourceOption) (*Bucket, error) {
	var resource Bucket
	err := ctx.ReadResource("volcengine:tos/bucket:Bucket", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Bucket resources.
type bucketState struct {
	// The user set of grant full control.
	AccountAcls []BucketAccountAcl `pulumi:"accountAcls"`
	// The name of the bucket.
	BucketName *string `pulumi:"bucketName"`
	// The create date of the TOS bucket.
	CreationDate *string `pulumi:"creationDate"`
	// The flag of enable tos version.
	EnableVersion *bool `pulumi:"enableVersion"`
	// The extranet endpoint of the TOS bucket.
	ExtranetEndpoint *string `pulumi:"extranetEndpoint"`
	// The intranet endpoint the TOS bucket.
	IntranetEndpoint *string `pulumi:"intranetEndpoint"`
	// The location of the TOS bucket.
	Location *string `pulumi:"location"`
	// The public acl control of object.Valid value is private|public-read|public-read-write|authenticated-read|bucket-owner-read.
	PublicAcl *string `pulumi:"publicAcl"`
	// The storage type of the object.Valid value is STANDARD|IA|ARCHIVE_FR.Default is STANDARD.
	StorageClass *string `pulumi:"storageClass"`
}

type BucketState struct {
	// The user set of grant full control.
	AccountAcls BucketAccountAclArrayInput
	// The name of the bucket.
	BucketName pulumi.StringPtrInput
	// The create date of the TOS bucket.
	CreationDate pulumi.StringPtrInput
	// The flag of enable tos version.
	EnableVersion pulumi.BoolPtrInput
	// The extranet endpoint of the TOS bucket.
	ExtranetEndpoint pulumi.StringPtrInput
	// The intranet endpoint the TOS bucket.
	IntranetEndpoint pulumi.StringPtrInput
	// The location of the TOS bucket.
	Location pulumi.StringPtrInput
	// The public acl control of object.Valid value is private|public-read|public-read-write|authenticated-read|bucket-owner-read.
	PublicAcl pulumi.StringPtrInput
	// The storage type of the object.Valid value is STANDARD|IA|ARCHIVE_FR.Default is STANDARD.
	StorageClass pulumi.StringPtrInput
}

func (BucketState) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketState)(nil)).Elem()
}

type bucketArgs struct {
	// The user set of grant full control.
	AccountAcls []BucketAccountAcl `pulumi:"accountAcls"`
	// The name of the bucket.
	BucketName string `pulumi:"bucketName"`
	// The flag of enable tos version.
	EnableVersion *bool `pulumi:"enableVersion"`
	// The public acl control of object.Valid value is private|public-read|public-read-write|authenticated-read|bucket-owner-read.
	PublicAcl *string `pulumi:"publicAcl"`
	// The storage type of the object.Valid value is STANDARD|IA|ARCHIVE_FR.Default is STANDARD.
	StorageClass *string `pulumi:"storageClass"`
}

// The set of arguments for constructing a Bucket resource.
type BucketArgs struct {
	// The user set of grant full control.
	AccountAcls BucketAccountAclArrayInput
	// The name of the bucket.
	BucketName pulumi.StringInput
	// The flag of enable tos version.
	EnableVersion pulumi.BoolPtrInput
	// The public acl control of object.Valid value is private|public-read|public-read-write|authenticated-read|bucket-owner-read.
	PublicAcl pulumi.StringPtrInput
	// The storage type of the object.Valid value is STANDARD|IA|ARCHIVE_FR.Default is STANDARD.
	StorageClass pulumi.StringPtrInput
}

func (BucketArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketArgs)(nil)).Elem()
}

type BucketInput interface {
	pulumi.Input

	ToBucketOutput() BucketOutput
	ToBucketOutputWithContext(ctx context.Context) BucketOutput
}

func (*Bucket) ElementType() reflect.Type {
	return reflect.TypeOf((**Bucket)(nil)).Elem()
}

func (i *Bucket) ToBucketOutput() BucketOutput {
	return i.ToBucketOutputWithContext(context.Background())
}

func (i *Bucket) ToBucketOutputWithContext(ctx context.Context) BucketOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketOutput)
}

// BucketArrayInput is an input type that accepts BucketArray and BucketArrayOutput values.
// You can construct a concrete instance of `BucketArrayInput` via:
//
//	BucketArray{ BucketArgs{...} }
type BucketArrayInput interface {
	pulumi.Input

	ToBucketArrayOutput() BucketArrayOutput
	ToBucketArrayOutputWithContext(context.Context) BucketArrayOutput
}

type BucketArray []BucketInput

func (BucketArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Bucket)(nil)).Elem()
}

func (i BucketArray) ToBucketArrayOutput() BucketArrayOutput {
	return i.ToBucketArrayOutputWithContext(context.Background())
}

func (i BucketArray) ToBucketArrayOutputWithContext(ctx context.Context) BucketArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketArrayOutput)
}

// BucketMapInput is an input type that accepts BucketMap and BucketMapOutput values.
// You can construct a concrete instance of `BucketMapInput` via:
//
//	BucketMap{ "key": BucketArgs{...} }
type BucketMapInput interface {
	pulumi.Input

	ToBucketMapOutput() BucketMapOutput
	ToBucketMapOutputWithContext(context.Context) BucketMapOutput
}

type BucketMap map[string]BucketInput

func (BucketMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Bucket)(nil)).Elem()
}

func (i BucketMap) ToBucketMapOutput() BucketMapOutput {
	return i.ToBucketMapOutputWithContext(context.Background())
}

func (i BucketMap) ToBucketMapOutputWithContext(ctx context.Context) BucketMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketMapOutput)
}

type BucketOutput struct{ *pulumi.OutputState }

func (BucketOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Bucket)(nil)).Elem()
}

func (o BucketOutput) ToBucketOutput() BucketOutput {
	return o
}

func (o BucketOutput) ToBucketOutputWithContext(ctx context.Context) BucketOutput {
	return o
}

// The user set of grant full control.
func (o BucketOutput) AccountAcls() BucketAccountAclArrayOutput {
	return o.ApplyT(func(v *Bucket) BucketAccountAclArrayOutput { return v.AccountAcls }).(BucketAccountAclArrayOutput)
}

// The name of the bucket.
func (o BucketOutput) BucketName() pulumi.StringOutput {
	return o.ApplyT(func(v *Bucket) pulumi.StringOutput { return v.BucketName }).(pulumi.StringOutput)
}

// The create date of the TOS bucket.
func (o BucketOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Bucket) pulumi.StringOutput { return v.CreationDate }).(pulumi.StringOutput)
}

// The flag of enable tos version.
func (o BucketOutput) EnableVersion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Bucket) pulumi.BoolPtrOutput { return v.EnableVersion }).(pulumi.BoolPtrOutput)
}

// The extranet endpoint of the TOS bucket.
func (o BucketOutput) ExtranetEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *Bucket) pulumi.StringOutput { return v.ExtranetEndpoint }).(pulumi.StringOutput)
}

// The intranet endpoint the TOS bucket.
func (o BucketOutput) IntranetEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *Bucket) pulumi.StringOutput { return v.IntranetEndpoint }).(pulumi.StringOutput)
}

// The location of the TOS bucket.
func (o BucketOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Bucket) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The public acl control of object.Valid value is private|public-read|public-read-write|authenticated-read|bucket-owner-read.
func (o BucketOutput) PublicAcl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Bucket) pulumi.StringPtrOutput { return v.PublicAcl }).(pulumi.StringPtrOutput)
}

// The storage type of the object.Valid value is STANDARD|IA|ARCHIVE_FR.Default is STANDARD.
func (o BucketOutput) StorageClass() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Bucket) pulumi.StringPtrOutput { return v.StorageClass }).(pulumi.StringPtrOutput)
}

type BucketArrayOutput struct{ *pulumi.OutputState }

func (BucketArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Bucket)(nil)).Elem()
}

func (o BucketArrayOutput) ToBucketArrayOutput() BucketArrayOutput {
	return o
}

func (o BucketArrayOutput) ToBucketArrayOutputWithContext(ctx context.Context) BucketArrayOutput {
	return o
}

func (o BucketArrayOutput) Index(i pulumi.IntInput) BucketOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Bucket {
		return vs[0].([]*Bucket)[vs[1].(int)]
	}).(BucketOutput)
}

type BucketMapOutput struct{ *pulumi.OutputState }

func (BucketMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Bucket)(nil)).Elem()
}

func (o BucketMapOutput) ToBucketMapOutput() BucketMapOutput {
	return o
}

func (o BucketMapOutput) ToBucketMapOutputWithContext(ctx context.Context) BucketMapOutput {
	return o
}

func (o BucketMapOutput) MapIndex(k pulumi.StringInput) BucketOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Bucket {
		return vs[0].(map[string]*Bucket)[vs[1].(string)]
	}).(BucketOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BucketInput)(nil)).Elem(), &Bucket{})
	pulumi.RegisterInputType(reflect.TypeOf((*BucketArrayInput)(nil)).Elem(), BucketArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BucketMapInput)(nil)).Elem(), BucketMap{})
	pulumi.RegisterOutputType(BucketOutput{})
	pulumi.RegisterOutputType(BucketArrayOutput{})
	pulumi.RegisterOutputType(BucketMapOutput{})
}